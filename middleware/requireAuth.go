package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/EspiraMarvin/go-crud-postgres/initializers"
	"github.com/EspiraMarvin/go-crud-postgres/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	logg "github.com/sirupsen/logrus"
)

func RequireAuth(c *gin.Context) {
	// get the cookie off req
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// validate/decode cookie
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user with token sub
		var user models.User
		// initializers.DB.First(&user, claims["sub"])
		// initializers.DB.First(&user, "uuid = ?", claims["sub"])
		initializers.DB.Where("uuid = ?", claims["sub"]).First(&user)

		if user.UUID == "" {
			logg.Info("USER WITH TOKEN NOT FOUND", user)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		logg.Info("USER WITH TOKEN FOUND", user)

		// attach to req
		c.Set("user", user)

		// continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if err != nil {
		log.Fatal(err)
	}

}
