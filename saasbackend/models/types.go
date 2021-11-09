package models

type ProductsResponse struct {
	Count    int64      `json:"count"`
	Products []*Product `json:"products"`
}

type Product struct {
	ProductId            string `json:"product_id"`
	ProductName          string `json:"product_name"`
	ProductPrice         int64  `json:"product_price"`
	ProductDiscountPrice int64  `json:"product_discount_price"`
	CouponCode           string `json:"coupon_code"`
}
