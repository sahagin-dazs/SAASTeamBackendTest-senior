package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"saasteamtest/saasbackend/domain"
	"saasteamtest/saasbackend/models"

	"github.com/go-chi/chi"
)

// CreateProduct creates a new product.
// POST /products
func CreateProduct(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("readAll: %w", err)
		}

		var product models.Product
		err = json.Unmarshal(body, &product)
		if err != nil {
			// Return an HTTP 400 if the request payload is malformed
			return Respond(w, "Payload is malformed", 400)
		}

		savedProduct, err := productService.Save(product)

		if err != nil {
			return fmt.Errorf("save: %w", err)
		}
		return RespondOK(w, savedProduct)
	}
}

// GetProductById gets a single product by its ID.
// GET /products/{product_id}
func GetProductById(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		productId := chi.URLParam(r, "product_id")

		product, err := productService.GetProductById(productId)
		if err != nil {
			// Return an HTTP 404 if the requested product is not found
			return Respond(w, "Product not found", 404)
		}

		return RespondOK(w, product)
	}
}

// GetAllProducts gets all of the products.
// GET /products
func GetAllProducts(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		products, err := productService.GetAllProducts()
		if err != nil {
			return fmt.Errorf("get all products: %w", err)
		}

		return RespondOK(w, products)
	}
}

// CalculatePrice calculated the price for the entire cart.
// POST /calculate-price
func CalculatePrice(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		// Get the JSON body from the inbound request
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("readAll: %w", err)
		}

		// Create new Cart object and unmarshal the JSON body into the object
		var cart models.Cart
		err = json.Unmarshal(body, &cart)
		if err != nil {
			// Return an HTTP 400 if the request payload is malformed
			return Respond(w, "Payload is malformed", 400)
		}

		// Call the CalculatePrice service
		calculatePrice, err := productService.CalculatePrice(cart)

		if err != nil {
			return fmt.Errorf("get total cost of cart: %w", err)
		}

		return RespondOK(w, calculatePrice)
	}
}
