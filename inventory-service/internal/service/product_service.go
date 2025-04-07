package service

import (
	"errors"
	"inventory-service/internal/repository"
)

func ReduceStock(productID, quantity int) error {
	product := repository.GetProductByID(productID)
	if product == nil {
		return errors.New("product not found")
	}

	if product.Stock < quantity {
		return errors.New("not enough stock")
	}

	product.Stock -= quantity
	repository.UpdateProduct(*product)
	return nil
}
func ApplyDiscount(productID int, discountPercentage float64) error {
	product := repository.GetProductByID(productID)
	if product == nil {
		return errors.New("product not found")
	}

	if discountPercentage < 0 || discountPercentage > 100 {
		return errors.New("invalid discount percentage")
	}

	product.Price = product.Price * (1 - discountPercentage/100)
	repository.UpdateProduct(*product)
	return nil
}
