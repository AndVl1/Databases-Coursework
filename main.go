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
	//e.GET("/students", controller.GetStudents)
	e.GET("/users", controller.GetAllUsers)
	e.GET("/users/:id", controller.GetUser)
	e.GET("/projects", controller.GetProjectsForUser)
	// POST
	e.POST("/login", controller.LoginUser)
	e.POST("/check", controller.CheckUser)
	e.POST("/register", controller.RegisterUser)
	e.POST("/projects/add", controller.AddProject)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
