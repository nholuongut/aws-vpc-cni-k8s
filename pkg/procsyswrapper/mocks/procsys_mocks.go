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
// Source: github.com/nholuongut/amazon-vpc-cni-k8s/pkg/procsyswrapper (interfaces: ProcSys)

// Package mock_procsyswrapper is a generated GoMock package.
package mock_procsyswrapper

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProcSys is a mock of ProcSys interface.
type MockProcSys struct {
	ctrl     *gomock.Controller
	recorder *MockProcSysMockRecorder
}

// MockProcSysMockRecorder is the mock recorder for MockProcSys.
type MockProcSysMockRecorder struct {
	mock *MockProcSys
}

// NewMockProcSys creates a new mock instance.
func NewMockProcSys(ctrl *gomock.Controller) *MockProcSys {
	mock := &MockProcSys{ctrl: ctrl}
	mock.recorder = &MockProcSysMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcSys) EXPECT() *MockProcSysMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockProcSys) Get(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProcSysMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProcSys)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockProcSys) Set(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockProcSysMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockProcSys)(nil).Set), arg0, arg1)
}
