package controllers

import (
	"fmt"
	"net/http"
	"stage1/models"
	"stage1/utilities"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetEditProjectController(c echo.Context) error {
	// Menangkap Id dari Query Params
	id, _:= strconv.Atoi(c.Param("id"))
	

	// Membuat Struct berdasarkan Id dari Query Params
	var projectDetails models.Projects
	
	for index, data := range DataProjects {
		if index == id {
			projectDetails = models.Projects {
				Id: id,
				ProjectName: data.ProjectName,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
				Description: data.Description,
				Technologies: utilities.GetTechnologiesValue(data.Technologies),
				Image: data.Image,
			}
			
		}
	} 
	
	
	projects := map[string]interface{}{
		"Projects": projectDetails,
		"IsCheckedReactJs": utilities.GetTechnologiesChecked(projectDetails.Technologies, "reactjs"),
		"IsCheckedNextJs": utilities.GetTechnologiesChecked(projectDetails.Technologies, "nextjs"),
		"IsCheckedNodeJs": utilities.GetTechnologiesChecked(projectDetails.Technologies, "nodejs"),
		"IsCheckedTypescript": utilities.GetTechnologiesChecked(projectDetails.Technologies, "typescript"),

	}

	
	var tmpl, err = template.ParseFiles("views/edit-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), projects)
	
}



func SendEditedProjectController(c echo.Context) error {
	// Menangkap Id dari Query Params
	id, _:= strconv.Atoi(c.Param("id"))

	// //Menangkap value image name
	// image, err := c.FormFile("upload-image")
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Image harus di upload Bos!"})
	// }

	// Menangkap value checkbox
	reactjs := c.FormValue("reactjs")
	nextjs := c.FormValue("nextjs")
	nodejs := c.FormValue("nodejs")
	typescript := c.FormValue("typescript")
	technologiesData := utilities.GetTechnologies(reactjs, nextjs, nodejs, typescript)


	// Membuat new Project
	duration := utilities.GetDuration(c.FormValue("start-date"), c.FormValue("end-date"))
	editedProject := models.Projects{
		ProjectName: c.FormValue("project-name"),
		StartDate: c.FormValue("start-date"),
		EndDate: c.FormValue("end-date"),
		Duration: duration,
		Description: c.FormValue("description"),
		Technologies: technologiesData,
		Image: "project-list1.png",
	}

	DataProjects[id] = editedProject
	fmt.Printf("Project with id: %d, Successfully Updated!\n", id)
	
	return c.Redirect(http.StatusMovedPermanently, "/")

}