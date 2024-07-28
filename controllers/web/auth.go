package web

import (
	"context"

	"github.com/labstack/echo"
	"github.com/viay1668/spectre/templates/common/userauth"
)

func RenderLogin(c echo.Context) error {
	loginComponent := userauth.Login()
	return loginComponent.Render(context.Background(), c.Response())
}

func RenderSignup(c echo.Context) error {
	signupComponent := userauth.Signup()
	return signupComponent.Render(context.Background(), c.Response())
}
