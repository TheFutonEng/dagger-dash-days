package zarf

import (
	"fmt"
	"main/src/pkg/dagger/types"
	"runtime"
)

var (
	arch, os   string
	wgetZarf   []string
	chmodZarf  []string
	zarfVersion []string
)

func init() {
	// Determine CPU architecture
	switch runtime.GOARCH {
	case "amd64":
		arch = "amd64"
	case "arm64":
		arch = "arm64"
	default:
		arch = "unknown"
	}

	// Construct the variables based on `arch` and `os`
	version := "v0.29.2"
	baseURL := "https://github.com/defenseunicorns/zarf/releases/download"
	filename := fmt.Sprintf("zarf_%s_%s_%s", version, "Linux", arch)

	wgetZarf = []string{"wget", fmt.Sprintf("%s/%s/%s", baseURL, version, filename)}
	chmodZarf = []string{"chmod", "+x", filename}
	zarfVersion = []string{"./" + filename, "version"}
}

// var (
// 	wgetZarf    = []string{"wget", "https://github.com/defenseunicorns/zarf/releases/download/v0.29.2/zarf_v0.29.2_Linux_arm64"}
// 	chmodZarf   = []string{"chmod", "+x", "zarf_v0.29.2_Linux_arm64"}
// 	zarfVersion = []string{"./zarf_v0.29.2_Linux_arm64", "version"}
	
// )

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
