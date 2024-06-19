// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/artefactual-sdps/enduro/internal/api/auth (interfaces: TokenVerifier)
//
// Generated by this command:
//
//	mockgen -typed -destination=./internal/api/auth/fake/mock_token_verifier.go -package=fake github.com/artefactual-sdps/enduro/internal/api/auth TokenVerifier
//

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	auth "github.com/artefactual-sdps/enduro/internal/api/auth"
	gomock "go.uber.org/mock/gomock"
)

// MockTokenVerifier is a mock of TokenVerifier interface.
type MockTokenVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockTokenVerifierMockRecorder
}

// MockTokenVerifierMockRecorder is the mock recorder for MockTokenVerifier.
type MockTokenVerifierMockRecorder struct {
	mock *MockTokenVerifier
}

// NewMockTokenVerifier creates a new mock instance.
func NewMockTokenVerifier(ctrl *gomock.Controller) *MockTokenVerifier {
	mock := &MockTokenVerifier{ctrl: ctrl}
	mock.recorder = &MockTokenVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenVerifier) EXPECT() *MockTokenVerifierMockRecorder {
	return m.recorder
}

// Verify mocks base method.
func (m *MockTokenVerifier) Verify(arg0 context.Context, arg1 string) (*auth.Claims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1)
	ret0, _ := ret[0].(*auth.Claims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockTokenVerifierMockRecorder) Verify(arg0, arg1 any) *MockTokenVerifierVerifyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockTokenVerifier)(nil).Verify), arg0, arg1)
	return &MockTokenVerifierVerifyCall{Call: call}
}

// MockTokenVerifierVerifyCall wrap *gomock.Call
type MockTokenVerifierVerifyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTokenVerifierVerifyCall) Return(arg0 *auth.Claims, arg1 error) *MockTokenVerifierVerifyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTokenVerifierVerifyCall) Do(f func(context.Context, string) (*auth.Claims, error)) *MockTokenVerifierVerifyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTokenVerifierVerifyCall) DoAndReturn(f func(context.Context, string) (*auth.Claims, error)) *MockTokenVerifierVerifyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
