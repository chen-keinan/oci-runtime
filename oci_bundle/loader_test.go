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

func TestReadFile(t *testing.T) {
	tests := []struct {
		name     string
		relPath  string
		wantData string
		wantErr  error
	}{
		{name: "readFile good path", relPath: "./fixture/config.json", wantData: "{\n  \"ociVersion\": \"1.0.0\",\n  \"root\": {\n    \"path\": \"path\",\n    \"redaonly\": true\n  }\n}\n", wantErr: nil},
		{name: "readFile bad path", relPath: "./fixture/config.json1", wantData: "", wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fp, err := filepath.Abs(tt.relPath)
			if err != nil {
				t.Error(err)
			}
			gotData, gotErr := ReadFile(fp)
			if tt.wantData != gotData {
				t.Errorf("TestReadFile(),not expected value")
			}
			if gotErr != nil && tt.wantErr != nil && tt.wantErr.Error() != gotErr.Error() {
				t.Errorf("TestReadFile(),not expected value")
			}
		})
	}
}
