package main

import (
	"os"
	"stage1/configs"
	"stage1/controllers"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	// Koneksi ke Database
	configs.DatabaseConnection()

	e := echo.New()

	// Session middleware
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	// Mengatur penanganan file static
	e.Static("/public", "public")

	// Daftar Routes GET
	e.GET("/", controllers.GetHomeController)
	e.GET("/register", controllers.GetRegisterController)
	e.GET("login", controllers.GetLoginController)
	e.GET("/projects/:id", controllers.GetProjectDetailController)
	e.GET("/add-project", controllers.GetAddProjectController)
	e.GET("/testimonials", controllers.GetTestimonialsController)
	e.GET("/contact", controllers.GetContactController)
	e.GET("/edit-project/:id", controllers.GetEditProjectController)

	// Daftar Routes POST
	e.POST("/register", controllers.PostRegisterController)
	e.POST("/login", controllers.PostLoginController)
	e.POST("/add-project", controllers.PostProjectController)
	e.POST("/edit-project/:id", controllers.PostEditedProjectController)
	e.POST("/delete-project/:id", controllers.DeleteProjectController)
	e.POST("/logout", controllers.PostLogoutController)

	// Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
