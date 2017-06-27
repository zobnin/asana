package commands

import (
	"github.com/codegangsta/cli"

	"asana/api"
	"asana/cache"
)

func Delete(c *cli.Context) {
	api.DeleteTask(cache.FindId("task", c.Args().First(), false))
}
