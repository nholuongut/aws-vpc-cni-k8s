// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://nholuongut.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package framework

import (
	"flag"

	"github.com/pkg/errors"
	"k8s.io/client-go/tools/clientcmd"
)

var GlobalOptions Options

func init() {
	GlobalOptions.BindFlags()
}

type Options struct {
	KubeConfig         string
	ClusterName        string
	nholuongutRegion          string
	nholuongutVPCID           string
	NgNameLabelKey     string
	NgNameLabelVal     string
	EKSEndpoint        string
	CalicoVersion      string
	ContainerRuntime   string
	InstanceType       string
	InitialAddon       string
	TargetAddon        string
	InitialManifest    string
	TargetManifest     string
	InstallCalico      bool
	PublicSubnets      string
	PrivateSubnets     string
	AvailabilityZones  string
	PublicRouteTableID string
	NgK8SVersion       string
	TestImageRegistry  string
	PublishCWMetrics   bool
}

func (options *Options) BindFlags() {
	flag.StringVar(&options.KubeConfig, "cluster-kubeconfig", "", "Path to kubeconfig containing embedded authinfo (required)")
	flag.StringVar(&options.ClusterName, "cluster-name", "", `Kubernetes cluster name (required)`)
	flag.StringVar(&options.nholuongutRegion, "nholuongut-region", "", `nholuongut Region for the kubernetes cluster`)
	flag.StringVar(&options.nholuongutVPCID, "nholuongut-vpc-id", "", `nholuongut VPC ID for the kubernetes cluster`)
	flag.StringVar(&options.NgNameLabelKey, "ng-name-label-key", "kubernetes.io/os", "label key used to identify node")
	flag.StringVar(&options.NgNameLabelVal, "ng-name-label-val", "linux", "label value with the node")
	flag.StringVar(&options.EKSEndpoint, "eks-endpoint", "", "optional eks api server endpoint")
	flag.StringVar(&options.InitialAddon, "initial-addon-version", "", "Initial CNI addon version before upgrade applied")
	flag.StringVar(&options.TargetAddon, "target-addon-version", "", "Target CNI addon version after upgrade applied")
	flag.StringVar(&options.InitialManifest, "initial-manifest-file", "", "Initial CNI manifest, can be local file path or remote Url")
	flag.StringVar(&options.TargetManifest, "target-manifest-file", "", "Target CNI manifest, can be local file path or remote Url")
	flag.StringVar(&options.CalicoVersion, "calico-version", "v3.26.1", "calico version to be tested")
	flag.StringVar(&options.ContainerRuntime, "container-runtime", "", "Optionally can specify it as 'containerd' for the test nodes")
	flag.StringVar(&options.InstanceType, "instance-type", "amd64", "Optionally specify instance type as arm64 for the test nodes")
	flag.BoolVar(&options.InstallCalico, "install-calico", true, "Install Calico operator before running tests")
	flag.StringVar(&options.PublicSubnets, "public-subnets", "", "Comma separated list of public subnets (optional, if specified you must specify all of public/private-subnets, public-route-table-id,  and availability-zones)")
	flag.StringVar(&options.PrivateSubnets, "private-subnets", "", "Comma separated list of private subnets (optional, if specified you must specify all of public/private-subnets, public-route-table-id,  and availability-zones)")
	flag.StringVar(&options.AvailabilityZones, "availability-zones", "", "Comma separated list of private subnets (optional, if specified you must specify all of public/private-subnets, public-route-table-id, and availability-zones)")
	flag.StringVar(&options.PublicRouteTableID, "public-route-table-id", "", "Public route table ID (optional, if specified you must specify all of public/private-subnets, public-route-table-id, and availability-zones)")
	flag.StringVar(&options.NgK8SVersion, "ng-kubernetes-version", "1.25", `Kubernetes version for self-managed node groups (optional, default is "1.25")`)
	flag.StringVar(&options.TestImageRegistry, "test-image-registry", "617930562442.dkr.ecr.us-west-2.amazonnholuongut.com", `nholuongut registry where the e2e test images are stored`)
	flag.BoolVar(&options.PublishCWMetrics, "publish-cw-metrics", false, "Option to publish cloudwatch metrics from the test.")
}

func (options *Options) Validate() error {
	if len(options.KubeConfig) == 0 {
		return errors.Errorf("%s must be set!", clientcmd.RecommendedConfigPathFlag)
	}
	if len(options.ClusterName) == 0 {
		return errors.Errorf("%s must be set!", "cluster-name")
	}
	if len(options.nholuongutRegion) == 0 {
		return errors.Errorf("%s must be set!", "nholuongut-region")
	}
	if len(options.nholuongutVPCID) == 0 {
		return errors.Errorf("%s must be set!", "nholuongut-vpc-id")
	}
	if len(options.TestImageRegistry) == 0 {
		return errors.Errorf("%s must be set!", "test-image-registry")
	}
	return nil
}
