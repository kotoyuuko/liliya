package main

import (
	"errors"
	"io/ioutil"
	"path"
	"strings"
)

// File stores config of source file
type File struct {
	Type string
	Name string
	Path string
}

var file File

func makeFile() error {
	if file.Type == "service" {
		file.Path = path.Join(project.Path, "src/service", file.Name+".go")
		return makeServiceFile()
	} else if file.Type == "model" {
		file.Path = path.Join(project.Path, "src/model", file.Name+".go")
		return makeModelFile()
	} else {
		return errors.New("unsupported file type")
	}
}

func makeServiceFile() error {
	content := strings.Replace(tplService, "github.com/kotoyuuko/liliya/src", replaceTargetPath(), -1)

	return ioutil.WriteFile(file.Path, []byte(content), 0644)
}

func makeModelFile() error {
	content := strings.Replace(tplModel, "github.com/kotoyuuko/liliya/src", replaceTargetPath(), -1)
	content = strings.Replace(content, "{model}", strings.Title(file.Name), -1)

	return ioutil.WriteFile(file.Path, []byte(content), 0644)
}
