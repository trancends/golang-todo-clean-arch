package controller

import (
	"fmt"
	"net/http"
	"todo-clean-arch/model/dto"
	"todo-clean-arch/shared/common"
	"todo-clean-arch/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUC usecase.AuthUseCase
	rg     *gin.RouterGroup
}

func (a *AuthController) loginHandler(c *gin.Context) {
	fmt.Println("hit loginHandler")
	var payload dto.AuthRequestDTO
	err := c.ShouldBindJSON(&payload)
	fmt.Println(payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// send payload to usecase
	response, err := a.authUC.Login(payload)
	fmt.Println(response)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("token", response.Token, 3600, "/", "localhost", false, true)
	common.SendSingleResponse(c, response, "login success")
}

func (a *AuthController) Route() {
	a.rg.POST("/login", a.loginHandler)
}

func NewAuthController(authUC usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{
		authUC: authUC,
		rg:     rg,
	}
}
