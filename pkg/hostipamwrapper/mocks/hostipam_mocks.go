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
// Source: github.com/nholuongut/amazon-vpc-cni-k8s/pkg/hostipamwrapper (interfaces: HostIpam)

// Package mock_hostipamwrapper is a generated GoMock package.
package mock_hostipamwrapper

import (
	reflect "reflect"

	types "github.com/containernetworking/cni/pkg/types"
	types100 "github.com/containernetworking/cni/pkg/types/100"
	gomock "github.com/golang/mock/gomock"
)

// MockHostIpam is a mock of HostIpam interface.
type MockHostIpam struct {
	ctrl     *gomock.Controller
	recorder *MockHostIpamMockRecorder
}

// MockHostIpamMockRecorder is the mock recorder for MockHostIpam.
type MockHostIpamMockRecorder struct {
	mock *MockHostIpam
}

// NewMockHostIpam creates a new mock instance.
func NewMockHostIpam(ctrl *gomock.Controller) *MockHostIpam {
	mock := &MockHostIpam{ctrl: ctrl}
	mock.recorder = &MockHostIpamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHostIpam) EXPECT() *MockHostIpamMockRecorder {
	return m.recorder
}

// ConfigureIface mocks base method.
func (m *MockHostIpam) ConfigureIface(arg0 string, arg1 *types100.Result) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureIface", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigureIface indicates an expected call of ConfigureIface.
func (mr *MockHostIpamMockRecorder) ConfigureIface(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureIface", reflect.TypeOf((*MockHostIpam)(nil).ConfigureIface), arg0, arg1)
}

// ExecAdd mocks base method.
func (m *MockHostIpam) ExecAdd(arg0 string, arg1 []byte) (types.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecAdd", arg0, arg1)
	ret0, _ := ret[0].(types.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecAdd indicates an expected call of ExecAdd.
func (mr *MockHostIpamMockRecorder) ExecAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecAdd", reflect.TypeOf((*MockHostIpam)(nil).ExecAdd), arg0, arg1)
}

// ExecCheck mocks base method.
func (m *MockHostIpam) ExecCheck(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecCheck", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecCheck indicates an expected call of ExecCheck.
func (mr *MockHostIpamMockRecorder) ExecCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecCheck", reflect.TypeOf((*MockHostIpam)(nil).ExecCheck), arg0, arg1)
}

// ExecDel mocks base method.
func (m *MockHostIpam) ExecDel(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecDel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecDel indicates an expected call of ExecDel.
func (mr *MockHostIpamMockRecorder) ExecDel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecDel", reflect.TypeOf((*MockHostIpam)(nil).ExecDel), arg0, arg1)
}
