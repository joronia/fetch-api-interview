package main

import (
	"fetch-api-interview/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	Controller := echo.New()

	Controller.GET("/receipts/:id/points", handlers.GetPointsForReceipt)

	Controller.POST("/receipts/process", handlers.ProcessReceipt)

	Controller.Logger.Fatal(Controller.Start(":8080"))
}
