package model

import (
	"database/sql"
	"encoding/json"
	"log"
	"reflect"
)

type NullInt32 sql.NullInt32

//easyjson:json
type Issue struct {
	Id                 int       `json:"issueId"`
	Name               string    `json:"name"`
	ProjectIssueNumber int       `json:"projectIssueNumber"`
	Description        string    `json:"description"`
	AuthorId           int       `json:"authorId"`
	AssigneeId         NullInt32 `json:"assigneeId"`
	ReleaseVersion     string    `json:"releaseVersion"`
	CreationDate       int64     `json:"creationDate"`
	Deadline           int64     `json:"deadline"`
	ProjectId          int       `json:"projectId"`
	StatusId           int       `json:"statusId"`
	LabelId            int       `json:"labelId"`
}

//easyjson:json
type Issues []*Issue

func (ni *NullInt32) Scan(value interface{}) error {
	var i sql.NullInt32
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt32{i.Int32, false}
	} else {
		*ni = NullInt32{i.Int32, true}
	}
	return nil
}

// MarshalJSON for NullInt32
// ni -> json
func (ni *NullInt32) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int32)
}

// UnmarshalJSON for NullInt32
// json -> ni
func (ni *NullInt32) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int32)
	ni.Valid = err == nil
	log.Println("UNMARSHALL", ni, err)
	return err
}
