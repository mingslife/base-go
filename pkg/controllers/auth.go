package controllers

import (
	"base-go/pkg/middleware"
	"base-go/pkg/models"
	"base-go/pkg/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

type AuthUser struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (c *AuthController) Login(ctx *gin.Context) {
	authUser := &AuthUser{}
	if err := ctx.BindJSON(authUser); err != nil {
		return
	}
	fmt.Println(utils.GenerateFromPassword(authUser.Password))
	if user := models.GetUserByUsernameAndPassword(authUser.Account, authUser.Password); user != nil {
		token, _, _ := middleware.JWT.GenerateToken(map[string]interface{}{
			"sub": user.ID,
		})
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect account or password"})
	}
}

func (c *AuthController) Logout(ctx *gin.Context) {}

func (c *AuthController) Refresh(ctx *gin.Context) {
	if token, _, err := middleware.JWT.RefreshToken(ctx); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func NewAuthController(r gin.IRouter) *AuthController {
	c := &AuthController{}
	r.Group("/auth").
		POST("/login", c.Login).
		POST("/logout", c.Logout).
		POST("/refresh", c.Refresh)
	return c
}

func GetAuthUser(ctx *gin.Context) {}
