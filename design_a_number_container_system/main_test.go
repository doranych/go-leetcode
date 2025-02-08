package design_a_number_container_system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("", func(t *testing.T) {
		nc := Constructor()
		assert.Equal(t, nc.Find(10), -1)
		nc.Change(2, 10)
		nc.Change(1, 10)
		nc.Change(3, 10)
		nc.Change(5, 10)
		assert.Equal(t, nc.Find(10), 1)
		nc.Change(1, 20)
		assert.Equal(t, nc.Find(10), 2)
	})
	t.Run("", func(t *testing.T) {
		nc := Constructor()
		nc.Change(1, 10)
		nc.Change(1, 10)
		assert.Equal(t, nc.Find(10), 1)
		nc.Change(1, 20)
		assert.Equal(t, nc.Find(10), -1)
	})
}
