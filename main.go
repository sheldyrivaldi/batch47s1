package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"stage1/controllers"

	"github.com/labstack/echo/v4"
)



func main() {
	e := echo.New()

	// Mengatur penanganan file static
	e.Static("/public", "public")

	// Daftar Routes GET
	e.GET("/", controllers.GetHomeController)
	e.GET("/project", controllers.GetProjectDetailController)
	e.GET("/add-project", controllers.GetAddProjectController)
	e.GET("/testimonials", controllers.GetTestimonialsController)
	e.GET("/contacts", controllers.GetContactController)

	//Daftar Routes POST
	e.POST("/add-project", controllers.AddProjectController)

	// Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e.Logger.Fatal(e.Start(":" + port))
}


func addProject(c echo.Context) error {
	name := c.FormValue("project-name")
	desc := c.FormValue("description")
	fmt.Println(name)
	fmt.Println(desc)


	// Source
	image, err := c.FormFile("upload-image")
	if err != nil {
		return err
	}
	src, err := image.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	image.Filename = "12";
	dst, err := os.Create("public/project-image/" + image.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", image.Filename, name, desc))
}