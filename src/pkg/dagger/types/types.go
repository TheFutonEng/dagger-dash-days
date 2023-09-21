package types

import (
	"context"

	"dagger.io/dagger"
)

type RunnerContainer struct {
	Runner *dagger.Container
	Ctx    context.Context
}

func NewRunnerContainer(ctx context.Context, runner *dagger.Container) *RunnerContainer {
	return &RunnerContainer{
		Runner: runner,
		Ctx:    ctx,
	}
}
