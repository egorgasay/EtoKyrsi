// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	entity "checkwork/internal/entity"
	repository "checkwork/internal/repository"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIStorage is a mock of IStorage interface.
type MockIStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIStorageMockRecorder
}

// MockIStorageMockRecorder is the mock recorder for MockIStorage.
type MockIStorageMockRecorder struct {
	mock *MockIStorage
}

// NewMockIStorage creates a new mock instance.
func NewMockIStorage(ctrl *gomock.Controller) *MockIStorage {
	mock := &MockIStorage{ctrl: ctrl}
	mock.recorder = &MockIStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIStorage) EXPECT() *MockIStorageMockRecorder {
	return m.recorder
}

// AddPullRequest mocks base method.
func (m *MockIStorage) AddPullRequest(link, student string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPullRequest", link, student)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPullRequest indicates an expected call of AddPullRequest.
func (mr *MockIStorageMockRecorder) AddPullRequest(link, student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPullRequest", reflect.TypeOf((*MockIStorage)(nil).AddPullRequest), link, student)
}

// ChangePassword mocks base method.
func (m *MockIStorage) ChangePassword(username, oldPassword, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", username, oldPassword, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockIStorageMockRecorder) ChangePassword(username, oldPassword, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockIStorage)(nil).ChangePassword), username, oldPassword, newPassword)
}

// CheckIsPending mocks base method.
func (m *MockIStorage) CheckIsPending(username string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIsPending", username)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIsPending indicates an expected call of CheckIsPending.
func (mr *MockIStorageMockRecorder) CheckIsPending(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIsPending", reflect.TypeOf((*MockIStorage)(nil).CheckIsPending), username)
}

// CheckPassword mocks base method.
func (m *MockIStorage) CheckPassword(username, password string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPassword", username, password)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPassword indicates an expected call of CheckPassword.
func (mr *MockIStorageMockRecorder) CheckPassword(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPassword", reflect.TypeOf((*MockIStorage)(nil).CheckPassword), username, password)
}

// CreateUser mocks base method.
func (m *MockIStorage) CreateUser(username, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", username, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIStorageMockRecorder) CreateUser(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIStorage)(nil).CreateUser), username, password)
}

// DeleteAccount mocks base method.
func (m *MockIStorage) DeleteAccount() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockIStorageMockRecorder) DeleteAccount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockIStorage)(nil).DeleteAccount))
}

// DeletePullRequest mocks base method.
func (m *MockIStorage) DeletePullRequest(student string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePullRequest", student)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePullRequest indicates an expected call of DeletePullRequest.
func (mr *MockIStorageMockRecorder) DeletePullRequest(student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePullRequest", reflect.TypeOf((*MockIStorage)(nil).DeletePullRequest), student)
}

// DeleteTask mocks base method.
func (m *MockIStorage) DeleteTask(num int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", num)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockIStorageMockRecorder) DeleteTask(num interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockIStorage)(nil).DeleteTask), num)
}

// Disconnect mocks base method.
func (m *MockIStorage) Disconnect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockIStorageMockRecorder) Disconnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockIStorage)(nil).Disconnect))
}

// GetTaskIDAndMsg mocks base method.
func (m *MockIStorage) GetTaskIDAndMsg(username string) (int, sql.NullString, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskIDAndMsg", username)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(sql.NullString)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTaskIDAndMsg indicates an expected call of GetTaskIDAndMsg.
func (mr *MockIStorageMockRecorder) GetTaskIDAndMsg(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskIDAndMsg", reflect.TypeOf((*MockIStorage)(nil).GetTaskIDAndMsg), username)
}

// GetTasks mocks base method.
func (m *MockIStorage) GetTasks() ([]entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks")
	ret0, _ := ret[0].([]entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockIStorageMockRecorder) GetTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockIStorage)(nil).GetTasks))
}

// GetTitle mocks base method.
func (m *MockIStorage) GetTitle(number int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTitle", number)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTitle indicates an expected call of GetTitle.
func (mr *MockIStorageMockRecorder) GetTitle(number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTitle", reflect.TypeOf((*MockIStorage)(nil).GetTitle), number)
}

// GetUsers mocks base method.
func (m *MockIStorage) GetUsers() ([]entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockIStorageMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockIStorage)(nil).GetUsers))
}

// GetWorks mocks base method.
func (m *MockIStorage) GetWorks() ([]repository.Work, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorks")
	ret0, _ := ret[0].([]repository.Work)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorks indicates an expected call of GetWorks.
func (mr *MockIStorageMockRecorder) GetWorks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorks", reflect.TypeOf((*MockIStorage)(nil).GetWorks))
}

// SetPending mocks base method.
func (m *MockIStorage) SetPending(username string, status int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPending", username, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPending indicates an expected call of SetPending.
func (mr *MockIStorageMockRecorder) SetPending(username, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPending", reflect.TypeOf((*MockIStorage)(nil).SetPending), username, status)
}

// SetVerdict mocks base method.
func (m *MockIStorage) SetVerdict(student, verdict string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVerdict", student, verdict)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVerdict indicates an expected call of SetVerdict.
func (mr *MockIStorageMockRecorder) SetVerdict(student, verdict interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVerdict", reflect.TypeOf((*MockIStorage)(nil).SetVerdict), student, verdict)
}

// UpdateTask mocks base method.
func (m *MockIStorage) UpdateTask(num int, title string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", num, title)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockIStorageMockRecorder) UpdateTask(num, title interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockIStorage)(nil).UpdateTask), num, title)
}

// UpdateUserScore mocks base method.
func (m *MockIStorage) UpdateUserScore(student string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserScore", student)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserScore indicates an expected call of UpdateUserScore.
func (mr *MockIStorageMockRecorder) UpdateUserScore(student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserScore", reflect.TypeOf((*MockIStorage)(nil).UpdateUserScore), student)
}
