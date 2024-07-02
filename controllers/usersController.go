package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/EspiraMarvin/go-crud-postgres/initializers"
	"github.com/EspiraMarvin/go-crud-postgres/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get email/pwd off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": "Failed to read body",
		})
		return
	}

	// hash the pwd
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": "Failed hash password",
		})
		return
	}

	// create user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": "Failed to create user",
		})
		return
	}

	// return user resp
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	// get email & pwd off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": "Failed to read body",
		})
		return
	}

	// look up requested user
	var user models.User // var to save the user queried
	// initializers.DB.Where("email = ?", body.Email).First(&user)
	initializers.DB.First(&user, "email = ?", body.Email)

	// compare req password and saved user pwd
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": "Invalid email or password",
		})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UUID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // 30 days
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": "Failed to create token",
		})
		return
	}

	//send jwt token back to user
	// c.JSON(http.StatusOK, gin.H{
	// "token": tokenString,
	// })

	// send as cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
