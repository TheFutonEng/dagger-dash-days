package types

import (
	"context"

	"dagger.io/dagger"
)

type RunnerContainer struct {
	Runner *dagger.Container
	Ctx    context.Context
	Client *dagger.Client
}

func NewRunnerContainer(ctx context.Context, runner *dagger.Container, client *dagger.Client) *RunnerContainer {
	return &RunnerContainer{
		Runner: runner,
		Ctx:    ctx,
		Client: client,
	}
}
