// Code generated by MockGen. DO NOT EDIT.
// Source: services/reverse/reverse.go

// Package reverse is a generated GoMock package.
package reverse

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReverseClient is a mock of Interface interface.
type MockReverseClient struct {
	ctrl     *gomock.Controller
	recorder *MockReverseClientMockRecorder
}

// MockReverseClientMockRecorder is the mock recorder for MockReverseClient.
type MockReverseClientMockRecorder struct {
	mock *MockReverseClient
}

// NewMockReverseClient creates a new mock instance.
func NewMockReverseClient(ctrl *gomock.Controller) *MockReverseClient {
	mock := &MockReverseClient{ctrl: ctrl}
	mock.recorder = &MockReverseClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReverseClient) EXPECT() *MockReverseClientMockRecorder {
	return m.recorder
}

// GetComponents mocks base method.
func (m *MockReverseClient) GetComponents(lat, lon float64, options CallOptions) ([]Component, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComponents", lat, lon, options)
	ret0, _ := ret[0].([]Component)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComponents indicates an expected call of GetComponents.
func (mr *MockReverseClientMockRecorder) GetComponents(lat, lon, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComponents", reflect.TypeOf((*MockReverseClient)(nil).GetComponents), lat, lon, options)
}

// GetComponentsWithContext mocks base method.
func (m *MockReverseClient) GetComponentsWithContext(ctx context.Context, lat, lon float64, options CallOptions) ([]Component, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComponentsWithContext", ctx, lat, lon, options)
	ret0, _ := ret[0].([]Component)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComponentsWithContext indicates an expected call of GetComponentsWithContext.
func (mr *MockReverseClientMockRecorder) GetComponentsWithContext(ctx, lat, lon, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComponentsWithContext", reflect.TypeOf((*MockReverseClient)(nil).GetComponentsWithContext), ctx, lat, lon, options)
}

// GetDisplayName mocks base method.
func (m *MockReverseClient) GetDisplayName(lat, lon float64, options CallOptions) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDisplayName", lat, lon, options)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDisplayName indicates an expected call of GetDisplayName.
func (mr *MockReverseClientMockRecorder) GetDisplayName(lat, lon, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDisplayName", reflect.TypeOf((*MockReverseClient)(nil).GetDisplayName), lat, lon, options)
}

// GetDisplayNameWithContext mocks base method.
func (m *MockReverseClient) GetDisplayNameWithContext(ctx context.Context, lat, lon float64, options CallOptions) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDisplayNameWithContext", ctx, lat, lon, options)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDisplayNameWithContext indicates an expected call of GetDisplayNameWithContext.
func (mr *MockReverseClientMockRecorder) GetDisplayNameWithContext(ctx, lat, lon, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDisplayNameWithContext", reflect.TypeOf((*MockReverseClient)(nil).GetDisplayNameWithContext), ctx, lat, lon, options)
}
