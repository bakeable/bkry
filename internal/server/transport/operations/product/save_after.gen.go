package product_operations

import (
	product "github.com/bakeable/bkry/internal/server/models/entities/product"
)

func afterSave(entity product.Product, editorID *string) {}