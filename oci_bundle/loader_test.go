package oci_bundle

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadBundle(t *testing.T) {
	fp, err := filepath.Abs("./fixture/")
	if err != nil {
		t.Error(err)
	}
	os.Setenv("CONTAINER_HOME", fp)
	fileData, err := LoadBundle("redis")
	if err != nil {
		t.Error("failed to load bundle")
	}
	if len(fileData) != 1 {
		t.Error(fmt.Sprintf("Load bundle want %d got %d", 1, len(fileData)))
	}
}
