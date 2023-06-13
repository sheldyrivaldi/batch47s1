package controllers

import (
	"net/http"
	"stage1/utilities"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetProjectDetailController(c echo.Context) error {
	// Menangkap Id dari Query Params
	id := c.Param("id")

	// Membuat Struct berdasarkan Id dari Query Params
	data, err := utilities.FindOneProject(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data.StartDate = data.StartDateTime.Format("01 Jan 2006")
	data.EndDate = data.EndDateTime.Format("01 Jan 2006")

	projects := map[string]interface{}{
		"Projects": data,
	}

	tmpl, err := template.ParseFiles("views/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}
