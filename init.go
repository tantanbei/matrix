package matrix

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
