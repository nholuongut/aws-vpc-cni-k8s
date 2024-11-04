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

package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/nholuongut/nholuongut-sdk-go/nholuongut"
	"github.com/nholuongut/nholuongut-sdk-go/service/cloudformation"

	"github.com/nholuongut/amazon-vpc-cni-k8s/pkg/vpc"
	"github.com/nholuongut/amazon-vpc-cni-k8s/test/framework"
	k8sUtils "github.com/nholuongut/amazon-vpc-cni-k8s/test/framework/resources/k8s/utils"
	"github.com/nholuongut/amazon-vpc-cni-k8s/test/framework/utils"
)

const (
	// Docker will be default, if not specified
	CONTAINERD                 = "containerd"
	CreateNodeGroupCFNTemplate = "/testdata/amazon-eks-nodegroup.yaml"
	NodeImageIdSSMParam        = "/nholuongut/service/eks/optimized-ami/%s/amazon-linux-2/recommended/image_id"
)

type NodeGroupProperties struct {
	// Required to verify the node is up and ready
	NgLabelKey string
	NgLabelVal string
	// ASG Size
	AsgSize       int
	NodeGroupName string
	// If custom networking is set then max pod
	// will be set on Kubelet extra arguments
	IsCustomNetworkingEnabled bool
	// Subnet where the node group will be created
	Subnet       []string
	InstanceType string
	KeyPairName  string

	// optional: specify container runtime
	ContainerRuntime string

	NodeImageId string
}

type ClusterVPCConfig struct {
	PublicSubnetList   []string
	AvailZones         []string
	PublicRouteTableID string
	PrivateSubnetList  []string
}

type nholuongutAuthMapRole struct {
	Groups   []string `yaml:"groups"`
	RoleArn  string   `yaml:"rolearn"`
	UserName string   `yaml:"username"`
}

