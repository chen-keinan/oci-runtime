package oci

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestChangeState(t *testing.T) {
	tests := []struct {
		name      string
		newState  ContainerState
		prevState []ContainerState
		home      string
		params    []string
		want      error
	}{
		{name: "change stat to create", newState: StateCreating, home: "../oci_bundle/fixture/", params: []string{"1234", "redis"}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathAbs, err := filepath.Abs(tt.home)
			if err != nil {
				t.Error(err)
			}
			os.Setenv("CONTAINER_HOME", pathAbs)
			got := changeState(tt.newState, tt.prevState, tt.params...)
			if tt.want != got {
				t.Errorf("TestChangeState(),not expected value")
			}
		})
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
