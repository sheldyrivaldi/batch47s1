package controllers

import (
	"net/http"
	"stage1/models"
	"stage1/utilities"
	"strconv"
	"text/template"

	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo/v4"
)

func GetProjectDetailController(c echo.Context) error {
	// Menangkap Id dari Query Params
	id, _ := strconv.Atoi(c.Param("id"))

	// Membuat Struct berdasarkan Id dari Query Params
	data, err := utilities.FindOneProject(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data.StartDate = data.StartDateTime.Format("01 Jan 2006")
	data.EndDate = data.EndDateTime.Format("01 Jan 2006")

	//Define session value to session model
	sess, _ := session.Get("session", c)
	var sessionData models.SessionData

	if sess.Values["isLoggedIn"] != true {
		sessionData.IsLoggedIn = false
	} else {
		sessionData.IsLoggedIn = sess.Values["isLoggedIn"].(bool)
		sessionData.Username = sess.Values["username"].(string)
	}

	projects := map[string]interface{}{
		"Projects":    data,
		"SessionData": sessionData,
	}

	tmpl, err := template.ParseFiles("views/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}
