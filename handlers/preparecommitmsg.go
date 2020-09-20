package handlers

import (
	"fisherman/commands/context"
	"fisherman/utils"
	"regexp"
)

// PrepareCommitMsgHandler is a execute function for prepare-commit-msg hook
func PrepareCommitMsgHandler(ctx *context.CommandContext, args []string) error {
	config := ctx.Config.Hooks
	info, err := ctx.GetGitInfo()
	utils.HandleCriticalError(err)

	c := config.PrepareCommitMsgHook
	if c != nil {
		message, isPresented := getPreparedMessage(c.Message, c.BranchRegExp, info.CurrentBranch)
		if isPresented {
			err = ctx.Files.Write(args[0], message)
			utils.HandleCriticalError(err)
		}
	}

	return nil
}

func getPreparedMessage(message, regexpString, branch string) (string, bool) {
	if !utils.IsEmpty(message) {
		if !utils.IsEmpty(regexpString) {
			return regexp.MustCompile(regexpString).
				ReplaceAllString(branch, message), true
		}

		return message, true
	}
	return "", false
}
