package api

import (
	"encoding/json"

	"asana/utils"
)

type User_t struct {
	Id         int
	Name       string
	Email      string
	Photo      map[string]string
	Workspaces []Base
}

func Me() User_t {
	var me map[string]User_t

	err := json.Unmarshal(Get("/api/1.0/users/me", nil), &me)
	utils.Check(err)
	return me["data"]
}

func User(userId string) User_t {
	var user map[string]User_t

	err := json.Unmarshal(Get("/api/1.0/users/"+userId, nil), &user)
	utils.Check(err)
	return user["data"]
}

func Users() []User_t {
    var users map[string][]User_t
    var usersA []User_t

    err := json.Unmarshal(Get("/api/1.0/users", nil), &users)
    utils.Check(err)

    for _, u := range users["data"] {
        usersA = append(usersA, u)
    }

    return usersA
}
