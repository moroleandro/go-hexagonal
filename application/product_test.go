package application_test

import (
	"testing"

	"github.com/moroleandro/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Coca Cola 2L"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}
