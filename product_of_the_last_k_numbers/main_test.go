package product_of_the_last_k_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductOfNumbers(t *testing.T) {
	obj := Constructor()
	obj.Add(3)
	obj.Add(0)
	obj.Add(2)
	obj.Add(5)
	obj.Add(4)
	assert.Equal(t, 20, obj.GetProduct(2))
	assert.Equal(t, 40, obj.GetProduct(3))
	assert.Equal(t, 0, obj.GetProduct(4))
	obj.Add(8)
	assert.Equal(t, 32, obj.GetProduct(2))
}
