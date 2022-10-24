package util

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Response ...
type Response map[string]interface{}

func generateResponse(data interface{}, message string) Response {
	return Response{
		"data":    data,
		"message": message,
	}
}

// Response200 success.....
func Response200(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Thành công"
	}
	return c.JSON(http.StatusOK, generateResponse(data, message))
}

// Response400 badrequest ...
func Response400(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Data invalid"
	}
	return c.JSON(http.StatusBadRequest, generateResponse(data, message))
}

// Response404 not found
func Response404(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Data not found"
	}
	return c.JSON(http.StatusNotFound, generateResponse(data, message))
}

func Response403(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Bạn không có quyền thực hiền hành động này"
	}
	return c.JSON(http.StatusNotFound, generateResponse(data, message))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
