package ejercicios

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasNextPosicionesPares(t *testing.T) {
	array := NewArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	iterator := array.PosicionesParesIterator()

	assert.True(t, iterator.HasNext(), "El iterador debería tener elementos")
}

func TestHasNextWhenIsEmptyPosicionesPares(t *testing.T) {
	array := NewArray([]int{})
	iterator := array.PosicionesParesIterator()

	assert.False(t, iterator.HasNext(), "El iterador no debería tener elementos")
}

func TestHasNextWhenEmptiedPosicionesPares(t *testing.T) {
	array := NewArray([]int{1, 2, 3})
	iterator := array.PosicionesParesIterator()

	for iterator.HasNext() {
		iterator.Next()
	}

	assert.False(t, iterator.HasNext(), "El iterador no debería tener elementos")
}

func TestNextPosicionesPares(t *testing.T) {
	array := NewArray([]int{1, 2, 3, 4, 5, 6, 7, 8})
	iterator := array.PosicionesParesIterator()

	assert.True(t, iterator.HasNext(), "El iterador debería tener elementos")
	value, _ := iterator.Next()
	assert.Equal(t, 1, value)

	assert.True(t, iterator.HasNext(), "El iterador debería tener elementos")
	value, _ = iterator.Next()
	assert.Equal(t, 3, value)

	assert.True(t, iterator.HasNext(), "El iterador debería tener elementos")
	value, _ = iterator.Next()
	assert.Equal(t, 5, value)

	assert.True(t, iterator.HasNext(), "El iterador debería tener elementos")
	value, _ = iterator.Next()
	assert.Equal(t, 7, value)

	assert.False(t, iterator.HasNext(), "El iterador no debería tener elementos")
}

func TestErrorsWhenEmptiedPosicionesPares(t *testing.T) {
	array := NewArray([]int{})
	iterator := array.PosicionesParesIterator()

	_, err := iterator.Next()
	expectedError := errors.New("no hay más elementos")

	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
}
