package main

import (
	"fmt"
)

func scanmtr(m int) (A [][]float64) {
	mtr := make([][]float64, m)
	for i := range mtr {
		mtr[i] = make([]float64, m)
	}
	t := 0.0
	fmt.Println("Введите матрицу:")
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&t)
			mtr[i][j] = t
		}
	}
	return mtr
}

func LU(mtr [][]float64, val []float64, m int) (A, B [][]float64, c []float64, d []float64) {
	L := make([][]float64, m)
	for i := range L {
		L[i] = make([]float64, m)
	}
	U := make([][]float64, m)
	for i := range U {
		U[i] = make([]float64, m)
	}

	y := make([]float64, m)
	x := make([]float64, m)
	t := 1
	var l, u float64 = 0, 0
	for i := 0; i < m; i++ {

		for j := 0; j < t; j++ {
			for k := 0; k <= j-1; k++ {
				l += L[i][k] * U[k][j]
			}

			L[i][j] = mtr[i][j] - l
			l = 0
		}
		for j := m - 1; j >= i; j-- {
			for k := 0; k <= i-1; k++ {
				u += (L[i][k] * U[k][j]) / L[i][i]
			}
			U[i][j] = mtr[i][j]/L[i][i] - u ///
			u = 0
		}
		t++

	}
	y[0] = val[0] / L[0][0]
	for i := 1; i < m; i++ {
		for k := 0; k <= i-1; k++ {
			u += (L[i][k] * y[k])
		}
		y[i] = (val[i] - u) / L[i][i]
		u = 0
	}
	x[m-1] = y[m-1]
	for i := m - 2; i >= 0; i-- {
		for k := i + 1; k <= m-1; k++ {
			u += U[i][k] * x[k]
		}
		x[i] = y[i] - u
		u = 0
	}

	return L, U, y, x
}

func prnt(mtr [][]float64, m int) {
	for i := 0; i < m; i++ {
		fmt.Println()
		for j := 0; j < m; j++ {
			fmt.Printf("%.4f ", mtr[i][j])
		}
	}
	fmt.Println()
}

func opposite(U, L [][]float64, m int) (x [][]float64) {
	l := 0.0

	X := make([][]float64, m)
	for i := range X {
		X[i] = make([]float64, m)
	}

	Y := make([][]float64, m)
	for i := range Y {
		Y[i] = make([]float64, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				X[i][i] = 1 / L[i][i]
				continue
			}
			for k := j; k < i; k++ {
				l += L[i][k] * X[k][j]
			}

			X[i][j] = -(l / L[i][i])
			l = 0
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := m - 1; j >= i; j-- {
			if i == j {
				Y[i][i] = 1 / U[i][i]
				continue
			}
			for k := j; k > i; k-- {
				l += U[i][k] * Y[k][j]
			}

			Y[i][j] = -(l / U[i][i])
			l = 0
		}
	}

	return mult(Y, X, m)
}

//func main() {
//	var m int
//	t := 0.0
//	fmt.Println("Введите размерность:")
//	fmt.Scan(&m)
//
//	val := make([]float64, m)
//
//	mtr := scanmtr(m)
//	fmt.Println("Введите правые части:")
//	for i := 0; i < m; i++ {
//		fmt.Scan(&t)
//		val[i] = t
//	}
//	prnt(mtr, m)
//	A, B, _, x := LU(mtr, val, m)
//	fmt.Println("Матрица L:")
//	prnt(A, m)
//	fmt.Println("Матрица U:")
//	prnt(B, m)
//
//	//fmt.Println(y)
//	fmt.Println()
//	fmt.Println("Результат:")
//	for i := 0; i < m; i++ {
//		fmt.Print(math.Round(x[i]), " ")
//	}
//	fmt.Println()
//	fmt.Println("Определитель:")
//	det := 1.0
//	for i := 0; i < m; i++ {
//		det *= A[i][i]
//	}
//	fmt.Println(math.Round(det))
//	fmt.Println("Обратная матрица:")
//	c := opposite(B, A, m)
//	prnt(c, m)
//	fmt.Println("Проверка A*A^(-1):")
//	prnt(mult(mtr, c, m), m)
//
//}
