package factory

import (
	"testing"
)

func TestConcreteCreator_CreateProduct(t *testing.T) {
	// Test case: Create a product
	t.Run("Create a product", func(t *testing.T) {
		creator := &ConcreteCreator{}
		product := creator.CreateProduct()

		if product == nil {
			t.Error("Expected a product, but got nil")
		}
	})

	// Test case: Create a product and use it
	t.Run("Create a product and use it", func(t *testing.T) {
		creator := &ConcreteCreator{}
		product := creator.CreateProduct()

		if product.Use() != "ConcreteProduct" {
			t.Errorf("Expected 'ConcreteProduct', but got %s", product.Use())
		}
	})
}
