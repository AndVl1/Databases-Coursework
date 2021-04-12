package model

//easyjson:json
type Status struct {
	StatusId   int `json:"labelId"`
	StatusName int `json:"labelName"`
}

//easyjson:json
type Statuses []*Label
