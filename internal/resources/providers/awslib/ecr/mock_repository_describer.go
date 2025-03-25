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

package ecr

import (
	context "context"

	types "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	mock "github.com/stretchr/testify/mock"
)

// MockRepositoryDescriber is an autogenerated mock type for the RepositoryDescriber type
type MockRepositoryDescriber struct {
	mock.Mock
}

type MockRepositoryDescriber_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepositoryDescriber) EXPECT() *MockRepositoryDescriber_Expecter {
	return &MockRepositoryDescriber_Expecter{mock: &_m.Mock}
}

// DescribeRepositories provides a mock function with given fields: ctx, repoNames, region
func (_m *MockRepositoryDescriber) DescribeRepositories(ctx context.Context, repoNames []string, region string) ([]types.Repository, error) {
	ret := _m.Called(ctx, repoNames, region)

	if len(ret) == 0 {
		panic("no return value specified for DescribeRepositories")
	}

	var r0 []types.Repository
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, string) ([]types.Repository, error)); ok {
		return rf(ctx, repoNames, region)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, string) []types.Repository); ok {
		r0 = rf(ctx, repoNames, region)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Repository)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, string) error); ok {
		r1 = rf(ctx, repoNames, region)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryDescriber_DescribeRepositories_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeRepositories'
type MockRepositoryDescriber_DescribeRepositories_Call struct {
	*mock.Call
}

// DescribeRepositories is a helper method to define mock.On call
//   - ctx context.Context
//   - repoNames []string
//   - region string
func (_e *MockRepositoryDescriber_Expecter) DescribeRepositories(ctx interface{}, repoNames interface{}, region interface{}) *MockRepositoryDescriber_DescribeRepositories_Call {
	return &MockRepositoryDescriber_DescribeRepositories_Call{Call: _e.mock.On("DescribeRepositories", ctx, repoNames, region)}
}

func (_c *MockRepositoryDescriber_DescribeRepositories_Call) Run(run func(ctx context.Context, repoNames []string, region string)) *MockRepositoryDescriber_DescribeRepositories_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string), args[2].(string))
	})
	return _c
}

func (_c *MockRepositoryDescriber_DescribeRepositories_Call) Return(_a0 []types.Repository, _a1 error) *MockRepositoryDescriber_DescribeRepositories_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryDescriber_DescribeRepositories_Call) RunAndReturn(run func(context.Context, []string, string) ([]types.Repository, error)) *MockRepositoryDescriber_DescribeRepositories_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepositoryDescriber creates a new instance of MockRepositoryDescriber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepositoryDescriber(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepositoryDescriber {
	mock := &MockRepositoryDescriber{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
