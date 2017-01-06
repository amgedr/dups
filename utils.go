package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

//walkfn is called by filepath.Walk for each file and dir
func walkfn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if excl.Exists(path) {
		return nil
	}

	if !info.IsDir() {
		if _, ok := bysize[info.Size()]; ok {
			bysize[info.Size()] = append(bysize[info.Size()], path)
		} else {
			bysize[info.Size()] = []string{path}
		}
	}

	return nil
}

//fileHash computes the MD5 hash of a file
func fileHash(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := md5.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
