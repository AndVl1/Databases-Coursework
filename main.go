package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/AndVl1/bugTrackerBackend/controller"
	"github.com/AndVl1/bugTrackerBackend/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	db := storage.NewDB()
	defer db.Close()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// GET
	e.GET("/", hello)

	e.GET("/users", controller.GetAllUsers)
	e.GET("/users/:id", controller.GetUser)
	e.GET("/projects", controller.GetProjectsForUser)
	e.GET("/issues/:assignee", controller.GetIssuesForAssignee)
	e.GET("/projects/:id/issues", controller.GetIssuesForProject)
	e.GET("/issues/:id/comments", controller.GetComments)
	e.GET("/projects/:id/users", controller.GetUsersForProject)
	// POST
	e.POST("/login", controller.LoginUser)
	e.POST("/check", controller.CheckUser)
	e.POST("/register", controller.RegisterUser)
	e.POST("/projects/add", controller.AddProject)
	e.POST("/projects/:id/issues/add", controller.AddIssue)
	e.POST("/issues/:id/comments/add", controller.AddComment)
	e.POST("/issues/:id/update", controller.UpdateIssue)

	e.POST("/projects/:id/adduser", controller.AddUserToProject)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "https://github.com/AndVl1/BugTracker-android")
}
