package db

import ""

type User struct {
	ID string `json:id`
	FirstName string `json:first_name`
	LlastName string `json:last_name`
	Email string `json:email`
	Password string `json:password`
}

var listUser = map[string]User{}