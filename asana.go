package main

import (
	"os"

	"github.com/codegangsta/cli"

	"asana/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "asana"
	app.Version = "0.2.0"
	app.Usage = "asana cui client ( https://github.com/memerelics/asana )"

	app.Commands = defs()
	app.Run(os.Args)
}

func defs() []cli.Command {
	return []cli.Command{
		{
			Name:  "config",
			Usage: "Asana configuration. Your settings will be saved in ~/.asana/config.yml",
			Action: func(c *cli.Context) {
				commands.Config(c)
			},
		},
		{
			Name:      "workspaces",
			ShortName: "w",
			Usage:     "Get workspaces",
			Action: func(c *cli.Context) {
				commands.Workspaces(c)
			},
		},
		{
			Name:      "projects",
			ShortName: "ps",
			Usage:     "Get projects",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "no-cache, n", Usage: "without cache"},
				cli.BoolFlag{Name: "refresh, r", Usage: "update cache"},
			},
			Action: func(c *cli.Context) {
				commands.Projects(c)
			},
		},
		{
			Name:      "users",
			ShortName: "u",
			Usage:     "Get users",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "no-cache, n", Usage: "without cache"},
				cli.BoolFlag{Name: "refresh, r", Usage: "update cache"},
			},
			Action: func(c *cli.Context) {
				commands.Users(c)
			},
		},
		{
			Name:      "tasks",
			ShortName: "ts",
			Usage:     "Get tasks",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "no-cache, n", Usage: "without cache"},
				cli.BoolFlag{Name: "refresh, r", Usage: "update cache"},
			},
			Action: func(c *cli.Context) {
				commands.Tasks(c)
			},
		},
		{
			Name:      "task",
			ShortName: "t",
			Usage:     "Get a task",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "verbose, v", Usage: "verbose output"},
			},
			Action: func(c *cli.Context) {
				commands.Task(c)
			},
		},
		{
			Name:      "create",
			ShortName: "c",
			Usage:     "Create a task",
			Action: func(c *cli.Context) {
				commands.Create(c)
			},
		},
		{
			Name:      "comment",
			ShortName: "cm",
			Usage:     "Post comment",
			Action: func(c *cli.Context) {
				commands.Comment(c)
			},
		},
		{
			Name:      "attach",
			ShortName: "a",
			Usage:     "Attach file",
			Action: func(c *cli.Context) {
				commands.Attach(c)
			},
		},
		{
			Name:      "assignee",
			ShortName: "as",
			Usage:     "Assignee task to user",
			Action: func(c *cli.Context) {
				commands.Assignee(c)
			},
		},
		{
			Name:  "done",
			Usage: "Complete task",
			Action: func(c *cli.Context) {
				commands.Done(c)
			},
		},
		{
			Name:  "due",
			Usage: "Set due date",
			Action: func(c *cli.Context) {
				commands.DueOn(c)
			},
		},
		{
			Name:  "delete",
			Usage: "Delete task",
			Action: func(c *cli.Context) {
				commands.Delete(c)
			},
		},
		{
			Name:      "browse",
			ShortName: "b",
			Usage:     "Open a task in the web browser",
			Action: func(c *cli.Context) {
				commands.Browse(c)
			},
		},
	}
}
