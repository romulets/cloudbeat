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

package dataprovider

import (
	beat "github.com/elastic/beats/v7/libbeat/beat"
	fetching "github.com/elastic/cloudbeat/internal/resources/fetching"
	mock "github.com/stretchr/testify/mock"
)

// MockCommonDataProvider is an autogenerated mock type for the CommonDataProvider type
type MockCommonDataProvider struct {
	mock.Mock
}

type MockCommonDataProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCommonDataProvider) EXPECT() *MockCommonDataProvider_Expecter {
	return &MockCommonDataProvider_Expecter{mock: &_m.Mock}
}

// EnrichEvent provides a mock function with given fields: event, resource
func (_m *MockCommonDataProvider) EnrichEvent(event *beat.Event, resource fetching.ResourceMetadata) error {
	ret := _m.Called(event, resource)

	if len(ret) == 0 {
		panic("no return value specified for EnrichEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*beat.Event, fetching.ResourceMetadata) error); ok {
		r0 = rf(event, resource)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommonDataProvider_EnrichEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnrichEvent'
type MockCommonDataProvider_EnrichEvent_Call struct {
	*mock.Call
}

// EnrichEvent is a helper method to define mock.On call
//   - event *beat.Event
//   - resource fetching.ResourceMetadata
func (_e *MockCommonDataProvider_Expecter) EnrichEvent(event interface{}, resource interface{}) *MockCommonDataProvider_EnrichEvent_Call {
	return &MockCommonDataProvider_EnrichEvent_Call{Call: _e.mock.On("EnrichEvent", event, resource)}
}

func (_c *MockCommonDataProvider_EnrichEvent_Call) Run(run func(event *beat.Event, resource fetching.ResourceMetadata)) *MockCommonDataProvider_EnrichEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*beat.Event), args[1].(fetching.ResourceMetadata))
	})
	return _c
}

func (_c *MockCommonDataProvider_EnrichEvent_Call) Return(_a0 error) *MockCommonDataProvider_EnrichEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommonDataProvider_EnrichEvent_Call) RunAndReturn(run func(*beat.Event, fetching.ResourceMetadata) error) *MockCommonDataProvider_EnrichEvent_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCommonDataProvider creates a new instance of MockCommonDataProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCommonDataProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCommonDataProvider {
	mock := &MockCommonDataProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
