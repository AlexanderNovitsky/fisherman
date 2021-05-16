package configuration

import (
	"fisherman/internal"
	"fisherman/utils"
	"fmt"
	"io"
	"regexp"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v3"
)

// TODO: Add new method in Rule interface to Decode rule from map[string]interface{} and
// try implement comman realization in base rule structure.
type Rule interface {
	GetType() string
	GetContition() string
	GetPosition() byte
	Check(internal.ExecutionContext, io.Writer) error
	Compile(map[string]interface{})
}

type ExtractVariable struct {
	Variable   string `yaml:"variable"`
	Expression string `yaml:"expression"`
}

type Variables = map[string]interface{}

type HookConfig struct {
	StaticVariables  map[string]string
	ExtractVariables []ExtractVariable
	Rules            []Rule
}

type hookConfigDef struct {
	StaticVariables  map[string]string `yaml:"variables,omitempty"`
	ExtractVariables []ExtractVariable `yaml:"extract-variables,omitempty"`
	Rules            []ruleDef         `yaml:"rules"`
}

func (c *HookConfig) UnmarshalYAML(value *yaml.Node) error {
	var def hookConfigDef
	err := value.Decode(&def)
	if err != nil {
		return err
	}

	c.ExtractVariables = def.ExtractVariables
	c.StaticVariables = def.StaticVariables

	for _, ruleDef := range def.Rules {
		c.Rules = append(c.Rules, ruleDef.Rule)
	}

	return nil
}

func (c *HookConfig) Compile(global Variables) (Variables, error) {
	variables := map[string]interface{}{}
	err := mergo.MergeWithOverwrite(&variables, global)
	if err != nil {
		return variables, err
	}

	if c.StaticVariables != nil {
		utils.FillTemplatesMap(c.StaticVariables, variables)

		interfaceMap := utils.StringMapToInterfaceMap(c.StaticVariables)
		err = mergo.MergeWithOverwrite(&variables, interfaceMap)
		if err != nil {
			return variables, err
		}
	}

	if c.ExtractVariables != nil {
		for _, value := range c.ExtractVariables {
			targetVar, ok := variables[value.Variable]
			if !ok {
				return nil, fmt.Errorf("variable '%s' is not defined", value.Variable)
			}

			extractedVariables, err := extract(targetVar.(string), value.Expression)
			if err != nil {
				return nil, err
			}

			err = mergo.MergeWithOverwrite(&variables, extractedVariables)
			if err != nil {
				return variables, err
			}
		}
	}

	for _, rule := range c.Rules {
		rule.Compile(variables)
	}

	return variables, nil
}

func extract(source, expression string) (map[string]interface{}, error) {
	variables := make(map[string]interface{})
	if !utils.IsEmpty(expression) && !utils.IsEmpty(source) {
		reg, err := regexp.Compile(expression)
		if err != nil {
			return nil, err
		}

		match := reg.FindStringSubmatch(source)
		if match == nil {
			return variables, nil
		}

		for i, name := range reg.SubexpNames() {
			if !utils.IsEmpty(name) {
				variables[name] = match[i]
			}
		}
	}

	return variables, nil
}
