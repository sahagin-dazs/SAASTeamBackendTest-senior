//+build before

package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

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
	myProduct["product_discount_price"] = 350
	myProduct["coupon_code"] = "food50"

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
	myProduct["product_discount_price"] = 840
	myProduct["coupon_code"] = "sport30"

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
	prod1["product_price"] = 500
	prod1["product_discount_price"] = 250
	prod1["coupon_code"] = "food50"

	prod2 := make(map[string]interface{})
	prod2["product_id"] = "2"
	prod2["product_name"] = "burrito"
	prod2["product_price"] = 700
	prod2["product_discount_price"] = 350
	prod2["coupon_code"] = "food50"

	prod3 := make(map[string]interface{})
	prod3["product_id"] = "3"
	prod3["product_name"] = "basketball"
	prod3["product_price"] = 1200
	prod3["product_discount_price"] = 840
	prod3["coupon_code"] = "sport30"

	prod4 := make(map[string]interface{})
	prod4["product_id"] = "4"
	prod4["product_name"] = "baseball"
	prod4["product_price"] = 900
	prod4["product_discount_price"] = 630
	prod4["coupon_code"] = "sport30"

	var productSlice []map[string]interface{}
	productSlice = append(productSlice, prod1, prod2, prod3, prod4)

	productResult := make(map[string]interface{})

	productResult["count"] = 4
	productResult["products"] = productSlice

	expectedResult, err := json.Marshal(productResult)
	if err != nil {
		t.Error(err)
	}

	testHelper.MapJSONBodyIsEqualString(t, rec.Body.String(), string(expectedResult))
}
