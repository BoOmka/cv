//go:build wasm
// +build wasm

package main

import (
	"fmt"

	"flag"

	"github.com/boomka/cv/app"
	"github.com/vugu/vugu"
	"github.com/vugu/vugu/domrender"
)

func main() {

	mountPoint := flag.String("mount-point", "#root", "The query selector for the mount point for the root component, if it is not a full HTML component")
	//mountPoint := flag.String("mount-point", ":root", "The query selector for the mount point for the root component, if it is not a full HTML component")
	flag.Parse()

	fmt.Printf("Entering main(), -mount-point=%q\n", *mountPoint)
	defer fmt.Printf("Exiting main()\n")

	renderer, err := domrender.New(*mountPoint, true)
	if err != nil {
		panic(err)
	}
	defer renderer.Release()

	buildEnv, err := vugu.NewBuildEnv(renderer.EventEnv())
	if err != nil {
		panic(err)
	}

	_, rootBuilder := app.VuguSetup(buildEnv, renderer.EventEnv(), &app.VuguSetupOptions{AutoReload: false})
	for ok := true; ok; ok = renderer.EventWait() {

		buildResults := buildEnv.RunBuild(rootBuilder)

		err = renderer.Render(buildResults)
		if err != nil {
			panic(err)
		}
	}
}
