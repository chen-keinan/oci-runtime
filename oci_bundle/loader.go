package oci_bundle

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func LoadBundle(bundleName string) ([]string, error) {
	filesData := make([]string, 0)
	bf, err := GetBundleFolder()
	if err != nil {
		return nil, err
	}
	bandlePath := path.Join(bf, bundleName)
	f, err := os.Open(fmt.Sprintf("%s.tgz", bandlePath))
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

func ReadFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
func GetBundleFolder() (string, error) {
	chome := os.Getenv("CONTAINER_HOME")
	if len(chome) > 0 {
		return chome, nil
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	bundleFolder := path.Join(dir, "bundles")
	if _, err := os.Stat(bundleFolder); os.IsNotExist(err) {
		err := os.Mkdir(bundleFolder, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return bundleFolder, err
}
