// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package backend

import (
	context "context"

	core "github.com/cschleiden/go-workflows/internal/core"
	history "github.com/cschleiden/go-workflows/internal/history"

	log "github.com/cschleiden/go-workflows/log"

	mock "github.com/stretchr/testify/mock"

	task "github.com/cschleiden/go-workflows/internal/task"
)

// MockBackend is an autogenerated mock type for the Backend type
type MockBackend struct {
	mock.Mock
}

// CancelWorkflowInstance provides a mock function with given fields: ctx, instance, event
func (_m *MockBackend) CancelWorkflowInstance(ctx context.Context, instance *core.WorkflowInstance, event *history.Event) error {
	ret := _m.Called(ctx, instance, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowInstance, *history.Event) error); ok {
		r0 = rf(ctx, instance, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteActivityTask provides a mock function with given fields: ctx, instance, activityID, event
func (_m *MockBackend) CompleteActivityTask(ctx context.Context, instance *core.WorkflowInstance, activityID string, event history.Event) error {
	ret := _m.Called(ctx, instance, activityID, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowInstance, string, history.Event) error); ok {
		r0 = rf(ctx, instance, activityID, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteWorkflowTask provides a mock function with given fields: ctx, taskID, instance, state, executedEvents, activityEvents, workflowEvents
func (_m *MockBackend) CompleteWorkflowTask(ctx context.Context, taskID string, instance *core.WorkflowInstance, state WorkflowState, executedEvents []history.Event, activityEvents []history.Event, workflowEvents []history.WorkflowEvent) error {
	ret := _m.Called(ctx, taskID, instance, state, executedEvents, activityEvents, workflowEvents)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *core.WorkflowInstance, WorkflowState, []history.Event, []history.Event, []history.WorkflowEvent) error); ok {
		r0 = rf(ctx, taskID, instance, state, executedEvents, activityEvents, workflowEvents)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateWorkflowInstance provides a mock function with given fields: ctx, event
func (_m *MockBackend) CreateWorkflowInstance(ctx context.Context, event history.WorkflowEvent) error {
	ret := _m.Called(ctx, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, history.WorkflowEvent) error); ok {
		r0 = rf(ctx, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtendActivityTask provides a mock function with given fields: ctx, activityID
func (_m *MockBackend) ExtendActivityTask(ctx context.Context, activityID string) error {
	ret := _m.Called(ctx, activityID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, activityID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtendWorkflowTask provides a mock function with given fields: ctx, taskID, instance
func (_m *MockBackend) ExtendWorkflowTask(ctx context.Context, taskID string, instance *core.WorkflowInstance) error {
	ret := _m.Called(ctx, taskID, instance)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *core.WorkflowInstance) error); ok {
		r0 = rf(ctx, taskID, instance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetActivityTask provides a mock function with given fields: ctx
func (_m *MockBackend) GetActivityTask(ctx context.Context) (*task.Activity, error) {
	ret := _m.Called(ctx)

	var r0 *task.Activity
	if rf, ok := ret.Get(0).(func(context.Context) *task.Activity); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*task.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowInstanceHistory provides a mock function with given fields: ctx, instance
func (_m *MockBackend) GetWorkflowInstanceHistory(ctx context.Context, instance *core.WorkflowInstance, lastSequenceID *int64) ([]history.Event, error) {
	ret := _m.Called(ctx, instance)

	var r0 []history.Event
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowInstance, *int64) []history.Event); ok {
		r0 = rf(ctx, instance, lastSequenceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]history.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.WorkflowInstance, *int64) error); ok {
		r1 = rf(ctx, instance, lastSequenceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowInstanceState provides a mock function with given fields: ctx, instance
func (_m *MockBackend) GetWorkflowInstanceState(ctx context.Context, instance *core.WorkflowInstance) (WorkflowState, error) {
	ret := _m.Called(ctx, instance)

	var r0 WorkflowState
	if rf, ok := ret.Get(0).(func(context.Context, *core.WorkflowInstance) WorkflowState); ok {
		r0 = rf(ctx, instance)
	} else {
		r0 = ret.Get(0).(WorkflowState)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.WorkflowInstance) error); ok {
		r1 = rf(ctx, instance)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowTask provides a mock function with given fields: ctx
func (_m *MockBackend) GetWorkflowTask(ctx context.Context) (*task.Workflow, error) {
	ret := _m.Called(ctx)

	var r0 *task.Workflow
	if rf, ok := ret.Get(0).(func(context.Context) *task.Workflow); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*task.Workflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logger provides a mock function with given fields:
func (_m *MockBackend) Logger() log.Logger {
	ret := _m.Called()

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func() log.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// SignalWorkflow provides a mock function with given fields: ctx, instanceID, event
func (_m *MockBackend) SignalWorkflow(ctx context.Context, instanceID string, event history.Event) error {
	ret := _m.Called(ctx, instanceID, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, history.Event) error); ok {
		r0 = rf(ctx, instanceID, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
