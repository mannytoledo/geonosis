package main

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
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

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Handler
func index(c *echo.Context) error {
	return c.Render(http.StatusOK, "welcome", "Joe")
	// return c.String(http.StatusOK, "Hello, World!\n")
}

func deploymentGet(c *echo.Context) error {
	return c.String(http.StatusOK, "Deployment GET\n")
}

func deploymentPost(c *echo.Context) error {
	return c.String(http.StatusOK, "Deployment POST\n")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Index("public/index.html")

	// Routes
	e.Get("/v1/deployment", deploymentGet)
	e.Post("/v1/deployment", deploymentPost)

	// Start server
	e.Run(":1323")
}
