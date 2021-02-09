package model

//easyjson:json
type User struct {
	Id       int    `json:"userId"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

//easyjson:json
type Users []*User
