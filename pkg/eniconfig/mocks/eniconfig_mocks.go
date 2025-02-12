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
// Source: github.com/nholuongut/amazon-vpc-cni-k8s/pkg/eniconfig (interfaces: ENIConfig)

// Package mock_eniconfig is a generated GoMock package.
package mock_eniconfig

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/nholuongut/amazon-vpc-cni-k8s/pkg/apis/crd/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockENIConfig is a mock of ENIConfig interface.
type MockENIConfig struct {
	ctrl     *gomock.Controller
	recorder *MockENIConfigMockRecorder
}

// MockENIConfigMockRecorder is the mock recorder for MockENIConfig.
type MockENIConfigMockRecorder struct {
	mock *MockENIConfig
}

// NewMockENIConfig creates a new mock instance.
func NewMockENIConfig(ctrl *gomock.Controller) *MockENIConfig {
	mock := &MockENIConfig{ctrl: ctrl}
	mock.recorder = &MockENIConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockENIConfig) EXPECT() *MockENIConfigMockRecorder {
	return m.recorder
}

// GetENIConfigName mocks base method.
func (m *MockENIConfig) GetENIConfigName(arg0 context.Context, arg1 client.Client) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetENIConfigName", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetENIConfigName indicates an expected call of GetENIConfigName.
func (mr *MockENIConfigMockRecorder) GetENIConfigName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetENIConfigName", reflect.TypeOf((*MockENIConfig)(nil).GetENIConfigName), arg0, arg1)
}

// MyENIConfig mocks base method.
func (m *MockENIConfig) MyENIConfig(arg0 client.Client) (*v1alpha1.ENIConfigSpec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MyENIConfig", arg0)
	ret0, _ := ret[0].(*v1alpha1.ENIConfigSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MyENIConfig indicates an expected call of MyENIConfig.
func (mr *MockENIConfigMockRecorder) MyENIConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MyENIConfig", reflect.TypeOf((*MockENIConfig)(nil).MyENIConfig), arg0)
}
