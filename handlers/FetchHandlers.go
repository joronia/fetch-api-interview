package handlers

import (
	"fetch-api-interview/models"
	"fetch-api-interview/responses"
	"fetch-api-interview/utilities"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Maps to store the Receipt objects and also to store the calculated points with key being the id returned from random int generator.
var m = make(map[string]models.Receipt)
var points = make(map[string]int64)

/*
Handler called by controller to Get Total points for Receipt ID returned from Processed Receipt handler.
*/
func GetPointsForReceipt(c echo.Context) error {
	Id := c.Param("id")

	// Check if the id exists
	if Points, Exists := points[Id]; Exists {
		PointsJson := responses.Points{
			Points: Points,
		}

		return c.JSON(http.StatusOK, PointsJson)
	}

	return c.JSON(http.StatusNotFound, responses.ErrorExceptionMessage{Description: "No receipt found for that ID."})
}

/*
Handler called by controller to process receipt recieved from JSON to bind to receipt model returning a ID for it.
*/
func ProcessReceipt(c echo.Context) error {
	r := new(models.Receipt)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ErrorExceptionMessage{Description: "The receipt is invalid."})
	}

	ReceiptId := strconv.Itoa(rand.Int())
	TotalPoints := utilities.CalculateTotalPoints(*r)

	//Store in memory using Map for easy access.
	m[ReceiptId] = *r
	points[ReceiptId] = TotalPoints

	IdJson := responses.ID{
		Id: ReceiptId,
	}

	return c.JSON(http.StatusOK, IdJson)
}
