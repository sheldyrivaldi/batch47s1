package controllers

import (
	"net/http"
	"stage1/models"
	"stage1/utilities"
	"text/template"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetHomeController(c echo.Context) error {
	// Get data dari database
	data, err := utilities.FindProjects()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Session
	sess, _ := session.Get("session", c)

	var sessionData models.SessionData

	if sess.Values["isLoggedIn"] != true {
		sessionData.IsLoggedIn = false
	} else {
		sessionData.IsLoggedIn = sess.Values["isLoggedIn"].(bool)
		sessionData.Username = sess.Values["username"].(string)
	}

	//Data
	projects := map[string]interface{}{
		"Projects":    data,
		"Status":      sess.Values["status"],
		"Message":     sess.Values["message"],
		"SessionData": sessionData,
	}

	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
}
