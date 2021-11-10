package models

// Create a struct pointer that will never be set
type omit *struct{}

// Changed the former struct to support the omitted fields
type ProductsResponseOmittedFields struct {
	Count    int64                  `json:"count"`
	Products []ProductOmittedFields `json:"products"`
}

// Wrap Product struct in a new struct and include fields to omit, specifying omitonempty to prevent the inevitably nil fields from projecting in the JSON object
// NOTE: I could have also done this by specifying the datatype as *struct instead of 'omit', but I wanted it to be clear what was going on
type ProductOmittedFields struct {
	*Product
	ProductDiscountPrice omit `json:"product_discount_price,omitempty"`
	CouponCode           omit `json:"coupon_code,omitempty"`
}

type Product struct {
	ProductId            string `json:"product_id"`
	ProductName          string `json:"product_name"`
	ProductType          string `json:"product_type"`
	ProductPrice         int64  `json:"product_price"`
	ProductDiscountPrice int64  `json:"product_discount_price"`
	CouponCode           string `json:"coupon_code"`
}
