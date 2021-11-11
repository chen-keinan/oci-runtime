package main

import (
	"flag"
	"fmt"
	"github.com/chen-keinan/oci-client/oci"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Print("missing args")
		os.Exit(1)
	}
	ociTool := oci.NewOciRuntime()
	switch args[0] {
	case "Create":
		if len(args) < 3 {
			fmt.Print("missing args")
			os.Exit(1)
		}
		err := ociTool.Create(args[0:2])
		if err != nil {
			fmt.Print(err)
		}
	case "Start":
		err := ociTool.Start(args[1])
		if err != nil {
			fmt.Print(err)
		}
	case "Kill":
		err := ociTool.Kill(args[1])
		if err != nil {
			fmt.Print(err)
		}
	case "Delete":
		err := ociTool.Delete(args[1])
		if err != nil {
			fmt.Print(err)
		}
	case "State":
		err := ociTool.State(args[1])
		if err != nil {
			fmt.Print(err)
		}
	}
}
