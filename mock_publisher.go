// Code generated by mockery v2.42.1. DO NOT EDIT.

package ctrl

import (
	context "context"

	condition "github.com/metal-automata/rivets/condition"

	mock "github.com/stretchr/testify/mock"
)

// MockPublisher is an autogenerated mock type for the Publisher type
type MockPublisher struct {
	mock.Mock
}

type MockPublisher_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPublisher) EXPECT() *MockPublisher_Expecter {
	return &MockPublisher_Expecter{mock: &_m.Mock}
}

// Publish provides a mock function with given fields: ctx, task, tsUpdateOnly
func (_m *MockPublisher) Publish(ctx context.Context, task *condition.Task[interface{}, interface{}], tsUpdateOnly bool) error {
	ret := _m.Called(ctx, task, tsUpdateOnly)

	if len(ret) == 0 {
		panic("no return value specified for Publish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *condition.Task[interface{}, interface{}], bool) error); ok {
		r0 = rf(ctx, task, tsUpdateOnly)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPublisher_Publish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Publish'
type MockPublisher_Publish_Call struct {
	*mock.Call
}

// Publish is a helper method to define mock.On call
//   - ctx context.Context
//   - task *condition.Task[interface{},interface{}]
//   - tsUpdateOnly bool
func (_e *MockPublisher_Expecter) Publish(ctx interface{}, task interface{}, tsUpdateOnly interface{}) *MockPublisher_Publish_Call {
	return &MockPublisher_Publish_Call{Call: _e.mock.On("Publish", ctx, task, tsUpdateOnly)}
}

func (_c *MockPublisher_Publish_Call) Run(run func(ctx context.Context, task *condition.Task[interface{}, interface{}], tsUpdateOnly bool)) *MockPublisher_Publish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*condition.Task[interface{}, interface{}]), args[2].(bool))
	})
	return _c
}

func (_c *MockPublisher_Publish_Call) Return(_a0 error) *MockPublisher_Publish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPublisher_Publish_Call) RunAndReturn(run func(context.Context, *condition.Task[interface{}, interface{}], bool) error) *MockPublisher_Publish_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPublisher creates a new instance of MockPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPublisher {
	mock := &MockPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
