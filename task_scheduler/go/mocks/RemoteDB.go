// Code generated by mockery v2.4.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	db "go.skia.org/infra/task_scheduler/go/db"

	time "time"

	types "go.skia.org/infra/task_scheduler/go/types"
)

// RemoteDB is an autogenerated mock type for the RemoteDB type
type RemoteDB struct {
	mock.Mock
}

// DeleteCommitComment provides a mock function with given fields: _a0, _a1
func (_m *RemoteDB) DeleteCommitComment(_a0 context.Context, _a1 *types.CommitComment) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.CommitComment) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTaskComment provides a mock function with given fields: _a0, _a1
func (_m *RemoteDB) DeleteTaskComment(_a0 context.Context, _a1 *types.TaskComment) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.TaskComment) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTaskSpecComment provides a mock function with given fields: _a0, _a1
func (_m *RemoteDB) DeleteTaskSpecComment(_a0 context.Context, _a1 *types.TaskSpecComment) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.TaskSpecComment) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCommentsForRepos provides a mock function with given fields: ctx, repos, from
func (_m *RemoteDB) GetCommentsForRepos(ctx context.Context, repos []string, from time.Time) ([]*types.RepoComments, error) {
	ret := _m.Called(ctx, repos, from)

	var r0 []*types.RepoComments
	if rf, ok := ret.Get(0).(func(context.Context, []string, time.Time) []*types.RepoComments); ok {
		r0 = rf(ctx, repos, from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.RepoComments)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, time.Time) error); ok {
		r1 = rf(ctx, repos, from)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobById provides a mock function with given fields: _a0, _a1
func (_m *RemoteDB) GetJobById(_a0 context.Context, _a1 string) (*types.Job, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Job
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Job); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobsFromDateRange provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *RemoteDB) GetJobsFromDateRange(_a0 context.Context, _a1 time.Time, _a2 time.Time, _a3 string) ([]*types.Job, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []*types.Job
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time, string) []*types.Job); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaskById provides a mock function with given fields: _a0, _a1
func (_m *RemoteDB) GetTaskById(_a0 context.Context, _a1 string) (*types.Task, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Task
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Task); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasksFromDateRange provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *RemoteDB) GetTasksFromDateRange(_a0 context.Context, _a1 time.Time, _a2 time.Time, _a3 string) ([]*types.Task, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []*types.Task
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time, string) []*types.Task); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifiedCommitCommentsCh provides a mock function with given fields: _a0
func (_m *RemoteDB) ModifiedCommitCommentsCh(_a0 context.Context) <-chan []*types.CommitComment {
	ret := _m.Called(_a0)

	var r0 <-chan []*types.CommitComment
	if rf, ok := ret.Get(0).(func(context.Context) <-chan []*types.CommitComment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan []*types.CommitComment)
		}
	}

	return r0
}

// ModifiedJobsCh provides a mock function with given fields: _a0
func (_m *RemoteDB) ModifiedJobsCh(_a0 context.Context) <-chan []*types.Job {
	ret := _m.Called(_a0)

	var r0 <-chan []*types.Job
	if rf, ok := ret.Get(0).(func(context.Context) <-chan []*types.Job); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan []*types.Job)
		}
	}

	return r0
}

// ModifiedTaskCommentsCh provides a mock function with given fields: _a0
func (_m *RemoteDB) ModifiedTaskCommentsCh(_a0 context.Context) <-chan []*types.TaskComment {
	ret := _m.Called(_a0)

	var r0 <-chan []*types.TaskComment
	if rf, ok := ret.Get(0).(func(context.Context) <-chan []*types.TaskComment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan []*types.TaskComment)
		}
	}

	return r0
}

// ModifiedTaskSpecCommentsCh provides a mock function with given fields: _a0
func (_m *RemoteDB) ModifiedTaskSpecCommentsCh(_a0 context.Context) <-chan []*types.TaskSpecComment {
	ret := _m.Called(_a0)

	var r0 <-chan []*types.TaskSpecComment
	if rf, ok := ret.Get(0).(func(context.Context) <-chan []*types.TaskSpecComment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan []*types.TaskSpecComment)
		}
	}

	return r0
}

// ModifiedTasksCh provides a mock function with given fields: _a0
func (_m *RemoteDB) ModifiedTasksCh(_a0 context.Context) <-chan []*types.Task {
	ret := _m.Called(_a0)

	var r0 <-chan []*types.Task
	if rf, ok := ret.Get(0).(func(context.Context) <-chan []*types.Task); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan []*types.Task)
		}
	}

	return r0
}

// PutCommitComment provides a mock function with given fields: _a0, _a1
func (_m *RemoteDB) PutCommitComment(_a0 context.Context, _a1 *types.CommitComment) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.CommitComment) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutTaskComment provides a mock function with given fields: _a0, _a1
func (_m *RemoteDB) PutTaskComment(_a0 context.Context, _a1 *types.TaskComment) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.TaskComment) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutTaskSpecComment provides a mock function with given fields: _a0, _a1
func (_m *RemoteDB) PutTaskSpecComment(_a0 context.Context, _a1 *types.TaskSpecComment) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.TaskSpecComment) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchJobs provides a mock function with given fields: _a0, _a1
func (_m *RemoteDB) SearchJobs(_a0 context.Context, _a1 *db.JobSearchParams) ([]*types.Job, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*types.Job
	if rf, ok := ret.Get(0).(func(context.Context, *db.JobSearchParams) []*types.Job); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Job)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *db.JobSearchParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchTasks provides a mock function with given fields: _a0, _a1
func (_m *RemoteDB) SearchTasks(_a0 context.Context, _a1 *db.TaskSearchParams) ([]*types.Task, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*types.Task
	if rf, ok := ret.Get(0).(func(context.Context, *db.TaskSearchParams) []*types.Task); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *db.TaskSearchParams) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
