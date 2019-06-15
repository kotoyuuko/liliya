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

	if err := createModelPath(); err != nil {
		return err
	}

	if err := createRouterPath(); err != nil {
		return err
	}

	if err := createServicePath(); err != nil {
		return err
	}

	if err := createUtilPath(); err != nil {
		return err
	}

	if err := createFileFromTemplate(tplMain); err != nil {
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

func createModelPath() error {
	modelPath := path.Join(project.Path, "src/model")

	if err := os.MkdirAll(modelPath, 0755); err != nil {
		return err
	}

	if err := createGitkeep(modelPath); err != nil {
		return err
	}

	if err := createFileFromTemplate(tplModelUser); err != nil {
		return err
	}

	return nil
}

func createRouterPath() error {
	routerPath := path.Join(project.Path, "src/router")

	if err := os.MkdirAll(routerPath, 0755); err != nil {
		return err
	}

	if err := createGitkeep(routerPath); err != nil {
		return err
	}

	if err := createFileFromTemplate(tplRouterRoutes); err != nil {
		return err
	}

	return nil
}

func createServicePath() error {
	servicePath := path.Join(project.Path, "src/service")

	if err := os.MkdirAll(servicePath, 0755); err != nil {
		return err
	}

	if err := createGitkeep(servicePath); err != nil {
		return err
	}

	if err := createFileFromTemplate(tplServiceRoot); err != nil {
		return err
	}

	return nil
}

func createUtilPath() error {
	utilPath := path.Join(project.Path, "src/util")

	if err := os.MkdirAll(utilPath, 0755); err != nil {
		return err
	}

	if err := createGitkeep(utilPath); err != nil {
		return err
	}

	if err := createUtilConfigPath(); err != nil {
		return err
	}

	if err := createUtilDAOPath(); err != nil {
		return err
	}

	return nil
}

func createUtilConfigPath() error {
	utilConfigPath := path.Join(project.Path, "src/util/config")

	if err := os.MkdirAll(utilConfigPath, 0755); err != nil {
		return err
	}

	if err := createGitkeep(utilConfigPath); err != nil {
		return err
	}

	if err := createFileFromTemplate(tplUtilConfig); err != nil {
		return err
	}

	return nil
}

func createUtilDAOPath() error {
	utilDAOPath := path.Join(project.Path, "src/util/dao")

	if err := os.MkdirAll(utilDAOPath, 0755); err != nil {
		return err
	}

	if err := createGitkeep(utilDAOPath); err != nil {
		return err
	}

	if err := createFileFromTemplate(tplUtilDAO); err != nil {
		return err
	}

	return nil
}
