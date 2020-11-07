package hooks_test

import (
	"fisherman/config/hooks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariables_GetFromBranch(t *testing.T) {
	tests := []struct {
		name              string
		variables         hooks.Variables
		branchName        string
		expectedVariables map[string]interface{}
		err               error
	}{
		{
			name:       "Parse single variable",
			branchName: "refs/heads/develop",
			err:        nil,
			expectedVariables: map[string]interface{}{
				"DEMO": "develop",
			},
			variables: hooks.Variables{FromBranch: "refs/heads/(?P<DEMO>.*)"},
		},
		{
			name:       "Parse multiple variables",
			branchName: "refs/heads/develop",
			err:        nil,
			expectedVariables: map[string]interface{}{
				"DEMO":   "develop",
				"ROOT":   "refs",
				"FOLDER": "heads",
			},
			variables: hooks.Variables{FromBranch: "(?P<ROOT>.*)/(?P<FOLDER>.*)/(?P<DEMO>.*)"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			variables, err := tt.variables.GetFromBranch(tt.branchName)
			assert.Equal(t, tt.err, err)
			if tt.err == nil {
				assert.Equal(t, tt.expectedVariables, variables)
			}
		})
	}
}

func TestVariables_GetFromTag(t *testing.T) {
	tests := []struct {
		name              string
		variables         hooks.Variables
		tagName           string
		expectedVariables map[string]interface{}
		err               string
	}{
		{
			name:    "correct FromLastTag expression",
			tagName: "refs/tags/v1.0.0",
			err:     "",
			expectedVariables: map[string]interface{}{
				"V": "v1.0.0",
			},
			variables: hooks.Variables{FromLastTag: "refs/tags/(?P<V>.*)"},
		},
		{
			name:              "not matched FromLastTag expression",
			tagName:           "refs/tags/v1.0.0",
			err:               "filed match 'refs/tags/v1.0.0' to expression 'xxx/tags/(?P<V>.*)'",
			expectedVariables: nil,
			variables:         hooks.Variables{FromLastTag: "xxx/tags/(?P<V>.*)"},
		},
		{
			name:              "incorrect FromLastTag expression",
			tagName:           "refs/tags/v1.0.0",
			err:               "error parsing regexp: missing closing ): `xxx/tags/(((?P<V>.*)`",
			expectedVariables: nil,
			variables:         hooks.Variables{FromLastTag: "xxx/tags/(((?P<V>.*)"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			variables, err := tt.variables.GetFromTag(tt.tagName)

			if len(tt.err) == 0 {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedVariables, variables)
			} else {
				assert.EqualError(t, err, tt.err)
			}
		})
	}
}
