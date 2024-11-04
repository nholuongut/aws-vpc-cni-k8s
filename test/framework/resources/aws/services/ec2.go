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

package services

import (
	"fmt"
	"strings"

	"github.com/nholuongut/nholuongut-sdk-go/nholuongut"
	"github.com/nholuongut/nholuongut-sdk-go/nholuongut/session"
	"github.com/nholuongut/nholuongut-sdk-go/service/ec2"
	"github.com/nholuongut/nholuongut-sdk-go/service/ec2/ec2iface"
)

type EC2 interface {
	DescribeInstanceType(instanceType string) ([]*ec2.InstanceTypeInfo, error)
	DescribeInstance(instanceID string) (*ec2.Instance, error)
	DescribeVPC(vpcID string) (*ec2.DescribeVpcsOutput, error)
	DescribeNetworkInterface(interfaceIDs []string) (*ec2.DescribeNetworkInterfacesOutput, error)
	AuthorizeSecurityGroupIngress(groupID string, protocol string, fromPort int, toPort int, cidrIP string, sourceSG bool) error
	RevokeSecurityGroupIngress(groupID string, protocol string, fromPort int, toPort int, cidrIP string, sourceSG bool) error
	AuthorizeSecurityGroupEgress(groupID string, protocol string, fromPort int, toPort int, cidrIP string) error
	RevokeSecurityGroupEgress(groupID string, protocol string, fromPort int, toPort int, cidrIP string) error
	AssociateVPCCIDRBlock(vpcId string, cidrBlock string) (*ec2.AssociateVpcCidrBlockOutput, error)
	TerminateInstance(instanceIDs []string) error
	DisAssociateVPCCIDRBlock(associationID string) error
	DescribeSubnet(subnetID string) (*ec2.DescribeSubnetsOutput, error)
	CreateSubnet(cidrBlock string, vpcID string, az string) (*ec2.CreateSubnetOutput, error)
	DeleteSubnet(subnetID string) error
	DescribeRouteTables(subnetID string) (*ec2.DescribeRouteTablesOutput, error)
	DescribeRouteTablesWithVPCID(vpcID string) (*ec2.DescribeRouteTablesOutput, error)
	CreateSecurityGroup(groupName string, description string, vpcID string) (*ec2.CreateSecurityGroupOutput, error)
	DeleteSecurityGroup(groupID string) error
	AssociateRouteTableToSubnet(routeTableId string, subnetID string) error
	CreateKey(keyName string) (*ec2.CreateKeyPairOutput, error)
	DeleteKey(keyName string) error
	DescribeKey(keyName string) (*ec2.DescribeKeyPairsOutput, error)
	ModifyNetworkInterfaceSecurityGroups(securityGroupIds []*string, networkInterfaceId *string) (*ec2.ModifyNetworkInterfaceAttributeOutput, error)
	DescribeAvailabilityZones() (*ec2.DescribeAvailabilityZonesOutput, error)
	CreateTags(resourceIds []string, tags []*ec2.Tag) (*ec2.CreateTagsOutput, error)
	DeleteTags(resourceIds []string, tags []*ec2.Tag) (*ec2.DeleteTagsOutput, error)
}

type defaultEC2 struct {
	ec2iface.EC2API
}

func (d *defaultEC2) DescribeInstanceType(instanceType string) ([]*ec2.InstanceTypeInfo, error) {
	describeInstanceTypeIp := &ec2.DescribeInstanceTypesInput{
		InstanceTypes: nholuongut.StringSlice([]string{instanceType}),
	}
	describeInstanceOp, err := d.EC2API.DescribeInstanceTypes(describeInstanceTypeIp)
	if err != nil {
		return nil, err
	}
	if len(describeInstanceOp.InstanceTypes) == 0 {
		return nil, fmt.Errorf("no instance type found in the output %s", instanceType)
	}
	return describeInstanceOp.InstanceTypes, nil
}

func (d *defaultEC2) DescribeAvailabilityZones() (*ec2.DescribeAvailabilityZonesOutput, error) {
	describeAvailabilityZonesInput := &ec2.DescribeAvailabilityZonesInput{}
	return d.EC2API.DescribeAvailabilityZones(describeAvailabilityZonesInput)
}

func (d *defaultEC2) ModifyNetworkInterfaceSecurityGroups(securityGroupIds []*string, networkInterfaceId *string) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	return d.EC2API.ModifyNetworkInterfaceAttribute(&ec2.ModifyNetworkInterfaceAttributeInput{
		NetworkInterfaceId: networkInterfaceId,
		Groups:             securityGroupIds,
	})
}

