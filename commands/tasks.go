package commands

import (
	"fmt"
	"net/url"

	"github.com/codegangsta/cli"

	"asana/api"
	"asana/cache"
)

func Tasks(c *cli.Context) {
	if c.Bool("no-cache") {
		fromAPI(false)
	} else {
		if cache.IsRefreshNeeded(cache.TasksFile()) || c.Bool("refresh") {
			fromAPI(true)
		} else {
			entries := cache.GetTasks()
			if entries != nil {
				for _, e := range entries {
					fmt.Printf("%2s [ %10s ] %s\n", e.Index, e.Date, e.Line)
				}
			} else {
				fromAPI(true)
			}
		}
	}
}

func fromAPI(saveCache bool) {
	tasks := api.Tasks(url.Values{}, false)
	if saveCache {
		cache.SaveTasks(tasks)
	}
	for i, t := range tasks {
		fmt.Printf("%2d [ %10s ] %s\n", i, t.Due_on, t.Name)
	}
}
