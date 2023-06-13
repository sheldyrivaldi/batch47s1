package controllers

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetContactController(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}
