package application_test

import (
	"testing"

	"github.com/codeedu/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Coca Cola 2L"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)
}
