package main

import (
	"ders-programi/database"
	"ders-programi/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
- öğrenci günlük planlarını kayıt edebilecek. // BİTTİ
- eklenen planın hangi gün ve saat aralığı kayıt edilecek. // BİTTİ
- planların iptal, bitti, yapılıyor gibi state durumları olacak (veri type’ları size bağlı, string veya integer olarak tutabilirsiniz.) BİTTİ
- planlar üzerinde güncelleme ve silme işlemleri yapılacak. // BİTTİ
- eklenen plan tarihinde ve saat aralığında başka bir plan olup olmadığını kontrol etme. // BİTTİ
- haftalık ve aylık listeleme seçenekleri olacak (bu madde isteğe bağlıdır, yapılması durumunda size artı katkı sağlar.) // BİTTİ
- öğrencilerin kayıt olması bilgilerini güncellemesi olacak (bu madde isteğe bağlıdır, yapılması durumunda size artı katkı sağlar.) // BİTTİ
*/

func main() {
	database.Connect()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.Router(e)

	e.Logger.Fatal(e.Start(":8080"))
}
