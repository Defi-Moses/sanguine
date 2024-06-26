// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	watcher "github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"

	prometheus "github.com/prometheus/client_golang/prometheus"
	mock "github.com/stretchr/testify/mock"
)

// BlockHeightWatcher is an autogenerated mock type for the BlockHeightWatcher type
type BlockHeightWatcher struct {
	mock.Mock
}

// ContextForCondition provides a mock function with given fields: parentCtx, check
func (_m *BlockHeightWatcher) ContextForCondition(parentCtx context.Context, check watcher.ConditionCheck) (context.Context, context.CancelFunc) {
	ret := _m.Called(parentCtx, check)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, watcher.ConditionCheck) context.Context); ok {
		r0 = rf(parentCtx, check)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 context.CancelFunc
	if rf, ok := ret.Get(1).(func(context.Context, watcher.ConditionCheck) context.CancelFunc); ok {
		r1 = rf(parentCtx, check)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(context.CancelFunc)
		}
	}

	return r0, r1
}

// GetMetrics provides a mock function with given fields: labels
func (_m *BlockHeightWatcher) GetMetrics(labels map[string]string) []prometheus.Collector {
	ret := _m.Called(labels)

	var r0 []prometheus.Collector
	if rf, ok := ret.Get(0).(func(map[string]string) []prometheus.Collector); ok {
		r0 = rf(labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]prometheus.Collector)
		}
	}

	return r0
}

// Subscribe provides a mock function with given fields:
func (_m *BlockHeightWatcher) Subscribe() <-chan uint64 {
	ret := _m.Called()

	var r0 <-chan uint64
	if rf, ok := ret.Get(0).(func() <-chan uint64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan uint64)
		}
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: ch
func (_m *BlockHeightWatcher) Unsubscribe(ch <-chan uint64) {
	_m.Called(ch)
}
