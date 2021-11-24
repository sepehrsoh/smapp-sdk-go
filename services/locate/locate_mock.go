// Code generated by MockGen. DO NOT EDIT.
// Source: services/locate/locate.go

// Package locate is a generated GoMock package.
package locate

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLocateClient is a mock of Interface interface.
type MockLocateClient struct {
	ctrl     *gomock.Controller
	recorder *MockLocateClientMockRecorder
}

// MockLocateClientMockRecorder is the mock recorder for MockLocateClient.
type MockLocateClientMockRecorder struct {
	mock *MockLocateClient
}

// NewMockLocateClient creates a new mock instance.
func NewMockLocateClient(ctrl *gomock.Controller) *MockLocateClient {
	mock := &MockLocateClient{ctrl: ctrl}
	mock.recorder = &MockLocateClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocateClient) EXPECT() *MockLocateClientMockRecorder {
	return m.recorder
}

// LocatePoints mocks base method.
func (m *MockLocateClient) LocatePoints(points []Point, options CallOptions) ([]Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocatePoints", points, options)
	ret0, _ := ret[0].([]Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocatePoints indicates an expected call of LocatePoints.
func (mr *MockLocateClientMockRecorder) LocatePoints(points, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocatePoints", reflect.TypeOf((*MockLocateClient)(nil).LocatePoints), points, options)
}

// LocatePointsWithContext mocks base method.
func (m *MockLocateClient) LocatePointsWithContext(ctx context.Context, points []Point, options CallOptions) ([]Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocatePointsWithContext", ctx, points, options)
	ret0, _ := ret[0].([]Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocatePointsWithContext indicates an expected call of LocatePointsWithContext.
func (mr *MockLocateClientMockRecorder) LocatePointsWithContext(ctx, points, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocatePointsWithContext", reflect.TypeOf((*MockLocateClient)(nil).LocatePointsWithContext), ctx, points, options)
}