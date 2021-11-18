//go:build after
// +build after

package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"saasteamtest/saasbackend/models"
	"testing"

	"github.com/go-chi/chi"
)

func TestCreateProduct(t *testing.T) {

	newProduct := models.Product{
		ProductName:          "volleyball",
		ProductType:          "sporting_good",
		ProductPrice:         750,
		ProductDiscountPrice: 525,
		CouponCode:           "sport30",
	}

	body, err := json.Marshal(newProduct)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/products", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/products", BaseHandler(CreateProduct(productService)))
	r.ServeHTTP(rec, req)

	newProduct.ProductId = 5

	expectedResult, err := json.Marshal(newProduct)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestGetProductById2(t *testing.T) {

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/2", nil)

	r := chi.NewRouter()
	r.Method("GET", "/products/{product_id}", BaseHandler(GetProductById(productService)))
	r.ServeHTTP(rec, req)

	myProduct := models.ProductResponse{
		ProductId:    2,
		ProductName:  "burrito",
		ProductType:  "food",
		ProductPrice: 700,
	}

	expectedResult, err := json.Marshal(myProduct)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestGetProductById3(t *testing.T) {

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/3", nil)

	r := chi.NewRouter()
	r.Method("GET", "/products/{product_id}", BaseHandler(GetProductById(productService)))
	r.ServeHTTP(rec, req)

	myProduct := models.ProductResponse{
		ProductId:    3,
		ProductName:  "basketball",
		ProductType:  "sporting_good",
		ProductPrice: 1200,
	}

	expectedResult, err := json.Marshal(myProduct)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestGetAllProducts(t *testing.T) {

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)

	r := chi.NewRouter()
	r.Method("GET", "/products", BaseHandler(GetAllProducts(productService)))
	r.ServeHTTP(rec, req)

	prod1 := models.ProductResponse{
		ProductId:    1,
		ProductName:  "banana",
		ProductType:  "food",
		ProductPrice: 500,
	}

	prod2 := models.ProductResponse{
		ProductId:    2,
		ProductName:  "burrito",
		ProductType:  "food",
		ProductPrice: 700,
	}

	prod3 := models.ProductResponse{
		ProductId:    3,
		ProductName:  "basketball",
		ProductType:  "sporting_good",
		ProductPrice: 1200,
	}

	prod4 := models.ProductResponse{
		ProductId:    4,
		ProductName:  "baseball",
		ProductType:  "sporting_good",
		ProductPrice: 900,
	}

	prod5 := models.ProductResponse{
		ProductId:    5,
		ProductName:  "volleyball",
		ProductType:  "sporting_good",
		ProductPrice: 750,
	}

	var productSlice []models.ProductResponse
	productSlice = append(productSlice, prod1, prod2, prod3, prod4, prod5)

	myProductResult := models.ProductsResponse{
		Count:    5,
		Products: productSlice,
	}

	expectedResult, err := json.Marshal(myProductResult)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice1(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := models.Cart{
		CartItems: []models.CartItem{
			{
				ProductId: 1,
				Quantity:  2,
			},
			{
				ProductId: 2,
				Quantity:  2,
			},
		},
	}

	body, err := json.Marshal(myPB)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculate-price", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/calculate-price", BaseHandler(CalculatePrice(productService)))
	r.ServeHTTP(rec, req)

	myPrices := models.CalculatePriceResponse{
		TotalObjects: 4,
		TotalCost:    2400,
	}

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice2(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := models.Cart{
		CartItems: []models.CartItem{
			{
				ProductId: 1,
				Quantity:  1,
			},
			{
				ProductId: 2,
				Quantity:  1,
			},
			{
				ProductId:  3,
				Quantity:   1,
				CouponCode: "sport30",
			},
			{
				ProductId:  4,
				Quantity:   1,
				CouponCode: "sport30",
			},
		},
	}

	body, err := json.Marshal(myPB)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculate-price", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/calculate-price", BaseHandler(CalculatePrice(productService)))
	r.ServeHTTP(rec, req)

	myPrices := models.CalculatePriceResponse{
		TotalObjects: 4,
		TotalCost:    2670,
	}

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice3(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := models.Cart{
		CartItems: []models.CartItem{
			{
				ProductId: 1,
				Quantity:  0,
			},
			{
				ProductId: 2,
				Quantity:  0,
			},
			{
				ProductId: 3,
				Quantity:  0,
			},
			{
				ProductId: 4,
				Quantity:  0,
			},
		},
	}

	body, err := json.Marshal(myPB)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculate-price", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/calculate-price", BaseHandler(CalculatePrice(productService)))
	r.ServeHTTP(rec, req)

	myPrices := models.CalculatePriceResponse{
		TotalObjects: 0,
		TotalCost:    0,
	}

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice4(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := models.Cart{
		CartItems: []models.CartItem{
			{
				ProductId: 1,
				Quantity:  100,
			},
			{
				ProductId:  2,
				Quantity:   100,
				CouponCode: "food50",
			},
			{
				ProductId: 3,
				Quantity:  0,
			},
			{
				ProductId: 4,
				Quantity:  0,
			},
		},
	}

	body, err := json.Marshal(myPB)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculate-price", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/calculate-price", BaseHandler(CalculatePrice(productService)))
	r.ServeHTTP(rec, req)

	myPrices := models.CalculatePriceResponse{
		TotalObjects: 200,
		TotalCost:    85000,
	}

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice5(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := models.Cart{
		CartItems: []models.CartItem{
			{
				ProductId: 1,
				Quantity:  50,
			},
			{
				ProductId:  21,
				Quantity:   50,
				CouponCode: "food50",
			},
		},
	}

	body, err := json.Marshal(myPB)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculate-price", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/calculate-price", BaseHandler(CalculatePrice(productService)))
	r.ServeHTTP(rec, req)

	myPrices := models.CalculatePriceResponse{
		TotalObjects: 50,
		TotalCost:    25000,
	}

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}
