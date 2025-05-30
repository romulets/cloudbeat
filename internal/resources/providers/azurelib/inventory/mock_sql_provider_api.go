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

// MockSQLProviderAPI is an autogenerated mock type for the SQLProviderAPI type
type MockSQLProviderAPI struct {
	mock.Mock
}

type MockSQLProviderAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSQLProviderAPI) EXPECT() *MockSQLProviderAPI_Expecter {
	return &MockSQLProviderAPI_Expecter{mock: &_m.Mock}
}

// GetSQLBlobAuditingPolicies provides a mock function with given fields: ctx, subID, resourceGroup, serverName
func (_m *MockSQLProviderAPI) GetSQLBlobAuditingPolicies(ctx context.Context, subID string, resourceGroup string, serverName string) ([]AzureAsset, error) {
	ret := _m.Called(ctx, subID, resourceGroup, serverName)

	if len(ret) == 0 {
		panic("no return value specified for GetSQLBlobAuditingPolicies")
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

// MockSQLProviderAPI_GetSQLBlobAuditingPolicies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSQLBlobAuditingPolicies'
type MockSQLProviderAPI_GetSQLBlobAuditingPolicies_Call struct {
	*mock.Call
}

// GetSQLBlobAuditingPolicies is a helper method to define mock.On call
//   - ctx context.Context
//   - subID string
//   - resourceGroup string
//   - serverName string
func (_e *MockSQLProviderAPI_Expecter) GetSQLBlobAuditingPolicies(ctx interface{}, subID interface{}, resourceGroup interface{}, serverName interface{}) *MockSQLProviderAPI_GetSQLBlobAuditingPolicies_Call {
	return &MockSQLProviderAPI_GetSQLBlobAuditingPolicies_Call{Call: _e.mock.On("GetSQLBlobAuditingPolicies", ctx, subID, resourceGroup, serverName)}
}

func (_c *MockSQLProviderAPI_GetSQLBlobAuditingPolicies_Call) Run(run func(ctx context.Context, subID string, resourceGroup string, serverName string)) *MockSQLProviderAPI_GetSQLBlobAuditingPolicies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockSQLProviderAPI_GetSQLBlobAuditingPolicies_Call) Return(_a0 []AzureAsset, _a1 error) *MockSQLProviderAPI_GetSQLBlobAuditingPolicies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSQLProviderAPI_GetSQLBlobAuditingPolicies_Call) RunAndReturn(run func(context.Context, string, string, string) ([]AzureAsset, error)) *MockSQLProviderAPI_GetSQLBlobAuditingPolicies_Call {
	_c.Call.Return(run)
	return _c
}

// ListSQLAdvancedThreatProtectionSettings provides a mock function with given fields: ctx, subID, resourceGroup, serverName
func (_m *MockSQLProviderAPI) ListSQLAdvancedThreatProtectionSettings(ctx context.Context, subID string, resourceGroup string, serverName string) ([]AzureAsset, error) {
	ret := _m.Called(ctx, subID, resourceGroup, serverName)

	if len(ret) == 0 {
		panic("no return value specified for ListSQLAdvancedThreatProtectionSettings")
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

// MockSQLProviderAPI_ListSQLAdvancedThreatProtectionSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSQLAdvancedThreatProtectionSettings'
type MockSQLProviderAPI_ListSQLAdvancedThreatProtectionSettings_Call struct {
	*mock.Call
}

// ListSQLAdvancedThreatProtectionSettings is a helper method to define mock.On call
//   - ctx context.Context
//   - subID string
//   - resourceGroup string
//   - serverName string
func (_e *MockSQLProviderAPI_Expecter) ListSQLAdvancedThreatProtectionSettings(ctx interface{}, subID interface{}, resourceGroup interface{}, serverName interface{}) *MockSQLProviderAPI_ListSQLAdvancedThreatProtectionSettings_Call {
	return &MockSQLProviderAPI_ListSQLAdvancedThreatProtectionSettings_Call{Call: _e.mock.On("ListSQLAdvancedThreatProtectionSettings", ctx, subID, resourceGroup, serverName)}
}

func (_c *MockSQLProviderAPI_ListSQLAdvancedThreatProtectionSettings_Call) Run(run func(ctx context.Context, subID string, resourceGroup string, serverName string)) *MockSQLProviderAPI_ListSQLAdvancedThreatProtectionSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockSQLProviderAPI_ListSQLAdvancedThreatProtectionSettings_Call) Return(_a0 []AzureAsset, _a1 error) *MockSQLProviderAPI_ListSQLAdvancedThreatProtectionSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSQLProviderAPI_ListSQLAdvancedThreatProtectionSettings_Call) RunAndReturn(run func(context.Context, string, string, string) ([]AzureAsset, error)) *MockSQLProviderAPI_ListSQLAdvancedThreatProtectionSettings_Call {
	_c.Call.Return(run)
	return _c
}

