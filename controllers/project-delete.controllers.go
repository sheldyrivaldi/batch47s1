package controllers

import (
	"fmt"
	"net/http"
	"os"
	"stage1/utilities"
	"strconv"

	"github.com/labstack/echo/v4"
)

func DeleteProjectController(c echo.Context) error {

	// Menangkap Id dari Query Params
	id, _ := strconv.Atoi(c.Param("id"))

	//Data Sebelumnya
	data, errDB := utilities.FindOneProject(id)
	if errDB != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errDB.Error()})
	}

	// Menghapus data dalam database
	err := utilities.DeleteProject(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	} else {
		err := os.Remove(fmt.Sprintf("public/images/uploads/%s", data.Image))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		} else {
			fmt.Println("Project successfully deleted!")
		}
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}
