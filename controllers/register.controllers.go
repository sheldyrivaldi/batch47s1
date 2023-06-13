package controllers

import (
	"fmt"
	"net/http"
	"stage1/utilities"
	"text/template"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetRegisterController(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"mesage": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func PostRegisterController(c echo.Context) error {
	errForm := c.Request().ParseForm()
	if errForm != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errForm.Error()})
	}

	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	err := utilities.InsertUser(username, email, string(passwordHash))
	if err != nil {
		utilities.RedirectWithMessage(c, "Register failed! Please try again.", false, "/register")
	} else {
		fmt.Println("User successfully registered!")
	}

	return utilities.RedirectWithMessage(c, "Register success!", true, "/login")
}
