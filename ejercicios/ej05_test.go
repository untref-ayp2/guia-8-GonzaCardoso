package ejercicios

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasNext(t *testing.T) {
	array := NewArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	iterator := array.ElementosParesIterator()

	assert.True(t, iterator.HasNext(), "El iterador debería tener elementos")
}

func TestHasNextWhenIsEmpty(t *testing.T) {
	array := NewArray([]int{})
	iterator := array.ElementosParesIterator()

	assert.False(t, iterator.HasNext(), "El iterador no debería tener elementos")
}

func TestHasNextWhenEmptied(t *testing.T) {
	array := NewArray([]int{1, 2, 3})
	iterator := array.ElementosParesIterator()

	for iterator.HasNext() {
		iterator.Next()
	}

	assert.False(t, iterator.HasNext(), "El iterador no debería tener elementos")
}

func TestNext(t *testing.T) {
	array := NewArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	iterator := array.ElementosParesIterator()

	assert.True(t, iterator.HasNext(), "El iterador debería tener elementos")
	value, _ := iterator.Next()
	assert.Equal(t, 2, value)

	assert.True(t, iterator.HasNext(), "El iterador debería tener elementos")
	value, _ = iterator.Next()
	assert.Equal(t, 4, value)

	assert.True(t, iterator.HasNext(), "El iterador debería tener elementos")
	value, _ = iterator.Next()
	assert.Equal(t, 6, value)

	assert.True(t, iterator.HasNext(), "El iterador debería tener elementos")
	value, _ = iterator.Next()
	assert.Equal(t, 8, value)

	assert.False(t, iterator.HasNext(), "El iterador no debería tener elementos")
}

func TestErrorsWhenEmptied(t *testing.T) {
	array := NewArray([]int{})
	iterator := array.ElementosParesIterator()

	_, err := iterator.Next()
	expectedError := errors.New("no hay más elementos")

	if assert.Error(t, err) {
		assert.Equal(t, expectedError, err)
	}
}
