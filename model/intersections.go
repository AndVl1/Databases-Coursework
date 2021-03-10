package model

//easyjson:json
type ProjectUser struct {
	UserId    uint64 `json:"userId"`
	ProjectId uint64 `json:"projectId"`
}
