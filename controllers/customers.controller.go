package controllers

import (
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
	"github.com/milkyway/gin_beginer/repositories"
	"github.com/milkyway/gin_beginer/services"
)

type CustomersController struct {
	CustomersService services.CustomerService
}

// constructor
// note : possibly using interface ? look chatGPT
func NewCustomersController(db_arg *sql.DB) CustomersController {
	return CustomersController{
		CustomersService: services.CustomerService{
			CustomersRepo: repositories.CustomersRepo{
				DB: db_arg,
			},
		},
	}
}

// NOTE: klo gak binding ke struct gak masalah gak dipake

func badRequestErrorResp(message string, err string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
		"err":     err,
	})
}

func unprocessableEntityErrorResp(message string, err string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": message,
		"err":     err,
	})
}

func (cc CustomersController) GetBookByRegistryNumber(ctx *gin.Context) {

	var response models.BodyReponseAPI = models.BodyReponseAPI{}

	cust_registry_param, ok := ctx.Params.Get("registry-cust")
	if !ok {
		badRequestErrorResp("failed get param", "registry-cust not found", ctx)
		return
	}

	result, err := cc.CustomersService.GetInfoByRegistry(ctx, cust_registry_param)
	if err != nil {
		unprocessableEntityErrorResp("err get customer by cust registry", err.Error(), ctx)
		return
	}

	response.Data = result
	response.Message = "success"
	ctx.JSON(http.StatusAccepted, gin.H{
		"body": response,
	})
}
