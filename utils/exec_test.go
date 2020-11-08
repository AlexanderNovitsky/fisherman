package utils_test

import (
	"fisherman/utils"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizePath(t *testing.T) {
	pingFullPath, err := exec.LookPath("ping")
	assert.NoError(t, err)

	tests := []struct {
		name     string
		binary   string
		expected string
	}{
		{
			name:     "binary not registered in PATH",
			binary:   filepath.Join("/", "demo", "not-exist-binary"),
			expected: filepath.Join("/", "demo", "not-exist-binary"),
		},
		{
			name:     "global defined commands",
			binary:   "ping",
			expected: "ping",
		},
		{
			name:     "binary registered in PATH",
			binary:   pingFullPath,
			expected: filepath.Base(pingFullPath),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := utils.NormalizePath(tt.binary)
			assert.Equal(t, tt.expected, path)
		})
	}
}
