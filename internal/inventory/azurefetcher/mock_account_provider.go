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

package azurefetcher

import (
	context "context"

	inventory "github.com/elastic/cloudbeat/internal/resources/providers/azurelib/inventory"
	mock "github.com/stretchr/testify/mock"
)

// mockAccountProvider is an autogenerated mock type for the accountProvider type
type mockAccountProvider struct {
	mock.Mock
}

type mockAccountProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *mockAccountProvider) EXPECT() *mockAccountProvider_Expecter {
	return &mockAccountProvider_Expecter{mock: &_m.Mock}
}

// ListSubscriptions provides a mock function with given fields: ctx
func (_m *mockAccountProvider) ListSubscriptions(ctx context.Context) ([]inventory.AzureAsset, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListSubscriptions")
	}

	var r0 []inventory.AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]inventory.AzureAsset, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []inventory.AzureAsset); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]inventory.AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAccountProvider_ListSubscriptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSubscriptions'
type mockAccountProvider_ListSubscriptions_Call struct {
	*mock.Call
}

// ListSubscriptions is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockAccountProvider_Expecter) ListSubscriptions(ctx interface{}) *mockAccountProvider_ListSubscriptions_Call {
	return &mockAccountProvider_ListSubscriptions_Call{Call: _e.mock.On("ListSubscriptions", ctx)}
}

func (_c *mockAccountProvider_ListSubscriptions_Call) Run(run func(ctx context.Context)) *mockAccountProvider_ListSubscriptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockAccountProvider_ListSubscriptions_Call) Return(_a0 []inventory.AzureAsset, _a1 error) *mockAccountProvider_ListSubscriptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAccountProvider_ListSubscriptions_Call) RunAndReturn(run func(context.Context) ([]inventory.AzureAsset, error)) *mockAccountProvider_ListSubscriptions_Call {
	_c.Call.Return(run)
	return _c
}

// ListTenants provides a mock function with given fields: ctx
func (_m *mockAccountProvider) ListTenants(ctx context.Context) ([]inventory.AzureAsset, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListTenants")
	}

	var r0 []inventory.AzureAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]inventory.AzureAsset, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []inventory.AzureAsset); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]inventory.AzureAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAccountProvider_ListTenants_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTenants'
type mockAccountProvider_ListTenants_Call struct {
	*mock.Call
}

// ListTenants is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockAccountProvider_Expecter) ListTenants(ctx interface{}) *mockAccountProvider_ListTenants_Call {
	return &mockAccountProvider_ListTenants_Call{Call: _e.mock.On("ListTenants", ctx)}
}

func (_c *mockAccountProvider_ListTenants_Call) Run(run func(ctx context.Context)) *mockAccountProvider_ListTenants_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockAccountProvider_ListTenants_Call) Return(_a0 []inventory.AzureAsset, _a1 error) *mockAccountProvider_ListTenants_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAccountProvider_ListTenants_Call) RunAndReturn(run func(context.Context) ([]inventory.AzureAsset, error)) *mockAccountProvider_ListTenants_Call {
	_c.Call.Return(run)
	return _c
}

// newMockAccountProvider creates a new instance of mockAccountProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAccountProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockAccountProvider {
	mock := &mockAccountProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
