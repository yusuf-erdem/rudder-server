// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/services/streammanager/kinesis (interfaces: KinesisClient)

// Package mock_kinesis is a generated GoMock package.
package mock_kinesis

import (
	reflect "reflect"

	kinesis "github.com/aws/aws-sdk-go/service/kinesis"
	gomock "github.com/golang/mock/gomock"
)

// MockKinesisClient is a mock of KinesisClient interface.
type MockKinesisClient struct {
	ctrl     *gomock.Controller
	recorder *MockKinesisClientMockRecorder
}

// MockKinesisClientMockRecorder is the mock recorder for MockKinesisClient.
type MockKinesisClientMockRecorder struct {
	mock *MockKinesisClient
}

// NewMockKinesisClient creates a new mock instance.
func NewMockKinesisClient(ctrl *gomock.Controller) *MockKinesisClient {
	mock := &MockKinesisClient{ctrl: ctrl}
	mock.recorder = &MockKinesisClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKinesisClient) EXPECT() *MockKinesisClientMockRecorder {
	return m.recorder
}

// PutRecord mocks base method.
func (m *MockKinesisClient) PutRecord(arg0 *kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutRecord", arg0)
	ret0, _ := ret[0].(*kinesis.PutRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutRecord indicates an expected call of PutRecord.
func (mr *MockKinesisClientMockRecorder) PutRecord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRecord", reflect.TypeOf((*MockKinesisClient)(nil).PutRecord), arg0)
}