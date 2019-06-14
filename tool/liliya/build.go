package main

import (
	"os"
	"os/exec"
	"path"
)

func buildApp() error {
	buildPath := path.Join(project.Path, "build")

	if err := os.MkdirAll(buildPath, 0755); err != nil {
		return err
	}

	mainPath := path.Join(project.Path, "src/main.go")
	outputPath := path.Join(buildPath, "app")

	cmd := exec.Command("go", "build", "-o", outputPath, mainPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
