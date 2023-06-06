package controllers

import (
	"fmt"
	"net/http"
	"stage1/models"
	"text/template"

	"github.com/labstack/echo/v4"
)

var Data []interface{}


func GetAddProjectController(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func AddProjectController(c echo.Context) error {

	//Menangkap value image name
	image, err := c.FormFile("upload-image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Image harus di upload Bos!"})
	}
	src, err := image.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Image harus di upload Bos!"})
	}
	defer src.Close()

	// Menangkap value checkbox
	var checkboxData []string
	reactjs := c.FormValue("reactjs")
	nextjs := c.FormValue("nextjs")
	nodejs := c.FormValue("nodejs")
	typescript := c.FormValue("typescript")
	
	if reactjs != "" {
		checkboxData = append(checkboxData, reactjs)
	}
	if nextjs != "" {
		checkboxData = append(checkboxData, nextjs)
	}
	if nodejs != "" {
		checkboxData = append(checkboxData, nodejs)		
	}
	if typescript != "" {
		checkboxData = append(checkboxData, typescript)
	}

	// Membuat new Project
	newProject := models.Projects{
		ProjectName: c.FormValue("project-name"),
		StartDate: c.FormValue("start-date"),
		EndDate: c.FormValue("end-date"),
		Description: c.FormValue("description"),
		Technologies: checkboxData,
		Image: image.Filename,


	}

	// Memasukan new Project ke Data
	Data = append(Data, newProject)


	// Print Data
	fmt.Println(Data)
	fmt.Println("ProjectName :", newProject.ProjectName)
	fmt.Println("StartDate :", newProject.StartDate)
	fmt.Println("EndDate :", newProject.EndDate)
	fmt.Println("Description :", newProject.Description)
	fmt.Println("Technologies :", newProject.Technologies)
	fmt.Println("Image :", newProject.Image)

	return c.Redirect(http.StatusMovedPermanently, "/add-project")
}