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

package awslib

import (
	context "context"

	aws "github.com/aws/aws-sdk-go-v2/aws"

	imds "github.com/aws/aws-sdk-go-v2/feature/ec2/imds"

	mock "github.com/stretchr/testify/mock"
)

// MockMetadataProvider is an autogenerated mock type for the MetadataProvider type
type MockMetadataProvider struct {
	mock.Mock
}

type MockMetadataProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMetadataProvider) EXPECT() *MockMetadataProvider_Expecter {
	return &MockMetadataProvider_Expecter{mock: &_m.Mock}
}

// GetMetadata provides a mock function with given fields: ctx, cfg
func (_m *MockMetadataProvider) GetMetadata(ctx context.Context, cfg aws.Config) (*imds.InstanceIdentityDocument, error) {
	ret := _m.Called(ctx, cfg)

	if len(ret) == 0 {
		panic("no return value specified for GetMetadata")
	}

	var r0 *imds.InstanceIdentityDocument
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, aws.Config) (*imds.InstanceIdentityDocument, error)); ok {
		return rf(ctx, cfg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, aws.Config) *imds.InstanceIdentityDocument); ok {
		r0 = rf(ctx, cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*imds.InstanceIdentityDocument)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, aws.Config) error); ok {
		r1 = rf(ctx, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetadataProvider_GetMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetadata'
type MockMetadataProvider_GetMetadata_Call struct {
	*mock.Call
}

// GetMetadata is a helper method to define mock.On call
//   - ctx context.Context
//   - cfg aws.Config
func (_e *MockMetadataProvider_Expecter) GetMetadata(ctx interface{}, cfg interface{}) *MockMetadataProvider_GetMetadata_Call {
	return &MockMetadataProvider_GetMetadata_Call{Call: _e.mock.On("GetMetadata", ctx, cfg)}
}

func (_c *MockMetadataProvider_GetMetadata_Call) Run(run func(ctx context.Context, cfg aws.Config)) *MockMetadataProvider_GetMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(aws.Config))
	})
	return _c
}

func (_c *MockMetadataProvider_GetMetadata_Call) Return(_a0 *imds.InstanceIdentityDocument, _a1 error) *MockMetadataProvider_GetMetadata_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetadataProvider_GetMetadata_Call) RunAndReturn(run func(context.Context, aws.Config) (*imds.InstanceIdentityDocument, error)) *MockMetadataProvider_GetMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMetadataProvider creates a new instance of MockMetadataProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMetadataProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMetadataProvider {
	mock := &MockMetadataProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
