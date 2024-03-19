package router

import (
	"net/http"

	auth "ders-programi/controller/auth"
	controller "ders-programi/controller/ders-programi"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// Kullanıcı işlemleri
	e.POST("/signup", auth.Signup)
	e.POST("/login", auth.Login)
	e.POST("/logout", auth.Logout)
	e.POST("/update", auth.UserUpdate)
	e.GET("/userInfo", auth.UserInfo)

	// Ders programı işlemleri
	e.GET("/programlar", controller.Programlar)
	e.GET("/kullaniciProgramlari", controller.KullanıcıProgramları)
	e.GET("/programlarAylik", controller.AylıkPlanlar)
	e.GET("/programlarHaftalik", controller.HaftalıkPlanlar)
	e.DELETE("/programKaldir", controller.ProgramSilme)
	e.POST("/programOlustur", controller.ProgramOlustur)
	e.PUT("/programGuncelle", controller.ProgramGuncelle)
	e.PUT("/iptal", controller.DurumIptal)
	e.PUT("/tamamlandi", controller.DurumBitti)
	e.PUT("/devam", controller.DurumDevam)

}
