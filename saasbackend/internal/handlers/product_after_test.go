//+build after

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

	newProduct := make(map[string]interface{})
	newProduct["product_name"] = "volleyball"
	newProduct["product_price"] = 750
	newProduct["product_type"] = "sporting_good"
	newProduct["product_discount_price"] = 525
	newProduct["coupon_code"] = "sport30"

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

	myProduct := make(map[string]interface{})
	myProduct["product_id"] = "2"
	myProduct["product_name"] = "burrito"
	myProduct["product_price"] = 700
	myProduct["product_type"] = "food"

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

	myProduct := make(map[string]interface{})
	myProduct["product_id"] = "3"
	myProduct["product_name"] = "basketball"
	myProduct["product_price"] = 1200
	myProduct["product_type"] = "sporting_good"

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

	prod1 := make(map[string]interface{})
	prod1["product_id"] = "1"
	prod1["product_name"] = "banana"
	prod1["product_type"] = "food"
	prod1["product_price"] = 500

	prod2 := make(map[string]interface{})
	prod2["product_id"] = "2"
	prod2["product_name"] = "burrito"
	prod2["product_type"] = "food"
	prod2["product_price"] = 700

	prod3 := make(map[string]interface{})
	prod3["product_id"] = "3"
	prod3["product_name"] = "basketball"
	prod3["product_type"] = "sporting_good"
	prod3["product_price"] = 1200

	prod4 := make(map[string]interface{})
	prod4["product_id"] = "4"
	prod4["product_name"] = "baseball"
	prod4["product_type"] = "sporting_good"
	prod4["product_price"] = 900

	var productSlice []map[string]interface{}
	productSlice = append(productSlice, prod1, prod2, prod3, prod4)

	myProductResult := make(map[string]interface{})
	myProductResult["count"] = 4
	myProductResult["products"] = productSlice

	expectedResult, err := json.Marshal(myProductResult)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice1(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := map[string]interface{}{}
	item1 := MSI{}
	item1["product_id"] = 1
	item1["quantity"] = 2
	item2 := MSI{}
	item2["product_id"] = 2
	item2["quantity"] = 2
	var shipmentItems1 []MSI
	shipmentItems1 = append(shipmentItems1, item1, item2)
	myPB["cart"] = shipmentItems1

	body, err := json.Marshal(myPB)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculate-price", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/calculate-price", BaseHandler(CalculatePrice(productService)))
	r.ServeHTTP(rec, req)

	myPrices := make(map[string]interface{})
	myPrices["total_objects"] = 4
	myPrices["total_cost"] = 2400

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice2(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := map[string]interface{}{}
	item1 := MSI{}
	item1["product_id"] = 1
	item1["quantity"] = 1
	item2 := MSI{}
	item2["product_id"] = 2
	item2["quantity"] = 1
	item3 := MSI{}
	item3["product_id"] = 3
	item3["quantity"] = 1
	item3["coupon_code"] = "sport30"
	item4 := MSI{}
	item4["product_id"] = 4
	item4["quantity"] = 1
	item4["coupon_code"] = "sport30"
	var shipmentItems1 []MSI
	shipmentItems1 = append(shipmentItems1, item1, item2, item3, item4)
	myPB["cart"] = shipmentItems1

	body, err := json.Marshal(myPB)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculate-price", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/calculate-price", BaseHandler(CalculatePrice(productService)))
	r.ServeHTTP(rec, req)

	myPrices := make(map[string]interface{})
	myPrices["total_objects"] = 4
	myPrices["total_cost"] = 2670

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice3(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := map[string]interface{}{}
	item1 := MSI{}
	item1["product_id"] = 1
	item1["quantity"] = 0
	item2 := MSI{}
	item2["product_id"] = 2
	item2["quantity"] = 0
	item3 := MSI{}
	item3["product_id"] = 3
	item3["quantity"] = 0
	item4 := MSI{}
	item4["product_id"] = 4
	item4["quantity"] = 0
	var shipmentItems1 []MSI
	shipmentItems1 = append(shipmentItems1, item1, item2, item3, item4)
	myPB["cart"] = shipmentItems1

	body, err := json.Marshal(myPB)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculate-price", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/calculate-price", BaseHandler(CalculatePrice(productService)))
	r.ServeHTTP(rec, req)

	myPrices := make(map[string]interface{})
	myPrices["total_objects"] = 0
	myPrices["total_cost"] = 0

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice4(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := map[string]interface{}{}
	item1 := MSI{}
	item1["product_id"] = 1
	item1["quantity"] = 100
	item2 := MSI{}
	item2["product_id"] = 2
	item2["quantity"] = 100
	item2["coupon_code"] = "food50"
	item3 := MSI{}
	item3["product_id"] = 3
	item3["quantity"] = 0
	item4 := MSI{}
	item4["product_id"] = 4
	item4["quantity"] = 0
	var shipmentItems1 []MSI
	shipmentItems1 = append(shipmentItems1, item1, item2, item3, item4)
	myPB["cart"] = shipmentItems1

	body, err := json.Marshal(myPB)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculate-price", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/calculate-price", BaseHandler(CalculatePrice(productService)))
	r.ServeHTTP(rec, req)

	myPrices := make(map[string]interface{})
	myPrices["total_objects"] = 200
	myPrices["total_cost"] = 85000

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}

func TestCalculatePrice5(t *testing.T) {

	// Create the object we will submit in the request body
	myPB := map[string]interface{}{}
	item1 := MSI{}
	item1["product_id"] = 1
	item1["quantity"] = 50
	item2 := MSI{}
	item2["product_id"] = 21
	item2["quantity"] = 50
	var shipmentItems1 []MSI
	shipmentItems1 = append(shipmentItems1, item1, item2)
	myPB["cart"] = shipmentItems1

	body, err := json.Marshal(myPB)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/calculate-price", bytes.NewReader(body))

	r := chi.NewRouter()
	r.Method("POST", "/calculate-price", BaseHandler(CalculatePrice(productService)))
	r.ServeHTTP(rec, req)

	myPrices := make(map[string]interface{})
	myPrices["total_objects"] = 50
	myPrices["total_cost"] = 25000

	expectedResult, err := json.Marshal(myPrices)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}
