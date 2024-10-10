package routes

import (
	"net/http"

	"example.com/rest_api/models"
	"example.com/rest_api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data !"})
		return
	}

	err = user.Save()

	if err !=nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user !"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message":"User created Successfully"})
}

func login(context *gin.Context){
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data !"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message":"Could not authenticate User"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not Authenticate the user"})
		return 
	}



	context.JSON(http.StatusOK, gin.H{"message":"Login Successful", "token":token} )

}