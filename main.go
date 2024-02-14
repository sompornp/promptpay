package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"somporn/promptpay/internal"
	"somporn/promptpay/internal/handler"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
		return
	}

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT_IN_SECOND"))
	route := handler.Route{
		Cache:           internal.NewCache(),
		TimeoutInSecond: int64(timeout),
	}

	r := gin.Default()

	r.POST("/createPromptpay", route.CreatePromptpay)
	r.POST("/confirmPromptpay", route.ConfirmPromptpay)

	r.Run()
}
