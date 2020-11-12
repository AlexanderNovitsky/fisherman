package handlers

import (
	"context"
	"errors"
	"fisherman/clicontext"
	"fisherman/config"
	"fisherman/config/hooks"
	"fisherman/infrastructure"
	"fisherman/mocks"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
)

func TestValidateMessageNotEmpty(t *testing.T) {
	err := errors.New("commit message should not be empty")
	testData := []struct {
		message  string
		notEmpty bool
		config   hooks.CommitMsgHookConfig
		err      error
	}{
		{message: "not empty string", notEmpty: true, err: nil},
		{message: "  not empty string", notEmpty: true, err: nil},
		{message: "", notEmpty: true, err: err},
		{message: "", notEmpty: false, err: nil},
		{message: "   ", notEmpty: true, err: err},
		{message: "   ", notEmpty: false, err: nil},
	}

	for _, tt := range testData {
		t.Run(tt.message, func(t *testing.T) {
			actualError := validateMessage(tt.message, &hooks.CommitMsgHookConfig{NotEmpty: tt.notEmpty})
			assertMultiError(t, actualError, tt.err)
		})
	}
}

func TestValidateMessageCommitPrefix(t *testing.T) {
	err := errors.New("commit message should have prefix '[prefix]'")
	config := hooks.CommitMsgHookConfig{MessagePrefix: "[prefix]"}

	testData := []struct {
		message string
		err     error
	}{
		{message: "[prefix] message", err: nil},
		{message: "message", err: err},
		{message: " [prefix] message", err: err},
		{message: "message[prefix]", err: err},
	}

	for _, tt := range testData {
		t.Run(tt.message, func(t *testing.T) {
			actualError := validateMessage(tt.message, &config)
			assertMultiError(t, actualError, tt.err)
		})
	}
}

func TestValidateMessageCommitSuffix(t *testing.T) {
	err := errors.New("commit message should have suffix '[suffix]'")
	config := hooks.CommitMsgHookConfig{MessageSuffix: "[suffix]"}

	testData := []struct {
		message string
		err     error
	}{
		{message: "[suffix] message", err: err},
		{message: "message", err: err},
		{message: "message [suffix] ", err: err},
		{message: "message [suffix]", err: nil},
	}

	for _, tt := range testData {
		t.Run(tt.message, func(t *testing.T) {
			actualError := validateMessage(tt.message, &config)
			assertMultiError(t, actualError, tt.err)
		})
	}
}

func TestValidateMessageCommitRegexp(t *testing.T) {
	testData := []struct {
		message string
		regexp  string
		err     error
	}{
		{message: "message", regexp: "", err: nil},
		{message: "message", regexp: "^[a-z]*$", err: nil},
		{
			message: "Message",
			regexp:  "^[a-z]*$",
			err:     errors.New("commit message should be matched regular expression '^[a-z]*$'"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.message, func(t *testing.T) {
			config := hooks.CommitMsgHookConfig{MessageRegexp: tt.regexp}
			actualError := validateMessage(tt.message, &config)
			assertMultiError(t, actualError, tt.err)
		})
	}
}

func assertMultiError(t *testing.T, multipleErrors *multierror.Error, expectedError error) {
	if expectedError != nil {
		assert.NotNil(t, multipleErrors)
		assert.Contains(t, multipleErrors.Errors, expectedError)
	} else {
		assert.Nil(t, multipleErrors)
	}
}

func TestCommitMsgHandler(t *testing.T) {
	var handler CommitMsgHandler

	tests := []struct {
		name string
		args []string
		err  error
	}{
		{name: "base test", args: []string{".git/MESSAGE"}, err: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := clicontext.NewContext(context.TODO(), clicontext.Args{
				Config: &config.DefaultConfig,
				Repository: mocks.NewRepositoryMock(t).
					GetCurrentBranchMock.Return("develop", nil).
					GetLastTagMock.Return("0.0.0", nil).
					GetUserMock.Return(infrastructure.User{}, nil),
				FileSystem: mocks.NewFileSystemMock(t).
					ReadMock.When(".git/MESSAGE").Then("[fisherman] test commit", nil),
				App: &clicontext.AppInfo{},
			})
			err := handler.Handle(ctx, tt.args)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestCommitMsgHandler_IsConfigured(t *testing.T) {
	var handler CommitMsgHandler

	tests := []struct {
		name     string
		config   *config.HooksConfig
		expected bool
	}{
		{
			name:     "empty structure",
			config:   &config.HooksConfig{},
			expected: false,
		},
		{
			name: "",
			config: &config.HooksConfig{
				CommitMsgHook: hooks.CommitMsgHookConfig{
					MessageRegexp: "test",
				},
			},
			expected: true,
		},
		{
			name: "",
			config: &config.HooksConfig{
				CommitMsgHook: hooks.CommitMsgHookConfig{
					NotEmpty: true,
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := handler.IsConfigured(tt.config)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestCommitMsgHandlerNoArgs(t *testing.T) {
	var handler CommitMsgHandler
	ctx := clicontext.NewContext(context.TODO(), clicontext.Args{
		Config: &config.DefaultConfig,
		App:    &clicontext.AppInfo{},
	})
	err := handler.Handle(ctx, []string{})
	assert.Error(t, err, "commit message file argument is not presented")
}
