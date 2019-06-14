package main

import (
	"os"
	"os/exec"
	"path"
)

func startApp() error {
	mainPath := path.Join(project.Path, "src/main.go")

	cmd := exec.Command("go", "run", mainPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
