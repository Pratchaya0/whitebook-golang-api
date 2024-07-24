package controllers

// import (
// 	"net/http"

// 	"github.com/Pratchaya0/whitebook-golang-api/dtos/responses"
// 	"github.com/Pratchaya0/whitebook-golang-api/entities"
// 	"github.com/Pratchaya0/whitebook-golang-api/services"
// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/crypto/bcrypt"
// )

// type LoginPayload struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type SignUpPayload struct {
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type LoginResponse struct {
// 	Token string        `json:"token"`
// 	ID    uint          `json:"id"`
// 	User  entities.User `json:"user"`
// 	Role  string
// }

// // @Summary Login
// // @Schemes
// // @Description ลงชื่อเข้าใช้ระบบ
// // @Tags Auth
// // @Accept json
// // @Produce json
// // @Success 200 {object} responses.Response{} "ok"
// // @Router /login [post]
// func Login(c *gin.Context) {
// 	var payload LoginPayload
// 	var user entities.User
// 	var userRole entities.UserRole

// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entities.DB().Raw("SELECT * FROM users WHERE email = ?", payload.Email).Preload("UserRole").Find(&user); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
// 		return
// 	}

// 	err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(payload.Password))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "password is incorrect"})
// 		return
// 	}

// 	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
// 	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
// 	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
// 	// ExpirationHours เป็นเวลาหมดอายุของ token

// 	jwtWrapper := services.JwtWrapper{
// 		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
// 		Issuer:          "AuthService",
// 		ExpirationHours: 24,
// 	}

// 	signedToken, err := jwtWrapper.GenerateToken(user.UserEmail)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
// 		return
// 	}

// 	if tx := entities.DB().Raw("SELECT * FROM user_roles WHERE user_role_name = ?", user.UserRoleId).First(&userRole); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "user role not found"})
// 		return
// 	}

// 	tokenResponse := LoginResponse{
// 		Token: signedToken,
// 		ID:    user.ID,
// 		User:  user,
// 		Role:  userRole.UserRoleName,
// 	}

// 	webResponse := responses.Response{
// 		Code:   http.StatusOK,
// 		Status: "Ok",
// 		Data:   tokenResponse,
// 	}

// 	c.JSON(http.StatusOK, webResponse)
// }

// // @Summary
// // @Description แก้ไขข้อมูลหนังสือ
// // @Tags Auth
// // @Accept json
// // @Produce json
// // @Success 200 {object} responses.Response{} "ok"
// // @Router /signup [post]
// func SignUp(c *gin.Context) {
// 	var payload SignUpPayload
// 	var user entities.User

// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
// 	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
// 		return
// 	}

// 	user.UserName = payload.Name
// 	user.UserEmail = payload.Email
// 	user.UserPassword = string(hashPassword)

// 	if err := entities.DB().Create(&user).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	webResponse := responses.Response{
// 		Code:   http.StatusOK,
// 		Status: "Ok",
// 		Data:   user,
// 	}

// 	c.JSON(http.StatusCreated, webResponse)
// }
