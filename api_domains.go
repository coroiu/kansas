package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func listDomains(c echo.Context) error {
	return c.JSON(200, configuration.Domains)
}

func createDomain(c echo.Context) error {
	configuration.Domains = append(configuration.Domains, DomainConfiguration{
		uuid.New().String(),
		"",
		"",
		1000,
		true,
	})
	saveConfiguration()
	return c.NoContent(204)
}

// func findDomain(c echo.Context) error {
// }

func updateDomain(c echo.Context) error {
	return c.NoContent(204)
}

func deleteDomain(c echo.Context) error {
	return c.NoContent(204)
}

func setupDomainAPIEndpoints(e *echo.Echo) {
	e.GET("/domains", listDomains)
	e.POST("/domains", createDomain)
	// e.GET("/domains/:id", findDomain)
	e.PUT("/domains/:id", updateDomain)
	e.DELETE("/domains/:id", deleteDomain)
}
