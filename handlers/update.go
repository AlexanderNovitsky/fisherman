package handlers

import (
	"fisherman/clicontext"
	"fisherman/constants"
	"fmt"
)

// UpdateHandler is a handler for update hook
func UpdateHandler(ctx *clicontext.CommandContext, args []string) error {
	return fmt.Errorf("this hook is not supported in version %s", constants.Version)
}
