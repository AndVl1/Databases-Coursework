package model

//easyjson:json
type Issue struct {
	Id                 int    `json:"issueId"`
	Name               string `json:"name"`
	ProjectIssueNumber int    `json:"projectIssueNumber"`
	Description        string `json:"description"`
	AuthorId           int    `json:"authorId"`
	AssigneeId         int    `json:"assigneeId"`
	ReleaseVersion     string `json:"releaseVersion"`
	CreationDate       int64  `json:"creationDate"`
	Deadline           int64  `json:"deadline"`
	ProjectId          int    `json:"projectId"`
	StatusId           int    `json:"statusId"`
	LabelId            int    `json:"labelId"`
}

//easyjson:json
type Issues []*Issue
