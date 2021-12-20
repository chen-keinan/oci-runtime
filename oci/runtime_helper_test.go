package oci

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestChangeState(t *testing.T) {
	filePath, err := filepath.Abs("../oci_bundle/fixture/1234")
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(filePath); err == nil {
		err = os.Remove(filePath)
		if err != nil {
			t.Error(err)
		}
	}
	tests := []struct {
		name      string
		newState  ContainerState
		prevState []ContainerState
		home      string
		params    []string
		want      error
	}{
		{name: "change stat to create", newState: StateCreating, prevState: []ContainerState{}, home: "../oci_bundle/fixture/", params: []string{"1234", "redis"}, want: nil},
		{name: "change stat to run", newState: StateRunning, prevState: []ContainerState{StateCreating}, home: "../oci_bundle/fixture/", params: []string{"1234", "redis"}, want: nil},
		{name: "change stat to stop", newState: StateStopped, prevState: []ContainerState{StateRunning}, home: "../oci_bundle/fixture/", params: []string{"1234", "redis"}, want: nil},
		{name: "change stat to delete", newState: StateDeleted, prevState: []ContainerState{StateStopped}, home: "../oci_bundle/fixture/", params: []string{"1234", "redis"}, want: nil},
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

func TestGetState(t *testing.T) {
	filePath, err := filepath.Abs("../oci_bundle/fixture/1234")
	if err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(filePath); err == nil {
		err = os.Remove(filePath)
		if err != nil {
			t.Error(err)
		}
	}
	tests := []struct {
		name      string
		newState  ContainerState
		prevState []ContainerState
		home      string
		cID       string
		params    []string
		want      []*State
		wantError error
	}{
		{name: "get state to no error", newState: StateCreating, prevState: []ContainerState{}, home: "../oci_bundle/fixture/", cID: "1234", params: []string{"1234", "redis"}, want: []*State{{ID: "1234"}}, wantError: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathAbs, err := filepath.Abs(tt.home)
			if err != nil {
				t.Error(err)
			}
			os.Setenv("CONTAINER_HOME", pathAbs)
			gotErr := changeState(tt.newState, tt.prevState, tt.params...)
			if gotErr != nil {
				t.Error(err)
			}
			got, gotErr := getState(tt.cID)
			if got != nil && got[0].ID != tt.want[0].ID {
				t.Errorf("TestGetState(),want %s got %s", tt.want[0].ID, got[0].ID)
			}
			if tt.wantError != gotErr {
				t.Errorf("TestGetState(),not expected value")
			}
		})
	}
}

func TestGetContainerFolder(t *testing.T) {
	os.Setenv("CONTAINER_HOME", "")
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

func TestPrintView(t *testing.T) {
	state := []*State{{Pid: 32342, ID: "1234", Status: "running", Bundle: "redis", Version: "1.0"}}
	err := printView(state)
	if err != nil {
		t.Errorf("TestPrintView: should not throw an error")
	}
}
