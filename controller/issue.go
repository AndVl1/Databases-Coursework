package controller

import (
	"context"
	"github.com/AndVl1/bugTrackerBackend/model"
	"github.com/AndVl1/bugTrackerBackend/storage"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func GetAllIssues(ctx echo.Context) error {
	bugs, _ := getRepoIssues()
	//if err != nil {
	//	return err
	//} TODO
	json, _ := bugs.MarshalJSON()

	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, json)
}

func GetIssuesForAssignee(ctx echo.Context) error {
	userId := ctx.Param("assignee")
	issues, _ := getRepoIssuesForAssignee(userId)
	json, _ := issues.MarshalJSON()

	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, json)
}

func GetIssuesForProject(ctx echo.Context) error {
	projectId := ctx.Param("id")
	issues, err := getRepoIssuesForProject(projectId)
	if err != nil {
		return ctx.String(http.StatusNotImplemented, err.Error())
	}
	json, _ := issues.MarshalJSON()

	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, json)
}

func GetUsersForProject(ctx echo.Context) error {
	projectId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	users, err := getRepoUsersForProject(projectId)
	if err != nil {
		log.Print(err)
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	json, _ := users.MarshalJSON()

	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, json)
}

func GetIssue(ctx echo.Context) error {
	id := ctx.FormValue("id")
	bug, _ := getRepoIssue(id)
	//if err != nil {
	//	return err
	//}
	json, _ := bug.MarshalJSON()
	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, json)
}

func AddIssue(ctx echo.Context) error {
	projectId := ctx.Param("id")
	issueJson := ctx.QueryParam("issue")
	log.Printf("%s", issueJson)
	issue := &model.Issue{}
	if err := issue.UnmarshalJSON([]byte(issueJson)); err != nil {
		return ctx.String(http.StatusNotImplemented, err.Error())
	}
	if err := insertIssue(issue, projectId); err != nil {
		return ctx.String(http.StatusNotImplemented, err.Error())
	}
	response, _ := issue.MarshalJSON()
	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, response)
}

func UpdateIssue(ctx echo.Context) error {
	issueJson := ctx.QueryParam("issue")
	newStatus, _ := strconv.Atoi(ctx.QueryParam("status"))
	issue := &model.Issue{}
	if err := issue.UnmarshalJSON([]byte(issueJson)); err != nil {
		return ctx.String(http.StatusNotImplemented, err.Error())
	}
	issue.StatusId = newStatus
	if err := updateIssue(issue); err != nil {
		log.Println(err.Error())
		return ctx.String(http.StatusNotImplemented, err.Error())
	}
	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, []byte(issueJson))
}

func updateIssue(issue *model.Issue) error {
	pool := storage.GetDBInstance()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return err
	}
	defer conn.Release()
	_ = conn.QueryRow(context.Background(),
		"UPDATE Issue SET statusId=$1 WHERE issueId=$2", issue.StatusId, issue.Id)

	return nil
}

func insertIssue(issue *model.Issue, projectId string) error {
	pool := storage.GetDBInstance()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return err
	}
	defer conn.Release()
	projectRow := conn.QueryRow(context.Background(),
		"SELECT issuesCount FROM Project WHERE projectId=$1",
		projectId)
	var issuesCount int
	log.Println("ISSUE", issue)
	if err = projectRow.Scan(&issuesCount); err != nil {
		return err
	}
	issuesCount++
	if issue.AssigneeId.Valid {
		row := conn.QueryRow(context.Background(),
			"INSERT INTO Issue ("+
				"name, description, statusId, authorId, projectIssueNumber, "+
				"labelId, releaseVersion, creationDate, deadline,"+
				"projectId, assigneeId"+
				") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING issueId",
			issue.Name, issue.Description, issue.StatusId, issue.AuthorId,
			issuesCount, issue.LabelId, issue.ReleaseVersion, issue.CreationDate,
			issue.Deadline, issue.ProjectId, issue.AssigneeId.Int32)
		var id int
		err = row.Scan(&id)
		issue.Id = id
		if err != nil {
			log.Printf("Unable to INSERT: %v\n", err)
			return err
		}
		row, err = conn.Query(context.Background(),
			"UPDATE Project SET issuesCount=$1 WHERE projectId=$2",
			issuesCount,
			projectId)
		if err != nil {
			log.Printf("Unable to UPDATE: %v\n", err)
			return err
		}
	} else {
		row := conn.QueryRow(context.Background(),
			"INSERT INTO Issue ("+
				"name, description, statusId, authorId, projectIssueNumber, "+
				"labelId, releaseVersion, creationDate, deadline,"+
				"projectId, assigneeId"+
				") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING issueId",
			issue.Name, issue.Description, issue.StatusId, issue.AuthorId,
			issuesCount, issue.LabelId, issue.ReleaseVersion, issue.CreationDate,
			issue.Deadline, issue.ProjectId)
		var id int
		err = row.Scan(&id)
		issue.Id = id
		if err != nil {
			log.Printf("Unable to INSERT: %v\n", err)
			return err
		}
		row, err = conn.Query(context.Background(),
			"UPDATE Project SET issuesCount=$1 WHERE projectId=$2",
			issuesCount,
			projectId)
		if err != nil {
			log.Printf("Unable to UPDATE: %v\n", err)
			return err
		}
	}
	return nil
}

