//go:build after
// +build after

package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

// shortcut for creating all the JSON objects we will send and receive.
type MSI map[string]interface{}

func TestCreateProduct(t *testing.T) {

	newProduct := MSI{
		"product_name":           "volleyball",
		"product_price":          750,
		"product_type":           "sporting_good",
		"product_discount_price": 525,
		"coupon_code":            "sport30",
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

	newProduct["product_id"] = "5"

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

	myProduct := MSI{
		"product_id":    "2",
		"product_name":  "burrito",
		"product_price": 700,
		"product_type":  "food",
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

	myProduct := MSI{
		"product_id":    "3",
		"product_name":  "basketball",
		"product_price": 1200,
		"product_type":  "sporting_good",
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

	prod1 := MSI{
		"product_id":    "1",
		"product_name":  "banana",
		"product_type":  "food",
		"product_price": 500,
	}

	prod2 := MSI{
		"product_id":    "2",
		"product_name":  "burrito",
		"product_type":  "food",
		"product_price": 700,
	}

	prod3 := MSI{
		"product_id":    "3",
		"product_name":  "basketball",
		"product_type":  "sporting_good",
		"product_price": 1200,
	}

	prod4 := MSI{
		"product_id":    "4",
		"product_name":  "baseball",
		"product_type":  "sporting_good",
		"product_price": 900,
	}

	var productSlice []MSI
	productSlice = append(productSlice, prod1, prod2, prod3, prod4)

	myProductResult := MSI{
		"count":    4,
		"products": productSlice,
	}

	expectedResult, err := json.Marshal(myProductResult)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice1(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := MSI{
		"cart": []interface{}{
			MSI{
				"product_id": "1",
				"quantity":   2,
			},
			MSI{
				"product_id": "2",
				"quantity":   2,
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

	myPrices := MSI{
		"total_objects": 4,
		"total_cost":    2400,
	}

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice2(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := MSI{
		"cart": []interface{}{
			MSI{
				"product_id": "1",
				"quantity":   1,
			},
			MSI{
				"product_id": "2",
				"quantity":   1,
			},
			MSI{
				"product_id":  "3",
				"quantity":    1,
				"coupon_code": "sport30",
			},
			MSI{
				"product_id":  "4",
				"quantity":    1,
				"coupon_code": "sport30",
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

	myPrices := MSI{
		"total_objects": 4,
		"total_cost":    2670,
	}

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice3(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := MSI{
		"cart": []interface{}{
			MSI{
				"product_id": "1",
				"quantity":   0,
			},
			MSI{
				"product_id": "2",
				"quantity":   0,
			},
			MSI{
				"product_id": "3",
				"quantity":   0,
			},
			MSI{
				"product_id": "4",
				"quantity":   0,
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

	myPrices := MSI{
		"total_objects": 0,
		"total_cost":    0,
	}

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice4(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := MSI{
		"cart": []interface{}{
			MSI{
				"product_id": "1",
				"quantity":   100,
			},
			MSI{
				"product_id":  "2",
				"quantity":    100,
				"coupon_code": "food50",
			},
			MSI{
				"product_id": "3",
				"quantity":   0,
			},
			MSI{
				"product_id": "4",
				"quantity":   0,
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

	myPrices := MSI{
		"total_objects": 200,
		"total_cost":    85000,
	}

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice5(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := MSI{
		"cart": []interface{}{
			MSI{
				"product_id": "1",
				"quantity":   50,
			},
			MSI{
				"product_id":  "21",
				"quantity":    50,
				"coupon_code": "food50",
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

	myPrices := MSI{
		"total_objects": 50,
		"total_cost":    25000,
	}

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}
