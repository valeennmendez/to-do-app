package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/valeennmendez/to-do/connection"
	"github.com/valeennmendez/to-do/models"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error to decode user",
		})
		return
	}

	/* 	if UserExist(user.Username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user exist",
		})
		return
	} */

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hashed password",
		})
		return
	}

	user.Password = string(passwordHashed)

	if err := connection.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
	})

}

func Login(c *gin.Context) {
	var credentials models.User

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to decode",
		})
		return
	}

	var user models.User

	if err := connection.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user or password",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid password",
		})
		return
	}

	session, _ := store.Get(c.Request, "session-name")
	session.Values["userID"] = user.ID

	pruebaID := session.Values["userID"]

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save session",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successfull",
		"userid":  pruebaID,
	})

}

func UserExist(username string) bool {

	if err := connection.DB.Where("username = ?", username).Error; err != nil {
		return false
	}
	return true

}
