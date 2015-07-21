package matrix

//addition the some matrix
//eg D = A + B + C
func Addition(m ...*matrix) *matrix {
	for i := 1; i < len(m); i++ {
		if m[i].col != m[i-1].col || m[i].row != m[i-1].row {
			panic("Matrixs have different size")
		}
	}
	mat := new(matrix)
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

//subtraction the matrix in order
//eg D = A - B - C
func Subtraction(m ...*matrix) *matrix {
	for i := 1; i < len(m); i++ {
		if m[i].col != m[i-1].col || m[i].row != m[i-1].row {
			panic("Matrixs have different size")
		}
	}
	mat := new(matrix)
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

//multiplication the double matrix
//eg C = A * B
func Multiplication(a *matrix, b *matrix) *matrix {
	if a.col != b.row {
		panic("Can not matrix multiplication!")
	}
	m := new(matrix)
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

//transpose this matrix
//the original matrix can not change
//eg B = AT(transpose)
func (m matrix) Transpose() *matrix {
	mT := new(matrix)
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

//negative this matrix
//the original matrix can not change
//eg B = -A
func (m *matrix) Negative() *matrix {
	mN := new(matrix)
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

//copy this matrix
//the original can not change
func (m *matrix) Copy() *matrix {
	mC := new(matrix)
	mC.col = m.col
	mC.row = m.row
	for i := 0; i < mC.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < mC.col; j++ {
			row = append(row, m.realData[i][j])
		}
		mC.realData = append(m.realData, row)
	}
	return mC
}

//the matrix add a float
//eg B = A + 0.5
func MatrixAddFloat(a *matrix, b float64) *matrix {
	m := new(matrix)
	m.col = a.col
	m.row = a.row
	for i := 0; i < m.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < m.col; j++ {
			row = append(row, a.realData[i][j]+b)
		}
		m.realData = append(m.realData, row)
	}
	return m
}

//augment the double matrix
//eg [a b]
func MatrixAugment(a *matrix, b *matrix) *matrix {
	if a.row != b.row {
		panic("Augment error!")
	}
	m := new(matrix)
	m.col = a.col + b.col
	m.row = a.row
	for i := 0; i < m.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < m.col; j++ {
			if j < a.col {
				row = append(row, a.realData[i][j])
			} else {
				row = append(row, b.realData[i][j-a.col])
			}
		}
		m.realData = append(m.realData, row)
	}
	return m
}

//stack the double matrix
//eg [ A
//     B ]
func MatrixStack(a *matrix, b *matrix) *matrix {
	if a.row != b.row {
		panic("Stack error!")
	}
	m := new(matrix)
	m.col = a.col
	m.row = a.row + b.row
	for i := 0; i < m.row; i++ {
		row := make([]float64, 0)
		for j := 0; j < m.col; j++ {
			if i < a.row {
				row = append(row, a.realData[i][j])
			} else {
				row = append(row, b.realData[i-a.row][j])
			}
		}
		m.realData = append(m.realData, row)
	}
	return m
}
