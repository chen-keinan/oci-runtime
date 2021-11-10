package oci_bundle

type Config struct {
	Version string `json:"ociVersion"`
	Root *Root `json:"root,omitempty"`
}

type Root struct {
	// Path is the absolute path to the container's root filesystem.
	Path string `json:"path"`
	// Readonly makes the root filesystem for the container readonly before the process is executed.
	Readonly bool `json:"readonly,omitempty"`
}

