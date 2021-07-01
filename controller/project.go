package controller

import (
	"context"
	"github.com/AndVl1/bugTrackerBackend/model"
	"github.com/AndVl1/bugTrackerBackend/storage"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetProjectsForUser(ctx echo.Context) error {
	projects, _ := getProjectsForUserRepo(ctx.FormValue("userId"))
	json, _ := projects.MarshalJSON()
	log.Println(string(json))
	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, json)
}

func AddProject(ctx echo.Context) error {
	project := &model.Project{
		Id:                 0,
		ProjectName:        ctx.FormValue("pname"),
		ProjectDescription: ctx.FormValue("pdesc"),
	}
	userId := ctx.FormValue("userId")
	id, err := insertProject(userId, project)
	if err != nil {
		log.Println(err)
		return err
	}
	project.Id = id
	var userJson, _ = project.MarshalJSON()
	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, userJson)
}

func AddUserToProject(ctx echo.Context) error {
	projectId := ctx.Param("id")
	userId := ctx.QueryParam("userid")
	if err := addUserToProject(userId, projectId); err != nil {
		return ctx.String(http.StatusBadRequest, "")
	}
	return ctx.String(http.StatusOK, "")
}

func addUserToProject(userId string, projectId string) error {
	pool := storage.GetDBInstance()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return err
	}
	defer conn.Release()
	row := conn.QueryRow(
		context.Background(), "INSERT INTO ProjectUser (userId, projectId) VALUES ($1, $2) RETURNING userId",
		userId, projectId)
	id := 1
	err = row.Scan(&id)
	if err != nil {
		log.Printf("Unable to INSERT: %v\n", err)
		return err
	}
	return nil
}

func insertProject(userId string, project *model.Project) (uint64, error) {
	pool := storage.GetDBInstance()
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return 0, err
	}
	defer conn.Release()
	projectRow := conn.QueryRow(
		context.Background(),
		"INSERT INTO Project(projectName, projectDescription) VALUES ($1, $2) RETURNING projectId",
		project.ProjectName, project.ProjectDescription)
	err = projectRow.Scan(&project.Id)
	if err != nil {
		log.Printf("AUnable to INSERT: %v\n", err)
		return 0, err
	}
	_, err = conn.Query(
		context.Background(),
		"INSERT INTO ProjectUser(projectId, userId) VALUES ($1, $2)",
		project.Id, userId)
	log.Println("Insert project")
	if err != nil {
		log.Printf("Unable to INSERT: %v\n", err)
		return 0, err
	}
	return project.Id, nil
}

func getProjectsForUserRepo(userId string) (model.Projects, error) {
	pool := storage.GetDBInstance()
	var projects model.Projects
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Unable to acquire a database connection: %v\n", err)
		return nil, err
	}
	defer conn.Release()
	rows, _ := conn.Query(context.Background(),
		"SELECT projectId, projectName, projectDescription, issuesCount "+
			"FROM ProjectUsersView WHERE userId=$1", userId)

	for rows.Next() {
		var project model.Project
		_ = rows.Scan(&project.Id, &project.ProjectName, &project.ProjectDescription, &project.IssuesCount)
		log.Println(project.ProjectName)
		if project.Id > 0 {
			projects = append(projects, &project)
		}
	}

	return projects, nil
}
