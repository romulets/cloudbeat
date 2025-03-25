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

	armsecurity "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"

	mock "github.com/stretchr/testify/mock"
)

// mockSecurityClientWrapper is an autogenerated mock type for the securityClientWrapper type
type mockSecurityClientWrapper struct {
	mock.Mock
}

type mockSecurityClientWrapper_Expecter struct {
	mock *mock.Mock
}

func (_m *mockSecurityClientWrapper) EXPECT() *mockSecurityClientWrapper_Expecter {
	return &mockSecurityClientWrapper_Expecter{mock: &_m.Mock}
}

// ListAutoProvisioningSettings provides a mock function with given fields: ctx, subID
func (_m *mockSecurityClientWrapper) ListAutoProvisioningSettings(ctx context.Context, subID string) ([]armsecurity.AutoProvisioningSettingsClientListResponse, error) {
	ret := _m.Called(ctx, subID)

	if len(ret) == 0 {
		panic("no return value specified for ListAutoProvisioningSettings")
	}

	var r0 []armsecurity.AutoProvisioningSettingsClientListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]armsecurity.AutoProvisioningSettingsClientListResponse, error)); ok {
		return rf(ctx, subID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []armsecurity.AutoProvisioningSettingsClientListResponse); ok {
		r0 = rf(ctx, subID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]armsecurity.AutoProvisioningSettingsClientListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, subID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockSecurityClientWrapper_ListAutoProvisioningSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAutoProvisioningSettings'
type mockSecurityClientWrapper_ListAutoProvisioningSettings_Call struct {
	*mock.Call
}

// ListAutoProvisioningSettings is a helper method to define mock.On call
//   - ctx context.Context
//   - subID string
func (_e *mockSecurityClientWrapper_Expecter) ListAutoProvisioningSettings(ctx interface{}, subID interface{}) *mockSecurityClientWrapper_ListAutoProvisioningSettings_Call {
	return &mockSecurityClientWrapper_ListAutoProvisioningSettings_Call{Call: _e.mock.On("ListAutoProvisioningSettings", ctx, subID)}
}

func (_c *mockSecurityClientWrapper_ListAutoProvisioningSettings_Call) Run(run func(ctx context.Context, subID string)) *mockSecurityClientWrapper_ListAutoProvisioningSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *mockSecurityClientWrapper_ListAutoProvisioningSettings_Call) Return(_a0 []armsecurity.AutoProvisioningSettingsClientListResponse, _a1 error) *mockSecurityClientWrapper_ListAutoProvisioningSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockSecurityClientWrapper_ListAutoProvisioningSettings_Call) RunAndReturn(run func(context.Context, string) ([]armsecurity.AutoProvisioningSettingsClientListResponse, error)) *mockSecurityClientWrapper_ListAutoProvisioningSettings_Call {
	_c.Call.Return(run)
	return _c
}

// ListSecurityContacts provides a mock function with given fields: ctx, subID
func (_m *mockSecurityClientWrapper) ListSecurityContacts(ctx context.Context, subID string) ([]armsecurity.ContactsClientListResponse, error) {
	ret := _m.Called(ctx, subID)

	if len(ret) == 0 {
		panic("no return value specified for ListSecurityContacts")
	}

	var r0 []armsecurity.ContactsClientListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]armsecurity.ContactsClientListResponse, error)); ok {
		return rf(ctx, subID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []armsecurity.ContactsClientListResponse); ok {
		r0 = rf(ctx, subID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]armsecurity.ContactsClientListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, subID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockSecurityClientWrapper_ListSecurityContacts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSecurityContacts'
type mockSecurityClientWrapper_ListSecurityContacts_Call struct {
	*mock.Call
}

// ListSecurityContacts is a helper method to define mock.On call
//   - ctx context.Context
//   - subID string
func (_e *mockSecurityClientWrapper_Expecter) ListSecurityContacts(ctx interface{}, subID interface{}) *mockSecurityClientWrapper_ListSecurityContacts_Call {
	return &mockSecurityClientWrapper_ListSecurityContacts_Call{Call: _e.mock.On("ListSecurityContacts", ctx, subID)}
}

func (_c *mockSecurityClientWrapper_ListSecurityContacts_Call) Run(run func(ctx context.Context, subID string)) *mockSecurityClientWrapper_ListSecurityContacts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *mockSecurityClientWrapper_ListSecurityContacts_Call) Return(_a0 []armsecurity.ContactsClientListResponse, _a1 error) *mockSecurityClientWrapper_ListSecurityContacts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockSecurityClientWrapper_ListSecurityContacts_Call) RunAndReturn(run func(context.Context, string) ([]armsecurity.ContactsClientListResponse, error)) *mockSecurityClientWrapper_ListSecurityContacts_Call {
	_c.Call.Return(run)
	return _c
}

// newMockSecurityClientWrapper creates a new instance of mockSecurityClientWrapper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockSecurityClientWrapper(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockSecurityClientWrapper {
	mock := &mockSecurityClientWrapper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
