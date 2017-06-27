package commands

import (
	"fmt"

	"github.com/codegangsta/cli"

	"asana/api"
	"asana/cache"
)

func Assignee(c *cli.Context) {
	taskId := cache.FindId("task", c.Args().First(), false)
	userId := cache.FindId("user", c.Args()[1], false)

	task := api.Update(taskId, "assignee", userId)

	fmt.Println("DONE! : " + task.Name)
}
