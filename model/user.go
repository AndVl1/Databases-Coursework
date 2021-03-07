package model

//easyjson:json
type User struct {
	Id       uint64 `json:"userId"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

//easyjson:json
type Users []*User
