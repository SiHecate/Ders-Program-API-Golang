package router

import (
	"net/http"

	auth "ders-programi/controller/auth"
	controller "ders-programi/controller/ders-programi"

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

	// Kullanıcı işlemleri
	e.POST("/signup", auth.Signup)
	e.POST("/login", auth.Login)
	e.POST("/logout", auth.Logout)
	e.GET("/userInfo", auth.UserInfo)

	// Ders programı işlemleri
	e.GET("/programlar", controller.Programlar)
	e.GET("/kullaniciProgramlari", controller.KullanıcıProgramları)
	e.DELETE("/programKaldir", controller.ProgramSilme)
	e.POST("/programOlustur", controller.ProgramOlustur)
	e.PUT("/programGuncelle", controller.ProgramGuncelle)

}
