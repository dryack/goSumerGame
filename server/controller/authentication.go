package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"goSumerGame/server/model"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Register(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createSendToken(context, *savedUser, http.StatusCreated)
}

// FIXME: this has stopped correctly returning the user's games
func Login(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.FindUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createSendToken(context, user, http.StatusCreated)
}

func createSendToken(c *gin.Context, user model.User, status int) {
	jwtSecret := []byte(os.Getenv("JWT_PRIVATE_KEY"))
	jwtExp, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	jwtDuration := time.Hour * 1 * time.Duration(jwtExp)
	jwtMaxAge := time.Now().Add(jwtDuration).Unix()

	jwt, err := generateJWT(user, jwtMaxAge, jwtSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", jwt, int(jwtDuration.Seconds()), "/", "localhost", false, true)

	c.JSON(status, gin.H{"token": jwt, "user": user})
}

func generateJWT(user model.User, maxAge int64, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		// using "eat" doesn't actually expire the token
		"exp": maxAge,
	})
	return token.SignedString(secret)
}
