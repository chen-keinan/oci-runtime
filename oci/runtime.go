package oci

import (
	"fmt"
	"strconv"
)

type Oci interface {
	State(args string) error
	Create(args []string) error
	Start(arg string) error
	Kill(arg string) error
	Delete(arg string) error
}

type OciRuntime struct {
}

func NewOciRuntime() Oci {
	return &OciRuntime{}
}
func (or OciRuntime) Create(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("must provide <container-id> and <path-to-bundle> args ")
	}
	err := changeState(StateCreating, []ContainerState{}, args[0], args[1])
	if err != nil {
		return err
	}
	return changeState(StateCreated, []ContainerState{StateCreating}, args[0], args[1])
}

func (or OciRuntime) Start(arg string) error {
	if len(arg) == 0 {
		return fmt.Errorf("must provide <container-id> arg")
	}
	return changeState(StateRunning, []ContainerState{StateStopped, StateCreated}, arg)
}
func (or OciRuntime) Kill(arg string) error {
	if len(arg) == 1 {
		return fmt.Errorf("must provide <container-id> arg")
	}
	return changeState(StateStopped, []ContainerState{StateRunning}, arg)
}
func (or OciRuntime) Delete(arg string) error {
	if len(arg) == 1 {
		return fmt.Errorf("must provide <container-id> arg")
	}
	return changeState(StateDeleted, []ContainerState{StateStopped}, arg)
}
func (or OciRuntime) State(arg string) error {
	if len(arg) == 1 {
		return fmt.Errorf("must provide <container-id> arg")
	}
	state, err := getState(arg)
	if err != nil || len(state) == 0 {
		state = []*State{{}}
		printView(state)
		return nil
	} else {
		for _, st := range state {
			st.PidString = strconv.Itoa(st.Pid)
		}
	}
	printView(state)
	return nil
}
