package ejercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArchivosMusica(t *testing.T) {
	original := NewFileMp3("el gato volador.mp3")
	adaptador := NewArchivoDeMusicaAdapter(*original)

	assert.Equal(t, "Reproduciendo archivo MP3. Nombre: el gato volador.mp3", adaptador.reproducir())
}
