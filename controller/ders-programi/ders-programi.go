package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

/*
- öğrenci günlük planlarını kayıt edebilecek.
- eklenen planın hangi gün ve saat aralığı kayıt edilecek.
- planların iptal, bitti, yapılıyor gibi state durumları olacak (veri type’ları size bağlı, string veya integer olarak tutabilirsiniz.)
- planlar üzerinde güncelleme ve silme işlemleri yapılacak.
- eklenen plan tarihinde ve saat aralığında başka bir plan olup olmadığını kontrol etme.
- haftalık ve aylık listeleme seçenekleri olacak (bu madde isteğe bağlıdır, yapılması durumunda size artı katkı sağlar.)
- öğrencilerin kayıt olması bilgilerini güncellemesi olacak (bu madde isteğe bağlıdır, yapılması durumunda size artı katkı sağlar.) // BİTTİ
*/

func ProgramOlustur(c echo.Context) error {
	var programRequest struct {
		Baslik          string    `json:"başlık"`
		Plan            string    `json:"plan"`
		Gun             time.Time `json:"gün"`
		BaslangicZamani time.Time `json:"başlangıç_zamanı"`
		BitisZamani     time.Time `json:"bitiş_zamanı"`
		Durum           string    `json:"durum"`
	}

	if err := c.Bind(&programRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Hatalı parametre: " + err.Error()})
	}

	err := zamanKontrol(programRequest.BaslangicZamani, programRequest.BitisZamani) if err != nil {
		return err
	}

	return nil
}

func zamanKontrol(başlangıç_zamanı time.Time, bitiş_zamanı time.Time) error {

}
