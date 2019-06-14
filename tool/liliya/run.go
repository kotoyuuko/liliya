package main

import (
	"errors"
	"os"
	"path"

	"github.com/urfave/cli"
)

func runCreate(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		return errors.New("required project name")
	}

	project.Name = ctx.Args()[0]
	pwd, _ := os.Getwd()
	project.Path = path.Join(pwd, project.Name)

	if err := createProject(); err != nil {
		return err
	}

	return nil
}

func runMake(ctx *cli.Context) error {
	if len(ctx.Args()) != 2 {
		return errors.New("required file type and name")
	}

	project.Path, _ = os.Getwd()

	file.Type = ctx.Args()[0]
	file.Name = ctx.Args()[1]

	if err := makeFile(); err != nil {
		return err
	}

	return nil
}

func runStart(ctx *cli.Context) error {
	project.Path, _ = os.Getwd()

	return startApp()
}

func runBuild(ctx *cli.Context) error {
	project.Path, _ = os.Getwd()

	return buildApp()
}
