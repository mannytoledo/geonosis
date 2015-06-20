package main

import (
  "net/http"

  "github.com/labstack/echo"
  mw "github.com/labstack/echo/middleware"
)

// Handler
func index(c *echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!\n")
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

  // Routes
  e.Get("/", index)
  e.Get("/v1/deployment", deploymentGet)
  e.Post("/v1/deployment", deploymentPost)

  // Start server
  e.Run(":1323")
}
