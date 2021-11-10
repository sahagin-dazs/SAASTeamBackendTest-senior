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

// send a whole YED shipment to this endpoint, return the vendor
func CreateProduct(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("readAll: %w", err)
		}

		var product models.Product
		err = json.Unmarshal(body, &product)
		if err != nil {
			return fmt.Errorf("unmarshal: %w", err)
		}

		savedProduct, err := productService.Save(product)

		if err != nil {
			return fmt.Errorf("save: %w", err)
		}
		return RespondOK(w, savedProduct)
	}
}

func GetProductById(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		productId := chi.URLParam(r, "product_id")

		product, err := productService.GetProductById(productId)
		if err != nil {
			return fmt.Errorf("get product by id: %w", err)
		}

		// Create a new object with the ProductOmittedFields struct to exclude the values from the returned payload
		productOmittedFields := models.ProductOmittedFields{
			Product: product,
		}

		return RespondOK(w, productOmittedFields)
	}
}

func GetAllProducts(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		products, err := productService.GetAllProducts()
		if err != nil {
			return fmt.Errorf("get all products: %w", err)
		}

		// Define empty object array
		var productsOmittedFields []models.ProductOmittedFields

		// Loop through all of the products we retrieved
		for _, product := range products {

			// Create a ProductOmittedFields object with the product value
			productOmittedFields := models.ProductOmittedFields{
				Product: product,
			}

			// Append that object to the object array
			productsOmittedFields = append(productsOmittedFields, productOmittedFields)
		}

		productsResponseOmittedFields := models.ProductsResponseOmittedFields{
			Products: productsOmittedFields,
			Count:    int64(len(products)),
		}

		return RespondOK(w, productsResponseOmittedFields)
	}
}

func CalculatePrice(productService domain.ProductServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}
