package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEjemplo(t *testing.T) {
	parte1 := NewParteSimple("tornillo", 5)
	parte2 := NewParteSimple("tuerca", 3)
	// parte3 := NewParteSimple("ruleman", 50)
	parte4 := NewParteSimple("varilla", 5)

	// tiene un tornillo, dos tuercas y una varilla
	ensamble := NewEnsamble()
	ensamble.AgregarParte(parte1)
	ensamble.AgregarParte(parte2)
	ensamble.AgregarParte(parte2)
	ensamble.AgregarParte(parte4)

	assert.Equal(t, float64(16), ensamble.ObtenerPrecio(), "El precio del ensamble debería ser 2")
	assert.Equal(t, "tornillo, tuerca, tuerca, varilla", ensamble.ObtenerDescripcion(), "La descripción del ensamble debería ser 'tornillo, tuerca, tuerca, varilla'")
}
