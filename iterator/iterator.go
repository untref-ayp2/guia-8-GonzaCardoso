package main

import (
	"errors"
	"fmt"
)

// Define una interfaz de Iterator que define los métodos necesarios para recorrer una matriz.
type Iterator interface {
	HasNext() bool
	Next() (int, error)
}

// Define una estructura de matriz que contiene los datos y las posiciones actuales.
type Matrix struct {
	data [][]int
}

type MatrixIterator struct {
	matrix        Matrix
	currentRow    int
	currentColumn int
}

// Implementa la interfaz Iterator para la matriz.
func (m *MatrixIterator) HasNext() bool {
	return m.currentRow < len(m.matrix.data) && m.currentColumn < len(m.matrix.data[m.currentRow])
}

// Implementa la interfaz Iterator para la matriz.
func (m *MatrixIterator) Next() (int, error) {
	if !m.HasNext() {
		return 0, errors.New("no hay más elementos")
	}

	value := m.matrix.data[m.currentRow][m.currentColumn]

	// Actualiza las posiciones actuales para el siguiente valor.
	m.currentColumn++
	if m.currentColumn >= len(m.matrix.data[m.currentRow]) {
		m.currentRow++
		m.currentColumn = 0
	}

	return value, nil
}

// El constructor de matriz crea una nueva matriz y la devuelve.
func NewMatrix(data [][]int) *Matrix {
	m := &Matrix{data: data}
	return m
}

// El método Iterator devuelve un iterador para la matriz.
func (m *Matrix) Iterator() *MatrixIterator {
	return &MatrixIterator{
		matrix:        *m,
		currentRow:    0,
		currentColumn: 0,
	}
}

func main() {
	// Crea una matriz de ejemplo.
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Crea un objeto de matriz y un iterador.
	matrix := NewMatrix(data)
	iterator := matrix.Iterator()

	// Utiliza el iterador para recorrer la matriz.
	for iterator.HasNext() {
		value, _ := iterator.Next()
		fmt.Println(value)
	}

	fmt.Println("--------------------")

	// Al pedir más elementos de los que hay, el iterador devuelve un error.
	// Por eso hay que utilizar HasNext() para asegurarse de que hay más elementos.
	val, err := iterator.Next()
	fmt.Println(val, err)
}