/*
   issueId             serial      NOT NULL,
   name                text        NOT NULL,
   projectIssueNumber  int         NOT NULL,
   description         text        NOT NULL,
   releaseVersion      int         NOT NULL,
   creationDate        date        NOT NULL,
   deadline            date        NOT NULL,
   assigneeId          int         NOT NULL,
   authorId            int         NOT NULL,
   projectId           int         NOT NULL,
   statusId            int         NOT NULL,
   -- in 'new', 'in progress', 'review', 'testing', 'ready', 'closed'
   labelId             int         NOT NULL,
*/

func getRepoIssuesForProject(projectId string) (model.Issues, error) {
	pool := storage.GetDBInstance()
	var issues model.Issues
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "SELECT issueId,"+
		"name, projectIssueNumber, description, releaseVersion, creationDate, deadline,"+
		"assigneeId, authorId, projectId, statusId, labelId FROM Issue WHERE projectId=$1", projectId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var issue model.Issue
		err = rows.Scan(
			&issue.Id,
			&issue.Name,
			&issue.ProjectIssueNumber,
			&issue.Description,
			&issue.ReleaseVersion,
			&issue.CreationDate,
			&issue.Deadline,
			&issue.AssigneeId,
			&issue.AuthorId,
			&issue.ProjectId,
			&issue.StatusId,
			&issue.LabelId)
		if err != nil {
			return nil, err
		}
		log.Println(issue)
		issues = append(issues, &issue)
	}
	return issues, nil
}

func getRepoUsersForProject(projectId int) (model.Users, error) {
	pool := storage.GetDBInstance()
	var users model.Users
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "SELECT userId, login, name "+
		"FROM ProjectUsersView WHERE projectId=$1", projectId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user model.User
		err = rows.Scan(
			&user.Id,
			&user.Login,
			&user.Name)
		if err != nil {
			return nil, err
		}
		log.Println(user)
		users = append(users, &user)
	}
	return users, nil
}

func getRepoIssuesForAssignee(userId string) (model.Issues, error) {
	pool := storage.GetDBInstance()
	var issues model.Issues
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return nil, err
	}
	defer conn.Release()
	rows, _ := conn.Query(context.Background(), "SELECT issueId,"+
		"name, projectIssueNumber, description, releaseVersion, creationDate, deadline,"+
		"assigneeId, authorId, projectId, statusId, labelId FROM Issue "+
		"WHERE assigneeId=$1", userId)
	for rows.Next() {
		var issue model.Issue
		_ = rows.Scan(
			&issue.Id,
			&issue.Name,
			&issue.ProjectIssueNumber,
			&issue.Description,
			&issue.ReleaseVersion,
			&issue.CreationDate,
			&issue.Deadline,
			&issue.AssigneeId,
			&issue.AuthorId,
			&issue.ProjectId,
			&issue.StatusId,
			&issue.LabelId)
		issues = append(issues, &issue)
	}
	return issues, nil
}

func getRepoIssue(id string) (model.Issue, error) {
	pool := storage.GetDBInstance()
	var bug model.Issue
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return model.Issue{}, err
	}
	defer conn.Release()
	row := conn.QueryRow(context.Background(), "SELECT * FROM Issue WHERE bugId=$1", id)
	_ = row.Scan(&bug.Id, &bug.Name, &bug.Description, &bug.StatusId, &bug.AuthorId)

	return bug, nil
}

func getRepoIssues() (model.Issues, error) {
	pool := storage.GetDBInstance()
	var bugs model.Issues
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return nil, err
	}
	defer conn.Release()
	rows, _ := conn.Query(context.Background(), `SELECT * FROM Issues`)
	for rows.Next() {
		var bug model.Issue
		_ = rows.Scan(&bug.Id, &bug.Name, &bug.Description, &bug.StatusId)
		log.Println(bug.Name)
		bugs = append(bugs, &bug)
	}

	return bugs, nil
}
