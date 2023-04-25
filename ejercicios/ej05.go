package ejercicios

import "errors"

type ElementosParesIterator struct {
	array     Array
	posActual int
}

func (a *ElementosParesIterator) HasNext() bool {
	posicion := a.posActual + 1
	haySiguiente := false
	for posicion < len(a.array.data) {
		if a.array.data[posicion]%2 == 0 {
			haySiguiente = true
			break
		}
		posicion++
	}
	return haySiguiente
}

func (a *ElementosParesIterator) Next() (int, error) {
	if !a.HasNext() {
		return 0, errors.New("no hay más elementos")
	}
	a.posActual++
	for a.array.data[a.posActual]%2 != 0 && a.HasNext() {
		a.posActual++
	}
	value := a.array.data[a.posActual]
	return value, nil
}

func (a *Array) ElementosParesIterator() *ElementosParesIterator {
	return &ElementosParesIterator{
		array:     *a,
		posActual: 0,
	}
}
