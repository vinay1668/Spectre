package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/viay1668/spectre/templates/common"
)

func HandleLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Here you would typically validate the credentials against a database
	// For this example, we'll use a simple check
	if username == "user@example.com" && password == "password" {
		//set session cookie
		jsonResponse := map[string]string{
			"username": username,
			"password": password,
		}
		return c.JSON(http.StatusOK, jsonResponse)

	} else {
		return common.LoginFailure().Render(c.Request().Context(), c.Response().Writer)
	}
}

func HandleSignup(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Here you would typically create a new user in your database
	// For this example, we'll use a simple check
	if username != "" && password != "" {
		return c.HTML(http.StatusOK, `<div id="signup-response" class="text-green-500">Signup successful!</div>`)
	} else {
		return c.HTML(http.StatusBadRequest, `<div id="signup-response" class="text-red-500">Invalid username or password</div>`)
	}
}
