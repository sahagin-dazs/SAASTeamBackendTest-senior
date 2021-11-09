package data

import (
	"errors"
	"saasteamtest/saasbackend/models"
)

type ProductHandle struct{}

func NewProductHandler() *ProductHandle {
	productHandle := ProductHandle{}
	return &productHandle
}

func (h *ProductHandle) Create(obj models.Product) (*models.Product, error) {
	obj.ProductId = "5"
	return &obj, nil
}

func (h *ProductHandle) ReadOne(q string) (*models.Product, error) {
	switch q {
	case "1":
		item := models.Product{ProductId: "1", ProductName: "banana", ProductPrice: 500, ProductDiscountPrice: 250, CouponCode: "food50"}
		return &item, nil
	case "2":
		item := models.Product{ProductId: "2", ProductName: "burrito", ProductPrice: 700, ProductDiscountPrice: 350, CouponCode: "food50"}
		return &item, nil
	case "3":
		item := models.Product{ProductId: "3", ProductName: "basketball", ProductPrice: 1200, ProductDiscountPrice: 840, CouponCode: "sport30"}
		return &item, nil
	case "4":
		item := models.Product{ProductId: "4", ProductName: "baseball", ProductPrice: 900, ProductDiscountPrice: 630, CouponCode: "sport30"}
		return &item, nil
	default:
		return nil, errors.New("no such product found")
	}
}

func (h *ProductHandle) Read() ([]*models.Product, error) {
	items := make([]*models.Product, 0)
	item1 := models.Product{ProductId: "1", ProductName: "banana", ProductPrice: 500, ProductDiscountPrice: 250, CouponCode: "food50"}
	items = append(items, &item1)
	item2 := models.Product{ProductId: "2", ProductName: "burrito", ProductPrice: 700, ProductDiscountPrice: 350, CouponCode: "food50"}
	items = append(items, &item2)
	item3 := models.Product{ProductId: "3", ProductName: "basketball", ProductPrice: 1200, ProductDiscountPrice: 840, CouponCode: "sport30"}
	items = append(items, &item3)
	item4 := models.Product{ProductId: "4", ProductName: "baseball", ProductPrice: 900, ProductDiscountPrice: 630, CouponCode: "sport30"}
	items = append(items, &item4)
	return items, nil
}
