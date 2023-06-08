package controllers

import (
	"fmt"
	"net/http"
	"stage1/models"
	"text/template"

	"stage1/utilities"

	"github.com/labstack/echo/v4"
)

// Dummy Data Project
var DataProjects = []models.Projects{
	{
		ProjectName: "Dumbways Mobile - 2021",
		StartDate: "2021-06-05",
		EndDate: "2021-07-05",
		Duration: utilities.GetDuration("2021-06-05", "2021-07-05"),
		Description: "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Doloribus, deserunt! Lorem, ipsum dolor sit amet consectetur adipisicing elit. Vero nostrum nam tempora quos eum velit quia at qui eos beatae?",
		Technologies: []string{"reactjs", "nextjs", "nodejs", "typescript"},
		Image: "project-list1.png",
	},
	{
		ProjectName: "Dumbways Website - 2022",
		StartDate: "2022-02-05",
		EndDate: "2022-07-05",
		Duration: utilities.GetDuration("2022-02-05", "2022-07-05"),
		Description: "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Doloribus, deserunt! Lorem, ipsum dolor sit amet consectetur adipisicing elit. Vero nostrum nam tempora quos eum velit quia at qui eos beatae?",
		Technologies: []string{"reactjs", "nextjs", "nodejs", "typescript"},
		Image: "project-list1.png",
	},
	{
		ProjectName: "Dumbways IOS App - 2023",
		StartDate: "2023-03-05",
		EndDate: "2023-06-05",
		Duration: utilities.GetDuration("2023-03-05", "2023-06-05"),
		Description: "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Doloribus, deserunt! Lorem, ipsum dolor sit amet consectetur adipisicing elit. Vero nostrum nam tempora quos eum velit quia at qui eos beatae?",
		Technologies: []string{"reactjs", "nextjs", "nodejs", "typescript"},
		Image: "project-list1.png",
	},
}


func GetAddProjectController(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func AddProjectController(c echo.Context) error {

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
	newProject := models.Projects{
		ProjectName: c.FormValue("project-name"),
		StartDate: c.FormValue("start-date"),
		EndDate: c.FormValue("end-date"),
		Duration: duration,
		Description: c.FormValue("description"),
		Technologies: technologiesData,
		Image: "project-list1.png",
	}

	DataProjects = append(DataProjects, newProject)
	fmt.Println("Project Successfully Created!")
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}