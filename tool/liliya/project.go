package main

import (
	"os"
	"path"
)

// Project stores config of project
type Project struct {
	Name string
	Path string
}

var project Project

func createProject() error {
	srcPath := path.Join(project.Path, "src")

	if err := os.MkdirAll(srcPath, 0755); err != nil {
		return err
	}

	if err := createConfigPath(); err != nil {
		return err
	}

	return nil
}

func createConfigPath() error {
	configPath := path.Join(project.Path, "src/config")

	if err := os.MkdirAll(configPath, 0755); err != nil {
		return err
	}

	if err := createFileFromTemplate(tplConfigAppIni); err != nil {
		return err
	}

	return nil
}
