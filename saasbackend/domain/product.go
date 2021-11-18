package domain

import (
	"fmt"
	"saasteamtest/saasbackend/models"
	"strings"
)

type ProductHandler interface {
	Create(models.Product) (*models.Product, error)
	ReadOne(int) (*models.Product, error)
	Read() ([]*models.Product, error)
}

// NOTE - I modified the interface to have the new return type with the omitted fields for Get Product(s) and Calculate Price
type ProductServiceInterface interface {
	CalculatePrice(cart models.Cart) (*models.CalculatePriceResponse, error)
	GetAllProducts() (*models.ProductsResponse, error)
	GetProductById(int) (*models.ProductResponse, error)
	Save(models.Product) (*models.Product, error)
}

type ProductService struct {
	productHandler ProductHandler
}

func NewProductService(p1 ProductHandler) ProductServiceInterface {
	return ProductService{
		productHandler: p1,
	}
}

func (ps ProductService) CalculatePrice(cart models.Cart) (*models.CalculatePriceResponse, error) {
	// Set up our variables
	totalCost := int(0)
	totalObjects := int(0)

	// Iterate throught the cart slice
	for _, item := range cart.CartItems {
		// Make sure a proper quantity was provided
		if item.Quantity <= 0 {
			continue
		}

		// Retrieve a product if it's valid, let us know if it isn't
		product, err := ps.productHandler.ReadOne(item.ProductId)

		// If a product isn't valid, move on to the next item
		if (err != nil) || (product.ProductId == 0) {
			continue
		}

		// If there are valid matching coupon codes for a given product ID (case insensitive), give the discounted price, otherwise give the normal price
		if (product.CouponCode != "") && (strings.EqualFold(product.CouponCode, item.CouponCode)) {
			totalCost += (product.ProductDiscountPrice * item.Quantity)
		} else {
			totalCost += (product.ProductPrice * item.Quantity)
		}

		totalObjects += item.Quantity
	}

	// Create the price calculation response object
	calculatePriceResponse := models.CalculatePriceResponse{
		TotalObjects: totalObjects,
		TotalCost:    totalCost,
	}

	return &calculatePriceResponse, nil
}

func (ps ProductService) GetAllProducts() (*models.ProductsResponse, error) {
	myProducts, err := ps.productHandler.Read()
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}

	// Create product response slice
	productResponse := make([]models.ProductResponse, 0)

	// Iterate through all of the products
	for _, myProduct := range myProducts {
		// Map a single product to a product response object
		product := models.ProductResponse{
			ProductId:    myProduct.ProductId,
			ProductName:  myProduct.ProductName,
			ProductType:  myProduct.ProductType,
			ProductPrice: myProduct.ProductPrice,
		}

		// Append the product response object to the product response object slice
		productResponse = append(productResponse, product)
	}

	// Create the get all products response object
	productsResponse := models.ProductsResponse{
		Products: productResponse,
		Count:    int(len(myProducts)),
	}

	return &productsResponse, nil
}

func (ps ProductService) GetProductById(productId int) (*models.ProductResponse, error) {
	myProduct, err := ps.productHandler.ReadOne(productId)
	if err != nil {
		return nil, fmt.Errorf("read one: %w", err)
	}

	// Create the get product response object
	productResponse := models.ProductResponse{
		ProductId:    myProduct.ProductId,
		ProductName:  myProduct.ProductName,
		ProductType:  myProduct.ProductType,
		ProductPrice: myProduct.ProductPrice,
	}

	return &productResponse, nil
}

func (ps ProductService) Save(product models.Product) (*models.Product, error) {
	myProduct := models.Product{
		ProductName:          product.ProductName,
		ProductType:          product.ProductType,
		ProductPrice:         product.ProductPrice,
		ProductDiscountPrice: product.ProductDiscountPrice,
		CouponCode:           product.CouponCode,
	}

	savedProduct, err := ps.productHandler.Create(myProduct)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return savedProduct, nil
}
