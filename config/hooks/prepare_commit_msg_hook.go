package hooks

import "fisherman/utils"

type PrepareCommitMsgHookConfig struct {
	Variables Variables `yaml:"variables,omitempty"`
	Message   string    `yaml:"message,omitempty"`
}

func (config *PrepareCommitMsgHookConfig) Compile(variables map[string]interface{}) {
	utils.FillTemplate(&config.Message, variables)
}

func (config *PrepareCommitMsgHookConfig) GetVarsSection() Variables {
	return config.Variables
}

func (*PrepareCommitMsgHookConfig) HasVars() bool {
	return true
}
