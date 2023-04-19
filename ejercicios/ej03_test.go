package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperacionUno(t *testing.T) {
	operacion := NewOperando(4)

	assert.Equal(t, "4", operacion.String())
	assert.Equal(t, 4, operacion.Resolver())
}

func TestOperacionDos(t *testing.T) {
	operacion := NewOperador("+", NewOperando(3), NewOperando(8))

	assert.Equal(t, "3+8", operacion.String())
	assert.Equal(t, 11, operacion.Resolver())
}

func TestOperacionTres(t *testing.T) {
	operacion := NewOperador("*", NewOperador("+", NewOperando(3), NewOperando(5)), NewOperando(14))

	assert.Equal(t, "3+5*14", operacion.String())
	assert.Equal(t, 112, operacion.Resolver())
}
