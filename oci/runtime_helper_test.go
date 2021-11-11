package oci

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestChangeState(t *testing.T) {
	err := changeState(StateCreating, []ContainerState{}, "1234", "../oci_bundle/fixture/redis")
	if err != nil {
		t.Error(err)
	}
}

func TestGetContainerFolder(t *testing.T) {
	got, err := getContainerFolder()
	if err != nil {
		t.Error(err)
	}
	homefolder, err := os.UserHomeDir()
	if err != nil {
		t.Error(err)
	}
	want := path.Join(homefolder, "containers")
	if got != want {
		t.Error(fmt.Sprintf("TestGetContainerFolder want %s got %s", want, got))
	}
}
