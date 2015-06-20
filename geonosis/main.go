package main

import (
	"os"
	"log"
	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
)

	dc "github.com/fsouza/go-dockerclient"
	// Template provides HTML template rendering
	Template struct {
		templates *template.Template
	}

)

// Handler
func index(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func createDeployment(c *echo.Context) error {
	return c.String(http.StatusOK, "Deployment POST\n")
}

func getDeployment(c *echo.Context) error {

	var retval bytes.Buffer

	// Init the client
	path := os.Getenv("DOCKER_CERT_PATH")
	ca := fmt.Sprintf("%s/ca.pem", path)
	cert := fmt.Sprintf("%s/cert.pem", path)
	key := fmt.Sprintf("%s/key.pem", path)

	docker, _ := dc.NewTLSClient(os.Getenv("DOCKER_HOST"), cert, key, ca)

	// Get images
	imgs, _ := docker.ListImages(dc.ListImagesOptions{All: false})
	for _, img := range imgs {
	    fmt.Println("ID: ", img.ID)
	    fmt.Println("RepoTags: ", img.RepoTags)
	    fmt.Println("Created: ", img.Created)
	    fmt.Println("Size: ", img.Size)
	    fmt.Println("VirtualSize: ", img.VirtualSize)
	    fmt.Println("ParentId: ", img.ParentID)
	}

	// Get only running containers
	containers, err := docker.ListContainers(dc.ListContainersOptions{All: false})
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range containers {
		retval.WriteString(c.ID)
		retval.WriteString("|")
		// retval.WriteString(c.Names)
		// retval.WriteString("\n")
	}

	return c.String(http.StatusOK, retval.String())
}

func updateDeployment(c *echo.Context) error {
	return c.String(http.StatusOK, "Deployment PATCH\n")
}

func deleteDeployment(c *echo.Context) error {
	return c.String(http.StatusOK, "Deployment DELETE\n")
}

func main() {
	// Echo instance
	e := echo.New()

	// Debug mode
  e.Use(mw.Logger())
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

	// Start server
	e.Run(":1323")
}
