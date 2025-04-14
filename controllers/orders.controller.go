package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milkyway/gin_beginer/models"
	"github.com/milkyway/gin_beginer/services"
)

type OrdersController struct {
	OrdersService *services.OrdersService
}

// constructor
// note : possibly using interface ? look chatGPT
func NewOrdersController(services *services.OrdersService) *OrdersController {
	return &OrdersController{
		OrdersService: services,
	}
}

// NOTE: klo gak binding ke struct gak masalah gak dipake

func (oc OrdersController) badRequestErrorResp(message string, err string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
		"err":     err,
	})
}

func (oc OrdersController) unprocessableEntityErrorResp(message string, err string, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": message,
		"err":     err,
	})
}

func (oc OrdersController) OrdersCheck(ctx *gin.Context) {

	var response models.BodyReponseAPI = models.BodyReponseAPI{}

	// cust_registry_param, ok := ctx.Params.Get("registry-cust")
	// if !ok {
	// 	badRequestErrorResp("failed get param", "registry-cust not found", ctx)
	// 	return
	// }

	resFetchBook, err := oc.OrdersService.BooksRepo.FetchBookByID(ctx, 2)
	if err != nil {
		oc.unprocessableEntityErrorResp("test1", err.Error(), ctx)
		return
	}

	resFetchCust, err := oc.OrdersService.CustomersRepo.FetchByRegistry(ctx, "mrm123")
	if err != nil {
		oc.unprocessableEntityErrorResp("test2", err.Error(), ctx)
		return
	}

	// anon slice of struct
	orderInfoReq := []services.OrderBookRequest{
		{
			ID:    "1",
			Title: "book title",
			Qty:   100,
		},
		{
			ID:    "10",
			Title: "book qwer",
			Qty:   1000,
		},
	}
	respCheck := oc.OrdersService.BuyBooks(orderInfoReq)

	response.Data = struct {
		Book      models.Books
		Cek       models.Customers
		ReqCheck  []services.OrderBookRequest
		RespCheck bool
	}{
		Book:      resFetchBook,
		Cek:       resFetchCust,
		ReqCheck:  orderInfoReq,
		RespCheck: respCheck,
	}
	response.Message = "success"
	ctx.JSON(http.StatusAccepted, gin.H{
		"body": response,
	})
}

func (oc OrdersController) OrderBooks(ctx *gin.Context) {
	var response models.BodyReponseAPI = models.BodyReponseAPI{}
	var order_req_info []services.OrderBookRequest

	if err := ctx.ShouldBindJSON(&order_req_info); err != nil {
		badRequestErrorResp("error request order book", err.Error(), ctx)
		return
	}
	// get request []req_data

	response.Data = order_req_info
	response.Message = "testing"
	ctx.JSON(http.StatusOK, gin.H{
		"body": response,
	})
}
