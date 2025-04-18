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

package ec2

import (
	context "context"

	awslib "github.com/elastic/cloudbeat/internal/resources/providers/awslib"

	mock "github.com/stretchr/testify/mock"

	types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// MockElasticCompute is an autogenerated mock type for the ElasticCompute type
type MockElasticCompute struct {
	mock.Mock
}

type MockElasticCompute_Expecter struct {
	mock *mock.Mock
}

func (_m *MockElasticCompute) EXPECT() *MockElasticCompute_Expecter {
	return &MockElasticCompute_Expecter{mock: &_m.Mock}
}

// DescribeNetworkAcl provides a mock function with given fields: ctx
func (_m *MockElasticCompute) DescribeNetworkAcl(ctx context.Context) ([]awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DescribeNetworkAcl")
	}

	var r0 []awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]awslib.AwsResource, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockElasticCompute_DescribeNetworkAcl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeNetworkAcl'
type MockElasticCompute_DescribeNetworkAcl_Call struct {
	*mock.Call
}

// DescribeNetworkAcl is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockElasticCompute_Expecter) DescribeNetworkAcl(ctx interface{}) *MockElasticCompute_DescribeNetworkAcl_Call {
	return &MockElasticCompute_DescribeNetworkAcl_Call{Call: _e.mock.On("DescribeNetworkAcl", ctx)}
}

func (_c *MockElasticCompute_DescribeNetworkAcl_Call) Run(run func(ctx context.Context)) *MockElasticCompute_DescribeNetworkAcl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockElasticCompute_DescribeNetworkAcl_Call) Return(_a0 []awslib.AwsResource, _a1 error) *MockElasticCompute_DescribeNetworkAcl_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockElasticCompute_DescribeNetworkAcl_Call) RunAndReturn(run func(context.Context) ([]awslib.AwsResource, error)) *MockElasticCompute_DescribeNetworkAcl_Call {
	_c.Call.Return(run)
	return _c
}

// DescribeSecurityGroups provides a mock function with given fields: ctx
func (_m *MockElasticCompute) DescribeSecurityGroups(ctx context.Context) ([]awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DescribeSecurityGroups")
	}

	var r0 []awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]awslib.AwsResource, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockElasticCompute_DescribeSecurityGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeSecurityGroups'
type MockElasticCompute_DescribeSecurityGroups_Call struct {
	*mock.Call
}

// DescribeSecurityGroups is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockElasticCompute_Expecter) DescribeSecurityGroups(ctx interface{}) *MockElasticCompute_DescribeSecurityGroups_Call {
	return &MockElasticCompute_DescribeSecurityGroups_Call{Call: _e.mock.On("DescribeSecurityGroups", ctx)}
}

func (_c *MockElasticCompute_DescribeSecurityGroups_Call) Run(run func(ctx context.Context)) *MockElasticCompute_DescribeSecurityGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockElasticCompute_DescribeSecurityGroups_Call) Return(_a0 []awslib.AwsResource, _a1 error) *MockElasticCompute_DescribeSecurityGroups_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockElasticCompute_DescribeSecurityGroups_Call) RunAndReturn(run func(context.Context) ([]awslib.AwsResource, error)) *MockElasticCompute_DescribeSecurityGroups_Call {
	_c.Call.Return(run)
	return _c
}

// DescribeVpcs provides a mock function with given fields: ctx
func (_m *MockElasticCompute) DescribeVpcs(ctx context.Context) ([]awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DescribeVpcs")
	}

	var r0 []awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]awslib.AwsResource, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockElasticCompute_DescribeVpcs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeVpcs'
type MockElasticCompute_DescribeVpcs_Call struct {
	*mock.Call
}

// DescribeVpcs is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockElasticCompute_Expecter) DescribeVpcs(ctx interface{}) *MockElasticCompute_DescribeVpcs_Call {
	return &MockElasticCompute_DescribeVpcs_Call{Call: _e.mock.On("DescribeVpcs", ctx)}
}

