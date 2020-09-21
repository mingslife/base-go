package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"base-go/pkg/models"
)

type UserController struct{}

func (c *UserController) GetMany(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	ctx.JSON(http.StatusOK, gin.H{
		"total": models.CountUsers(),
		"rows":  models.GetUsers(limit, page),
	})
}

func (c *UserController) GetOne(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.JSON(http.StatusOK, models.GetUser(uint(id)))
}

func (c *UserController) Create(ctx *gin.Context) {
	var v models.User
	if err := ctx.BindJSON(&v); err == nil {
		if e := v.Save(); e == nil {
			ctx.JSON(http.StatusCreated, v)
		} else {
			HandleError(ctx, e.Error())
		}
	} else {
		glog.Error(err.Error())
	}
}

func (c *UserController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var v models.User
	ctx.BindJSON(&v)
	v.ID = uint(id)
	if e := v.Update(); e == nil {
		ctx.JSON(http.StatusCreated, v)
	} else {
		HandleError(ctx, e.Error())
	}
}

func (c *UserController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var v models.User
	v.ID = uint(id)
	if e := v.Delete(); e == nil {
		ctx.JSON(http.StatusNoContent, nil)
	} else {
		HandleError(ctx, e.Error())
	}
}

func NewUserController(r gin.IRouter) *UserController {
	c := &UserController{}
	r.Group("/users").
		GET("", c.GetMany).
		GET("/:id", c.GetOne).
		POST("", c.Create).
		PUT("/:id", c.Update).
		DELETE("/:id", c.Delete).
		OPTIONS("/:id", nil)
	return c
}
