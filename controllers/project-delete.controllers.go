package controllers

import (
	"fmt"
	"net/http"
	"stage1/utilities"

	"github.com/labstack/echo/v4"
)

func DeleteProjectController(c echo.Context) error {

	// Menangkap Id dari Query Params
	id := c.Param("id")

	// Menghapus data dalam database
	err := utilities.DeleteProject(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	} else {
		fmt.Println("Project successfully deleted!")
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}
