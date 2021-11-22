package rules

import (
	"fisherman/internal"
	"fisherman/internal/utils"
	pkgutils "fisherman/pkg/utils"
	"io"
	"os/exec"

	"github.com/hashicorp/go-multierror"
)

const ExecType = "exec"

type Exec struct {
	BaseRule `yaml:",inline"`
	Name     string            `yaml:"name"`
	Env      map[string]string `yaml:"env"`
	Output   bool              `yaml:"output"`
	Dir      string            `yaml:"dir"`
	Commands []CommandDef      `yaml:"commands"`
}

type CommandDef struct {
	Program string            `yaml:"program"`
	Args    []string          `yaml:"args"`
	Env     map[string]string `yaml:"env"`
	Output  bool              `yaml:"output"`
	Dir     string            `yaml:"dir"`
}

func (command *CommandDef) Compile(variables map[string]interface{}) {
	utils.FillTemplate(&command.Dir, variables)
	utils.FillTemplate(&command.Program, variables)
	utils.FillTemplatesArray(command.Args, variables)
	utils.FillTemplatesMap(command.Env, variables)
}

func (rule *Exec) GetPosition() byte {
	return Scripts
}

func (rule *Exec) GetPrefix() string {
	if utils.IsEmpty(rule.Name) {
		return ExecType
	}

	return rule.Name
}

func (rule *Exec) Check(ctx internal.ExecutionContext, output io.Writer) error {
	globalEnv := pkgutils.MergeEnvs(ctx.Env(), rule.Env)

	var resultError *multierror.Error
	for _, commandDef := range rule.Commands {
		command := exec.CommandContext(ctx, commandDef.Program, commandDef.Args...) // nolint gosec
		command.Env = pkgutils.MergeEnvs(globalEnv, commandDef.Env)
		command.Dir = utils.FirstNotEmpty(commandDef.Dir, rule.Dir, ctx.Cwd())

		// TODO: Add custom encoding for different shell
		command.Stdout = output
		command.Stderr = output

		err := command.Run()
		if err != nil {
			resultError = multierror.Append(resultError, err)
		}
	}

	return resultError.ErrorOrNil()
}

func (rule *Exec) Compile(variables map[string]interface{}) {
	rule.BaseRule.Compile(variables)
	utils.FillTemplate(&rule.Dir, variables)
	utils.FillTemplatesMap(rule.Env, variables)
	for _, command := range rule.Commands {
		command.Compile(variables)
	}
}
