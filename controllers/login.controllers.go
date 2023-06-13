package controllers

import (
	"net/http"
	"stage1/utilities"
	"text/template"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetLoginController(c echo.Context) error {
	sess, _ := session.Get("session", c)

	data := map[string]interface{}{
		"Status":  sess.Values["status"],
		"Message": sess.Values["message"],
	}

	//Menghapus data yang dikirim dari register jika ada
	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	tmpl, err := template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func PostLoginController(c echo.Context) error {
	// Menampung value input login
	err := c.Request().ParseForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Mencari data user berdasarkan email
	user, err := utilities.FindOneUser(email)
	if err != nil {
		return utilities.RedirectWithMessage(c, "Invalid email!", false, "/login")
	}

	// Compare password input dengan database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return utilities.RedirectWithMessage(c, "Invalid password!", false, "/login")
	}

	// Membuat session
	sess, _ := session.Get("session", c)

	sess.Options.MaxAge = 21600 // 6 Jam
	sess.Values["message"] = "Successfully logged in!"
	sess.Values["status"] = true
	sess.Values["id"] = user.Id
	sess.Values["username"] = user.Username
	sess.Values["email"] = user.Email
	sess.Values["isLoggedIn"] = true

	sess.Save(c.Request(), c.Response())

	// Redirect to home
	return c.Redirect(http.StatusMovedPermanently, "/")
}
