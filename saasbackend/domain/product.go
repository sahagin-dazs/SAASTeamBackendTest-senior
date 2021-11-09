package domain

import (
	"fmt"
	"saasteamtest/saasbackend/models"
)

type ProductHandler interface {
	Create(models.Product) (*models.Product, error)
	ReadOne(string) (*models.Product, error)
	Read() ([]*models.Product, error)
}

type ProductServiceInterface interface {
	CalculatePrice() (*int64, error)
	GetAllProducts() ([]*models.Product, error)
	GetProductById(string) (*models.Product, error)
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

func (ps ProductService) CalculatePrice() (*int64, error) {
	cost := int64(0)
	return &cost, nil
}

func (ps ProductService) GetAllProducts() ([]*models.Product, error) {
	myProducts, err := ps.productHandler.Read()
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	return myProducts, nil
}

func (ps ProductService) GetProductById(productId string) (*models.Product, error) {
	myProduct, err := ps.productHandler.ReadOne(productId)
	if err != nil {
		return nil, fmt.Errorf("read one: %w", err)
	}
	return myProduct, nil
}

func (ps ProductService) Save(product models.Product) (*models.Product, error) {
	myProduct := models.Product{
		ProductName:          product.ProductName,
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