// Create self managed node group stack
func CreateAndWaitTillSelfManagedNGReady(f *framework.Framework, properties NodeGroupProperties) error {
	templatePath := utils.GetProjectRoot() + CreateNodeGroupCFNTemplate
	templateBytes, err := os.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("failed to read from %s, %v", templatePath, err)
	}
	template := string(templateBytes)

	describeClusterOutput, err := f.CloudServices.EKS().DescribeCluster(f.Options.ClusterName)
	if err != nil {
		return fmt.Errorf("failed to describe cluster %s: %v", f.Options.ClusterName, err)
	}

	var bootstrapArgs = fmt.Sprintf("--apiserver-endpoint %s --b64-cluster-ca %s",
		*describeClusterOutput.Cluster.Endpoint, *describeClusterOutput.Cluster.CertificateAuthority.Data)
	var kubeletExtraArgs = fmt.Sprintf("--node-labels=%s=%s", properties.NgLabelKey, properties.NgLabelVal)

	if properties.IsCustomNetworkingEnabled {
		limit, _ := vpc.GetInstance(properties.InstanceType)
		maxPods := (limit.ENILimit-1)*(limit.IPv4Limit-1) + 2

		bootstrapArgs += " --use-max-pods false"
		kubeletExtraArgs += fmt.Sprintf(" --max-pods=%d", maxPods)
	}

	containerRuntime := properties.ContainerRuntime
	if containerRuntime != "" {
		bootstrapArgs += fmt.Sprintf(" --container-runtime %s", containerRuntime)
	}

	asgSizeString := strconv.Itoa(properties.AsgSize)

	createNgStackParams := []*cloudformation.Parameter{
		{
			ParameterKey:   nholuongut.String("ClusterName"),
			ParameterValue: nholuongut.String(f.Options.ClusterName),
		},
		{
			ParameterKey:   nholuongut.String("VpcId"),
			ParameterValue: nholuongut.String(f.Options.nholuongutVPCID),
		},
		{
			ParameterKey:   nholuongut.String("Subnets"),
			ParameterValue: nholuongut.String(strings.Join(properties.Subnet, ",")),
		},
		{
			ParameterKey:   nholuongut.String("ClusterControlPlaneSecurityGroup"),
			ParameterValue: describeClusterOutput.Cluster.ResourcesVpcConfig.SecurityGroupIds[0],
		},
		{
			ParameterKey:   nholuongut.String("NodeGroupName"),
			ParameterValue: nholuongut.String(properties.NodeGroupName),
		},
		{
			ParameterKey:   nholuongut.String("NodeImageIdSSMParam"),
			ParameterValue: nholuongut.String(fmt.Sprintf(NodeImageIdSSMParam, f.Options.NgK8SVersion)),
		},
		{
			ParameterKey:   nholuongut.String("NodeAutoScalingGroupMinSize"),
			ParameterValue: nholuongut.String(asgSizeString),
		},
		{
			ParameterKey:   nholuongut.String("NodeAutoScalingGroupDesiredCapacity"),
			ParameterValue: nholuongut.String(asgSizeString),
		},
		{
			ParameterKey:   nholuongut.String("NodeAutoScalingGroupMaxSize"),
			ParameterValue: nholuongut.String(asgSizeString),
		},
		{
			ParameterKey:   nholuongut.String("NodeInstanceType"),
			ParameterValue: nholuongut.String(properties.InstanceType),
		},
		{
			ParameterKey:   nholuongut.String("BootstrapArguments"),
			ParameterValue: nholuongut.String(fmt.Sprintf("%s --kubelet-extra-args '%s'", bootstrapArgs, kubeletExtraArgs)),
		},
		{
			ParameterKey:   nholuongut.String("KeyName"),
			ParameterValue: nholuongut.String(properties.KeyPairName),
		},
		{
			ParameterKey:   nholuongut.String("DisableIMDSv1"),
			ParameterValue: nholuongut.String("true"),
		},
	}

	if properties.NodeImageId != "" {
		createNgStackParams = append(createNgStackParams, &cloudformation.Parameter{
			ParameterKey:   nholuongut.String("NodeImageId"),
			ParameterValue: nholuongut.String(properties.NodeImageId),
		})
	}

	describeStackOutput, err := f.CloudServices.CloudFormation().
		WaitTillStackCreated(properties.NodeGroupName, createNgStackParams, template)
	if err != nil {
		return fmt.Errorf("failed to create node group cfn stack: %v", err)
	}

	var nodeInstanceRole string
	for _, stackOutput := range describeStackOutput.Stacks[0].Outputs {
		if *stackOutput.OutputKey == "NodeInstanceRole" {
			nodeInstanceRole = *stackOutput.OutputValue
		}
	}

	if nodeInstanceRole == "" {
		return fmt.Errorf("failed to find node instance role in stack %+v", describeStackOutput)
	}

	// Update the nholuongut Auth Config with the Node Instance Role
	nholuongutAuth, err := f.K8sResourceManagers.ConfigMapManager().
		GetConfigMap("kube-system", "nholuongut-auth")
	if err != nil {
		return fmt.Errorf("failed to find nholuongut-auth configmap: %v", err)
	}

	updatednholuongutAuth := nholuongutAuth.DeepCopy()
	authMapRole := []nholuongutAuthMapRole{
		{
			Groups:   []string{"system:bootstrappers", "system:nodes"},
			RoleArn:  nodeInstanceRole,
			UserName: "system:node:{{EC2PrivateDNSName}}",
		},
	}
	yamlBytes, err := yaml.Marshal(authMapRole)

	updatednholuongutAuth.Data["mapRoles"] = updatednholuongutAuth.Data["mapRoles"] + string(yamlBytes)

	err = f.K8sResourceManagers.ConfigMapManager().UpdateConfigMap(nholuongutAuth, updatednholuongutAuth)
	if err != nil {
		return fmt.Errorf("failed to update the auth config with new node's instance role: %v", err)
	}

	// Wait till the node group have joined the cluster and are ready
	err = f.K8sResourceManagers.NodeManager().
		WaitTillNodesReady(properties.NgLabelKey, properties.NgLabelVal, properties.AsgSize)
	if err != nil {
		return fmt.Errorf("failed to list nodegroup with label key %s:%v: %v",
			properties.NgLabelKey, properties.NgLabelVal, err)
	}

	return nil
}

func DeleteAndWaitTillSelfManagedNGStackDeleted(f *framework.Framework, properties NodeGroupProperties) error {
	err := f.CloudServices.CloudFormation().WaitTillStackDeleted(properties.NodeGroupName)
	if err != nil {
		return fmt.Errorf("failed to delete node group cfn stack: %v", err)
	}
	return nil
}

