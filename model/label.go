package model

//easyjson:json
type Label struct {
	LabelId   int `json:"labelId"`
	LabelName int `json:"labelName"`
}

//easyjson:json
type Labels []*Label
