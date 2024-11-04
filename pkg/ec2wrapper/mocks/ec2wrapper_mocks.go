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
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nholuongut/amazon-vpc-cni-k8s/pkg/ec2wrapper (interfaces: EC2)

// Package mock_ec2wrapper is a generated GoMock package.
package mock_ec2wrapper

import (
	context "context"
	reflect "reflect"

	request "github.com/nholuongut/nholuongut-sdk-go/nholuongut/request"
	ec2 "github.com/nholuongut/nholuongut-sdk-go/service/ec2"
	gomock "github.com/golang/mock/gomock"
)

// MockEC2 is a mock of EC2 interface.
type MockEC2 struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MockRecorder
}

// MockEC2MockRecorder is the mock recorder for MockEC2.
type MockEC2MockRecorder struct {
	mock *MockEC2
}

// NewMockEC2 creates a new mock instance.
func NewMockEC2(ctrl *gomock.Controller) *MockEC2 {
	mock := &MockEC2{ctrl: ctrl}
	mock.recorder = &MockEC2MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEC2) EXPECT() *MockEC2MockRecorder {
	return m.recorder
}

// AssignIpv6AddressesWithContext mocks base method.
func (m *MockEC2) AssignIpv6AddressesWithContext(arg0 context.Context, arg1 *ec2.AssignIpv6AddressesInput, arg2 ...request.Option) (*ec2.AssignIpv6AddressesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssignIpv6AddressesWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.AssignIpv6AddressesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignIpv6AddressesWithContext indicates an expected call of AssignIpv6AddressesWithContext.
func (mr *MockEC2MockRecorder) AssignIpv6AddressesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignIpv6AddressesWithContext", reflect.TypeOf((*MockEC2)(nil).AssignIpv6AddressesWithContext), varargs...)
}

// AssignPrivateIpAddressesWithContext mocks base method.
func (m *MockEC2) AssignPrivateIpAddressesWithContext(arg0 context.Context, arg1 *ec2.AssignPrivateIpAddressesInput, arg2 ...request.Option) (*ec2.AssignPrivateIpAddressesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssignPrivateIpAddressesWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.AssignPrivateIpAddressesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignPrivateIpAddressesWithContext indicates an expected call of AssignPrivateIpAddressesWithContext.
func (mr *MockEC2MockRecorder) AssignPrivateIpAddressesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignPrivateIpAddressesWithContext", reflect.TypeOf((*MockEC2)(nil).AssignPrivateIpAddressesWithContext), varargs...)
}

// AttachNetworkInterfaceWithContext mocks base method.
func (m *MockEC2) AttachNetworkInterfaceWithContext(arg0 context.Context, arg1 *ec2.AttachNetworkInterfaceInput, arg2 ...request.Option) (*ec2.AttachNetworkInterfaceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AttachNetworkInterfaceWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.AttachNetworkInterfaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachNetworkInterfaceWithContext indicates an expected call of AttachNetworkInterfaceWithContext.
func (mr *MockEC2MockRecorder) AttachNetworkInterfaceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachNetworkInterfaceWithContext", reflect.TypeOf((*MockEC2)(nil).AttachNetworkInterfaceWithContext), varargs...)
}

// CreateNetworkInterfaceWithContext mocks base method.
func (m *MockEC2) CreateNetworkInterfaceWithContext(arg0 context.Context, arg1 *ec2.CreateNetworkInterfaceInput, arg2 ...request.Option) (*ec2.CreateNetworkInterfaceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNetworkInterfaceWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.CreateNetworkInterfaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNetworkInterfaceWithContext indicates an expected call of CreateNetworkInterfaceWithContext.
func (mr *MockEC2MockRecorder) CreateNetworkInterfaceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetworkInterfaceWithContext", reflect.TypeOf((*MockEC2)(nil).CreateNetworkInterfaceWithContext), varargs...)
}

// CreateTagsWithContext mocks base method.
func (m *MockEC2) CreateTagsWithContext(arg0 context.Context, arg1 *ec2.CreateTagsInput, arg2 ...request.Option) (*ec2.CreateTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTagsWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.CreateTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTagsWithContext indicates an expected call of CreateTagsWithContext.
func (mr *MockEC2MockRecorder) CreateTagsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTagsWithContext", reflect.TypeOf((*MockEC2)(nil).CreateTagsWithContext), varargs...)
}

// DeleteNetworkInterfaceWithContext mocks base method.
func (m *MockEC2) DeleteNetworkInterfaceWithContext(arg0 context.Context, arg1 *ec2.DeleteNetworkInterfaceInput, arg2 ...request.Option) (*ec2.DeleteNetworkInterfaceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteNetworkInterfaceWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DeleteNetworkInterfaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNetworkInterfaceWithContext indicates an expected call of DeleteNetworkInterfaceWithContext.
func (mr *MockEC2MockRecorder) DeleteNetworkInterfaceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetworkInterfaceWithContext", reflect.TypeOf((*MockEC2)(nil).DeleteNetworkInterfaceWithContext), varargs...)
}

// DescribeInstanceTypesWithContext mocks base method.
func (m *MockEC2) DescribeInstanceTypesWithContext(arg0 context.Context, arg1 *ec2.DescribeInstanceTypesInput, arg2 ...request.Option) (*ec2.DescribeInstanceTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstanceTypesWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstanceTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstanceTypesWithContext indicates an expected call of DescribeInstanceTypesWithContext.
func (mr *MockEC2MockRecorder) DescribeInstanceTypesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstanceTypesWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeInstanceTypesWithContext), varargs...)
}

