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

package builder

import (
	context "context"

	beat "github.com/elastic/beats/v7/libbeat/beat"

	evaluator "github.com/elastic/cloudbeat/internal/evaluator"

	mock "github.com/stretchr/testify/mock"
)

// MockTransformer is an autogenerated mock type for the Transformer type
type MockTransformer struct {
	mock.Mock
}

type MockTransformer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTransformer) EXPECT() *MockTransformer_Expecter {
	return &MockTransformer_Expecter{mock: &_m.Mock}
}

// CreateBeatEvents provides a mock function with given fields: ctx, data
func (_m *MockTransformer) CreateBeatEvents(ctx context.Context, data evaluator.EventData) ([]beat.Event, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for CreateBeatEvents")
	}

	var r0 []beat.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, evaluator.EventData) ([]beat.Event, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, evaluator.EventData) []beat.Event); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]beat.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, evaluator.EventData) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTransformer_CreateBeatEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBeatEvents'
type MockTransformer_CreateBeatEvents_Call struct {
	*mock.Call
}

// CreateBeatEvents is a helper method to define mock.On call
//   - ctx context.Context
//   - data evaluator.EventData
func (_e *MockTransformer_Expecter) CreateBeatEvents(ctx interface{}, data interface{}) *MockTransformer_CreateBeatEvents_Call {
	return &MockTransformer_CreateBeatEvents_Call{Call: _e.mock.On("CreateBeatEvents", ctx, data)}
}

func (_c *MockTransformer_CreateBeatEvents_Call) Run(run func(ctx context.Context, data evaluator.EventData)) *MockTransformer_CreateBeatEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(evaluator.EventData))
	})
	return _c
}

func (_c *MockTransformer_CreateBeatEvents_Call) Return(_a0 []beat.Event, _a1 error) *MockTransformer_CreateBeatEvents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTransformer_CreateBeatEvents_Call) RunAndReturn(run func(context.Context, evaluator.EventData) ([]beat.Event, error)) *MockTransformer_CreateBeatEvents_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTransformer creates a new instance of MockTransformer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTransformer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTransformer {
	mock := &MockTransformer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
