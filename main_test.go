package main // replace with your package name

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"somporn/promptpay/internal"
	"somporn/promptpay/internal/handler"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Set(key string, value interface{}) {
	m.Called(key, value)
}

func (m *MockCache) Get(key string) (interface{}, bool) {
	args := m.Called(key)
	return args.Get(0), args.Bool(1)
}

func TestCreatePromptpay(t *testing.T) {
	// Setting up the router
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockCache := new(MockCache)
	route := &handler.Route{Cache: internal.NewCache(), TimeoutInSecond: 10}
	r.POST("/createPromptpay", route.CreatePromptpay)

	// Mocking Cache.Set
	mockCache.On("Set", mock.Anything, mock.Anything).Return()

	// Creating a request body
	requestBody, _ := json.Marshal(handler.PromptpayCreateRequest{
		Input:      "input data",
		CustomerId: "customer123",
	})

	req, _ := http.NewRequest(http.MethodPost, "/createPromptpay", bytes.NewBuffer(requestBody))
	resp := httptest.NewRecorder()

	// Perform the test
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	// Additional assertions can be added here, like checking the response body
}

func TestConfirmPromptpay(t *testing.T) {
	// Similar setup to TestCreatePromptpay
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockCache := new(MockCache)
	route := &handler.Route{Cache: internal.NewCache(), TimeoutInSecond: 10}
	r.POST("/confirmPromptpay", route.ConfirmPromptpay)

	// Mocking Cache.Get to simulate finding a promptpay ID
	mockCache.On("Get", "existing-id").Return(time.Now().Unix(), true)

	// Creating a request body
	requestBody, _ := json.Marshal(handler.PromptpayConfirmRequest{
		PromptpayId: "existing-id",
		Amount:      100,
	})

	req, _ := http.NewRequest(http.MethodPost, "/confirmPromptpay", bytes.NewBuffer(requestBody))
	resp := httptest.NewRecorder()

	// Perform the test
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Additional assertions can be added here
}
