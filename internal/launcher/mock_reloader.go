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

package launcher

import (
	config "github.com/elastic/elastic-agent-libs/config"
	mock "github.com/stretchr/testify/mock"
)

// MockReloader is an autogenerated mock type for the Reloader type
type MockReloader struct {
	mock.Mock
}

type MockReloader_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReloader) EXPECT() *MockReloader_Expecter {
	return &MockReloader_Expecter{mock: &_m.Mock}
}

// Channel provides a mock function with no fields
func (_m *MockReloader) Channel() <-chan *config.C {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Channel")
	}

	var r0 <-chan *config.C
	if rf, ok := ret.Get(0).(func() <-chan *config.C); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *config.C)
		}
	}

	return r0
}

// MockReloader_Channel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Channel'
type MockReloader_Channel_Call struct {
	*mock.Call
}

// Channel is a helper method to define mock.On call
func (_e *MockReloader_Expecter) Channel() *MockReloader_Channel_Call {
	return &MockReloader_Channel_Call{Call: _e.mock.On("Channel")}
}

func (_c *MockReloader_Channel_Call) Run(run func()) *MockReloader_Channel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockReloader_Channel_Call) Return(_a0 <-chan *config.C) *MockReloader_Channel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReloader_Channel_Call) RunAndReturn(run func() <-chan *config.C) *MockReloader_Channel_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with no fields
func (_m *MockReloader) Stop() {
	_m.Called()
}

// MockReloader_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockReloader_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockReloader_Expecter) Stop() *MockReloader_Stop_Call {
	return &MockReloader_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockReloader_Stop_Call) Run(run func()) *MockReloader_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockReloader_Stop_Call) Return() *MockReloader_Stop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockReloader_Stop_Call) RunAndReturn(run func()) *MockReloader_Stop_Call {
	_c.Run(run)
	return _c
}

// NewMockReloader creates a new instance of MockReloader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReloader(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReloader {
	mock := &MockReloader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
