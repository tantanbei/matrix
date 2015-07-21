package matrix

import (
	"testing"
)

func TestNewMZ(T *testing.T) {
	m := NewMatrixZeros(4)
	if m.realData[0][0] != 0 || m.realData[3][2] != 0 {
		T.Error("zeros error!")
	}

	m1 := NewMatrixZeros(4, 5)
	if m1.col != 5 || m1.row != 4 {
		T.Error("zeros error")
	}
}

func TestNewMO(T *testing.T) {
	m := NewMatrixOnes(4)
	if m.realData[0][0] != 1 || m.realData[3][2] != 1 {
		T.Error("ones error!")
	}

	m1 := NewMatrixOnes(4, 5)
	if m1.col != 5 || m1.row != 4 || m1.realData[3][3] != 1 {
		T.Error("ones error")
	}
}

func TestNewME(T *testing.T) {
	m := NewMatrixEye(4)
	if m.col != 4 || m.row != 4 || m.realData[3][2] != 0 {
		T.Error("eye error")
	}

	m1 := NewMatrixEye(4, 5)
	if m1.col != 5 || m1.row != 4 || m1.realData[3][3] != 1 {
		T.Error("eye error")
	}
}

func TestMatrixMultiplication(T *testing.T) {
	dataA := [][]float64{{1, 1, 1}, {1, 1, 1}}
	a := NewMatrix(dataA)
	dataB := [][]float64{{1, 1}, {1, 1}, {1, 1}}
	b := NewMatrix(dataB)

	c := Multiplication(a, b)

	if c.col != 2 || c.row != 2 || c.realData[1][1] != 3 {
		T.Error("multi error")
	}
}

func TestMatrixMatrixAddition(T *testing.T) {
	dataA := [][]float64{{1.5, 1}, {1.2, 0.8}}
	a := NewMatrix(dataA)
	dataB := [][]float64{{1.1, 0.9}, {1.5, 1.9}}
	b := NewMatrix(dataB)

	c := Addition(a, b)

	if c.col != 2 || c.row != 2 || c.realData[1][1] != 2.7 {
		T.Error("add error")
	}
}

func TestMatrixSub(T *testing.T) {
	dataA := [][]float64{{1.5, 1}, {1.2, 0.8}}
	a := NewMatrix(dataA)
	dataB := [][]float64{{1.1, 0.9}, {1.5, 1.9}}
	b := NewMatrix(dataB)

	c := Subtraction(a, b)

	if c.col != 2 || c.row != 2 {
		T.Error("sub error")
	}
}

func TestTranM(T *testing.T) {
	dataA := [][]float64{{1}, {2}, {3}, {4}}
	a := NewMatrix(dataA)

	aT := a.Transpose()
	if aT.col != 4 || aT.row != 1 {
		T.Error("tran error")
	}
}

func TestNM(T *testing.T) {
	dataA := [][]float64{{1}, {2}, {3}, {4}}
	a := NewMatrix(dataA)

	aN := a.Negative()
	if aN.col != 1 || aN.row != 4 || aN.realData[0][0] != -1 {
		T.Error("negative error")
	}
}

func TestMutliAdd(T *testing.T) {
	a := NewMatrixEye(4)
	b := NewMatrixOnes(4)
	c := NewMatrixZeros(4)
	d := Addition(a, b, c)
	if d.col != 4 || d.row != 4 || d.realData[3][3] != 2 {
		T.Error("muladd error")
	}
}

func TestMutliSub(T *testing.T) {
	a := NewMatrixEye(4)
	b := NewMatrixOnes(4)
	c := NewMatrixZeros(4)
	d := Subtraction(a, b, c)
	if d.col != 4 || d.row != 4 || d.realData[1][2] != -1 {
		T.Error("mulsub error")
	}
}

func TestAddFloat(T *testing.T) {
	a := NewMatrixOnes(2)
	b := MatrixAddFloat(a, -1.9)
	if b.col != 2 || b.row != 2 {
		T.Error("addfloat error")
	}
}

func TestAugment(T *testing.T) {
	a := NewMatrixEye(3)
	b := MatrixAddFloat(a, 1)
	m := MatrixAugment(a, b)
	if m.col != 6 || m.row != 3 {
		T.Error("augment error")
	}
}

func TestStack(T *testing.T) {
	a := NewMatrixZeros(2)
	b := NewMatrixOnes(2)
	m := MatrixStack(a, b)
	if m.col != 2 || m.row != 4 {
		T.Error("stack error")
	}
}

func TestCopy(T *testing.T) {
	a := NewMatrixOnes(2)
	m := a.Copy()
	if m.col != 2 || m.row != 2 || m.realData[0][0] != 1 {
		T.Error("zeros error")
	}
}

func TestNormal(T *testing.T) {
	a := NewMatrixNormal(2, 4)
	if a.col != 4 || a.row != 2 {
		T.Error("normal error")
	}
}

func TestDiagonal(T *testing.T) {
	a := []float64{1, 2, 3, 4}
	m := NewMatrixDiagonal(a)
	if m.col != 4 || m.row != 4 || m.realData[3][3] != 4 {
		T.Error("dia error")
	}
}

func TestByMatlab(T *testing.T) {
	m := NewMatrixByMatlab("[1 2 3;4 5 6]")
	if m.col != 3 || m.row != 2 || m.realData[0][0] != 1 {
		T.Error("mat error")
	}

	m1 := NewMatrixByMatlab("[1.5,2.6;3.06,4.33;5.5,6]")
	if m1.col != 2 || m1.row != 3 || m1.realData[1][1] != 4.33 {
		T.Error("mat error")
	}
}
