package commands

import (
	"fmt"
	"github.com/codegangsta/cli"

    "asana/api"
)

func Workspaces(c *cli.Context) {
	for _, w := range api.Me().Workspaces {
		fmt.Printf("%16d %s\n", w.Id, w.Name)
	}
}
