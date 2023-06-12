package main

import (
	"os"
	"stage1/configs"
	"stage1/controllers"

	"github.com/labstack/echo/v4"
)



func main() {
	// Koneksi ke Database
	configs.DatabaseConnection()

	e := echo.New()

	// Mengatur penanganan file static
	e.Static("/public", "public")

	// Daftar Routes GET
	e.GET("/", controllers.GetHomeController)
	e.GET("/projects/:id", controllers.GetProjectDetailController)
	e.GET("/add-project", controllers.GetAddProjectController)
	e.GET("/testimonials", controllers.GetTestimonialsController)
	e.GET("/contact", controllers.GetContactController)
	e.GET("/edit-project/:id", controllers.GetEditProjectController)

	// Daftar Routes POST
	e.POST("/add-project", controllers.AddProjectController)
	e.POST("/edit-project/:id", controllers.SendEditedProjectController)
	e.POST("/delete-project/:id", controllers.DeleteProjectController)


	// Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	e.Logger.Fatal(e.Start(":" + port))
}