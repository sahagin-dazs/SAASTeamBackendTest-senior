package data

import (
	"fmt"
	"os"
	"saasteamtest/saasbackend/models"

	"github.com/hashicorp/go-memdb"
)

type ProductHandle struct {
	db      *memdb.MemDB
	err     error
	counter int
}

func NewProductHandler() *ProductHandle {
	productHandle := ProductHandle{}

	// Initialize database
	productHandle.counter = 5
	productHandle.db, productHandle.err = productHandle.startDatabase()
	if productHandle.err != nil {
		fmt.Println("failure to start database:", productHandle.err)
		os.Exit(1)
	}
	return &productHandle
}

func (h *ProductHandle) startDatabase() (*memdb.MemDB, error) {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"product": {
				Name: "product",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:         "id",
						Unique:       true,
						AllowMissing: false,
						Indexer:      &memdb.IntFieldIndex{Field: "ProductId"},
					},
					"product_name": {
						Name:         "product_name",
						Unique:       false,
						AllowMissing: false,
						Indexer:      &memdb.StringFieldIndex{Field: "ProductName"},
					},
					"product_type": {
						Name:         "product_type",
						Unique:       false,
						AllowMissing: false,
						Indexer:      &memdb.StringFieldIndex{Field: "ProductType"},
					},
					"product_price": {
						Name:         "product_price",
						Unique:       false,
						AllowMissing: false,
						Indexer:      &memdb.IntFieldIndex{Field: "ProductPrice"},
					},
					"product_discount_price": {
						Name:         "product_discount_price",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.IntFieldIndex{Field: "ProductDiscountPrice"},
					},
					"coupon_code": {
						Name:         "coupon_code",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "CouponCode"},
					},
				},
			},
		},
	}

	// Create a new database
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, fmt.Errorf("create database: %w", err)
	}

	// Create a write transaction
	txn := db.Txn(true)

	// Prepare test data
	items := make([]*models.Product, 0)
	item1 := models.Product{ProductId: 1, ProductName: "banana", ProductType: "food", ProductPrice: 500, ProductDiscountPrice: 250, CouponCode: "food50"}
	items = append(items, &item1)
	item2 := models.Product{ProductId: 2, ProductName: "burrito", ProductType: "food", ProductPrice: 700, ProductDiscountPrice: 350, CouponCode: "food50"}
	items = append(items, &item2)
	item3 := models.Product{ProductId: 3, ProductName: "basketball", ProductType: "sporting_good", ProductPrice: 1200, ProductDiscountPrice: 840, CouponCode: "sport30"}
	items = append(items, &item3)
	item4 := models.Product{ProductId: 4, ProductName: "baseball", ProductType: "sporting_good", ProductPrice: 900, ProductDiscountPrice: 630, CouponCode: "sport30"}
	items = append(items, &item4)

	// Insert the test data
	for _, p := range items {
		if err := txn.Insert("product", p); err != nil {
			return nil, fmt.Errorf("insert test data: %w", err)
		}
	}

	// Commit the transaction
	txn.Commit()

	return db, nil
}

func (h *ProductHandle) Create(obj models.Product) (*models.Product, error) {
	// Create write transaction
	var txn = h.db.Txn(true)

	// Assign the autonumber for product_id
	obj.ProductId = h.counter
	h.counter++

	// Insert into the database
	if err := txn.Insert("product", &obj); err != nil {
		return nil, fmt.Errorf("insert record: %w", err)
	}

	// Commit the transaction
	txn.Commit()

	return &obj, nil
}

func (h *ProductHandle) ReadOne(q int) (*models.Product, error) {
	// Create read-only transaction
	var txn = h.db.Txn(false)
	defer txn.Abort()

	// Get single product by unique id
	myProduct, err := txn.First("product", "id", q)
	if err != nil {
		return nil, fmt.Errorf("get product: %w", err)
	}

	// Attempt to cast the db result to the Product model
	product, ok := myProduct.(*models.Product)
	if !ok {
		return nil, fmt.Errorf("error casting db result as product")
	}

	return product, nil
}

func (h *ProductHandle) Read() ([]*models.Product, error) {
	// Create read-only transaction
	var txn = h.db.Txn(false)
	defer txn.Abort()

	// Get all products
	product, err := txn.Get("product", "id")
	if err != nil {
		return nil, fmt.Errorf("get products: %w", err)
	}

	// Create empty Product slice
	var products []*models.Product

	// Iterate through records and add to products slice
	for obj := product.Next(); obj != nil; obj = product.Next() {
		products = append(products, obj.(*models.Product))
	}

	return products, nil
}
