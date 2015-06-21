package main

import (
	"fmt"
	dc "github.com/fsouza/go-dockerclient"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type (
	// Template provides HTML template rendering
	Template struct {
		templates *template.Template
	}

	// user struct {
	// 	ID   string `json:"id"`
	// 	Name string `json:"name"`
	// }
)

// Template provides HTML template rendering
func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Handler
func createDeployment(c *echo.Context) error {
	return c.String(http.StatusOK, "Deployment POST\n")
}

func getDeployment(c *echo.Context) error {

	// Init the client
	path := os.Getenv("DOCKER_CERT_PATH")
	ca := fmt.Sprintf("%s/ca.pem", path)
	cert := fmt.Sprintf("%s/cert.pem", path)
	key := fmt.Sprintf("%s/key.pem", path)

	docker, _ := dc.NewTLSClient(os.Getenv("DOCKER_HOST"), cert, key, ca)

	// Get only running containers
	containers, err := docker.ListContainers(dc.ListContainersOptions{All: false})
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, containers)
}

func updateDeployment(c *echo.Context) error {
	return c.String(http.StatusOK, "Deployment PATCH\n")
}

func deleteDeployment(c *echo.Context) error {
	return c.String(http.StatusOK, "Deployment DELETE\n")
}

func main() {

	// Check for required variables
	docker_host := os.Getenv("DOCKER_HOST")
	docker_cert_path := os.Getenv("DOCKER_CERT_PATH")

	if len(strings.TrimSpace(docker_host)) == 0 {
	  panic("Please set DOCKER_HOST!")
	}

	if len(strings.TrimSpace(docker_cert_path)) == 0 {
		panic("Please set DOCKER_CERT_PATH")
	}

	// Echo instance
	e := echo.New()

	// Debug mode
	e.SetDebug(true)

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Index("public/index.html")

	// Deployment Routes
	e.Post("/v1/deployments", createDeployment)
	e.Get("/v1/deployments", getDeployment)
	e.Get("/v1/deployments/:id", getDeployment)
	e.Patch("/v1/deployments/:id", updateDeployment)
	e.Delete("/v1/deployments/:id", deleteDeployment)

	// Serve Index
	e.Index("public/index.html")

	// Serve Static Assets
	e.Static("/assets", "public/assets/")

	// Start server
	e.Run(":1323")
}
