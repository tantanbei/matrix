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

	c := Multiplication(a, b)

	PrintMatrix(a, b)
	PrintMatrix(c)
}

func TestMatrixMatrixAddition(T *testing.T) {
	dataA := [][]float64{{1.5, 1}, {1.2, 0.8}}
	a := NewMatrix(dataA)
	dataB := [][]float64{{1.1, 0.9}, {1.5, 1.9}}
	b := NewMatrix(dataB)

	c := Addition(a, b)

	PrintMatrix(c)
}

func TestMatrixSub(T *testing.T) {
	dataA := [][]float64{{1.5, 1}, {1.2, 0.8}}
	a := NewMatrix(dataA)
	dataB := [][]float64{{1.1, 0.9}, {1.5, 1.9}}
	b := NewMatrix(dataB)

	c := Subtraction(a, b)

	PrintMatrix(c)
}

func TestTranM(T *testing.T) {
	dataA := [][]float64{{1}, {2}, {3}, {4}}
	a := NewMatrix(dataA)

	aT := a.Transpose()
	PrintMatrix(a, aT)
}

func TestNM(T *testing.T) {
	dataA := [][]float64{{1}, {2}, {3}, {4}}
	a := NewMatrix(dataA)

	aN := a.Negative()
	PrintMatrix(a, aN)
}

func TestMutliAdd(T *testing.T) {
	a := NewMatrixEye(4)
	b := NewMatrixOnes(4)
	c := NewMatrixZeros(4)
	d := Addition(a, b, c)
	PrintMatrix(d)
}

func TestMutliSub(T *testing.T) {
	a := NewMatrixEye(4)
	b := NewMatrixOnes(4)
	c := NewMatrixZeros(4)
	d := Subtraction(a, b, c)
	PrintMatrix(d)
}

func TestAddFloat(T *testing.T) {
	a := NewMatrixOnes(2)
	b := MatrixAddFloat(a, -1.9)
	PrintMatrix(b)
}

func TestAugment(T *testing.T) {
	a := NewMatrixEye(3)
	b := MatrixAddFloat(a, 1)
	m := MatrixAugment(a, b)
	PrintMatrix(m)
}

func TestStack(T *testing.T) {
	a := NewMatrixZeros(2)
	b := NewMatrixOnes(2)
	m := MatrixStack(a, b)
	PrintMatrix(m)
}

func TestCopy(T *testing.T) {
	a := NewMatrixOnes(2)
	m := a.Copy()
	PrintMatrix(m)
}

func TestNormal(T *testing.T) {
	a := NewMatrixNormal(2, 4)
	PrintMatrix(a)
}

func TestDiagonal(T *testing.T) {
	a := []float64{1, 2, 3, 4}
	m := NewMatrixDiagonal(a)
	PrintMatrix(m)
}

func TestByMatlab(T *testing.T) {
	m := NewMatrixByMatlab("[1 2 3;4 5 6]")
	m1 := NewMatrixByMatlab("[1,2;3,4;5,6]")
	PrintMatrix(m, m1)
}
