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

package k8s

import (
	clog "github.com/elastic/cloudbeat/internal/infra/clog"
	client_gokubernetes "k8s.io/client-go/kubernetes"

	kubernetes "github.com/elastic/elastic-agent-autodiscover/kubernetes"

	mock "github.com/stretchr/testify/mock"
)

// MockClientGetterAPI is an autogenerated mock type for the ClientGetterAPI type
type MockClientGetterAPI struct {
	mock.Mock
}

type MockClientGetterAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientGetterAPI) EXPECT() *MockClientGetterAPI_Expecter {
	return &MockClientGetterAPI_Expecter{mock: &_m.Mock}
}

// GetClient provides a mock function with given fields: log, kubeConfig, options
func (_m *MockClientGetterAPI) GetClient(log *clog.Logger, kubeConfig string, options kubernetes.KubeClientOptions) (client_gokubernetes.Interface, error) {
	ret := _m.Called(log, kubeConfig, options)

	if len(ret) == 0 {
		panic("no return value specified for GetClient")
	}

	var r0 client_gokubernetes.Interface
	var r1 error
	if rf, ok := ret.Get(0).(func(*clog.Logger, string, kubernetes.KubeClientOptions) (client_gokubernetes.Interface, error)); ok {
		return rf(log, kubeConfig, options)
	}
	if rf, ok := ret.Get(0).(func(*clog.Logger, string, kubernetes.KubeClientOptions) client_gokubernetes.Interface); ok {
		r0 = rf(log, kubeConfig, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(client_gokubernetes.Interface)
		}
	}

	if rf, ok := ret.Get(1).(func(*clog.Logger, string, kubernetes.KubeClientOptions) error); ok {
		r1 = rf(log, kubeConfig, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClientGetterAPI_GetClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetClient'
type MockClientGetterAPI_GetClient_Call struct {
	*mock.Call
}

// GetClient is a helper method to define mock.On call
//   - log *clog.Logger
//   - kubeConfig string
//   - options kubernetes.KubeClientOptions
func (_e *MockClientGetterAPI_Expecter) GetClient(log interface{}, kubeConfig interface{}, options interface{}) *MockClientGetterAPI_GetClient_Call {
	return &MockClientGetterAPI_GetClient_Call{Call: _e.mock.On("GetClient", log, kubeConfig, options)}
}

func (_c *MockClientGetterAPI_GetClient_Call) Run(run func(log *clog.Logger, kubeConfig string, options kubernetes.KubeClientOptions)) *MockClientGetterAPI_GetClient_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*clog.Logger), args[1].(string), args[2].(kubernetes.KubeClientOptions))
	})
	return _c
}

func (_c *MockClientGetterAPI_GetClient_Call) Return(_a0 client_gokubernetes.Interface, _a1 error) *MockClientGetterAPI_GetClient_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClientGetterAPI_GetClient_Call) RunAndReturn(run func(*clog.Logger, string, kubernetes.KubeClientOptions) (client_gokubernetes.Interface, error)) *MockClientGetterAPI_GetClient_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClientGetterAPI creates a new instance of MockClientGetterAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientGetterAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientGetterAPI {
	mock := &MockClientGetterAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
