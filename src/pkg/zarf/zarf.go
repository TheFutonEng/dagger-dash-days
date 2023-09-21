package zarf

import (
	"fmt"
	"main/src/pkg/dagger/types"
)

var (
	wgetZarf    = []string{"wget", "https://github.com/defenseunicorns/zarf/releases/download/v0.29.2/zarf_v0.29.2_Linux_arm64"}
	chmodZarf   = []string{"chmod", "+x", "zarf_v0.29.2_Linux_arm64"}
	zarfVersion = []string{"./zarf_v0.29.2_Linux_arm64", "version"}
)

func SetupZarf(runner *types.RunnerContainer) {
	container := runner.Runner
	container = container.WithWorkdir("/bin").WithExec(wgetZarf)
	container = container.WithExec(chmodZarf)

	out, err := container.WithExec(zarfVersion).Stdout(runner.Ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Using Zarf Version: " + out)
}
