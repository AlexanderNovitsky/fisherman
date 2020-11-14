package hooks

import (
	"fisherman/actions"
	c "fisherman/config/hooks"
	"fisherman/internal"
	h "fisherman/internal/handling"
	v "fisherman/internal/validation"
)

func PrepareCommitMsg(factory internal.CtxFactory, conf c.PrepareCommitMsgHookConfig, extr v.VariablesExtractor) *h.HookHandler {
	variables, err := extr.Variables(conf.Variables)
	if err != nil {
		panic(err)
	}

	conf.Compile(variables)

	return h.NewHookHandler(
		factory,
		[]h.Action{
			func(ctx internal.SyncContext) (bool, error) {
				return actions.PrepareMessage(ctx, conf.Message)
			},
		},
		NoSyncValidators,
		NoAsyncValidators,
		NoAfterActions,
	)
}
