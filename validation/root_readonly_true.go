package main

import (
	"os"
	"runtime"

	"github.com/opencontainers/runtime-tools/validation/util"
)

func main() {
	if "windows" == runtime.GOOS {
		util.Skip("non-Windows root.readonly test", runtime.GOOS)
		os.Exit(0)
	}

	g := util.GetDefaultGenerator()
	g.SetRootReadonly(true)
	err := util.RuntimeInsideValidate(g, nil)
	if err != nil {
		util.Fatal(err)
	}
}
