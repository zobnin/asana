package commands

import (
	"fmt"
	"regexp"

	"github.com/codegangsta/cli"

	"asana/api"
	"asana/cache"
)

func Task(c *cli.Context) {
	t, stories := api.Task(cache.FindId("task", c.Args().First(), true), c.Bool("verbose"))

	fmt.Printf("[ %s ] %s\n", t.Due_on, t.Name)

	showTags(t.Tags)

	fmt.Printf("\n%s\n", t.Notes)

	if stories != nil {
		fmt.Println("\n----------------------------------------\n")
		for _, s := range stories {
			s.Text = findAndReplaceUser(s.Text)
			fmt.Printf("%s\n", s)
		}
	}
}

func findAndReplaceUser(s string) string {
	userUrlRegexp, _ := regexp.Compile(`https://app.asana.com/0/\d+/\d+`)
	userIdRegexp, _ := regexp.Compile(`\d+$`)

	userUrl := userUrlRegexp.FindString(s)
	if userUrl != "" {
		userId := userIdRegexp.FindString(userUrl)

		// FIXME Workoroud: userId don't match real user Id
		ns := userUrlRegexp.ReplaceAllString(s, "To user: "+userId+"")
		return ns
	}
	return s
}

func showTags(tags []api.Base) {
	if len(tags) > 0 {
		fmt.Print("  Tags: ")
		for i, tag := range tags {
			print(tag.Name)
			if len(tags) != 1 && i != (len(tags)-1) {
				print(", ")
			}
		}
		println("")
	}
}
