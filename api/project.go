package api

import (
	"asana/config"
	"asana/utils"
	"encoding/json"
	"net/url"
	"strconv"
)

type Project_t struct {
	Name           string
	Id             int
	Owner          Base
	Current_status string
	Due_date       string
	Created_at     string
	Modified_at    string
	Archived       bool
	Public         bool
	Members        []Base
	Followers      []Base
	Color          string
	Notes          string
	Workspace      Base
	Team           Base
}

func Projects(params url.Values, withCompleted bool) []Project_t {
	var projects map[string][]Project_t
	var projectsA []Project_t

	params.Add("workspace", strconv.Itoa(config.Load().Workspace))
	params.Add("archived", "false")

	err := json.Unmarshal(Get("/api/1.0/projects", params), &projects)
	utils.Check(err)

	for _, p := range projects["data"] {
		projectsA = append(projectsA, p)
	}

	return projectsA
}

func Project(projectId string, verbose bool) Project_t {
	var project map[string]Project_t

	err := json.Unmarshal(Get("/api/1.0/projects/"+projectId, nil), project)
	utils.Check(err)

	return project["data"]
}
