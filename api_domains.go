package main

import (
	"fmt"

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
	fmt.Println(c.Param("id"))
	return c.NoContent(204)
}

func deleteDomain(c echo.Context) error {
	i := configuration.Domains.findIndex(c.Param("id"))
	if i >= 0 {
		configuration.Domains = append(configuration.Domains[:i], configuration.Domains[i+1:]...)
		return c.NoContent(204)
	}
	return c.NoContent(404)
}

func setupDomainAPIEndpoints(e *echo.Echo) {
	e.GET("/domains", listDomains)
	e.POST("/domains", createDomain)
	// e.GET("/domains/:id", findDomain)
	e.PUT("/domains/:id", updateDomain)
	e.DELETE("/domains/:id", deleteDomain)
}
