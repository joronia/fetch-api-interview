package main

import (
	"encoding/json"
	"fetch-api-interview/handlers"
	"fetch-api-interview/responses"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockRecieptTarget              = `{"retailer":"Target","purchaseDate":"2022-01-01","purchaseTime":"13:01","items":[{"shortDescription":"Mountain Dew 12PK","price":"6.49"},{"shortDescription":"Emils Cheese Pizza","price":"12.25"},{"shortDescription":"Knorr Creamy Chicken","price":"1.26"},{"shortDescription":"Doritos Nacho Cheese","price":"3.35"},{"shortDescription":"   Klarbrunn 12-PK 12 FL OZ  ","price":"12.00"}],"total":"35.35"}`
	mockRecieptMM                  = `{"retailer":"M&M Corner Market","purchaseDate":"2022-03-20","purchaseTime":"14:33","items":[{"shortDescription":"Gatorade","price":"2.25"},{"shortDescription":"Gatorade","price":"2.25"},{"shortDescription":"Gatorade","price":"2.25"},{"shortDescription":"Gatorade","price":"2.25"}],"total":"9.00"}`
	jsonExpectedResultId           = `{"id":'12'}`
	jsonExpectedPointsResultTarget = `{"points":28}`
	jsonExpectedPointsResultMM     = `{"points":109}`
)

func TestProcessReceiptTarget(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/receipts/process", strings.NewReader(mockRecieptTarget))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	data := responses.ID{}

	// Assertions
	if assert.NoError(t, handlers.ProcessReceipt(c)) {
		json.Unmarshal([]byte(rec.Body.Bytes()), &data)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.IsType(t, jsonExpectedResultId, rec.Body.String())
	}

	ReqPoints := httptest.NewRequest(http.MethodGet, "/", nil)
	RecPoints := httptest.NewRecorder()
	Context := e.NewContext(ReqPoints, RecPoints)
	Context.SetPath("/receipts/:id/points")
	Context.SetParamNames("id")
	Context.SetParamValues(data.Id)

	if assert.NoError(t, handlers.GetPointsForReceipt(Context)) {
		assert.Equal(t, http.StatusOK, RecPoints.Code)
		assert.Equal(t, jsonExpectedPointsResultTarget, strings.ReplaceAll(RecPoints.Body.String(), "\n", ""))
	}

}

func TestProcessReceiptMM(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/receipts/process", strings.NewReader(mockRecieptMM))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	data := responses.ID{}

	// Assertions
	if assert.NoError(t, handlers.ProcessReceipt(c)) {
		json.Unmarshal([]byte(rec.Body.Bytes()), &data)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.IsType(t, jsonExpectedResultId, rec.Body.String())
	}

	ReqPoints := httptest.NewRequest(http.MethodGet, "/", nil)
	RecPoints := httptest.NewRecorder()
	Context := e.NewContext(ReqPoints, RecPoints)
	Context.SetPath("/receipts/:id/points")
	Context.SetParamNames("id")
	Context.SetParamValues(data.Id)

	if assert.NoError(t, handlers.GetPointsForReceipt(Context)) {
		assert.Equal(t, http.StatusOK, RecPoints.Code)
		assert.Equal(t, jsonExpectedPointsResultMM, strings.ReplaceAll(RecPoints.Body.String(), "\n", ""))
	}
}
