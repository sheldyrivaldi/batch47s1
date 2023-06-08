package controllers

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func GetTestimonialsController(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}