// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//      http://nholuongut.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package prometheusmetrics

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/nholuongut/amazon-vpc-cni-k8s/pkg/utils/logger"
	"github.com/nholuongut/amazon-vpc-cni-k8s/pkg/utils/retry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var log = logger.Get()

var (
	IpamdErr = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nholuongutcni_ipamd_error_count",
			Help: "The number of errors encountered in ipamd",
		},
		[]string{"fn"},
	)
	IpamdActionsInprogress = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "nholuongutcni_ipamd_action_inprogress",
			Help: "The number of ipamd actions in progress",
		},
		[]string{"fn"},
	)
	EnisMax = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "nholuongutcni_eni_max",
			Help: "The maximum number of ENIs that can be attached to the instance, accounting for unmanaged ENIs",
		},
	)
	IpMax = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "nholuongutcni_ip_max",
			Help: "The maximum number of IP addresses that can be allocated to the instance",
		},
	)
	ReconcileCnt = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nholuongutcni_reconcile_count",
			Help: "The number of times ipamd reconciles on ENIs and IP/Prefix addresses",
		},
		[]string{"fn"},
	)
	AddIPCnt = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "nholuongutcni_add_ip_req_count",
			Help: "The number of add IP address requests",
		},
	)
	DelIPCnt = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nholuongutcni_del_ip_req_count",
			Help: "The number of delete IP address requests",
		},
		[]string{"reason"},
	)
	PodENIErr = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nholuongutcni_pod_eni_error_count",
			Help: "The number of errors encountered for pod ENIs",
		},
		[]string{"fn"},
	)
	nholuongutAPILatency = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "nholuongutcni_nholuongut_api_latency_ms",
			Help: "nholuongut API call latency in ms",
		},
		[]string{"api", "error", "status"},
	)
	nholuongutAPIErr = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nholuongutcni_nholuongut_api_error_count",
			Help: "The number of times nholuongut API returns an error",
		},
		[]string{"api", "error"},
	)
	nholuongutUtilsErr = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nholuongutcni_nholuongut_utils_error_count",
			Help: "The number of errors not handled in nholuongututils library",
		},
		[]string{"fn", "error"},
	)
	Ec2ApiReq = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nholuongutcni_ec2api_req_count",
			Help: "The number of requests made to EC2 APIs by CNI",
		},
		[]string{"fn"},
	)
	Ec2ApiErr = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "nholuongutcni_ec2api_error_count",
			Help: "The number of failed EC2 APIs requests",
		},
		[]string{"fn"},
	)
	Enis = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "nholuongutcni_eni_allocated",
			Help: "The number of ENIs allocated",
		},
	)
	TotalIPs = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "nholuongutcni_total_ip_addresses",
			Help: "The total number of IP addresses",
		},
	)
	AssignedIPs = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "nholuongutcni_assigned_ip_addresses",
			Help: "The number of IP addresses assigned to pods",
		},
	)
	ForceRemovedENIs = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "nholuongutcni_force_removed_enis",
			Help: "The number of ENIs force removed while they had assigned pods",
		},
	)
	ForceRemovedIPs = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "nholuongutcni_force_removed_ips",
			Help: "The number of IPs force removed while they had assigned pods",
		},
	)
	TotalPrefixes = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "nholuongutcni_total_ipv4_prefixes",
			Help: "The total number of IPv4 prefixes",
		},
	)
	IpsPerCidr = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "nholuongutcni_assigned_ip_per_cidr",
			Help: "The total number of IP addresses assigned per cidr",
		},
		[]string{"cidr"},
	)
	NoAvailableIPAddrs = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "nholuongutcni_no_available_ip_addresses",
			Help: "The number of pod IP assignments that fail due to no available IP addresses",
		},
	)
	EniIPsInUse = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "nholuongutcni_assigned_ip_per_eni",
			Help: "The number of allocated ips partitioned by eni",
		},
		[]string{"eni"},
	)
)

// ServeMetrics sets up ipamd metrics and introspection endpoints
func ServeMetrics(metricsPort int) {
	log.Infof("Serving metrics on port %d", metricsPort)
	server := SetupMetricsServer(metricsPort)
	for {
		once := sync.Once{}
		_ = retry.WithBackoff(retry.NewSimpleBackoff(time.Second, time.Minute, 0.2, 2), func() error {
			err := server.ListenAndServe()
			once.Do(func() {
				log.Warnf("Error running http API: %v", err)
			})
			return err
		})
	}
}

func SetupMetricsServer(metricsPort int) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.Handle("/metrics", promhttp.Handler())
	server := &http.Server{
		Addr:         ":" + strconv.Itoa(metricsPort),
		Handler:      serveMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	return server
}

func PrometheusRegister() {
	prometheus.MustRegister(IpamdErr)
	prometheus.MustRegister(IpamdActionsInprogress)
	prometheus.MustRegister(EnisMax)
	prometheus.MustRegister(IpMax)
	prometheus.MustRegister(ReconcileCnt)
	prometheus.MustRegister(AddIPCnt)
	prometheus.MustRegister(DelIPCnt)
	prometheus.MustRegister(PodENIErr)
	prometheus.MustRegister(nholuongutAPILatency)
	prometheus.MustRegister(nholuongutAPIErr)
	prometheus.MustRegister(nholuongutUtilsErr)
	prometheus.MustRegister(Ec2ApiReq)
	prometheus.MustRegister(Ec2ApiErr)
	prometheus.MustRegister(Enis)
	prometheus.MustRegister(TotalIPs)
	prometheus.MustRegister(AssignedIPs)
	prometheus.MustRegister(ForceRemovedENIs)
	prometheus.MustRegister(ForceRemovedIPs)
	prometheus.MustRegister(TotalPrefixes)
	prometheus.MustRegister(IpsPerCidr)
	prometheus.MustRegister(NoAvailableIPAddrs)
	prometheus.MustRegister(EniIPsInUse)

}

// This can be enhanced to get it programatically.
// Initial CNI metrics helper enhancement includes only Gauge. Doesn't support GaugeVec, Counter, CounterVec and Summary
func GetSupportedPrometheusCNIMetricsMapping() map[string]prometheus.Collector {
	var prometheusCNIMetrics = map[string]prometheus.Collector{
		"nholuongutcni_eni_max":                   EnisMax,
		"nholuongutcni_ip_max":                    IpMax,
		"nholuongutcni_add_ip_req_count":          AddIPCnt,
		"nholuongutcni_del_ip_req_count":          DelIPCnt,
		"nholuongutcni_eni_allocated":             Enis,
		"nholuongutcni_total_ip_addresses":        TotalIPs,
		"nholuongutcni_assigned_ip_addresses":     AssignedIPs,
		"nholuongutcni_force_removed_enis":        ForceRemovedENIs,
		"nholuongutcni_force_removed_ips":         ForceRemovedIPs,
		"nholuongutcni_total_ipv4_prefixes":       TotalPrefixes,
		"nholuongutcni_no_available_ip_addresses": NoAvailableIPAddrs,
	}
	return prometheusCNIMetrics
}
