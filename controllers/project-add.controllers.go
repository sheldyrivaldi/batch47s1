package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"stage1/models"
	"stage1/utilities"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetAddProjectController(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func PostProjectController(c echo.Context) error {
	// Session
	sess, _ := session.Get("session", c)

	var sessionData models.SessionData

	if sess.Values["isLoggedIn"] != true {
		sessionData.IsLoggedIn = false
	} else {
		sessionData.IsLoggedIn = sess.Values["isLoggedIn"].(bool)
		sessionData.UserId = sess.Values["id"].(int)
		sessionData.Username = sess.Values["username"].(string)
	}

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
	userId := sessionData.UserId

	// Insert New Project to Database
	err := utilities.InsertProject(projectName, startDate, endDate, description, technologies, image, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	} else {
		fmt.Println("Project successfully created!")
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}
