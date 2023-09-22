package dagger

import (
	"context"
	"main/src/pkg/dagger/types"
	"os"

	"dagger.io/dagger"
)

func SetupDagger() *types.RunnerContainer {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}

	runner := client.Container().
		From("golang:1.20")

	return types.NewRunnerContainer(ctx, runner, client)
}
