package controllers

import (
	"net/http"

	"challange11/database"
	"challange11/helpers"
	"challange11/models"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContenType(ctx)
	user := models.User{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	err := db.Debug().Create(&user).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors":  "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"full_name": user.FullName,
	})
}

func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContenType(ctx)
	user := models.User{}
	var password string

	if contentType == appJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	password = user.Password

	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email / password",
			"data":    user,
			"pasd":    password,
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email / password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email, user.Level)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
