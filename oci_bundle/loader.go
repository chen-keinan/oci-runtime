package oci_bundle

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func LoadBundle(pathToBundle string) ([]string, error) {
	filesData := make([]string, 0)
	f, err := os.Open(pathToBundle)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Print(err.Error())
		}
	}()

	gzReader, err := GetReader(f)
	if err != nil {
		return nil, err
	}
	rc := tar.NewReader(gzReader)
	for {
		archiveEntry, err := rc.Next()
		if err == io.EOF {
			break
		}
		if archiveEntry.Name == "config.json" {
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, err
			}
			filesData = append(filesData, string(data))
		}
	}
	return filesData, nil
}

//GetReader return gzip reader
//accept io.reader
func GetReader(reader io.Reader) (io.ReadCloser, error) {
	fileReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	return fileReader, nil
}
