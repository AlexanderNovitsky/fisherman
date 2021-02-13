package main

import (
	"context"
	"fisherman/commands"
	"fisherman/commands/handle"
	"fisherman/commands/initialize"
	"fisherman/commands/remove"
	"fisherman/commands/version"
	"fisherman/configuration"
	"fisherman/infrastructure/filesystem"
	"fisherman/infrastructure/log"
	"fisherman/infrastructure/shell"
	"fisherman/infrastructure/vcs"
	"fisherman/internal"
	"fisherman/internal/expression"
	"fisherman/internal/hookfactory"
	"fisherman/internal/runner"
	"fisherman/utils"
	"os"
	"os/user"
)

const fatalExitCode = 1

func main() {
	defer utils.PanicInterceptor(os.Exit, fatalExitCode)

	usr, err := user.Current()
	utils.HandleCriticalError(err)

	cwd, err := os.Getwd()
	utils.HandleCriticalError(err)

	executable, err := os.Executable()
	utils.HandleCriticalError(err)

	fileSystem := filesystem.NewLocalFileSystem()

	configLoader := configuration.NewLoader(usr, cwd, fileSystem)
	configFiles, err := configLoader.FindConfigFiles()
	utils.HandleCriticalError(err)

	config, err := configLoader.Load(configFiles)
	utils.HandleCriticalError(err)

	log.Configure(config.Output)

	ctx := context.Background()
	sysShell := shell.NewShell(os.Stdout, cwd, config.DefaultShell)
	repository := vcs.NewGitRepository(cwd)

	engine := expression.NewExpressionEngine(config.GlobalVariables)

	ctxFactory := internal.NewCtxFactory(ctx, fileSystem, sysShell, repository)
	hookFactory := hookfactory.NewFactory(engine, config.Hooks)

	appInfo := internal.AppInfo{
		Executable: executable,
		Cwd:        cwd,
		Configs:    configFiles,
	}

	commands := []commands.CliCommand{
		initialize.NewCommand(fileSystem, &appInfo, usr),
		handle.NewCommand(hookFactory, ctxFactory, &config.Hooks, &appInfo),
		remove.NewCommand(fileSystem, &appInfo, usr),
		version.NewCommand(),
	}

	instance := runner.NewRunner(commands)
	if err = instance.Run(os.Args[1:]); err != nil {
		panic(err)
	}
}