func GetClusterVPCConfig(f *framework.Framework) (*ClusterVPCConfig, error) {
	clusterConfig := &ClusterVPCConfig{
		PublicSubnetList:  []string{},
		AvailZones:        []string{},
		PrivateSubnetList: []string{},
	}

	if len(f.Options.PublicSubnets) > 0 {
		clusterConfig.PublicSubnetList = strings.Split(f.Options.PublicSubnets, ",")
	}
	if len(f.Options.PrivateSubnets) > 0 {
		clusterConfig.PrivateSubnetList = strings.Split(f.Options.PrivateSubnets, ",")
	}
	if len(f.Options.AvailabilityZones) > 0 {
		clusterConfig.AvailZones = strings.Split(f.Options.AvailabilityZones, ",")
	}
	if f.Options.PublicRouteTableID != "" {
		clusterConfig.PublicRouteTableID = f.Options.PublicRouteTableID
	}

	// user provided the info so we don't need to look it up
	if clusterConfig.PublicRouteTableID != "" && len(clusterConfig.PublicSubnetList) > 0 && len(clusterConfig.AvailZones) > 0 {
		return clusterConfig, nil
	}

	if clusterConfig.PublicRouteTableID != "" || len(clusterConfig.PublicSubnetList) > 0 ||
		len(clusterConfig.PrivateSubnetList) > 0 || len(clusterConfig.AvailZones) > 0 {
		return nil, fmt.Errorf("partial configuration, if supplying config via flags you need to provide at least public route table ID, public subnet list and availibility zone list")
	}

	describeClusterOutput, err := f.CloudServices.EKS().DescribeCluster(f.Options.ClusterName)
	if err != nil {
		return nil, fmt.Errorf("failed to describe cluster %s: %v", f.Options.ClusterName, err)
	}

	for _, subnet := range describeClusterOutput.Cluster.ResourcesVpcConfig.SubnetIds {
		describeRouteOutput, err := f.CloudServices.EC2().DescribeRouteTables(*subnet)
		if err != nil {
			return nil, fmt.Errorf("failed to describe subnet %s: %v", *subnet, err)
		}

		isPublic := false
		for _, route := range describeRouteOutput.RouteTables[0].Routes {
			if route.GatewayId != nil && strings.Contains(*route.GatewayId, "igw-") {
				isPublic = true
				clusterConfig.PublicSubnetList = append(clusterConfig.PublicSubnetList, *subnet)
				clusterConfig.PublicRouteTableID = *describeRouteOutput.RouteTables[0].RouteTableId
			}
		}
		if !isPublic {
			clusterConfig.PrivateSubnetList = append(clusterConfig.PrivateSubnetList, *subnet)
		}
	}

	uniqueAZ := map[string]bool{}
	for _, subnet := range clusterConfig.PublicSubnetList {
		describeSubnet, err := f.CloudServices.EC2().DescribeSubnet(subnet)
		if err != nil {
			return nil, fmt.Errorf("failed to describe the subnet %s: %v", subnet, err)
		}
		if ok := uniqueAZ[*describeSubnet.Subnets[0].AvailabilityZone]; !ok {
			uniqueAZ[*describeSubnet.Subnets[0].AvailabilityZone] = true
			clusterConfig.AvailZones =
				append(clusterConfig.AvailZones, *describeSubnet.Subnets[0].AvailabilityZone)
		}
	}

	return clusterConfig, nil
}

func TerminateInstances(f *framework.Framework) error {
	nodeList, err := f.K8sResourceManagers.NodeManager().GetNodes(f.Options.NgNameLabelKey, f.Options.NgNameLabelVal)
	if err != nil {
		return fmt.Errorf("failed to get list of nodes created: %v", err)
	}

	var instanceIDs []string
	for _, node := range nodeList.Items {
		instanceIDs = append(instanceIDs, k8sUtils.GetInstanceIDFromNode(node))
	}

	err = f.CloudServices.EC2().TerminateInstance(instanceIDs)
	if err != nil {
		return fmt.Errorf("failed to terminate instances: %v", err)
	}

	// Wait for instances to be replaced
	time.Sleep(time.Minute * 8)
	return nil
}