func (d *defaultEC2) DescribeInstance(instanceID string) (*ec2.Instance, error) {
	describeInstanceInput := &ec2.DescribeInstancesInput{
		InstanceIds: nholuongut.StringSlice([]string{instanceID}),
	}
	describeInstanceOutput, err := d.EC2API.DescribeInstances(describeInstanceInput)
	if err != nil {
		return nil, err
	}
	if describeInstanceOutput == nil || len(describeInstanceOutput.Reservations) == 0 ||
		len(describeInstanceOutput.Reservations[0].Instances) == 0 {
		return nil, fmt.Errorf("failed to find instance %s", instanceID)
	}
	return describeInstanceOutput.Reservations[0].Instances[0], nil
}

func (d *defaultEC2) AuthorizeSecurityGroupIngress(groupID string, protocol string, fromPort int, toPort int, cidrIP string, sourceSG bool) error {
	var ipv4Ranges []*ec2.IpRange
	var ipv6Ranges []*ec2.Ipv6Range
	var ipPermissions *ec2.IpPermission
	if !sourceSG {
		if strings.Contains(cidrIP, ":") {
			ipv6Ranges = []*ec2.Ipv6Range{
				{
					CidrIpv6: nholuongut.String(cidrIP),
				},
			}
		} else {
			ipv4Ranges = []*ec2.IpRange{
				{
					CidrIp: nholuongut.String(cidrIP),
				},
			}
		}

		ipPermissions = &ec2.IpPermission{
			FromPort:   nholuongut.Int64(int64(fromPort)),
			ToPort:     nholuongut.Int64(int64(toPort)),
			IpProtocol: nholuongut.String(protocol),
			IpRanges:   ipv4Ranges,
			Ipv6Ranges: ipv6Ranges,
		}
	} else {
		ipPermissions = &ec2.IpPermission{
			FromPort:   nholuongut.Int64(int64(fromPort)),
			ToPort:     nholuongut.Int64(int64(toPort)),
			IpProtocol: nholuongut.String(protocol),
			UserIdGroupPairs: []*ec2.UserIdGroupPair{
				{
					GroupId: nholuongut.String(cidrIP),
				},
			},
		}
	}
	authorizeSecurityGroupIngressInput := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:       nholuongut.String(groupID),
		IpPermissions: []*ec2.IpPermission{ipPermissions},
	}
	_, err := d.EC2API.AuthorizeSecurityGroupIngress(authorizeSecurityGroupIngressInput)
	return err
}

func (d *defaultEC2) RevokeSecurityGroupIngress(groupID string, protocol string, fromPort int, toPort int, cidrIP string, sourceSG bool) error {
	var ipv4Ranges []*ec2.IpRange
	var ipv6Ranges []*ec2.Ipv6Range
	var ipPermissions *ec2.IpPermission
	if !sourceSG {
		if strings.Contains(cidrIP, ":") {
			ipv6Ranges = []*ec2.Ipv6Range{
				{
					CidrIpv6: nholuongut.String(cidrIP),
				},
			}
		} else {
			ipv4Ranges = []*ec2.IpRange{
				{
					CidrIp: nholuongut.String(cidrIP),
				},
			}
		}

		ipPermissions = &ec2.IpPermission{
			FromPort:   nholuongut.Int64(int64(fromPort)),
			ToPort:     nholuongut.Int64(int64(toPort)),
			IpProtocol: nholuongut.String(protocol),
			IpRanges:   ipv4Ranges,
			Ipv6Ranges: ipv6Ranges,
		}
	} else {
		ipPermissions = &ec2.IpPermission{
			FromPort:   nholuongut.Int64(int64(fromPort)),
			ToPort:     nholuongut.Int64(int64(toPort)),
			IpProtocol: nholuongut.String(protocol),
			UserIdGroupPairs: []*ec2.UserIdGroupPair{
				{
					GroupId: nholuongut.String(cidrIP),
				},
			},
		}
	}
	revokeSecurityGroupIngressInput := &ec2.RevokeSecurityGroupIngressInput{
		GroupId:       nholuongut.String(groupID),
		IpPermissions: []*ec2.IpPermission{ipPermissions},
	}
	_, err := d.EC2API.RevokeSecurityGroupIngress(revokeSecurityGroupIngressInput)
	return err
}

