package controllers

import (
	"net/http"
	"stage1/models"
	"stage1/utilities"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetProjectDetailController(c echo.Context) error {
	// Menangkap Id dari Query Params
	id, _:= strconv.Atoi(c.Param("id"))
	

	// Membuat Struct berdasarkan Id dari Query Params
	var projectDetails models.Projects
	
	for index, data := range DataProjects {
		if index == id {
			projectDetails = models.Projects {
					ProjectName: data.ProjectName,
					StartDate: utilities.GetDurationFormat(data.StartDate),
					EndDate: utilities.GetDurationFormat(data.EndDate),
					Duration: data.Duration,
					Description: data.Description,
					Technologies: utilities.GetTechnologiesValue(data.Technologies),
					Image: data.Image,
			}
			
		}
	} 
	
	
	projects := map[string]interface{}{
		"Projects": projectDetails,

	}
	
	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), projects)
}