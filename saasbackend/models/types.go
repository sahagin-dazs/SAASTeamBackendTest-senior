package models

// ProductsResponse object containing a ProductResponse array
type ProductsResponse struct {
	Count    int               `json:"count"`
	Products []ProductResponse `json:"products"`
}

// ProductResponse object containing the fields we want the API consumer to receive
type ProductResponse struct {
	ProductId    int    `json:"product_id,omitempty"`
	ProductName  string `json:"product_name,omitempty"`
	ProductType  string `json:"product_type,omitempty"`
	ProductPrice int    `json:"product_price,omitempty"`
}

// Product object including ProductType
type Product struct {
	ProductId            int    `json:"product_id"`
	ProductName          string `json:"product_name"`
	ProductType          string `json:"product_type"`
	ProductPrice         int    `json:"product_price"`
	ProductDiscountPrice int    `json:"product_discount_price"`
	CouponCode           string `json:"coupon_code"`
}

// Request object containing slice of CartItems
type Cart struct {
	CartItems []CartItem `json:"cart"`
}

// A specific cart item
type CartItem struct {
	ProductId  int    `json:"product_id"`
	Quantity   int    `json:"quantity"`
	CouponCode string `json:"coupon_code,omitempty"`
}

// Calculated price response object
type CalculatePriceResponse struct {
	TotalObjects int `json:"total_objects"`
	TotalCost    int `json:"total_cost"`
}
