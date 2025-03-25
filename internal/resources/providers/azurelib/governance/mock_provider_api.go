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

package governance

import (
	context "context"

	cycle "github.com/elastic/cloudbeat/internal/resources/fetching/cycle"
	mock "github.com/stretchr/testify/mock"
)

// MockProviderAPI is an autogenerated mock type for the ProviderAPI type
type MockProviderAPI struct {
	mock.Mock
}

type MockProviderAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProviderAPI) EXPECT() *MockProviderAPI_Expecter {
	return &MockProviderAPI_Expecter{mock: &_m.Mock}
}

// GetSubscriptions provides a mock function with given fields: ctx, cycleMetadata
func (_m *MockProviderAPI) GetSubscriptions(ctx context.Context, cycleMetadata cycle.Metadata) (map[string]Subscription, error) {
	ret := _m.Called(ctx, cycleMetadata)

	if len(ret) == 0 {
		panic("no return value specified for GetSubscriptions")
	}

	var r0 map[string]Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cycle.Metadata) (map[string]Subscription, error)); ok {
		return rf(ctx, cycleMetadata)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cycle.Metadata) map[string]Subscription); ok {
		r0 = rf(ctx, cycleMetadata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, cycle.Metadata) error); ok {
		r1 = rf(ctx, cycleMetadata)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProviderAPI_GetSubscriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptions'
type MockProviderAPI_GetSubscriptions_Call struct {
	*mock.Call
}

// GetSubscriptions is a helper method to define mock.On call
//   - ctx context.Context
//   - cycleMetadata cycle.Metadata
func (_e *MockProviderAPI_Expecter) GetSubscriptions(ctx interface{}, cycleMetadata interface{}) *MockProviderAPI_GetSubscriptions_Call {
	return &MockProviderAPI_GetSubscriptions_Call{Call: _e.mock.On("GetSubscriptions", ctx, cycleMetadata)}
}

func (_c *MockProviderAPI_GetSubscriptions_Call) Run(run func(ctx context.Context, cycleMetadata cycle.Metadata)) *MockProviderAPI_GetSubscriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cycle.Metadata))
	})
	return _c
}

func (_c *MockProviderAPI_GetSubscriptions_Call) Return(_a0 map[string]Subscription, _a1 error) *MockProviderAPI_GetSubscriptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProviderAPI_GetSubscriptions_Call) RunAndReturn(run func(context.Context, cycle.Metadata) (map[string]Subscription, error)) *MockProviderAPI_GetSubscriptions_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProviderAPI creates a new instance of MockProviderAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProviderAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProviderAPI {
	mock := &MockProviderAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
