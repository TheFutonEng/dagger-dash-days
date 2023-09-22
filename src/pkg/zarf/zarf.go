package zarf

import (
	"fmt"
	"main/src/pkg/dagger/types"
	"main/src/pkg/message"
	"os"
	"runtime"
)

var (
	arch        string
	wgetZarf    []string
	chmodZarf   []string
	zarfVersion []string
	filename    string
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
	filename = fmt.Sprintf("zarf_%s_%s_%s", version, "Linux", arch)

	wgetZarf = []string{"wget", fmt.Sprintf("%s/%s/%s", baseURL, version, filename)}
	chmodZarf = []string{"chmod", "+x", filename}
	zarfVersion = []string{"./" + filename, "version"}
}

func SetupZarf(runner *types.RunnerContainer) {
	container := runner.Runner
	container = container.WithWorkdir("./project").WithExec(wgetZarf).WithExec(chmodZarf).WithExec(zarfVersion)
	out, err := container.Stdout(runner.Ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Using Zarf Version: " + out)

	out, err = container.
		WithDirectory(".", runner.Client.Host().Directory(".")).
		WithExec([]string{"ls", "zarf.yaml"}).
		Stdout(runner.Ctx)
	if err != nil {
		message.TellUserOnFail("zarf.yaml not found", "Please create a zarf.yaml")
		os.Exit(1)
	}
	message.Info(out + " found. Building Zarf package...")
	out, err = container.WithDirectory(".", runner.Client.Host().Directory(".")).WithExec([]string{"./" + filename, "package", "create", "--confirm"}).Stdout(runner.Ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
