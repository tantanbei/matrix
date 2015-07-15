package matrix

import (
	"testing"
)

func TestNewMZ(T *testing.T) {
	m := NewMatrixZeros(4)
	PrintMatrix(m)

	m1 := NewMatrixZeros(4, 5)
	PrintMatrix(m1)
}

func TestNewMO(T *testing.T) {
	m := NewMatrixOnes(4)
	PrintMatrix(m)

	m1 := NewMatrixOnes(4, 5)
	PrintMatrix(m1)
}

func TestNewME(T *testing.T) {
	m := NewMatrixEye(4)
	PrintMatrix(m)

	m1 := NewMatrixEye(4, 5)
	PrintMatrix(m1)
}

func TestMatrixMultiplication(T *testing.T) {
	dataA := [][]float64{{1, 1, 1}, {1, 1, 1}}
	a := NewMatrix(dataA)
	dataB := [][]float64{{1, 1}, {1, 1}, {1, 1}}
	b := NewMatrix(dataB)

	c := MatrixMultiplication(a, b)

	PrintMatrix(a, b)
	PrintMatrix(c)
}

func TestMatrixMatrixAddition(T *testing.T) {
	dataA := [][]float64{{1.5, 1}, {1.2, 0.8}}
	a := NewMatrix(dataA)
	dataB := [][]float64{{1.1, 0.9}, {1.5, 1.9}}
	b := NewMatrix(dataB)

	c := MatrixAddition(a, b)

	PrintMatrix(c)
}

func TestMatrixSub(T *testing.T) {
	dataA := [][]float64{{1.5, 1}, {1.2, 0.8}}
	a := NewMatrix(dataA)
	dataB := [][]float64{{1.1, 0.9}, {1.5, 1.9}}
	b := NewMatrix(dataB)

	c := MatrixSubtraction(a, b)

	PrintMatrix(c)
}

func TestTranM(T *testing.T) {
	dataA := [][]float64{{1}, {2}, {3}, {4}}
	a := NewMatrix(dataA)

	aT := a.TransposeMatrix()
	PrintMatrix(a, aT)
}
