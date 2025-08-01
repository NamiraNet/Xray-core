// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NamiraNet/xray-core/common/log (interfaces: Handler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	log "github.com/NamiraNet/xray-core/common/log"
)

// LogHandler is a mock of Handler interface.
type LogHandler struct {
	ctrl     *gomock.Controller
	recorder *LogHandlerMockRecorder
}

// LogHandlerMockRecorder is the mock recorder for LogHandler.
type LogHandlerMockRecorder struct {
	mock *LogHandler
}

// NewLogHandler creates a new mock instance.
func NewLogHandler(ctrl *gomock.Controller) *LogHandler {
	mock := &LogHandler{ctrl: ctrl}
	mock.recorder = &LogHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *LogHandler) EXPECT() *LogHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *LogHandler) Handle(arg0 log.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Handle", arg0)
}

// Handle indicates an expected call of Handle.
func (mr *LogHandlerMockRecorder) Handle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*LogHandler)(nil).Handle), arg0)
}
