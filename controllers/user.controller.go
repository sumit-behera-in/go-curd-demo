package controllers

import (
	"go-curd-demo/models"
	"go-curd-demo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if error := ctx.ShouldBindJSON(&user); error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": error.Error()})
		return
	}

	if error := uc.UserService.CreateUser(&user); error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Sucess"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, error := uc.UserService.GetUser(username)
	if error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user) // TODO: Implement
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	userSlice, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userSlice) // TODO: Implement
}

func (uc *UserController) UpateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := uc.UserService.UpateUser(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, "Sucessfully updated") // TODO: Implement
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "") // TODO: Implement
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/getall", uc.GetAll)
	userRoute.GET("/get/:name", uc.GetUser)
	userRoute.PATCH("/update", uc.UpateUser)
	userRoute.DELETE("/delete", uc.DeleteUser)
}
