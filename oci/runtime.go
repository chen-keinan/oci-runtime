package oci

import "fmt"

type Oci interface {
	Create(args []string) error
	Run(args []string) error
	Kill(args []string) error
	Delete(args []string) error
}

type OciRuntime struct {
}

func NewOciRuntime() Oci {
	return &OciRuntime{}
}
func (or OciRuntime) Create(args []string) error {
	if len (args) != 2 {
		return fmt.Errorf("must provide <container-id> <path-to-bundle> args only")
	}
	return nil
}
func (or OciRuntime) Run(args []string) error {
	return nil
}
func (or OciRuntime) Kill(args []string) error {
	return nil
}
func (or OciRuntime) Delete(args []string) error {
	return nil
}

