package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
	"net/http"
	"time"
)

type Response struct {
	Data  interface{}    `json:"data,omitempty"`
	Error *ErrorResponse `json:"error,omitempty"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// PromptpayCreateRequest defines request for creating a promptpay initiate transaction
type PromptpayCreateRequest struct {
	Input      string `json:"input" binding:"required"`
	CustomerId string `json:"customerId" binding:"required"`
}

// PromptpayCreateResponse defines response for creating a promptpay initiate transaction
type PromptpayCreateResponse struct {
	PromptpayId string `json:"promptpayId"`
	Name        string `json:"name"`
	ExpiredTs   int64  `json:"expiredTs"`
}

// PromptpayConfirmRequest defines request for creating a promptpay confirm transaction
type PromptpayConfirmRequest struct {
	PromptpayId string `json:"promptpayId" binding:"required"`
	Amount      int64  `json:"amount" binding:"required"`
}

// PromptpayConfirmResponse defines response for creating a promptpay confirm transaction
type PromptpayConfirmResponse struct {
	Message string `json:"message"`
}

func (route *Route) CreatePromptpay(c *gin.Context) {
	var request PromptpayCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: &ErrorResponse{
				Code:    "400",
				Message: "Bad request",
			}})
		return
	}

	// this is for mocking purpose
	fake := faker.New()

	response := PromptpayCreateResponse{
		PromptpayId: uuid.New().String(),
		Name:        fake.Person().Name(),
		ExpiredTs:   time.Now().Add(time.Second * time.Duration(route.TimeoutInSecond)).Unix(),
	}

	route.Cache.Set(response.PromptpayId, response.ExpiredTs)

	c.JSON(http.StatusOK, Response{
		Data: response,
	})
}

func (route *Route) ConfirmPromptpay(c *gin.Context) {
	var request PromptpayConfirmRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Error: &ErrorResponse{
				Code:    "400",
				Message: "Bad request",
			}})
		return
	}

	// validate amount must be greater than 100; last 2 digits are satang
	if request.Amount < 100 {
		c.JSON(http.StatusBadRequest, Response{
			Error: &ErrorResponse{
				Code:    "400",
				Message: "Amount must be greater than 100 (last 2 digits are satang)",
			}})
		return
	}

	ts, found := route.Cache.Get(request.PromptpayId)
	if !found {
		c.JSON(http.StatusBadRequest, Response{
			Error: &ErrorResponse{
				Code:    "400",
				Message: "PromptpayId not found",
			}})
		return
	}

	// check PromptpayId expired
	if time.Now().Unix()-ts.(int64) > route.TimeoutInSecond {
		c.JSON(http.StatusBadRequest, Response{
			Error: &ErrorResponse{
				Code:    "400",
				Message: "PromptpayId expired",
			}})
		return
	}

	// in real case, we should call bank API to confirm the transaction

	response := PromptpayConfirmResponse{
		Message: "Promptpay confirmed successfully",
	}

	c.JSON(http.StatusOK, response)
}
