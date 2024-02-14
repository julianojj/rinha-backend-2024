package main

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/rinha-backend-2024/internal/core/services"
	"github.com/julianojj/rinha-backend-2024/internal/infra/api/middlewares"
	"github.com/julianojj/rinha-backend-2024/internal/infra/repository"
)

func main() {
	r := gin.Default()

	db, err := sql.Open("postgres", os.Getenv("DB_URI"))
	if err != nil {
		panic(err)
	}

	maxConnections, err := strconv.Atoi(os.Getenv("MAX_CONNECTIONS"))
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxConnections)
	db.SetMaxIdleConns(maxConnections)
	db.SetConnMaxIdleTime(5 * time.Minute)

	customerRepository := repository.NewCustomerRepositoryDatabase(db)
	transactionRepository := repository.NewTransactionRepositoryDatabase(db)
	transactionService := services.NewTransactionService(customerRepository, transactionRepository)

	r.Use(middlewares.ErrorHandler())

	r.POST("/clientes/:id/transacoes", func(ctx *gin.Context) {
		var input *services.ProcessTransactionInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(422, err)
			return
		}
		customerId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(422, err)
			return
		}
		input.CustomerId = int64(customerId)
		output, err := transactionService.ProcessTransaction(input)
		if err == nil {
			ctx.JSON(200, output)
			return
		}
		ctx.Error(err)
	})

	r.GET("/clientes/:id/extrato", func(ctx *gin.Context) {
		customerId, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(422, err)
			return
		}
		output, err := transactionService.RequestStatement(int64(customerId))
		if err == nil {
			ctx.JSON(200, output)
			return
		}
		ctx.Error(err)
	})

	r.Run(":8080")
}
