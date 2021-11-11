package oci

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chen-keinan/oci-client/oci_bundle"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
)

func ChangeState(newStatus ContainerState, oldStatus []ContainerState, params ...string) error {
	if len(params) < 1 {
		return fmt.Errorf("failed to create container missing params")
	}
	cFolder, err := GetContainerFolder()
	filePath := path.Join(cFolder, params[0])
	switch newStatus {
	case StateCreating:
		return CreatingContainer(params, err, filePath)
	case StateRunning, StateStopped, StateCreated:
		err = ChangeContainerStates(newStatus, oldStatus, filePath)
		if err != nil {
			return err
		}
	case StateDeleted:
		err = DeleteContainer(oldStatus, filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func ChangeContainerStates(newStatus ContainerState, oldStatus []ContainerState, filePath string) error {
	state, err := oci_bundle.ReadFile(filePath)
	if err != nil {
		return err
	}
	var st State
	err = json.Unmarshal([]byte(state), &st)
	if err != nil {
		return err
	}
	err = matchOldState(oldStatus, state, st)
	if err != nil {
		return err
	}
	st.Status = newStatus
	b, err := json.Marshal(&st)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, b, 0)
	if err != nil {
		return err
	}
	return nil
}

func DeleteContainer(oldStatus []ContainerState, filePath string) error {
	state, err := oci_bundle.ReadFile(filePath)
	if err != nil {
		return err
	}
	var st State
	err = json.Unmarshal([]byte(state), &st)
	if err != nil {
		return err
	}
	err = matchOldState(oldStatus, state, st)
	if err != nil {
		return err
	}
	err = os.Remove(filePath)
	if err != nil {
		return err
	}
	return nil
}

func matchOldState(oldStatus []ContainerState, state string, st State) error {
	if len(oldStatus) == 0 && len(state) > 0 {
		return fmt.Errorf("container has already created")
	}
	err := json.Unmarshal([]byte(state), &st)
	if err != nil {
		return err
	}
	for _, ost := range oldStatus {
		if st.Status == ost {
			break
		}
		return fmt.Errorf("container state %s cannot be changed from state %s", st.Status, oldStatus)
	}
	return nil
}

func CreatingContainer(params []string, err error, filePath string) error {
	if len(params) < 2 {
		return fmt.Errorf("failed to create container missing params")
	}
	// check if container has been created already
	if _, err := os.Stat(filePath); !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to create container contianer file already exist")
	}
	fileData, err := oci_bundle.LoadBundle(params[1])
	if err != nil {
		return fmt.Errorf("failed to create container")
	}
	if len(fileData) < 1 {
		return fmt.Errorf("failed to create container bundle file is missing")
	}
	newSt := State{Version: "1.0", ID: params[0], Status: StateCreating, Bundle: params[1], Pid: rand.Int()}
	b, err := json.Marshal(&newSt)
	if err != nil {
		return fmt.Errorf("failed to create container")
	}
	err = ioutil.WriteFile(filePath, b, 0777)
	if err != nil {
		return fmt.Errorf("failed to create container")
	}
	return nil
}

func GetContainerFolder() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	containerFolder := path.Join(dir, "containers")
	if _, err := os.Stat(containerFolder); os.IsNotExist(err) {
		err := os.Mkdir(containerFolder, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return containerFolder, err
}


func GetState(containerID string) (*State, error) {
	cf, err := GetContainerFolder()
	if err != nil {
		return nil, err
	}
	fPath := path.Join(cf, containerID)
	stData, err := oci_bundle.ReadFile(fPath)
	if err != nil {
		return nil, err
	}
	var st State
	err = json.Unmarshal([]byte(stData), &st)
	if err != nil {
		return nil, err
	}
	return &st, nil
}
