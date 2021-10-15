package shell

import (
	"context"
	"fisherman/internal/utils"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/go-errors/errors"
)

type SystemShell struct {
	cwd          string
	defaultShell string
}

func NewShell(options ...shellOption) *SystemShell {
	sh := SystemShell{
		defaultShell: PlatformDefaultShell,
		cwd:          "",
	}

	for _, option := range options {
		option(&sh)
	}

	return &sh
}

func (sh *SystemShell) Exec(ctx context.Context, output io.Writer, shell string, script *Script) error {
	envList := os.Environ()
	for key, value := range script.env {
		envList = append(envList, fmt.Sprintf("%s=%s", key, value))
	}

	config, err := getShellWrapConfiguration(utils.GetOrDefault(shell, sh.defaultShell))
	if err != nil {
		return errors.Errorf("failed to get shell configuration: %w", err)
	}

	command := exec.CommandContext(ctx, config.Path, config.Args...) // nolint gosec
	command.Env = envList
	command.Dir = utils.GetOrDefault(script.dir, sh.cwd)
	command.Stdout = output
	command.Stderr = output

	stdin, err := command.StdinPipe()
	if err != nil {
		return errors.Errorf("failed to get stdin pipe for communication with the process: %w", err)
	}

	go startWriter(stdin, script.commands, config)

	script.duration, err = ExecWithTime(command.Run)
	if err != nil {
		return errors.Errorf("script completed with an error: %w", err)
	}

	return nil
}

func writeCommandLine(stdin io.Writer, command string) {
	if !utils.IsEmpty(command) {
		fmt.Fprintln(stdin, command)
	}
}

func startWriter(stdin io.WriteCloser, commands []string, config WrapConfiguration) {
	defer stdin.Close()

	writeCommandLine(stdin, config.Init)
	for _, commandLine := range commands {
		writeCommandLine(stdin, commandLine)
		writeCommandLine(stdin, config.PostCommand)
	}

	writeCommandLine(stdin, config.Dispose)
}
