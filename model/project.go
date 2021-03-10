package model

//easyjson:json
type Project struct {
	Id                 uint64 `json:"projectId"`
	ProjectName        string `json:"projectName"`
	ProjectDescription string `json:"projectDescription"`
}

//easyjson:json
type Projects []*Project
