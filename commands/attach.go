package commands

import (
	"fmt"

	"github.com/codegangsta/cli"

	"asana/api"
	"asana/cache"
)

func Attach(c *cli.Context) {
	taskId := cache.FindId("task", c.Args().First(), true)
	api.Attach(taskId, c.Args()[1])
	fmt.Println("File uploaded")
}
