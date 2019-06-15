package main

import (
	"io/ioutil"
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

	if err := createGitignore(); err != nil {
		return err
	}

	if err := createConfigPath(); err != nil {
		return err
	}

	if err := createDatabasePath(); err != nil {
		return err
	}

	if err := createLogPath(); err != nil {
		return err
	}

	return nil
}

func createGitkeep(folder string) error {
	filePath := path.Join(folder, ".gitkeep")

	return ioutil.WriteFile(filePath, []byte(""), 0644)
}

func createGitignore() error {
	return createFileFromTemplate(tplGitignore)
}

func createConfigPath() error {
	configPath := path.Join(project.Path, "src/config")

	if err := os.MkdirAll(configPath, 0755); err != nil {
		return err
	}

	if err := createGitkeep(configPath); err != nil {
		return err
	}

	if err := createFileFromTemplate(tplConfigAppIni); err != nil {
		return err
	}

	return nil
}

func createDatabasePath() error {
	databasePath := path.Join(project.Path, "src/database")

	if err := os.MkdirAll(databasePath, 0755); err != nil {
		return err
	}

	if err := createGitkeep(databasePath); err != nil {
		return err
	}

	if err := createFileFromTemplate(tplDatabaseMigration); err != nil {
		return err
	}

	if err := createFileFromTemplate(tplDatabaseSeeder); err != nil {
		return err
	}

	return nil
}

func createLogPath() error {
	logPath := path.Join(project.Path, "src/log")

	if err := os.MkdirAll(logPath, 0755); err != nil {
		return err
	}

	if err := createGitkeep(logPath); err != nil {
		return err
	}

	return nil
}
