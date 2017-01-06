package main

import (
	"path/filepath"
	"strings"
)

type Extensions []string

func (extensions *Extensions) String() string {
	return strings.Join(*extensions, ", ")
}

func (extensions *Extensions) Set(val string) error {
	*extensions = append(*extensions, val)
	return nil
}

//Exists checks if the extension of the file exists in the slice
func (extensions Extensions) Exists(filename string) bool {
	var ext string
	if len(filepath.Ext(filename)) > 1 {
		//Ext returns the extension starting with a period
		ext = filepath.Ext(filename)[1:]
	} else {
		return false
	}

	exists := false
	for _, e := range extensions {
		if e == ext {
			exists = true
			break
		}
	}

	return exists
}
