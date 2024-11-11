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

package elb

import (
	context "context"

	awslib "github.com/elastic/cloudbeat/internal/resources/providers/awslib"

	mock "github.com/stretchr/testify/mock"

	types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
)

// MockLoadBalancerDescriber is an autogenerated mock type for the LoadBalancerDescriber type
type MockLoadBalancerDescriber struct {
	mock.Mock
}

type MockLoadBalancerDescriber_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLoadBalancerDescriber) EXPECT() *MockLoadBalancerDescriber_Expecter {
	return &MockLoadBalancerDescriber_Expecter{mock: &_m.Mock}
}

// DescribeAllLoadBalancers provides a mock function with given fields: _a0
func (_m *MockLoadBalancerDescriber) DescribeAllLoadBalancers(_a0 context.Context) ([]awslib.AwsResource, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DescribeAllLoadBalancers")
	}

	var r0 []awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]awslib.AwsResource, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []awslib.AwsResource); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLoadBalancerDescriber_DescribeAllLoadBalancers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeAllLoadBalancers'
type MockLoadBalancerDescriber_DescribeAllLoadBalancers_Call struct {
	*mock.Call
}

// DescribeAllLoadBalancers is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *MockLoadBalancerDescriber_Expecter) DescribeAllLoadBalancers(_a0 interface{}) *MockLoadBalancerDescriber_DescribeAllLoadBalancers_Call {
	return &MockLoadBalancerDescriber_DescribeAllLoadBalancers_Call{Call: _e.mock.On("DescribeAllLoadBalancers", _a0)}
}

func (_c *MockLoadBalancerDescriber_DescribeAllLoadBalancers_Call) Run(run func(_a0 context.Context)) *MockLoadBalancerDescriber_DescribeAllLoadBalancers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockLoadBalancerDescriber_DescribeAllLoadBalancers_Call) Return(_a0 []awslib.AwsResource, _a1 error) *MockLoadBalancerDescriber_DescribeAllLoadBalancers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLoadBalancerDescriber_DescribeAllLoadBalancers_Call) RunAndReturn(run func(context.Context) ([]awslib.AwsResource, error)) *MockLoadBalancerDescriber_DescribeAllLoadBalancers_Call {
	_c.Call.Return(run)
	return _c
}

// DescribeLoadBalancers provides a mock function with given fields: ctx, balancersNames
func (_m *MockLoadBalancerDescriber) DescribeLoadBalancers(ctx context.Context, balancersNames []string) ([]types.LoadBalancerDescription, error) {
	ret := _m.Called(ctx, balancersNames)

	if len(ret) == 0 {
		panic("no return value specified for DescribeLoadBalancers")
	}

	var r0 []types.LoadBalancerDescription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]types.LoadBalancerDescription, error)); ok {
		return rf(ctx, balancersNames)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []types.LoadBalancerDescription); ok {
		r0 = rf(ctx, balancersNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.LoadBalancerDescription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, balancersNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLoadBalancerDescriber_DescribeLoadBalancers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeLoadBalancers'
type MockLoadBalancerDescriber_DescribeLoadBalancers_Call struct {
	*mock.Call
}

// DescribeLoadBalancers is a helper method to define mock.On call
//   - ctx context.Context
//   - balancersNames []string
func (_e *MockLoadBalancerDescriber_Expecter) DescribeLoadBalancers(ctx interface{}, balancersNames interface{}) *MockLoadBalancerDescriber_DescribeLoadBalancers_Call {
	return &MockLoadBalancerDescriber_DescribeLoadBalancers_Call{Call: _e.mock.On("DescribeLoadBalancers", ctx, balancersNames)}
}

func (_c *MockLoadBalancerDescriber_DescribeLoadBalancers_Call) Run(run func(ctx context.Context, balancersNames []string)) *MockLoadBalancerDescriber_DescribeLoadBalancers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *MockLoadBalancerDescriber_DescribeLoadBalancers_Call) Return(_a0 []types.LoadBalancerDescription, _a1 error) *MockLoadBalancerDescriber_DescribeLoadBalancers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLoadBalancerDescriber_DescribeLoadBalancers_Call) RunAndReturn(run func(context.Context, []string) ([]types.LoadBalancerDescription, error)) *MockLoadBalancerDescriber_DescribeLoadBalancers_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLoadBalancerDescriber creates a new instance of MockLoadBalancerDescriber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLoadBalancerDescriber(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLoadBalancerDescriber {
	mock := &MockLoadBalancerDescriber{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
