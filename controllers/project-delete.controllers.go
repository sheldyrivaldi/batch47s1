package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func DeleteProjectController (c echo.Context) error{

	// Menangkap Id dari Query Params
	id, _:= strconv.Atoi(c.Param("id"))
	
	DataProjects = append(DataProjects[:id], DataProjects[id+1:]...)
	fmt.Printf("Project with id: %d, Successfully Deleted!\n", id)

	return c.Redirect(http.StatusMovedPermanently, "/")
}