// DescribeInstancesWithContext mocks base method.
func (m *MockEC2) DescribeInstancesWithContext(arg0 context.Context, arg1 *ec2.DescribeInstancesInput, arg2 ...request.Option) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstancesWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstancesWithContext indicates an expected call of DescribeInstancesWithContext.
func (mr *MockEC2MockRecorder) DescribeInstancesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstancesWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeInstancesWithContext), varargs...)
}

// DescribeNetworkInterfacesPagesWithContext mocks base method.
func (m *MockEC2) DescribeNetworkInterfacesPagesWithContext(arg0 context.Context, arg1 *ec2.DescribeNetworkInterfacesInput, arg2 func(*ec2.DescribeNetworkInterfacesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNetworkInterfacesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DescribeNetworkInterfacesPagesWithContext indicates an expected call of DescribeNetworkInterfacesPagesWithContext.
func (mr *MockEC2MockRecorder) DescribeNetworkInterfacesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNetworkInterfacesPagesWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeNetworkInterfacesPagesWithContext), varargs...)
}

// DescribeNetworkInterfacesWithContext mocks base method.
func (m *MockEC2) DescribeNetworkInterfacesWithContext(arg0 context.Context, arg1 *ec2.DescribeNetworkInterfacesInput, arg2 ...request.Option) (*ec2.DescribeNetworkInterfacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNetworkInterfacesWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeNetworkInterfacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeNetworkInterfacesWithContext indicates an expected call of DescribeNetworkInterfacesWithContext.
func (mr *MockEC2MockRecorder) DescribeNetworkInterfacesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNetworkInterfacesWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeNetworkInterfacesWithContext), varargs...)
}

// DescribeSubnetsWithContext mocks base method.
func (m *MockEC2) DescribeSubnetsWithContext(arg0 context.Context, arg1 *ec2.DescribeSubnetsInput, arg2 ...request.Option) (*ec2.DescribeSubnetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSubnetsWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeSubnetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSubnetsWithContext indicates an expected call of DescribeSubnetsWithContext.
func (mr *MockEC2MockRecorder) DescribeSubnetsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSubnetsWithContext", reflect.TypeOf((*MockEC2)(nil).DescribeSubnetsWithContext), varargs...)
}

// DetachNetworkInterfaceWithContext mocks base method.
func (m *MockEC2) DetachNetworkInterfaceWithContext(arg0 context.Context, arg1 *ec2.DetachNetworkInterfaceInput, arg2 ...request.Option) (*ec2.DetachNetworkInterfaceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DetachNetworkInterfaceWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.DetachNetworkInterfaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachNetworkInterfaceWithContext indicates an expected call of DetachNetworkInterfaceWithContext.
func (mr *MockEC2MockRecorder) DetachNetworkInterfaceWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachNetworkInterfaceWithContext", reflect.TypeOf((*MockEC2)(nil).DetachNetworkInterfaceWithContext), varargs...)
}

// ModifyNetworkInterfaceAttributeWithContext mocks base method.
func (m *MockEC2) ModifyNetworkInterfaceAttributeWithContext(arg0 context.Context, arg1 *ec2.ModifyNetworkInterfaceAttributeInput, arg2 ...request.Option) (*ec2.ModifyNetworkInterfaceAttributeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ModifyNetworkInterfaceAttributeWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.ModifyNetworkInterfaceAttributeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyNetworkInterfaceAttributeWithContext indicates an expected call of ModifyNetworkInterfaceAttributeWithContext.
func (mr *MockEC2MockRecorder) ModifyNetworkInterfaceAttributeWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyNetworkInterfaceAttributeWithContext", reflect.TypeOf((*MockEC2)(nil).ModifyNetworkInterfaceAttributeWithContext), varargs...)
}

// UnassignIpv6AddressesWithContext mocks base method.
func (m *MockEC2) UnassignIpv6AddressesWithContext(arg0 context.Context, arg1 *ec2.UnassignIpv6AddressesInput, arg2 ...request.Option) (*ec2.UnassignIpv6AddressesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnassignIpv6AddressesWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.UnassignIpv6AddressesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnassignIpv6AddressesWithContext indicates an expected call of UnassignIpv6AddressesWithContext.
func (mr *MockEC2MockRecorder) UnassignIpv6AddressesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnassignIpv6AddressesWithContext", reflect.TypeOf((*MockEC2)(nil).UnassignIpv6AddressesWithContext), varargs...)
}

// UnassignPrivateIpAddressesWithContext mocks base method.
func (m *MockEC2) UnassignPrivateIpAddressesWithContext(arg0 context.Context, arg1 *ec2.UnassignPrivateIpAddressesInput, arg2 ...request.Option) (*ec2.UnassignPrivateIpAddressesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnassignPrivateIpAddressesWithContext", varargs...)
	ret0, _ := ret[0].(*ec2.UnassignPrivateIpAddressesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnassignPrivateIpAddressesWithContext indicates an expected call of UnassignPrivateIpAddressesWithContext.
func (mr *MockEC2MockRecorder) UnassignPrivateIpAddressesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnassignPrivateIpAddressesWithContext", reflect.TypeOf((*MockEC2)(nil).UnassignPrivateIpAddressesWithContext), varargs...)
}