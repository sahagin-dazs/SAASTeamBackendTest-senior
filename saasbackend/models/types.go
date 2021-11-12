package models

// ProductsResponse object containing a ProductResponse array
type ProductsResponse struct {
	Count    int64             `json:"count"`
	Products []ProductResponse `json:"products"`
}

// ProductResponse object containing the fields we want the API consumer to receive
type ProductResponse struct {
	ProductId    string `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductType  string `json:"product_type"`
	ProductPrice int64  `json:"product_price"`
}

// Product object including ProductType
type Product struct {
	ProductId            string `json:"product_id"`
	ProductName          string `json:"product_name"`
	ProductType          string `json:"product_type"`
	ProductPrice         int64  `json:"product_price"`
	ProductDiscountPrice int64  `json:"product_discount_price"`
	CouponCode           string `json:"coupon_code"`
}

// Request object containing cart details
type Cart struct {
	Cart []struct {
		ProductID  string `json:"product_id"`
		Quantity   int64  `json:"quantity"`
		CouponCode string `json:"coupon_code,omitempty"`
	} `json:"cart"`
}

// Calculated price response object
type CalculatePriceResponse struct {
	TotalObjects int64 `json:"total_objects"`
	TotalCost    int64 `json:"total_cost"`
}
