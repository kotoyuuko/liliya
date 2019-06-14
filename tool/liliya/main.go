package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Liliya"
	app.Usage = "Liliya's code generator."
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "create a new project.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "name",
					Value:       "",
					Usage:       "project name",
					Destination: &project.Name,
				},
			},
			Action: runCreate,
		},
		{
			Name:  "make",
			Usage: "make source code file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "type",
					Value:       "",
					Usage:       "file type",
					Destination: &file.Type,
				},
				cli.StringFlag{
					Name:        "name",
					Value:       "",
					Usage:       "file name",
					Destination: &file.Name,
				},
			},
			Action: runMake,
		},
		{
			Name:   "start",
			Usage:  "start liliya application",
			Action: runStart,
		},
		{
			Name:   "build",
			Usage:  "build liliya application",
			Action: runBuild,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
