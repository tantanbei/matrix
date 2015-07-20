package matrix

import (
	"fmt"
)

type Matrix struct {
	realData [][]float64
	row      int
	col      int
}

func (m *Matrix) GetRow() int {
	return m.row
}

func (m *Matrix) GetCol() int {
	return m.col
}

func (m *Matrix) GetRealData() [][]float64 {
	return m.realData
}

//Print the matrix
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
