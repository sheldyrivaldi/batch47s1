package controllers

import (
	"fmt"
	"net/http"
	"stage1/utilities"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetEditProjectController(c echo.Context) error {
	// Menangkap Id dari Query Params
	id := c.Param("id")
	

	// Get data dari database berdasarkan Id
	result := utilities.FindOneProject(id)

	// Membuat Map untuk dikirim export ke html
	projects := map[string]interface{}{
		"Projects":            result,
		"IsCheckedReactJs":    utilities.GetTechnologiesChecked(result.Technologies, "reactjs"),
		"IsCheckedNextJs":     utilities.GetTechnologiesChecked(result.Technologies, "nextjs"),
		"IsCheckedNodeJs":     utilities.GetTechnologiesChecked(result.Technologies, "nodejs"),
		"IsCheckedTypescript": utilities.GetTechnologiesChecked(result.Technologies, "typescript"),
	}

	// Render templates
	var tmpl, err = template.ParseFiles("views/edit-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}


	return tmpl.Execute(c.Response(), projects)
	
}



func SendEditedProjectController(c echo.Context) error {
	// Menangkap Id dari Query Params
	id := c.Param("id")

	
	// Menangkap value checkbox
	reactjs := c.FormValue("reactjs")
	nextjs := c.FormValue("nextjs")
	nodejs := c.FormValue("nodejs")
	typescript := c.FormValue("typescript")
	technologiesData := utilities.GetTechnologies(reactjs, nextjs, nodejs, typescript)

	// Membuat New Project
		projectName := c.FormValue("project-name")
		startDate := c.FormValue("start-date")
		endDate := c.FormValue("end-date")
		description := c.FormValue("description")
		technologies := technologiesData
		image := "project-list1.png"
	

	// Insert New Project to Database
	utilities.UpdateProject(id, projectName, startDate, endDate, description, technologies, image)

	fmt.Println("Project successfully updated!")
	
	return c.Redirect(http.StatusMovedPermanently, "/")

}