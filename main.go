package main

import (
	"flag"
	"fmt"
	"github.com/chen-keinan/oci-runtime/oci"
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
	case "create":
		if len(args) < 3 {
			fmt.Print("missing args")
			os.Exit(1)
		}
		err := ociTool.Create(args[1:3])
		if err != nil {
			fmt.Print(err)
		}
	case "start":
		err := ociTool.Start(args[1])
		if err != nil {
			fmt.Print(err)
		}
	case "kill":
		err := ociTool.Kill(args[1])
		if err != nil {
			fmt.Print(err)
		}
	case "delete":
		err := ociTool.Delete(args[1])
		if err != nil {
			fmt.Print(err)
		}
	case "state":
		err := ociTool.State(args[1])
		if err != nil {
			fmt.Print(err)
		}
	default:
		fmt.Printf("operation: %s is not supported", args[0])
	}
}
