package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/IrvanWijayaSardam/mynotesapp/dto"
	"github.com/IrvanWijayaSardam/mynotesapp/entity"
	"github.com/IrvanWijayaSardam/mynotesapp/helper"
	"github.com/IrvanWijayaSardam/mynotesapp/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PagerController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	FindStatusById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type pagerController struct {
	pagerService service.PagerService
	jwtService   service.JWTService
}

func NewPagerController(pagerserv service.PagerService, jwtServ service.JWTService) PagerController {
	return &pagerController{
		pagerService: pagerserv,
		jwtService:   jwtServ,
	}
}

func (c *pagerController) All(context *gin.Context) {
	var pager []entity.Pager = c.pagerService.All()
	res := helper.BuildResponse(true, "OK!", pager)
	context.JSON(http.StatusOK, res)
}

func (c *pagerController) FindById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("user_id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No Parameter ID was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var pager entity.Pager = c.pagerService.FindByID(id)
	if (pager == entity.Pager{}) {
		res := helper.BuildErrorResponse("Data Not Found", "No Data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", pager)
		context.JSON(http.StatusOK, res)
	}
}

func (c *pagerController) Insert(context *gin.Context) {
	var pagerCreateDTO dto.PagerCreateDTO
	errDTO := context.ShouldBind(&pagerCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			pagerCreateDTO.UserID = convertedUserID
		}
		result := c.pagerService.Insert(pagerCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *pagerController) FindStatusById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("user_id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No Parameter ID was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var pager entity.Pager = c.pagerService.FindByID(id)
	if (pager == entity.Pager{}) {
		res := helper.BuildErrorResponse("Data Not Found", "No Data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", true)
		context.JSON(http.StatusOK, res)
	}
}

func (c *pagerController) Update(context *gin.Context) {
	var pagerUpdateDTO dto.PagerUpdateDTO
	errDTO := context.ShouldBind(&pagerUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["userid"])
	if c.pagerService.IsAllowedToEdit(userID, pagerUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			pagerUpdateDTO.UserID = id
		}
		result := c.pagerService.Update(pagerUpdateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *pagerController) Delete(context *gin.Context) {
	var pager entity.Pager
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	pager.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["userid"])
	if c.pagerService.IsAllowedToEdit(userID, pager.ID) {
		c.pagerService.Delete(pager)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *pagerController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["userid"])
	return id
}
