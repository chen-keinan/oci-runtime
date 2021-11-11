package oci_bundle

import (
	"fmt"
	"testing"
)

func TestLoadBundle(t *testing.T) {
	fileData, err := LoadBundle("./fixture/redis")
	if err != nil {
		t.Error("failed to load bundle")
	}
	if len(fileData) != 1 {
		t.Error(fmt.Sprintf("Load bundle want %d got %d",1,len(fileData)))
	}
}
