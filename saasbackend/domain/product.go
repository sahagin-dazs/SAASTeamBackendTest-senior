package domain

import (
	"fmt"
	"saasteamtest/saasbackend/models"
	"strings"
)

type ProductHandler interface {
	Create(models.Product) (*models.Product, error)
	ReadOne(string) (*models.Product, error)
	Read() ([]*models.Product, error)
}

// NOTE - I modified the interface to have the new return type with the omitted fields for Get Product(s) and Calculate Price
type ProductServiceInterface interface {
	CalculatePrice(cart models.Cart) (*models.CalculatePriceResponse, error)
	GetAllProducts() (*models.ProductsOmittedFields, error)
	GetProductById(string) (*models.ProductOmittedFields, error)
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
	totalCost := int64(0)
	totalObjects := int64(0)

	// Iterate throught the cart object
	for _, item := range cart.Cart {
		// Retrieve a product if it's valid, let us know if it isn't
		product, invalidProduct := ps.productHandler.ReadOne(item.ProductID)

		// If a product isn't valid or the quantity is invalid, remove it from the total objects and move on to the next item
		if (invalidProduct != nil) || (item.Quantity <= 0) {
			continue
		}

		// If there are valid matching coupon codes for a given product ID (case insensitive), give the discounted price, otherwise give the normal price
		if (len(item.CouponCode) != 0) && (len(product.CouponCode) != 0) && (strings.EqualFold(item.CouponCode, product.CouponCode)) {
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

func (ps ProductService) GetAllProducts() (*models.ProductsOmittedFields, error) {
	myProducts, err := ps.productHandler.Read()
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}

	// Define ProductOmittedFields object array
	var productsOmittedFields []models.ProductOmittedFields

	// Iterate through all products
	for _, product := range myProducts {

		// Create a ProductOmittedFields object and assign the current product to it
		productOmittedFields := models.ProductOmittedFields{
			Product: product,
		}

		// Append that object to the object array
		productsOmittedFields = append(productsOmittedFields, productOmittedFields)
	}

	// Create the get all products response with omitted fields object
	productsResponseOmittedFields := models.ProductsOmittedFields{
		Products: productsOmittedFields,
		Count:    int64(len(myProducts)),
	}

	return &productsResponseOmittedFields, nil
}

func (ps ProductService) GetProductById(productId string) (*models.ProductOmittedFields, error) {
	myProduct, err := ps.productHandler.ReadOne(productId)
	if err != nil {
		return nil, fmt.Errorf("read one: %w", err)
	}

	// Create the get product response with omitted fields object
	productResponseOmittedFields := models.ProductOmittedFields{
		Product: myProduct,
	}

	return &productResponseOmittedFields, nil
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
