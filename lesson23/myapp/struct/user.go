package _struct

type Name struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type User struct {
	FullName Name   `json:"name"`
	Age      string `json:"age"`
	Hobby    string `json:"hobby"`
}
