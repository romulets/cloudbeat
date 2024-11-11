// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mockery v2.42.0. DO NOT EDIT.

package awslib

import (
	context "context"

	aws "github.com/aws/aws-sdk-go-v2/aws"

	mock "github.com/stretchr/testify/mock"
)

// MockEKSClusterNameProviderAPI is an autogenerated mock type for the EKSClusterNameProviderAPI type
type MockEKSClusterNameProviderAPI struct {
	mock.Mock
}

type MockEKSClusterNameProviderAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEKSClusterNameProviderAPI) EXPECT() *MockEKSClusterNameProviderAPI_Expecter {
	return &MockEKSClusterNameProviderAPI_Expecter{mock: &_m.Mock}
}

// GetClusterName provides a mock function with given fields: ctx, cfg, instanceId
func (_m *MockEKSClusterNameProviderAPI) GetClusterName(ctx context.Context, cfg aws.Config, instanceId string) (string, error) {
	ret := _m.Called(ctx, cfg, instanceId)

	if len(ret) == 0 {
		panic("no return value specified for GetClusterName")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, aws.Config, string) (string, error)); ok {
		return rf(ctx, cfg, instanceId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, aws.Config, string) string); ok {
		r0 = rf(ctx, cfg, instanceId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, aws.Config, string) error); ok {
		r1 = rf(ctx, cfg, instanceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEKSClusterNameProviderAPI_GetClusterName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClusterName'
type MockEKSClusterNameProviderAPI_GetClusterName_Call struct {
	*mock.Call
}

// GetClusterName is a helper method to define mock.On call
//   - ctx context.Context
//   - cfg aws.Config
//   - instanceId string
func (_e *MockEKSClusterNameProviderAPI_Expecter) GetClusterName(ctx interface{}, cfg interface{}, instanceId interface{}) *MockEKSClusterNameProviderAPI_GetClusterName_Call {
	return &MockEKSClusterNameProviderAPI_GetClusterName_Call{Call: _e.mock.On("GetClusterName", ctx, cfg, instanceId)}
}

func (_c *MockEKSClusterNameProviderAPI_GetClusterName_Call) Run(run func(ctx context.Context, cfg aws.Config, instanceId string)) *MockEKSClusterNameProviderAPI_GetClusterName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(aws.Config), args[2].(string))
	})
	return _c
}

func (_c *MockEKSClusterNameProviderAPI_GetClusterName_Call) Return(_a0 string, _a1 error) *MockEKSClusterNameProviderAPI_GetClusterName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockEKSClusterNameProviderAPI_GetClusterName_Call) RunAndReturn(run func(context.Context, aws.Config, string) (string, error)) *MockEKSClusterNameProviderAPI_GetClusterName_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEKSClusterNameProviderAPI creates a new instance of MockEKSClusterNameProviderAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEKSClusterNameProviderAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEKSClusterNameProviderAPI {
	mock := &MockEKSClusterNameProviderAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
