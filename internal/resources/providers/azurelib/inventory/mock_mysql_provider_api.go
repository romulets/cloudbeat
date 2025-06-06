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

// Code generated by mockery v2.53.3. DO NOT EDIT.

package inventory

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockMysqlProviderAPI is an autogenerated mock type for the MysqlProviderAPI type
type MockMysqlProviderAPI struct {
	mock.Mock
}

type MockMysqlProviderAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMysqlProviderAPI) EXPECT() *MockMysqlProviderAPI_Expecter {
	return &MockMysqlProviderAPI_Expecter{mock: &_m.Mock}
}

// GetFlexibleTLSVersionConfiguration provides a mock function with given fields: ctx, subID, resourceGroup, serverName
func (_m *MockMysqlProviderAPI) GetFlexibleTLSVersionConfiguration(ctx context.Context, subID string, resourceGroup string, serverName string) ([]AzureAsset, error) {
	ret := _m.Called(ctx, subID, resourceGroup, serverName)

	if len(ret) == 0 {
		panic("no return value specified for GetFlexibleTLSVersionConfiguration")
	}

	var r0 []AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) ([]AzureAsset, error)); ok {
		return rf(ctx, subID, resourceGroup, serverName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) []AzureAsset); ok {
		r0 = rf(ctx, subID, resourceGroup, serverName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, subID, resourceGroup, serverName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMysqlProviderAPI_GetFlexibleTLSVersionConfiguration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFlexibleTLSVersionConfiguration'
type MockMysqlProviderAPI_GetFlexibleTLSVersionConfiguration_Call struct {
	*mock.Call
}

// GetFlexibleTLSVersionConfiguration is a helper method to define mock.On call
//   - ctx context.Context
//   - subID string
//   - resourceGroup string
//   - serverName string
func (_e *MockMysqlProviderAPI_Expecter) GetFlexibleTLSVersionConfiguration(ctx interface{}, subID interface{}, resourceGroup interface{}, serverName interface{}) *MockMysqlProviderAPI_GetFlexibleTLSVersionConfiguration_Call {
	return &MockMysqlProviderAPI_GetFlexibleTLSVersionConfiguration_Call{Call: _e.mock.On("GetFlexibleTLSVersionConfiguration", ctx, subID, resourceGroup, serverName)}
}

func (_c *MockMysqlProviderAPI_GetFlexibleTLSVersionConfiguration_Call) Run(run func(ctx context.Context, subID string, resourceGroup string, serverName string)) *MockMysqlProviderAPI_GetFlexibleTLSVersionConfiguration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockMysqlProviderAPI_GetFlexibleTLSVersionConfiguration_Call) Return(_a0 []AzureAsset, _a1 error) *MockMysqlProviderAPI_GetFlexibleTLSVersionConfiguration_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMysqlProviderAPI_GetFlexibleTLSVersionConfiguration_Call) RunAndReturn(run func(context.Context, string, string, string) ([]AzureAsset, error)) *MockMysqlProviderAPI_GetFlexibleTLSVersionConfiguration_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMysqlProviderAPI creates a new instance of MockMysqlProviderAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMysqlProviderAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMysqlProviderAPI {
	mock := &MockMysqlProviderAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
