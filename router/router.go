package router

import (
	"net/http"

	auth "ders-programi/controller/auth"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	// Ana sayfaya "Hello, World!" yanıtı veren bir GET rotası tanımlama
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Başka bir GET rotası tanımlama
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.POST("auth/signup", auth.Signup)
	e.POST("auth/login", auth.Login)
	e.POST("auth/logout", auth.Logout)
	e.GET("auth/userInfo", auth.UserInfo)

}
