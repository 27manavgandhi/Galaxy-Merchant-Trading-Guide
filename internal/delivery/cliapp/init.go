package cliapp

import (
	"io"

	"github.com/27manavgandhi/galaxy-merchant-trading/internal/model"
	"github.com/27manavgandhi/galaxy-merchant-trading/internal/usecase"
	"github.com/urfave/cli/v2"
)

var cmd *CommandLine

type CommandLine struct {
	App     *cli.App
	Param   model.Param
	Usecase *usecase.Usecase
	Writer  io.Writer
}

func New(w io.Writer, usecase *usecase.Usecase) *CommandLine {
	if cmd != nil {
		cmd.Usecase = usecase
		cmd.Writer = w
		return cmd
	}

	cmd = &CommandLine{
		Usecase: usecase,
		Writer:  w,
	}

	app := &cli.App{
		Name:        model.AppName,
		Description: model.AppDescription,
		Usage:       model.AppUsage,
		Commands: []*cli.Command{
			{
				Name: "run",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "file",
						Aliases:     []string{"f"},
						Usage:       "Load input from `FILE`",
						Destination: &cmd.Param.File,
					},
					&cli.BoolFlag{
						Name:        "verbose",
						Aliases:     []string{"v"},
						Usage:       "Add error message on output",
						Destination: &cmd.Param.Verbose,
					},
					&cli.BoolFlag{
						Name:        "custom-unit",
						Aliases:     []string{"custom"},
						Usage:       "Remove validation to check unit other than: Silver, Gold, Iron",
						Destination: &cmd.Param.CustomUnit,
					},
				},
				Action: cmd.Run,
			},
		},
	}

	cmd.App = app
	return cmd
}
