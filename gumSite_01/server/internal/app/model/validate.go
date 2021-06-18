package model

func ValidateProduct(product *Product) bool {
	if product.Cost < 0 {
		return false
	}
	if product.Name == "" {
		return false
	}
	if product.Description == "" {
		return false
	}
	return true
}
