package middlewares

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadImage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("upload-image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer src.Close()

		tempFile, err := ioutil.TempFile("public/images/uploads", "image-*.png")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer tempFile.Close()

		if _, err = io.Copy(tempFile, src); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		data := tempFile.Name()
		filename := data[22:] // public/images/uploads/

		c.Set("dataFile", filename)
		return next(c)
	}
}
