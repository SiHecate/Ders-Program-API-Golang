package controllers

import (
	"ders-programi/database"
	"ders-programi/model"
	"ders-programi/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Kullancıı kayıt fonksiyonu
func Signup(c echo.Context) error {
	var signupRequest struct {
		Name             string `json:"name"`
		Lastname         string `json:"lastname"`
		Email            string `json:"email"`
		Password         string `json:"password"`
		Password_confirm string `json:"password_confirm"`
	}

	signupRequest.Name = strings.TrimSpace(signupRequest.Name)
	signupRequest.Lastname = strings.TrimSpace(signupRequest.Lastname)
	signupRequest.Email = strings.TrimSpace(signupRequest.Email)

	if err := c.Bind(&signupRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Hatalı parametre: " + err.Error()})
	}

	if signupRequest.Password != signupRequest.Password_confirm {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Şifreler uyuşmuyor"})
	}

	var existingUser model.User
	if err := database.Conn.Where("email = ?", signupRequest.Email).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Kayıt olunmaya çalışan kullanıcı mevcut"})
	}

	hashedPassword, err := utils.GenerateHashPassword(signupRequest.Password)
	if err != nil {
		return err
	}

	newUser := model.User{
		Name:     signupRequest.Name,
		Lastname: signupRequest.Lastname,
		Email:    signupRequest.Email,
		Password: hashedPassword,
	}

	if err := database.Conn.Create(&newUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "User oluşturma database tarafında başarısız"})
	}

	return c.JSON(http.StatusCreated, newUser)
}

// Kullanıcı girişi fonksiyonu
func Login(c echo.Context) error {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Hatalı parametre: " + err.Error()})
	}

	var existingUser model.User
	database.Conn.Where("email = ?", loginRequest.Email).First(&existingUser)

	if existingUser.ID == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Kullanıcı bulunamadı."})
	}

	if err := comparePasswords(existingUser.Password, loginRequest.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid password",
		})
	}

	// Generate JWT token
	tokenString, err := generateJWT(existingUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "JWT Token oluşturulamadı"})
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Minute * 60),
		HttpOnly: true,
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Kullanıcı girişi başarılı!",
	})
}

func generateJWT(user model.User) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &model.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			Subject:   user.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}
	var jwtKey = []byte("my_secret_key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func comparePasswords(hashedPassword string, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err
}

func Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Çıkış yapıldı",
	})
}

func UserInfo(c echo.Context) error {
	userId := c.Get("issuer")

	var user model.User
	if err := database.Conn.First(&user, userId).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "User not found",
		})
	}

	return c.JSON(http.StatusOK, user)
}