func (d *defaultEC2) AuthorizeSecurityGroupEgress(groupID string, protocol string, fromPort int, toPort int, cidrIP string) error {
	var ipv4Ranges []*ec2.IpRange
	var ipv6Ranges []*ec2.Ipv6Range
	if strings.Contains(cidrIP, ":") {
		ipv6Ranges = []*ec2.Ipv6Range{
			{
				CidrIpv6: nholuongut.String(cidrIP),
			},
		}
	} else {
		ipv4Ranges = []*ec2.IpRange{
			{
				CidrIp: nholuongut.String(cidrIP),
			},
		}
	}

	ipPermissions := &ec2.IpPermission{
		FromPort:   nholuongut.Int64(int64(fromPort)),
		ToPort:     nholuongut.Int64(int64(toPort)),
		IpProtocol: nholuongut.String(protocol),
		IpRanges:   ipv4Ranges,
		Ipv6Ranges: ipv6Ranges,
	}
	authorizeSecurityGroupEgressInput := &ec2.AuthorizeSecurityGroupEgressInput{
		GroupId:       nholuongut.String(groupID),
		IpPermissions: []*ec2.IpPermission{ipPermissions},
	}
	_, err := d.EC2API.AuthorizeSecurityGroupEgress(authorizeSecurityGroupEgressInput)
	return err
}

func (d *defaultEC2) RevokeSecurityGroupEgress(groupID string, protocol string, fromPort int, toPort int, cidrIP string) error {
	var ipv4Ranges []*ec2.IpRange
	var ipv6Ranges []*ec2.Ipv6Range
	if strings.Contains(cidrIP, ":") {
		ipv6Ranges = []*ec2.Ipv6Range{
			{
				CidrIpv6: nholuongut.String(cidrIP),
			},
		}
	} else {
		ipv4Ranges = []*ec2.IpRange{
			{
				CidrIp: nholuongut.String(cidrIP),
			},
		}
	}

	ipPermissions := &ec2.IpPermission{
		FromPort:   nholuongut.Int64(int64(fromPort)),
		ToPort:     nholuongut.Int64(int64(toPort)),
		IpProtocol: nholuongut.String(protocol),
		IpRanges:   ipv4Ranges,
		Ipv6Ranges: ipv6Ranges,
	}
	revokeSecurityGroupEgressInput := &ec2.RevokeSecurityGroupEgressInput{
		GroupId:       nholuongut.String(groupID),
		IpPermissions: []*ec2.IpPermission{ipPermissions},
	}
	_, err := d.EC2API.RevokeSecurityGroupEgress(revokeSecurityGroupEgressInput)
	return err
}

func (d *defaultEC2) DescribeNetworkInterface(interfaceIDs []string) (*ec2.DescribeNetworkInterfacesOutput, error) {
	describeNetworkInterfaceInput := &ec2.DescribeNetworkInterfacesInput{
		NetworkInterfaceIds: nholuongut.StringSlice(interfaceIDs),
	}

	return d.EC2API.DescribeNetworkInterfaces(describeNetworkInterfaceInput)
}

func (d *defaultEC2) AssociateVPCCIDRBlock(vpcId string, cidrBlock string) (*ec2.AssociateVpcCidrBlockOutput, error) {
	associateVPCCidrBlockInput := &ec2.AssociateVpcCidrBlockInput{
		CidrBlock: nholuongut.String(cidrBlock),
		VpcId:     nholuongut.String(vpcId),
	}

	return d.EC2API.AssociateVpcCidrBlock(associateVPCCidrBlockInput)
}

func (d *defaultEC2) DisAssociateVPCCIDRBlock(associationID string) error {
	disassociateVPCCidrBlockInput := &ec2.DisassociateVpcCidrBlockInput{
		AssociationId: nholuongut.String(associationID),
	}

	_, err := d.EC2API.DisassociateVpcCidrBlock(disassociateVPCCidrBlockInput)
	return err
}

func (d *defaultEC2) CreateSubnet(cidrBlock string, vpcID string, az string) (*ec2.CreateSubnetOutput, error) {
	createSubnetInput := &ec2.CreateSubnetInput{
		AvailabilityZone: nholuongut.String(az),
		CidrBlock:        nholuongut.String(cidrBlock),
		VpcId:            nholuongut.String(vpcID),
	}
	return d.EC2API.CreateSubnet(createSubnetInput)
}

func (d *defaultEC2) DescribeSubnet(subnetID string) (*ec2.DescribeSubnetsOutput, error) {
	describeSubnetInput := &ec2.DescribeSubnetsInput{
		SubnetIds: nholuongut.StringSlice([]string{subnetID}),
	}
	return d.EC2API.DescribeSubnets(describeSubnetInput)
}

