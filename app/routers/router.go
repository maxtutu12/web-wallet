package routers

import (
	"web-wallet/app/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()

	r.POST("/api/login", api.Login)

	account := r.Group("/api/account")
	{
		account.POST("/create", api.CreateAccount)
		account.POST("/private_key", api.DerivePrivateKey)
		account.GET("/balance", api.GetBalance)
	}

	transaction := r.Group("/api/transaction")
	{
		transaction.POST("send", api.SendTransaction)
		transaction.GET("/gas_price", api.GetGasPrice)
	}

	return r
}
