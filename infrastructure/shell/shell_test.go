package shell

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemShell_Exec(t *testing.T) {
	sh := NewShell()

	tests := []struct {
		name           string
		commands       []string
		env            *map[string]string
		expectedStdout string
		err            error
	}{
		{
			name:           "should return 1,2",
			commands:       []string{"echo 1", "echo 2"},
			env:            &map[string]string{"demo": "demo"},
			err:            nil,
			expectedStdout: fmt.Sprintf("1%s2%s", LineBreak, LineBreak),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sh.Exec(tt.commands, tt.env, true)
			assert.Equal(t, tt.expectedStdout, result.Output)
			assert.Equal(t, 0, result.ExitCode)
			assert.NoError(t, result.Error)
		})
	}
}
