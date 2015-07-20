package matrix

import (
	"math/rand"
	"strconv"
)

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

func NewMatrixNormal(row, col int) *Matrix {
	if row < 1 || col < 1 {
		panic("create random matrix error!")
	}
	m := new(Matrix)
	m.col = col
	m.row = row
	for i := 0; i < m.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < m.col; j++ {
			row = append(row, rand.NormFloat64())
		}
		m.realData = append(m.realData, row)
	}
	return m
}

func NewMatrixDiagonal(data []float64) *Matrix {
	if len(data) < 1 {
		panic("create diagonal matrix error!")
	}
	m := new(Matrix)
	m.col = len(data)
	m.row = len(data)
	for i := 0; i < m.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < m.col; j++ {
			if i == j {
				row = append(row, data[i])
			} else {
				row = append(row, 0)
			}
		}
		m.realData = append(m.realData, row)
	}
	return m
}

func NewMatrixByMatlab(cmd string) *Matrix {
	rn := []rune(cmd)
	if rn[0] != '[' || rn[len(rn)-1] != ']' {
		panic("create matrix by matlab error!")
	}
	index_start := 1
	index_end := 1

	m := new(Matrix)
	row := make([]float64, 0)
	for i := 1; i < len(rn); i++ {
		switch rn[i] {
		case ' ', ',':
			if index_end == index_start {
				panic("syntax error!")
			}
			num, err := strconv.ParseFloat(string(rn[index_start:index_end]), 64)
			if err != nil {
				panic("syntax error")
			}
			row = append(row, num)
			index_end = i + 1
			index_start = i + 1
		case ';', ']':
			if index_end == index_start {
				panic("syntax error!")
			}
			num, err := strconv.ParseFloat(string(rn[index_start:index_end]), 64)
			if err != nil {
				panic("syntax error")
			}
			row = append(row, num)
			index_end = i + 1
			index_start = i + 1

			m.col = len(row)
			m.row++

			m.realData = append(m.realData, row)
			row = make([]float64, 0)
		default:
			index_end++
		}
	}
	return m
}
