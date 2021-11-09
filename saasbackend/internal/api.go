package internal

import (
	"net/http"
	"saasteamtest/saasbackend/data"
	"saasteamtest/saasbackend/domain"
	"saasteamtest/saasbackend/internal/handlers"

	"github.com/go-chi/chi"
)

func RouterInitializer() *chi.Mux {
	productHandler := data.NewProductHandler()
	productService := domain.NewProductService(productHandler)

	r := chi.NewRouter()

	r.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
	})

	r.Route("/products", func(r chi.Router) {
		r.Method("GET", "/", handlers.BaseHandler(handlers.GetAllProducts(productService)))
		r.Method("POST", "/", handlers.BaseHandler(handlers.CreateProduct(productService)))
		r.Route("/{product_id}", func(r chi.Router) {
			r.Method("GET", "/", handlers.BaseHandler(handlers.GetProductById(productService)))
		})
	})
	r.Route("/calculate-price", func(r chi.Router) {
		r.Method("POST", "/", handlers.BaseHandler(handlers.CalculatePrice(productService)))
	})

	return r
}
