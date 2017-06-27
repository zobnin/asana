package commands

import (
	"fmt"

	"github.com/codegangsta/cli"

	"asana/api"
	"asana/cache"
)

func Users(c *cli.Context) {
	if c.Bool("no-cache") {
		usersFromAPI(false)
	} else {
		if cache.IsRefreshNeeded(cache.UsersFile()) || c.Bool("refresh") {
			usersFromAPI(true)
		} else {
			entries := cache.GetUsers()
			if entries != nil {
				for _, e := range entries {
					fmt.Printf("%2s %16s %s\n", e.Index, e.Id, e.Line)
				}
			} else {
				usersFromAPI(true)
			}
		}
	}
}

func usersFromAPI(saveCache bool) {
	users := api.Users()
	if saveCache {
		cache.SaveUsers(users)
	}
	for i, u := range users {
		fmt.Printf("%2d %16d %s\n", i, u.Id, u.Name)
	}
}
