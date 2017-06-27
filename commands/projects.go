package commands

import (
	"fmt"
	"net/url"

	"github.com/codegangsta/cli"

	"asana/api"
)

func Projects(c *cli.Context) {
	projects := api.Projects(url.Values{}, false)

	for i, p := range projects {
		fmt.Printf("%2d [ %10s ] %s\n", i, p.Due_date, p.Name)
	}
}
