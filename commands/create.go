package commands

import (
	"github.com/codegangsta/cli"

	"asana/api"
)

func Create(c *cli.Context) {
	api.CreateTask(c.Args().First())
}