func (d *defaultEC2) DescribeRouteTablesWithVPCID(vpcID string) (*ec2.DescribeRouteTablesOutput, error) {
	describeRouteTableInput := &ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{
				Name:   nholuongut.String("vpc-id"),
				Values: nholuongut.StringSlice([]string{vpcID}),
			},
		},
	}
	return d.EC2API.DescribeRouteTables(describeRouteTableInput)
}

func (d *defaultEC2) DeleteSubnet(subnetID string) error {
	deleteSubnetInput := &ec2.DeleteSubnetInput{
		SubnetId: nholuongut.String(subnetID),
	}
	_, err := d.EC2API.DeleteSubnet(deleteSubnetInput)
	return err
}

func (d *defaultEC2) DescribeRouteTables(subnetID string) (*ec2.DescribeRouteTablesOutput, error) {
	describeRouteTableInput := &ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{
				Name:   nholuongut.String("association.subnet-id"),
				Values: nholuongut.StringSlice([]string{subnetID}),
			},
		},
	}
	return d.EC2API.DescribeRouteTables(describeRouteTableInput)
}

func (d *defaultEC2) AssociateRouteTableToSubnet(routeTableId string, subnetID string) error {
	associateRouteTableInput := &ec2.AssociateRouteTableInput{
		RouteTableId: nholuongut.String(routeTableId),
		SubnetId:     nholuongut.String(subnetID),
	}
	_, err := d.EC2API.AssociateRouteTable(associateRouteTableInput)
	return err
}

func (d *defaultEC2) DeleteSecurityGroup(groupID string) error {
	deleteSecurityGroupInput := &ec2.DeleteSecurityGroupInput{
		GroupId: nholuongut.String(groupID),
	}

	_, err := d.EC2API.DeleteSecurityGroup(deleteSecurityGroupInput)
	return err
}

func (d *defaultEC2) CreateSecurityGroup(groupName string, description string, vpcID string) (*ec2.CreateSecurityGroupOutput, error) {
	createSecurityGroupInput := &ec2.CreateSecurityGroupInput{
		Description: nholuongut.String(description),
		GroupName:   nholuongut.String(groupName),
		VpcId:       nholuongut.String(vpcID),
	}

	return d.EC2API.CreateSecurityGroup(createSecurityGroupInput)
}

func (d *defaultEC2) CreateKey(keyName string) (*ec2.CreateKeyPairOutput, error) {
	createKeyInput := &ec2.CreateKeyPairInput{
		KeyName: nholuongut.String(keyName),
	}
	return d.EC2API.CreateKeyPair(createKeyInput)
}

func (d *defaultEC2) DeleteKey(keyName string) error {
	deleteKeyPairInput := &ec2.DeleteKeyPairInput{
		KeyName: nholuongut.String(keyName),
	}
	_, err := d.EC2API.DeleteKeyPair(deleteKeyPairInput)
	return err
}

func (d *defaultEC2) DescribeKey(keyName string) (*ec2.DescribeKeyPairsOutput, error) {
	keyPairInput := &ec2.DescribeKeyPairsInput{
		KeyNames: []*string{
			&keyName,
		},
	}
	return d.EC2API.DescribeKeyPairs(keyPairInput)
}

func (d *defaultEC2) TerminateInstance(instanceIDs []string) error {
	terminateInstanceInput := &ec2.TerminateInstancesInput{
		DryRun:      nil,
		InstanceIds: nholuongut.StringSlice(instanceIDs),
	}
	_, err := d.EC2API.TerminateInstances(terminateInstanceInput)
	return err
}

func (d *defaultEC2) DescribeVPC(vpcID string) (*ec2.DescribeVpcsOutput, error) {
	describeVPCInput := &ec2.DescribeVpcsInput{
		VpcIds: nholuongut.StringSlice([]string{vpcID}),
	}
	return d.EC2API.DescribeVpcs(describeVPCInput)
}

func (d *defaultEC2) CreateTags(resourceIds []string, tags []*ec2.Tag) (*ec2.CreateTagsOutput, error) {
	input := &ec2.CreateTagsInput{
		Resources: nholuongut.StringSlice(resourceIds),
		Tags:      tags,
	}
	return d.EC2API.CreateTags(input)
}

func (d *defaultEC2) DeleteTags(resourceIds []string, tags []*ec2.Tag) (*ec2.DeleteTagsOutput, error) {
	input := &ec2.DeleteTagsInput{
		Resources: nholuongut.StringSlice(resourceIds),
		Tags:      tags,
	}
	return d.EC2API.DeleteTags(input)
}

func NewEC2(session *session.Session) EC2 {
	return &defaultEC2{
		EC2API: ec2.New(session),
	}
}
