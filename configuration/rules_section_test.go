package configuration_test

import (
	"fisherman/configuration"
	"fisherman/internal/rules"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestRulesSection_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name   string
		source string
		rules  []configuration.Rule
	}{
		{
			name: "commit-message rule",
			source: `
rules:
  - type: commit-message
    condition: 1 == 1
    prefix: message-prefix
    suffix: message-suffix
    regexp: message-prefix
`,
			rules: []configuration.Rule{
				&rules.CommitMessage{
					BaseRule: rules.BaseRule{
						Type:      rules.CommitMessageType,
						Condition: "1 == 1",
					},
					Prefix: "message-prefix",
					Suffix: "message-suffix",
					Regexp: "message-regexp",
				},
			},
		},
		{
			name: "suppress-commit-files rule",
			source: `
rules:
  - type: suppress-commit-files
    condition: 1 == 1
    globs: ["glob1", "glob2", "glob3"]
    remove-from-index: true
`,
			rules: []configuration.Rule{
				&rules.SuppressCommitFiles{
					BaseRule: rules.BaseRule{
						Type:      rules.SuppressCommitFilesType,
						Condition: "1 == 1",
					},
					Globs:           []string{"glob1", "glob2", "glob3"},
					RemoveFromIndex: true,
				},
			},
		},
		{
			name: "prepare-message rule",
			source: `
rules:
  - type: prepare-message
    condition: 1 == 1
    message: "test message"
`,
			rules: []configuration.Rule{
				&rules.PrepareMessage{
					BaseRule: rules.BaseRule{
						Type:      rules.PrepareMessageType,
						Condition: "1 == 1",
					},
					Message: "test message",
				},
			},
		},
		{
			name: "shell-script rule",
			source: `
rules:
  - type: shell-script
    condition: 1 == 1
    name: "test name"
    commands: ["command1", "command2"]
`,
			rules: []configuration.Rule{
				&rules.ShellScript{
					BaseRule: rules.BaseRule{
						Type:      rules.ShellScriptType,
						Condition: "1 == 1",
					},
					Name:     "test name",
					Commands: []string{"command1", "command2"},
				},
			},
		},
		{
			name: "add-to-index rule",
			source: `
rules:
  - type: add-to-index
    condition: 1 == 1
    globs:
    - glob: demo.go
      required: true
    - glob: test.go
      required: false
`,
			rules: []configuration.Rule{
				&rules.AddToIndex{
					BaseRule: rules.BaseRule{
						Type:      rules.AddToIndexType,
						Condition: "1 == 1",
					},
					Globs: []rules.Glob{
						{Glob: "demo.go", IsRequired: true},
						{Glob: "test.go", IsRequired: false},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := yaml.NewDecoder(strings.NewReader(tt.source))
			decoder.KnownFields(true)

			var section configuration.RulesSection

			err := decoder.Decode(&section)

			assert.NoError(t, err)
			assert.ObjectsAreEqual(configuration.RulesSection{tt.rules}, section)
		})
	}
}

func TestRulesSection_UnmarshalYAML_Error(t *testing.T) {
	tests := []struct {
		name          string
		source        string
		expectedError string
	}{
		{
			name: "unknown markup",
			source: `
rules:
  type: unknown
`,
			expectedError: "yaml: unmarshal errors:\n  line 3: cannot unmarshal !!map into []configuration.ruleDef",
		},
		{
			name: "no type section",
			source: `
rules:
  - condition: 1 == 1
`,
			expectedError: "required property 'type' not defined",
		},
		{
			name: "no type section",
			source: `
rules:
  - type: unknown-type
`,
			expectedError: "type unknown-type is not supported",
		},
	}

	for _, tt := range tests {
		reader := strings.NewReader(tt.source)
		decoder := yaml.NewDecoder(reader)
		decoder.KnownFields(true)

		var section configuration.RulesSection

		err := decoder.Decode(&section)

		assert.EqualError(t, err, tt.expectedError)
	}
}

func TestRulesSection_UnmarshalYAML_InlineStructure(t *testing.T) {
	type testType struct {
		configuration.RulesSection `yaml:",inline"`
		CustomFiled                string `yaml:"custom-filed,omitempty"`
		OtherCustomFiled           int    `yaml:"other-custom-filed,omitempty"`
	}

	source := `
rules:
  - type: commit-message
    condition: 1 == 1
    prefix: message-prefix
    suffix: message-suffix
    regexp: message-prefix
custom-filed: custom-filed-value
other-custom-filed: 11
`
	reader := strings.NewReader(source)
	decoder := yaml.NewDecoder(reader)
	decoder.KnownFields(true)

	var testStructure testType

	err := decoder.Decode(&testStructure)

	assert.NoError(t, err)
	assert.ObjectsAreEqual(
		testType{
			RulesSection: configuration.RulesSection{
				Rules: []configuration.Rule{
					&rules.SuppressCommitFiles{
						BaseRule: rules.BaseRule{
							Type:      rules.SuppressCommitFilesType,
							Condition: "1 == 1",
						},
						Globs:           []string{"glob1", "glob2", "glob3"},
						RemoveFromIndex: true,
					},
				},
			},
			CustomFiled:      "custom-filed-value",
			OtherCustomFiled: 11,
		},
		testStructure,
	)
}
