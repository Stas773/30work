package models

type User struct {
	Age     int      `json:"age"`
	Name    string   `json:"name"`
	Id      string   `json:"id"`
	Friends []string `json:"friends"`
}
