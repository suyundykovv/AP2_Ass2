package service

import (
	"errors"
	"inventory-service/internal/repository"
)

func ReduceStock(productID, quantity int) error {
	// Получение продукта по ID
	product := repository.GetProductByID(productID)
	if product == nil {
		return errors.New("product not found")
	}

	// Проверка доступного количества на складе
	if product.Stock < quantity {
		return errors.New("not enough stock")
	}

	// Уменьшение запаса
	product.Stock -= quantity

	// Обновление продукта в репозитории
	err := repository.UpdateProduct(productID, *product)
	if err != nil {
		return err
	}

	return nil
}
func ApplyDiscount(productID int, discountPercentage float64) error {
	// Получение продукта по ID
	product := repository.GetProductByID(productID)
	if product == nil {
		return errors.New("product not found")
	}

	// Проверка валидности процента скидки
	if discountPercentage < 0 || discountPercentage > 100 {
		return errors.New("invalid discount percentage")
	}

	// Применение скидки
	product.Price = product.Price * (1 - discountPercentage/100)

	// Обновление продукта в репозитории
	err := repository.UpdateProduct(productID, *product)
	if err != nil {
		return err
	}

	return nil
}
