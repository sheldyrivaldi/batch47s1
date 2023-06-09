package controllers

import (
	"fmt"
	"net/http"
	"stage1/utilities"

	"github.com/labstack/echo/v4"
)

func DeleteProjectController (c echo.Context) error{

	// Menangkap Id dari Query Params
	id := c.Param("id")

	// Menghapus data dalam database
	utilities.DeleteProject(id)

	fmt.Println("Project successfully deleted!")

	return c.Redirect(http.StatusMovedPermanently, "/")
}