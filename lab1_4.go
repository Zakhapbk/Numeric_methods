package main

import (
	"math"
)

func mult(L, U [][]float64, m int) (B [][]float64) {

	A := make([][]float64, m)
	for i := range A {
		A[i] = make([]float64, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < m; k++ {
				A[i][j] += L[i][k] * U[k][j]
			}
		}

	}
	return A
}

func trans(A [][]float64, m int) (b [][]float64) {

	a := make([][]float64, m)
	for i := range a {
		a[i] = make([]float64, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			a[i][j] = A[j][i]
		}
	}
	return a
}

func yak(a [][]float64, m int, eps float64) (v []float64, c [][]float64, it int) {
	t := 1.0
	i := 1
	var phi float64
	var I, J = 0, 1
	U := make([][]float64, m)
	for i := 0; i < m; i++ {
		U[i] = make([]float64, m)
	}
	Ur := make([][]float64, m)
	for i := 0; i < m; i++ {
		Ur[i] = make([]float64, m)
	}
	Ul := make([][]float64, m)
	for i := 0; i < m; i++ {
		Ul[i] = make([]float64, m)
	}

	for ; t > eps; i++ {
		t = 0
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				Ul[i][j] = U[i][j]
			}
		}
		max := math.Abs(a[0][1])
		I, J = 0, 1
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				U[i][j] = 0
			}
			U[i][i] = 1
		}

		for i := 0; i < m; i++ {
			for j := i; j < m; j++ {
				if (i != j) && (math.Abs(a[i][j]) > max) {
					max = math.Abs(a[i][j])
					I = i
					J = j
				}
			}

		}
		phi = 0.5 * math.Atan(2*a[I][J]/(a[I][I]-a[J][J]))
		U[I][J] = -math.Sin(phi)
		U[J][I] = math.Sin(phi)
		U[I][I], U[J][J] = math.Cos(phi), math.Cos(phi)

		a = mult(trans(U, m), mult(a, U, m), m)
		for i := 0; i < m; i++ {
			for j := i; j < m; j++ {
				if i == j {
					continue
				}
				t += a[i][j] * a[i][j]
			}
		}
		t = math.Sqrt(t)
		if i != 1 {
			Ur = mult(Ur, U, m)
		} else {
			for i := 0; i < m; i++ {
				for j := 0; j < m; j++ {
					Ur[i][j] = U[i][j]
				}
			}

		}
	}
	sval := make([]float64, m)
	sv := make([][]float64, m)
	for i := range sv {
		sv[i] = make([]float64, m)
	}

	for i := 0; i < m; i++ {
		sval[i] = a[i][i]
	}
	return sval, Ur, i
}

//func main() {
//	var m int
//	var eps float64
//	fmt.Println("Введите точность:")
//	fmt.Scan(&eps)
//	//t := 0.0
//	fmt.Println("Введите размерность:")
//	fmt.Scan(&m)
//	mtr := scanmtr(m)
//	var v, sv, _ = yak(mtr, m, eps)
//	fmt.Println()
//	fmt.Println("Собственные значения")
//	for _, val := range v {
//		fmt.Print(val, " ")
//	}
//	fmt.Println()
//	fmt.Println("Собственные векторы")
//	prnt(sv, m)
//
//}
