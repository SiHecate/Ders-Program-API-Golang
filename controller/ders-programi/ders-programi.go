package controller

import (
	controller "ders-programi/controller/auth"
	"ders-programi/database"
	"ders-programi/model"
	"errors"
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

	// Zaman kontrolü yapılacak
	err := zamanKontrol(programRequest.BaslangicZamani, programRequest.BitisZamani)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Zaman kontrol hatası: " + err.Error()})
	}

	UserID, err := controller.GetUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
	}

	newPlan := model.Plan{
		UserID:          UserID,
		Baslik:          programRequest.Baslik,
		Plan:            programRequest.Plan,
		Gun:             programRequest.Gun,
		BaslangicZamani: programRequest.BaslangicZamani,
		BitisZamani:     programRequest.BitisZamani,
		Durum:           programRequest.Durum,
	}

	if err := database.Conn.Create(&newPlan).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Plan oluşturma database tarafında başarısız."})
	}

	// Başarılı bir şekilde tamamlandı mesajını döndür
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Program başarıyla oluşturuldu", "program": newPlan})
}

func zamanKontrol(baslangic_zamanı time.Time, bitis_zamanı time.Time) error {
	var existingPlans []model.Plan
	if err := database.Conn.Find(&existingPlans).Error; err != nil {
		return err
	}

	for _, plan := range existingPlans {
		if baslangic_zamanı.Before(plan.BitisZamani) && bitis_zamanı.After(plan.BaslangicZamani) {
			return errors.New("zaman aralıkları çakışıyor")
		}
	}

	if baslangic_zamanı.After(bitis_zamanı) {
		return errors.New("başlangıç zamanı, bitiş zamanından sonra olamaz")
	}

	return nil
}
