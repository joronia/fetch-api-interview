package utilities

import (
	"fetch-api-interview/models"
	"math"
	"strings"
	"time"
	"unicode"
)

/*
Main funtion used to calculate the receipt total points from the Receipt model from controller.
*/
func CalculateTotalPoints(r models.Receipt) int64 {

	var TotalPoints int64
	Retailer := r.Retailer

	//Count the alphanumeric characters in the retailer name
	for _, c := range Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			TotalPoints++
		}
	}

	//Is total a round dollar amount with no cents?
	if r.Total == float64(int(r.Total)) {
		TotalPoints += 50
	}

	//Is Total a multiple of 0.25
	if math.Mod(r.Total, 0.25) == 0 {
		TotalPoints += 25
	}

	// Add 5 points for every two items on the receipt.
	TotalPoints += int64((len(r.Items) / 2) * 5)

	/*If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the
	nearest integer. The result is the number of points earned.*/
	for _, item := range r.Items {
		NewDescription := strings.TrimSpace(item.ShortDescription)
		LengthOfNewDescription := len(NewDescription)

		if LengthOfNewDescription%3 == 0 {
			TotalPoints += int64(math.Ceil(item.Price * 0.2))
		}
	}

	//6 points if the day in the purchase date is odd.
	t, err := time.Parse("2006-01-02", r.PurchaseDate)

	if err == nil {
		//lets get the day and check if it is an odd day
		day := t.Day()

		if (day % 2) != 0 {
			TotalPoints += 6
		}
	}

	//10 points if the time of purchase is after 2:00pm and before 4:00pm
	h, error := time.Parse("15:04", r.PurchaseTime)

	if error == nil {
		// Check if the hour is between 2 and 4 PM
		Hour := h.Hour()
		if Hour >= 14 && Hour < 16 {
			TotalPoints += 10
		}
	}

	return TotalPoints
}
