package matrix

import (
	"fmt"
)

func MatrixAddition(m ...*Matrix) *Matrix {
	for i := 1; i < len(m); i++ {
		if m[i].col != m[i-1].col || m[i].row != m[i-1].row {
			panic("Matrixs have different size")
		}
	}
	mat := new(Matrix)
	mat.col = m[0].col
	mat.row = m[0].row

	for i := 0; i < mat.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < mat.col; j++ {
			var num float64
			num = 0
			for k := 0; k < len(m); k++ {
				num = num + m[k].realData[i][j]
			}
			row = append(row, num)
		}
		mat.realData = append(mat.realData, row)
	}
	return mat
}

func MatrixSubtraction(m ...*Matrix) *Matrix {
	for i := 1; i < len(m); i++ {
		if m[i].col != m[i-1].col || m[i].row != m[i-1].row {
			panic("Matrixs have different size")
		}
	}
	mat := new(Matrix)
	mat.col = m[0].col
	mat.row = m[0].row

	for i := 0; i < mat.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < mat.col; j++ {
			var num float64
			num = m[0].realData[i][j]
			for k := 1; k < len(m); k++ {
				num = num - m[k].realData[i][j]
			}
			row = append(row, num)
		}
		mat.realData = append(mat.realData, row)
	}
	return mat
}

func MatrixMultiplication(a *Matrix, b *Matrix) *Matrix {
	if a.col != b.row {
		panic("Can not matrix multiplication!")
	}
	m := new(Matrix)
	m.row = a.row
	m.col = b.col

	var num float64

	for i := 0; i < m.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < m.col; j++ {
			num = 0
			for k := 0; k < a.col; k++ {
				num = a.realData[i][k]*b.realData[k][i] + num
			}
			row = append(row, num)
		}
		m.realData = append(m.realData, row)
	}
	return m
}

func PrintMatrix(m ...*Matrix) {
	for _, mat := range m {
		for i := 0; i < mat.row; i++ {
			for j := 0; j < mat.col; j++ {
				fmt.Print(mat.realData[i][j], "	")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func (m Matrix) Transpose() *Matrix {
	mT := new(Matrix)
	mT.col = m.row
	mT.row = m.col
	for j := 0; j < m.col; j++ {
		row := make([]float64, 0)
		for i := 0; i < m.row; i++ {
			row = append(row, m.realData[i][j])
		}
		mT.realData = append(mT.realData, row)
	}
	return mT
}

func (m Matrix) Negative() *Matrix {
	mN := new(Matrix)
	mN.col = m.col
	mN.row = m.row
	for i := 0; i < m.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < m.col; j++ {
			row = append(row, -m.realData[i][j])
		}
		mN.realData = append(mN.realData, row)
	}
	return mN
}