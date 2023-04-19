package ejercicios

type Iterator interface {
	HasNext() bool
	Next() (int, error)
}

type Array struct {
	data []int
}

func NewArray(data []int) *Array {
	a := &Array{data: data}
	return a
}
