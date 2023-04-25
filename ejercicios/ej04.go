package ejercicios

import "errors"

type PosicionesParesIterator struct {
	array     Array
	posActual int
}

func (a *PosicionesParesIterator) HasNext() bool {
	return a.posActual+1 < len(a.array.data)
}

func (a *PosicionesParesIterator) Next() (int, error) {
	if !a.HasNext() {
		return 0, errors.New("no hay más elementos")
	}
	value := a.array.data[a.posActual]
	a.posActual = a.posActual + 2
	return value, nil
}

func (a *Array) PosicionesParesIterator() *PosicionesParesIterator {

	return &PosicionesParesIterator{
		array:     *a,
		posActual: 0,
	}
}
