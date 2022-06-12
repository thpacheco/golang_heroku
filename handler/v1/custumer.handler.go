package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/thpacheco/golang_heroku/common/obj"
	"github.com/thpacheco/golang_heroku/common/response"
	"github.com/thpacheco/golang_heroku/dto"
	"github.com/thpacheco/golang_heroku/service"
)

type CustumerHandler interface {
	All(ctx *gin.Context)
	Createcustumer(ctx *gin.Context)
	Updatecustumer(ctx *gin.Context)
	Deletecustumer(ctx *gin.Context)
	FindOnecustumerByID(ctx *gin.Context)
}

type custumerHandler struct {
	custumerService service.CustumerService
	jwtService      service.JWTService
}

func NewCustumerHandler(custumerService service.CustumerService, jwtService service.JWTService) CustumerHandler {
	return &custumerHandler{
		custumerService: custumerService,
		jwtService:      jwtService,
	}
}

func (c *custumerHandler) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	custumers, err := c.custumerService.All(userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", custumers)
	ctx.JSON(http.StatusOK, response)
}

func (c *custumerHandler) Createcustumer(ctx *gin.Context) {
	var createcustumerReq dto.CreateCustumerRequest
	createcustumerReq.DataInicio = time.Now().Local().String()

	err := ctx.ShouldBind(&createcustumerReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.custumerService.CreateCustumer(createcustumerReq, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

func (c *custumerHandler) FindOnecustumerByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.custumerService.FindOneCustumerByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *custumerHandler) Deletecustumer(ctx *gin.Context) {
	id := ctx.Param("id")

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	err := c.custumerService.DeleteCustumer(id, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

func (c *custumerHandler) Updatecustumer(ctx *gin.Context) {
	updatecustumerRequest := dto.UpdateCustumerRequest{}
	err := ctx.ShouldBind(&updatecustumerRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updatecustumerRequest.ID = id
	custumer, err := c.custumerService.UpdateCustumer(updatecustumerRequest, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", custumer)
	ctx.JSON(http.StatusOK, response)

}