// ListSQLEncryptionProtector provides a mock function with given fields: ctx, subID, resourceGroup, serverName
func (_m *MockSQLProviderAPI) ListSQLEncryptionProtector(ctx context.Context, subID string, resourceGroup string, serverName string) ([]AzureAsset, error) {
	ret := _m.Called(ctx, subID, resourceGroup, serverName)

	if len(ret) == 0 {
		panic("no return value specified for ListSQLEncryptionProtector")
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

// MockSQLProviderAPI_ListSQLEncryptionProtector_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSQLEncryptionProtector'
type MockSQLProviderAPI_ListSQLEncryptionProtector_Call struct {
	*mock.Call
}

// ListSQLEncryptionProtector is a helper method to define mock.On call
//   - ctx context.Context
//   - subID string
//   - resourceGroup string
//   - serverName string
func (_e *MockSQLProviderAPI_Expecter) ListSQLEncryptionProtector(ctx interface{}, subID interface{}, resourceGroup interface{}, serverName interface{}) *MockSQLProviderAPI_ListSQLEncryptionProtector_Call {
	return &MockSQLProviderAPI_ListSQLEncryptionProtector_Call{Call: _e.mock.On("ListSQLEncryptionProtector", ctx, subID, resourceGroup, serverName)}
}

func (_c *MockSQLProviderAPI_ListSQLEncryptionProtector_Call) Run(run func(ctx context.Context, subID string, resourceGroup string, serverName string)) *MockSQLProviderAPI_ListSQLEncryptionProtector_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockSQLProviderAPI_ListSQLEncryptionProtector_Call) Return(_a0 []AzureAsset, _a1 error) *MockSQLProviderAPI_ListSQLEncryptionProtector_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSQLProviderAPI_ListSQLEncryptionProtector_Call) RunAndReturn(run func(context.Context, string, string, string) ([]AzureAsset, error)) *MockSQLProviderAPI_ListSQLEncryptionProtector_Call {
	_c.Call.Return(run)
	return _c
}

// ListSQLFirewallRules provides a mock function with given fields: ctx, subID, resourceGroup, serverName
func (_m *MockSQLProviderAPI) ListSQLFirewallRules(ctx context.Context, subID string, resourceGroup string, serverName string) ([]AzureAsset, error) {
	ret := _m.Called(ctx, subID, resourceGroup, serverName)

	if len(ret) == 0 {
		panic("no return value specified for ListSQLFirewallRules")
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

// MockSQLProviderAPI_ListSQLFirewallRules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSQLFirewallRules'
type MockSQLProviderAPI_ListSQLFirewallRules_Call struct {
	*mock.Call
}

// ListSQLFirewallRules is a helper method to define mock.On call
//   - ctx context.Context
//   - subID string
//   - resourceGroup string
//   - serverName string
func (_e *MockSQLProviderAPI_Expecter) ListSQLFirewallRules(ctx interface{}, subID interface{}, resourceGroup interface{}, serverName interface{}) *MockSQLProviderAPI_ListSQLFirewallRules_Call {
	return &MockSQLProviderAPI_ListSQLFirewallRules_Call{Call: _e.mock.On("ListSQLFirewallRules", ctx, subID, resourceGroup, serverName)}
}

func (_c *MockSQLProviderAPI_ListSQLFirewallRules_Call) Run(run func(ctx context.Context, subID string, resourceGroup string, serverName string)) *MockSQLProviderAPI_ListSQLFirewallRules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockSQLProviderAPI_ListSQLFirewallRules_Call) Return(_a0 []AzureAsset, _a1 error) *MockSQLProviderAPI_ListSQLFirewallRules_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSQLProviderAPI_ListSQLFirewallRules_Call) RunAndReturn(run func(context.Context, string, string, string) ([]AzureAsset, error)) *MockSQLProviderAPI_ListSQLFirewallRules_Call {
	_c.Call.Return(run)
	return _c
}

// ListSQLTransparentDataEncryptions provides a mock function with given fields: ctx, subID, resourceGroup, serverName
func (_m *MockSQLProviderAPI) ListSQLTransparentDataEncryptions(ctx context.Context, subID string, resourceGroup string, serverName string) ([]AzureAsset, error) {
	ret := _m.Called(ctx, subID, resourceGroup, serverName)

	if len(ret) == 0 {
		panic("no return value specified for ListSQLTransparentDataEncryptions")
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

// MockSQLProviderAPI_ListSQLTransparentDataEncryptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSQLTransparentDataEncryptions'
type MockSQLProviderAPI_ListSQLTransparentDataEncryptions_Call struct {
	*mock.Call
}

// ListSQLTransparentDataEncryptions is a helper method to define mock.On call
//   - ctx context.Context
//   - subID string
//   - resourceGroup string
//   - serverName string
func (_e *MockSQLProviderAPI_Expecter) ListSQLTransparentDataEncryptions(ctx interface{}, subID interface{}, resourceGroup interface{}, serverName interface{}) *MockSQLProviderAPI_ListSQLTransparentDataEncryptions_Call {
	return &MockSQLProviderAPI_ListSQLTransparentDataEncryptions_Call{Call: _e.mock.On("ListSQLTransparentDataEncryptions", ctx, subID, resourceGroup, serverName)}
}

func (_c *MockSQLProviderAPI_ListSQLTransparentDataEncryptions_Call) Run(run func(ctx context.Context, subID string, resourceGroup string, serverName string)) *MockSQLProviderAPI_ListSQLTransparentDataEncryptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockSQLProviderAPI_ListSQLTransparentDataEncryptions_Call) Return(_a0 []AzureAsset, _a1 error) *MockSQLProviderAPI_ListSQLTransparentDataEncryptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSQLProviderAPI_ListSQLTransparentDataEncryptions_Call) RunAndReturn(run func(context.Context, string, string, string) ([]AzureAsset, error)) *MockSQLProviderAPI_ListSQLTransparentDataEncryptions_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSQLProviderAPI creates a new instance of MockSQLProviderAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSQLProviderAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSQLProviderAPI {
	mock := &MockSQLProviderAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
