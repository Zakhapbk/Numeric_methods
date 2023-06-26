package main

import (
	"fmt"
	"math"
)

func iter(a [][]float64, val []float64, m int, eps float64) (A []float64, r int) {
	var max, count = 0.0, 0.0
	i := 1
	xl := make([]float64, m)
	xn := make([]float64, m)
	bet := make([]float64, m)
	alp := make([][]float64, m)
	for i := 0; i < m; i++ {
		alp[i] = make([]float64, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i == j {
				continue
			}
			alp[i][j] = -a[i][j] / a[i][i]
		}
		bet[i] = val[i] / a[i][i]
	}
	for i := 0; i < m; i++ {
		max += math.Abs(alp[0][i])
	}

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			count += math.Abs(alp[i][j])
		}
		count = 0
	}
	if count > 1 {
		fmt.Println(count)
		panic("Условие сходимости не выполнено")
	}
	for i := 0; i < m; i++ {
		xl[i] = bet[i]
	}
	e := eps + 1
	for ; e > eps; i++ {
		xn = sumv(bet, multv(alp, xl, m), m)
		e = norm(xn, xl, m)
		for j := 0; j < m; j++ {
			xl[j] = xn[j]
		}
	}

	return xn, i
}

func zeyd(a [][]float64, val []float64, m int, eps float64) (A []float64, r int) {
	i := 1
	xl := make([]float64, m)
	xn := make([]float64, m)
	bet := make([]float64, m)
	alp := make([][]float64, m)
	for i := 0; i < m; i++ {
		alp[i] = make([]float64, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i == j {
				continue
			}
			alp[i][j] = -a[i][j] / a[i][i]
		}
		bet[i] = val[i] / a[i][i]
	}
	for i := 0; i < m; i++ {
		xl[i] = bet[i]
	}

	e := eps + 1
	for ; e > eps; i++ {
		for i := 0; i < m; i++ {
			xn[i] = bet[i]
			for j := m - 1; j >= i; j-- {
				xn[i] += alp[i][j] * xl[j]
			}
			for j := 0; j < i; j++ {
				xn[i] += alp[i][j] * xn[j]
				if (i == m-1) && (j == i-1) {
					xn[i] += alp[i][j+1] * xn[j+1]
				}
			}
		}
		e = norm(xn, xl, m)
		for i := 0; i < m; i++ {
			xl[i] = xn[i]
		}
	}

	return xn, i
}

func multv(mtr [][]float64, v []float64, m int) (r []float64) {
	tmp := 0.0
	res := make([]float64, m)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			tmp += mtr[i][j] * v[j]
		}
		res[i] = tmp
		tmp = 0
	}
	return res
}

func sumv(a, b []float64, m int) (c []float64) {
	res := make([]float64, m)
	for i := 0; i < m; i++ {
		res[i] = a[i] + b[i]
	}
	return res
}
func norm(a, b []float64, m int) (c float64) {
	var res, tmp = 0.0, 0.0
	for i := 0; i < m; i++ {
		tmp += (a[i] - b[i]) * (a[i] - b[i])
	}
	res = math.Sqrt(tmp)
	return res
}

//
//func main() {
//	var m int
//	var eps float64
//	fmt.Println("Введите точность:")
//	fmt.Scan(&eps)
//	t := 0.0
//	fmt.Println("Введите размерность:")
//	fmt.Scan(&m)
//	mtr := scanmtr(m)
//	val := make([]float64, m)
//	fmt.Println("Введите правые части:")
//	for i := 0; i < m; i++ {
//		fmt.Scan(&t)
//		val[i] = t
//	}
//
//	fmt.Println("Метод итераций:")
//	var vec, tt = iter(mtr, val, m, eps) // mtr, val, m, eps
//	for i := 0; i < m; i++ {
//		fmt.Println(vec[i])
//	}
//	fmt.Println("Итерации:")
//	fmt.Println(tt)
//
//	fmt.Println("Метод Зейделя:")
//	var vec1, tt1 = zeyd(mtr, val, m, eps) // mtr, val, m, eps
//	for i := 0; i < m; i++ {
//		fmt.Println(vec1[i])
//	}
//	fmt.Println("Итерации:")
//	fmt.Println(tt1)
//}
