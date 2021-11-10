package oci_bundle

import (
	"fmt"
	"testing"
)

func TestLoadBundle(t *testing.T) {
	fileData, err := LoadBundle("./fixture/redis.tgz")
	if err != nil {
		t.Error("failed to load bundle")
	}
	fmt.Print(fileData)
}
