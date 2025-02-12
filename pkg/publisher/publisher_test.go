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

package publisher

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/nholuongut/amazon-vpc-cni-k8s/pkg/utils/logger"
	"github.com/nholuongut/nholuongut-sdk-go/nholuongut"
	"github.com/nholuongut/nholuongut-sdk-go/service/cloudwatch"
	"github.com/nholuongut/nholuongut-sdk-go/service/cloudwatch/cloudwatchiface"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const (
	testClusterID       = "TEST_CLUSTER_ID"
	testMetricOne       = "TEST_METRIC_ONE"
	testMonitorDuration = time.Millisecond * 10
)

func TestCloudWatchPublisherWithNoIMDS(t *testing.T) {
	log := getCloudWatchLog()
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	region := "us-west-2"
	clusterID := testClusterID

	cw, err := New(ctx, region, clusterID, log)
	assert.NoError(t, err)
	assert.NotNil(t, cw)
}

func TestCloudWatchPublisherWithSingleDatum(t *testing.T) {
	cloudwatchPublisher := getCloudWatchPublisher(t)

	testCloudwatchMetricDatum := &cloudwatch.MetricDatum{
		MetricName: nholuongut.String(testMetricOne),
		Unit:       nholuongut.String(cloudwatch.StandardUnitNone),
		Value:      nholuongut.Float64(1.0),
	}

	cloudwatchPublisher.Publish(testCloudwatchMetricDatum)
	assert.Len(t, cloudwatchPublisher.localMetricData, 1)
	assert.EqualValues(t, cloudwatchPublisher.localMetricData[0], testCloudwatchMetricDatum)

	cloudwatchPublisher.pushLocal()
	assert.Empty(t, cloudwatchPublisher.localMetricData)
}

func TestCloudWatchPublisherWithMultipleDatum(t *testing.T) {
	cloudwatchPublisher := getCloudWatchPublisher(t)

	var metricDataPoints []*cloudwatch.MetricDatum

	for i := 0; i < 10; i++ {
		metricName := "TEST_METRIC_" + strconv.Itoa(i)
		testCloudwatchMetricDatum := &cloudwatch.MetricDatum{
			MetricName: nholuongut.String(metricName),
			Unit:       nholuongut.String(cloudwatch.StandardUnitNone),
			Value:      nholuongut.Float64(1.0),
		}
		metricDataPoints = append(metricDataPoints, testCloudwatchMetricDatum)
	}

	cloudwatchPublisher.Publish(metricDataPoints...)
	assert.Len(t, cloudwatchPublisher.localMetricData, 10)
	cloudwatchPublisher.pushLocal()

	assert.Empty(t, cloudwatchPublisher.localMetricData)
}

func TestCloudWatchPublisherWithGreaterThanMaxDatapoints(t *testing.T) {
	cloudwatchPublisher := getCloudWatchPublisher(t)

	var metricDataPoints []*cloudwatch.MetricDatum

	for i := 0; i < 30; i++ {
		metricName := "TEST_METRIC_" + strconv.Itoa(i)
		testCloudwatchMetricDatum := &cloudwatch.MetricDatum{
			MetricName: nholuongut.String(metricName),
			Unit:       nholuongut.String(cloudwatch.StandardUnitNone),
			Value:      nholuongut.Float64(1.0),
		}
		metricDataPoints = append(metricDataPoints, testCloudwatchMetricDatum)
	}

	cloudwatchPublisher.Publish(metricDataPoints...)
	assert.Len(t, cloudwatchPublisher.localMetricData, 30)
	cloudwatchPublisher.pushLocal()

	assert.Empty(t, cloudwatchPublisher.localMetricData)
}

func TestCloudWatchPublisherWithGreaterThanMaxDatapointsAndStop(t *testing.T) {
	cloudwatchPublisher := getCloudWatchPublisher(t)

	var metricDataPoints []*cloudwatch.MetricDatum
	for i := 0; i < 30; i++ {
		metricName := "TEST_METRIC_" + strconv.Itoa(i)
		testCloudwatchMetricDatum := &cloudwatch.MetricDatum{
			MetricName: nholuongut.String(metricName),
			Unit:       nholuongut.String(cloudwatch.StandardUnitNone),
			Value:      nholuongut.Float64(1.0),
		}
		metricDataPoints = append(metricDataPoints, testCloudwatchMetricDatum)
	}

	cloudwatchPublisher.Publish(metricDataPoints...)
	assert.Len(t, cloudwatchPublisher.localMetricData, 30)

	go cloudwatchPublisher.monitor(testMonitorDuration)

	// Delays added to prevent test flakiness
	<-time.After(5 * testMonitorDuration)
	cloudwatchPublisher.Stop()
	<-time.After(5 * testMonitorDuration)

	assert.Empty(t, cloudwatchPublisher.localMetricData)
}

