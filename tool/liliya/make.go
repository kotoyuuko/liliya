package main

import (
	"io/ioutil"
	"path"
)

// File stores config of source file
type File struct {
	Type string
	Name string
}

var file File

func createFileFromTemplate(tpl Template) error {
	filePath := path.Join(project.Path, tpl.Path)

	return ioutil.WriteFile(filePath, []byte(tpl.Content), 0644)
}

func makeFile() error {
	return nil
}
