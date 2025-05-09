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

package add_cluster_id

import mock "github.com/stretchr/testify/mock"

// MockClusterHelper is an autogenerated mock type for the ClusterHelper type
type MockClusterHelper struct {
	mock.Mock
}

type MockClusterHelper_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClusterHelper) EXPECT() *MockClusterHelper_Expecter {
	return &MockClusterHelper_Expecter{mock: &_m.Mock}
}

// ClusterId provides a mock function with no fields
func (_m *MockClusterHelper) ClusterId() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ClusterId")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockClusterHelper_ClusterId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClusterId'
type MockClusterHelper_ClusterId_Call struct {
	*mock.Call
}

// ClusterId is a helper method to define mock.On call
func (_e *MockClusterHelper_Expecter) ClusterId() *MockClusterHelper_ClusterId_Call {
	return &MockClusterHelper_ClusterId_Call{Call: _e.mock.On("ClusterId")}
}

func (_c *MockClusterHelper_ClusterId_Call) Run(run func()) *MockClusterHelper_ClusterId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClusterHelper_ClusterId_Call) Return(_a0 string) *MockClusterHelper_ClusterId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClusterHelper_ClusterId_Call) RunAndReturn(run func() string) *MockClusterHelper_ClusterId_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClusterHelper creates a new instance of MockClusterHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClusterHelper(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClusterHelper {
	mock := &MockClusterHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
