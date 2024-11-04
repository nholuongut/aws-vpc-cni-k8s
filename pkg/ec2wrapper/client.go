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

package ec2wrapper

import (
	"github.com/nholuongut/nholuongut-sdk-go/nholuongut"
	"github.com/nholuongut/nholuongut-sdk-go/nholuongut/request"
	"github.com/nholuongut/nholuongut-sdk-go/nholuongut/session"
	ec2svc "github.com/nholuongut/nholuongut-sdk-go/service/ec2"
)

// EC2 is the EC2 wrapper interface
type EC2 interface {
	CreateNetworkInterfaceWithContext(ctx nholuongut.Context, input *ec2svc.CreateNetworkInterfaceInput, opts ...request.Option) (*ec2svc.CreateNetworkInterfaceOutput, error)
	DescribeInstancesWithContext(ctx nholuongut.Context, input *ec2svc.DescribeInstancesInput, opts ...request.Option) (*ec2svc.DescribeInstancesOutput, error)
	DescribeInstanceTypesWithContext(ctx nholuongut.Context, input *ec2svc.DescribeInstanceTypesInput, opts ...request.Option) (*ec2svc.DescribeInstanceTypesOutput, error)
	AttachNetworkInterfaceWithContext(ctx nholuongut.Context, input *ec2svc.AttachNetworkInterfaceInput, opts ...request.Option) (*ec2svc.AttachNetworkInterfaceOutput, error)
	DeleteNetworkInterfaceWithContext(ctx nholuongut.Context, input *ec2svc.DeleteNetworkInterfaceInput, opts ...request.Option) (*ec2svc.DeleteNetworkInterfaceOutput, error)
	DetachNetworkInterfaceWithContext(ctx nholuongut.Context, input *ec2svc.DetachNetworkInterfaceInput, opts ...request.Option) (*ec2svc.DetachNetworkInterfaceOutput, error)
	AssignPrivateIpAddressesWithContext(ctx nholuongut.Context, input *ec2svc.AssignPrivateIpAddressesInput, opts ...request.Option) (*ec2svc.AssignPrivateIpAddressesOutput, error)
	UnassignPrivateIpAddressesWithContext(ctx nholuongut.Context, input *ec2svc.UnassignPrivateIpAddressesInput, opts ...request.Option) (*ec2svc.UnassignPrivateIpAddressesOutput, error)
	AssignIpv6AddressesWithContext(ctx nholuongut.Context, input *ec2svc.AssignIpv6AddressesInput, opts ...request.Option) (*ec2svc.AssignIpv6AddressesOutput, error)
	UnassignIpv6AddressesWithContext(ctx nholuongut.Context, input *ec2svc.UnassignIpv6AddressesInput, opts ...request.Option) (*ec2svc.UnassignIpv6AddressesOutput, error)
	DescribeNetworkInterfacesWithContext(ctx nholuongut.Context, input *ec2svc.DescribeNetworkInterfacesInput, opts ...request.Option) (*ec2svc.DescribeNetworkInterfacesOutput, error)
	ModifyNetworkInterfaceAttributeWithContext(ctx nholuongut.Context, input *ec2svc.ModifyNetworkInterfaceAttributeInput, opts ...request.Option) (*ec2svc.ModifyNetworkInterfaceAttributeOutput, error)
	CreateTagsWithContext(ctx nholuongut.Context, input *ec2svc.CreateTagsInput, opts ...request.Option) (*ec2svc.CreateTagsOutput, error)
	DescribeNetworkInterfacesPagesWithContext(ctx nholuongut.Context, input *ec2svc.DescribeNetworkInterfacesInput, fn func(*ec2svc.DescribeNetworkInterfacesOutput, bool) bool, opts ...request.Option) error
	DescribeSubnetsWithContext(ctx nholuongut.Context, input *ec2svc.DescribeSubnetsInput, opts ...request.Option) (*ec2svc.DescribeSubnetsOutput, error)
}

// New creates a new EC2 wrapper
func New(sess *session.Session) EC2 {
	return ec2svc.New(sess)
}
