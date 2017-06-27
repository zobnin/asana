package commands

import (
	"github.com/codegangsta/cli"

	"asana/api"
	"asana/cache"
)

func Done(c *cli.Context) {
	api.Update(cache.FindId("task", c.Args().First(), false), "completed", "true")
}
