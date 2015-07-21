package matrix

import (
	"fmt"
)

type matrix struct {
	realData [][]float64
	row      int
	col      int
}

func (m *matrix) GetRow() int {
	return m.row
}

func (m *matrix) GetCol() int {
	return m.col
}

func (m *matrix) GetRealData() [][]float64 {
	return m.realData
}

func (m *matrix) GetSize() (row, col int) {
	row = m.row
	col = m.col
	return
}

func (m *matrix) NumElements() int {
	return m.row * m.col
}

func (m *matrix) Nil() bool {
	return m == nil
}

//Print the matrix
func PrintMatrix(m ...*matrix) {
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
