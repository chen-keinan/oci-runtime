package oci

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
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
		return fmt.Errorf("must provide <container-id> <path-to-bundle> args only")
	}
	err := ChangeState(StateCreating, []ContainerState{}, args[0], args[1])
	if err != nil {
		return err
	}
	return ChangeState(StateCreated, []ContainerState{StateCreating}, args[0], args[1])
}

func (or OciRuntime) Start(arg string) error {
	if len(arg) == 0 {
		return fmt.Errorf("must provide <container-id> args only")
	}
	return ChangeState(StateRunning, []ContainerState{StateStopped, StateCreated}, arg)
}
func (or OciRuntime) Kill(arg string) error {
	if len(arg) == 1 {
		return fmt.Errorf("must provide <container-id> args only")
	}
	return ChangeState(StateStopped, []ContainerState{StateRunning}, arg)
}
func (or OciRuntime) Delete(arg string) error {
	if len(arg) == 1 {
		return fmt.Errorf("must provide <container-id> args only")
	}
	return ChangeState(StateDeleted, []ContainerState{StateStopped}, arg)
}
func (or OciRuntime) State(arg string) error {
	if len(arg) == 1 {
		return fmt.Errorf("must provide <container-id> args only")
	}
	state, err := GetState(arg)
	var itoa string
	if err != nil {
		state = &State{}
		itoa = ""
	} else {
		itoa = strconv.Itoa(state.Pid)
	}
	data := [][]string{
		{state.ID, string(state.Status), state.Bundle, itoa, state.Version},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status", "Bundle", "PID", "Version"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
	return nil
}