func TestCloudWatchPublisherWithSingleDatumWithError(t *testing.T) {
	derivedContext, cancel := context.WithCancel(context.TODO())

	mockCloudWatch := mockCloudWatchClient{mockPutMetricDataError: errors.New("test error")}

	cloudwatchPublisher := &cloudWatchPublisher{
		ctx:              derivedContext,
		cancel:           cancel,
		cloudwatchClient: mockCloudWatch,
		clusterID:        testClusterID,
		localMetricData:  make([]*cloudwatch.MetricDatum, 0, localMetricDataSize),
		log:              getCloudWatchLog(),
	}

	testCloudwatchMetricDatum := &cloudwatch.MetricDatum{
		MetricName: nholuongut.String(testMetricOne),
		Unit:       nholuongut.String(cloudwatch.StandardUnitNone),
		Value:      nholuongut.Float64(1.0),
	}

	cloudwatchPublisher.Publish(testCloudwatchMetricDatum)
	assert.Len(t, cloudwatchPublisher.localMetricData, 1)
	assert.EqualValues(t, cloudwatchPublisher.localMetricData[0], testCloudwatchMetricDatum)

	cloudwatchPublisher.pushLocal()
	assert.Empty(t, cloudwatchPublisher.localMetricData)
}

func TestGetCloudWatchMetricNamespace(t *testing.T) {
	cloudwatchPublisher := getCloudWatchPublisher(t)

	testNamespace := cloudwatchPublisher.getCloudWatchMetricNamespace()
	assert.Equal(t, nholuongut.StringValue(testNamespace), cloudwatchMetricNamespace)
}

func TestGetCloudWatchMetricDatumDimensions(t *testing.T) {
	cloudwatchPublisher := getCloudWatchPublisher(t)

	expectedCloudwatchDimensions := []*cloudwatch.Dimension{
		{
			Name:  nholuongut.String(clusterIDDimension),
			Value: nholuongut.String(testClusterID),
		},
	}
	testCloudwatchDimensions := cloudwatchPublisher.getCloudWatchMetricDatumDimensions()

	assert.Equal(t, testCloudwatchDimensions, expectedCloudwatchDimensions)
}

func TestGetCloudWatchMetricDatumDimensionsWithMissingClusterID(t *testing.T) {
	cloudwatchPublisher := &cloudWatchPublisher{log: getCloudWatchLog()}

	expectedCloudwatchDimensions := []*cloudwatch.Dimension{
		{
			Name:  nholuongut.String(clusterIDDimension),
			Value: nholuongut.String(""),
		},
	}
	testCloudwatchDimensions := cloudwatchPublisher.getCloudWatchMetricDatumDimensions()

	assert.Equal(t, testCloudwatchDimensions, expectedCloudwatchDimensions)
}

func TestPublishWithNoData(t *testing.T) {
	cloudwatchPublisher := &cloudWatchPublisher{log: getCloudWatchLog()}

	testMetricDataPoints := []*cloudwatch.MetricDatum{}

	cloudwatchPublisher.Publish(testMetricDataPoints...)
	assert.Empty(t, cloudwatchPublisher.localMetricData)
}

func TestPushWithMissingData(t *testing.T) {
	cloudwatchPublisher := &cloudWatchPublisher{log: getCloudWatchLog()}
	testMetricDataPoints := []*cloudwatch.MetricDatum{}

	cloudwatchPublisher.push(testMetricDataPoints)
	assert.Empty(t, cloudwatchPublisher.localMetricData)
}

func TestMin(t *testing.T) {
	a, b := 1, 2

	minimum := min(a, b)
	assert.Equal(t, minimum, a)

	minimum = min(b, a)
	assert.Equal(t, minimum, a)
}

// mockCloudWatchClient is used to facilitate testing
type mockCloudWatchClient struct {
	cloudwatchiface.CloudWatchAPI
	mockPutMetricDataError error
}

func (m mockCloudWatchClient) PutMetricData(input *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	return &cloudwatch.PutMetricDataOutput{}, m.mockPutMetricDataError
}

func getCloudWatchLog() logger.Logger {
	logConfig := logger.Configuration{
		LogLevel:    "Debug",
		LogLocation: "stdout",
	}
	return logger.New(&logConfig)
}

func getCloudWatchPublisher(t *testing.T) *cloudWatchPublisher {
	// Setup context
	derivedContext, cancel := context.WithCancel(context.TODO())

	return &cloudWatchPublisher{
		ctx:              derivedContext,
		cancel:           cancel,
		cloudwatchClient: mockCloudWatchClient{},
		clusterID:        testClusterID,
		localMetricData:  make([]*cloudwatch.MetricDatum, 0, localMetricDataSize),
		log:              getCloudWatchLog(),
	}
}
