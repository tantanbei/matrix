package matrix

import (
	"fmt"
)

type Laser struct {
	dataFF float64
	dataFR float64
	dataFL float64
	dataRR float64
	dataLL float64
}

type Matrix struct {
	realData [][]float64
	row      int
	col      int
}

func NewMatrix(data [][]float64) *Matrix {
	m := new(Matrix)
	m.realData = data
	m.row = len(data)
	m.col = len(data[0])
	return m
}

func NewMatrixZeros(arg ...int) *Matrix {
	if len(arg) > 2 {
		panic("imput error!")
	}
	m := new(Matrix)
	if len(arg) == 1 {
		for i := 0; i < arg[0]; i++ {
			row := make([]float64, 0)
			for j := 0; j < arg[0]; j++ {
				row = append(row, 0)
			}
			m.realData = append(m.realData, row)
		}
	} else {
		for i := 0; i < arg[0]; i++ {
			row := make([]float64, 0)
			for j := 0; j < arg[1]; j++ {
				row = append(row, 0)
			}
			m.realData = append(m.realData, row)
		}
	}
	m.row = len(m.realData)
	m.col = len(m.realData[0])
	return m
}

func NewMatrixOnes(arg ...int) *Matrix {
	if len(arg) > 2 {
		panic("imput error!")
	}
	m := new(Matrix)
	if len(arg) == 1 {
		for i := 0; i < arg[0]; i++ {
			row := make([]float64, 0)
			for j := 0; j < arg[0]; j++ {
				row = append(row, 1)
			}
			m.realData = append(m.realData, row)
		}
	} else {
		for i := 0; i < arg[0]; i++ {
			row := make([]float64, 0)
			for j := 0; j < arg[1]; j++ {
				row = append(row, 1)
			}
			m.realData = append(m.realData, row)
		}
	}
	m.row = len(m.realData)
	m.col = len(m.realData[0])
	return m
}

func NewMatrixEye(arg ...int) *Matrix {
	if len(arg) > 2 {
		panic("imput error!")
	}
	m := new(Matrix)
	if len(arg) == 1 {
		for i := 0; i < arg[0]; i++ {
			row := make([]float64, 0)
			for j := 0; j < arg[0]; j++ {
				if j == i {
					row = append(row, 1)
				} else {
					row = append(row, 0)
				}
			}
			m.realData = append(m.realData, row)
		}
	} else {
		for i := 0; i < arg[0]; i++ {
			row := make([]float64, 0)
			for j := 0; j < arg[1]; j++ {
				if j == i {
					row = append(row, 1)
				} else {
					row = append(row, 0)
				}
			}
			m.realData = append(m.realData, row)
		}
	}
	m.row = len(m.realData)
	m.col = len(m.realData[0])
	return m
}

func MatrixAddition(a *Matrix, b *Matrix) *Matrix {
	if a.col != b.col || a.row != b.row {
		panic("Matrix addition error!")
	}
	m := new(Matrix)
	m.col = a.col
	m.row = a.row
	for i := 0; i < a.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < a.col; j++ {
			row = append(row, a.realData[i][j]+b.realData[i][j])
		}
		m.realData = append(m.realData, row)
	}
	return m
}

func MatrixSubtraction(a *Matrix, b *Matrix) *Matrix {
	if a.col != b.col || a.row != b.row {
		panic("Matrix addition error!")
	}
	m := new(Matrix)
	m.col = a.col
	m.row = a.row
	for i := 0; i < a.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < a.col; j++ {
			row = append(row, a.realData[i][j]-b.realData[i][j])
		}
		m.realData = append(m.realData, row)
	}
	return m
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
