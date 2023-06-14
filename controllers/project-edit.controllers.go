package controllers

import (
	"fmt"
	"net/http"
	"os"
	"stage1/utilities"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetEditProjectController(c echo.Context) error {
	// Menangkap Id dari Query Params
	id, _ := strconv.Atoi(c.Param("id"))

	// Get data dari database berdasarkan Id
	data, err := utilities.FindOneProject(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Membuat Map untuk dikirim export ke html
	projects := map[string]interface{}{
		"Projects":            data,
		"IsCheckedReactJs":    utilities.GetTechnologiesChecked(data.Technologies, "reactjs"),
		"IsCheckedNextJs":     utilities.GetTechnologiesChecked(data.Technologies, "nextjs"),
		"IsCheckedNodeJs":     utilities.GetTechnologiesChecked(data.Technologies, "nodejs"),
		"IsCheckedTypescript": utilities.GetTechnologiesChecked(data.Technologies, "typescript"),
	}

	// Render templates
	tmpl, err := template.ParseFiles("views/edit-project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}

func PostEditedProjectController(c echo.Context) error {
	// Menangkap Id dari Query Params

	id, _ := strconv.Atoi(c.Param("id"))

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
	image := c.Get("dataFile").(string)

	//Data Sebelumnya
	data, errDB := utilities.FindOneProject(id)
	if errDB != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errDB.Error()})
	}

	// Insert Updated Project to Database
	err := utilities.UpdateProject(id, projectName, startDate, endDate, description, technologies, image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	} else {
		err := os.Remove(fmt.Sprintf("public/images/uploads/%s", data.Image))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		} else {
			fmt.Println("Project successfully updated!")
		}
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}
