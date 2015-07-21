package matrix

import (
	"math/rand"
	"strconv"
)

//initial a new matrix by [][]float64
func NewMatrix(data [][]float64) *matrix {
	m := new(matrix)
	m.realData = data
	m.row = len(data)
	m.col = len(data[0])
	return m
}

//initial a zeros matrix
//arg[0] : row of the matrix
//arg[1] : col of the matrix
//if arg only one , there is a n*n matrix
func NewMatrixZeros(arg ...int) *matrix {
	if len(arg) > 2 {
		panic("imput error!")
	}
	m := new(matrix)
	if len(arg) == 1 {
		if arg[0] < 1 {
			panic("imput error!")
		}
		m.row = arg[0]
		m.col = arg[0]
	} else {
		if arg[0] < 1 || arg[1] < 1 {
			panic("imput error!")
		}
		m.row = arg[0]
		m.col = arg[1]
	}
	for i := 0; i < m.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < m.col; j++ {
			row = append(row, 0)
		}
		m.realData = append(m.realData, row)
	}
	return m
}

//initial a ones matrix
//arg[0] : row of the matrix
//arg[1] : col of the matrix
//if arg only one , there is a n*n matrix
func NewMatrixOnes(arg ...int) *matrix {
	if len(arg) > 2 {
		panic("imput error!")
	}
	m := new(matrix)
	if len(arg) == 1 {
		if arg[0] < 1 {
			panic("imput error!")
		}
		m.row = arg[0]
		m.col = arg[0]
	} else {
		if arg[0] < 1 || arg[1] < 1 {
			panic("imput error!")
		}
		m.row = arg[0]
		m.col = arg[1]
	}
	for i := 0; i < m.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < m.col; j++ {
			row = append(row, 1)
		}
		m.realData = append(m.realData, row)
	}
	return m
}

//initial a eye matrix
//arg[0] : row of the matrix
//arg[1] : col of the matrix
//if arg only one , there is a n*n matrix
func NewMatrixEye(arg ...int) *matrix {
	if len(arg) > 2 {
		panic("imput error!")
	}
	m := new(matrix)
	if len(arg) == 1 {
		if arg[0] < 1 {
			panic("imput error!")
		}
		m.row = arg[0]
		m.col = arg[0]
	} else {
		if arg[0] < 1 || arg[1] < 1 {
			panic("imput error!")
		}
		m.row = arg[0]
		m.col = arg[1]
	}
	for i := 0; i < m.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < m.col; j++ {
			if j == i {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		m.realData = append(m.realData, row)
	}
	return m
}

//initial a normal matrix
//arg[0] : row of the matrix
//arg[1] : col of the matrix
//if arg only one , there is a n*n matrix
func NewMatrixNormal(arg ...int) *matrix {
	if len(arg) > 2 {
		panic("imput error!")
	}
	m := new(matrix)
	if len(arg) == 1 {
		if arg[0] < 1 {
			panic("imput error!")
		}
		m.col = arg[0]
		m.row = arg[0]
	} else {
		if arg[0] < 1 || arg[1] < 1 {
			panic("imput error!")
		}
		m.row = arg[0]
		m.col = arg[1]
	}
	for i := 0; i < m.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < m.col; j++ {
			row = append(row, rand.NormFloat64())
		}
		m.realData = append(m.realData, row)
	}
	return m
}

//initial a diagonal matrix by []floate64
func NewMatrixDiagonal(data []float64) *matrix {
	if len(data) < 1 {
		panic("create diagonal matrix error!")
	}
	m := new(matrix)
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

//initial a matrix by matlab protocal
//eg [1 2;2 3;3 4]
//   [1,2;2,3;3,4]
func NewMatrixByMatlab(cmd string) *matrix {
	rn := []rune(cmd)
	if rn[0] != '[' || rn[len(rn)-1] != ']' {
		panic("create matrix by matlab error!")
	}
	index_start := 1
	index_end := 1

	m := new(matrix)
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
