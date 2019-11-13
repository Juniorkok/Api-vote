package model

type User struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"-"`
}

var listUser = map[string]User{}