func (_c *MockElasticCompute_DescribeVpcs_Call) Run(run func(ctx context.Context)) *MockElasticCompute_DescribeVpcs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockElasticCompute_DescribeVpcs_Call) Return(_a0 []awslib.AwsResource, _a1 error) *MockElasticCompute_DescribeVpcs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockElasticCompute_DescribeVpcs_Call) RunAndReturn(run func(context.Context) ([]awslib.AwsResource, error)) *MockElasticCompute_DescribeVpcs_Call {
	_c.Call.Return(run)
	return _c
}

// GetEbsEncryptionByDefault provides a mock function with given fields: ctx
func (_m *MockElasticCompute) GetEbsEncryptionByDefault(ctx context.Context) ([]awslib.AwsResource, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetEbsEncryptionByDefault")
	}

	var r0 []awslib.AwsResource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]awslib.AwsResource, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []awslib.AwsResource); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]awslib.AwsResource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockElasticCompute_GetEbsEncryptionByDefault_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEbsEncryptionByDefault'
type MockElasticCompute_GetEbsEncryptionByDefault_Call struct {
	*mock.Call
}

// GetEbsEncryptionByDefault is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockElasticCompute_Expecter) GetEbsEncryptionByDefault(ctx interface{}) *MockElasticCompute_GetEbsEncryptionByDefault_Call {
	return &MockElasticCompute_GetEbsEncryptionByDefault_Call{Call: _e.mock.On("GetEbsEncryptionByDefault", ctx)}
}

func (_c *MockElasticCompute_GetEbsEncryptionByDefault_Call) Run(run func(ctx context.Context)) *MockElasticCompute_GetEbsEncryptionByDefault_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockElasticCompute_GetEbsEncryptionByDefault_Call) Return(_a0 []awslib.AwsResource, _a1 error) *MockElasticCompute_GetEbsEncryptionByDefault_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockElasticCompute_GetEbsEncryptionByDefault_Call) RunAndReturn(run func(context.Context) ([]awslib.AwsResource, error)) *MockElasticCompute_GetEbsEncryptionByDefault_Call {
	_c.Call.Return(run)
	return _c
}

// GetRouteTableForSubnet provides a mock function with given fields: ctx, region, subnetId, vpcId
func (_m *MockElasticCompute) GetRouteTableForSubnet(ctx context.Context, region string, subnetId string, vpcId string) (types.RouteTable, error) {
	ret := _m.Called(ctx, region, subnetId, vpcId)

	if len(ret) == 0 {
		panic("no return value specified for GetRouteTableForSubnet")
	}

	var r0 types.RouteTable
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (types.RouteTable, error)); ok {
		return rf(ctx, region, subnetId, vpcId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) types.RouteTable); ok {
		r0 = rf(ctx, region, subnetId, vpcId)
	} else {
		r0 = ret.Get(0).(types.RouteTable)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, region, subnetId, vpcId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockElasticCompute_GetRouteTableForSubnet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRouteTableForSubnet'
type MockElasticCompute_GetRouteTableForSubnet_Call struct {
	*mock.Call
}

// GetRouteTableForSubnet is a helper method to define mock.On call
//   - ctx context.Context
//   - region string
//   - subnetId string
//   - vpcId string
func (_e *MockElasticCompute_Expecter) GetRouteTableForSubnet(ctx interface{}, region interface{}, subnetId interface{}, vpcId interface{}) *MockElasticCompute_GetRouteTableForSubnet_Call {
	return &MockElasticCompute_GetRouteTableForSubnet_Call{Call: _e.mock.On("GetRouteTableForSubnet", ctx, region, subnetId, vpcId)}
}

func (_c *MockElasticCompute_GetRouteTableForSubnet_Call) Run(run func(ctx context.Context, region string, subnetId string, vpcId string)) *MockElasticCompute_GetRouteTableForSubnet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockElasticCompute_GetRouteTableForSubnet_Call) Return(_a0 types.RouteTable, _a1 error) *MockElasticCompute_GetRouteTableForSubnet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockElasticCompute_GetRouteTableForSubnet_Call) RunAndReturn(run func(context.Context, string, string, string) (types.RouteTable, error)) *MockElasticCompute_GetRouteTableForSubnet_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockElasticCompute creates a new instance of MockElasticCompute. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockElasticCompute(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockElasticCompute {
	mock := &MockElasticCompute{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
