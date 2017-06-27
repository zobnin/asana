package commands

import (
	"os/exec"
	"strconv"

	"github.com/codegangsta/cli"

	"asana/cache"
	"asana/config"
	"asana/utils"
)

func Browse(c *cli.Context) {
	taskId := cache.FindId("task", c.Args().First(), true)
	url := "https://app.asana.com/0/" + strconv.Itoa(config.Load().Workspace) + "/" + taskId
	launcher, err := utils.BrowserLauncher()
	utils.Check(err)
	cmd := exec.Command(launcher, url)
	err = cmd.Start()
	utils.Check(err)
}
