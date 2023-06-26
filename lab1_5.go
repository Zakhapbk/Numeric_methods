package main

import (
	"fmt"
	"math"
)

func sign(a float64) (b float64) {
	if a > 0 {
		return 1
	} else if a == 0 {
		return 0
	} else {
		return -1
	}
}

func VVT(v []float64, m int) (a [][]float64) {
	r := make([][]float64, m)
	for i := range r {
		r[i] = make([]float64, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			r[i][j] = 2 * v[i] * v[j]
		}
	}
	return r
}

func Emtr(m int) (A [][]float64) {
	E := make([][]float64, m)
	for i := 0; i < m; i++ {
		E[i] = make([]float64, m)
	}
	for i := 0; i < m; i++ {
		E[i][i] = 1
	}
	return E
}

func compl_sq(A, B, C float64) (r1, r2 complex128) {

	d := B*B - 4*C*A
	z1 := complex(-B/(2*A), math.Sqrt(-d)/2/A)
	z2 := complex(-B/(2*A), -math.Sqrt(-d)/2/A)
	return z1, z2
}

func VTV(v []float64, m int) (a float64) {
	res := 0.0
	for i := 0; i < m; i++ {
		res += v[i] * v[i]
	}
	return res
}

func m_minus(a, b [][]float64, m int) (c [][]float64) {
	r := make([][]float64, m)
	for i := range r {
		r[i] = make([]float64, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			r[i][j] = a[i][j] - b[i][j]
		}
	}
	return r
}

func qr(A [][]float64, m int) (R1, R2 [][]float64) {
	H, Q := make([][]float64, m), make([][]float64, m)
	Hl := Emtr(m)
	v := make([]float64, m)
	for i := 0; i < m; i++ {
		H[i] = make([]float64, m)
		Q[i] = make([]float64, m)
	}
	sum := 0.0
	for i := 0; i < m-1; i++ {
		for k := i; k < m; k++ {
			sum += A[k][i] * A[k][i]
		}
		sum = math.Sqrt(sum)
		for j := i; j < m; j++ {
			if i == j {
				v[j] = A[i][i] + sign(A[i][i])*sum
			} else {
				v[j] = A[j][i]
			}
		}
		H = VVT(v, m)
		for i := 0; i < m; i++ {
			for j := 0; j < m; j++ {
				H[i][j] = H[i][j] / VTV(v, m)
			}
		}
		H = m_minus(Emtr(m), H, m)
		A = mult(H, A, m)
		Q = mult(Hl, H, m)
		Hl = H
		sum = 0
		for i := range v {
			v[i] = 0
		}
	}

	return Q, A
}

func values(Q, R [][]float64, m int, eps float64) {
	A := make([][]float64, m)
	for i := range A {
		A[i] = make([]float64, m)
	}
	count := 0
	for {
		A = mult(R, Q, m)
		if math.Sqrt(A[1][0]*A[1][0]+A[2][0]*A[2][0]) < eps && math.Abs(A[2][1]) < eps {
			fmt.Println("Coбственные значения:", A[0][0], A[1][1], A[2][2])
			break
		} else if math.Sqrt(A[1][0]*A[1][0]+A[2][0]*A[2][0]) < eps && math.Abs(A[2][1]) > eps {
			B := -(A[1][1] + A[2][2])
			C := A[1][1]*A[2][2] - A[1][2]*A[2][1]
			z1, z2 := compl_sq(1.0, B, C)
			fmt.Printf("Coбственные значения: %f  %v  %v", A[0][0], z1, z2)
			break
		} else if math.Abs(A[2][1]) < eps {
			B := -(A[0][0] + A[1][1])
			C := A[1][1]*A[0][0] - A[0][1]*A[1][0]
			z1, z2 := compl_sq(1.0, B, C)
			fmt.Printf("Coбственные значения: %f  %v  %v", A[2][2], z1, z2)
			break
		}
		count++
		Q, R = qr(A, m)
	}
	fmt.Println()
	fmt.Println("Число итераций:", count)
}

//func main() {
//	var m int
//	var eps float64
//	fmt.Println("Введите точность:")
//	fmt.Scan(&eps)
//
//	fmt.Println("Введите размерность:")
//	fmt.Scan(&m)
//	mtr := scanmtr(m)
//	Q, A := qr(mtr, m)
//	prnt(mult(Q, A, m), m)
//	values(Q, A, m, eps)
//}
