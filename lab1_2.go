package main

func prog(mtr [][]float64, val []float64, m int) (X, P, Q []float64) {
	v := make([]float64, m)
	u := make([]float64, m)
	x := make([]float64, m)
	v[0] = mtr[0][1] / (-mtr[0][0])
	u[0] = val[0] / mtr[0][0]
	for i := 1; i < m-1; i++ {
		v[i] = mtr[i][i+1] / (-mtr[i][i] - mtr[i][i-1]*v[i-1])
		u[i] = (mtr[i][i-1]*u[i-1] - val[i]) / (-mtr[i][i] - mtr[i][i-1]*v[i-1])
	}
	u[m-1] = (mtr[m-1][m-2]*u[m-2] - val[m-1]) / (-mtr[m-1][m-1] - mtr[m-1][m-2]*v[m-2])
	x[m-1] = u[m-1]
	for i := m - 1; i > 0; i-- {
		x[i-1] = v[i-1]*x[i] + u[i-1]
	}
	return x, v, u
}

//func main() {
//	var m int
//	t := 0.0
//	fmt.Println("Введите размерность:")
//	fmt.Scan(&m)
//	mtr := make([][]float64, m)
//	val := make([]float64, m)
//
//	for i := range mtr {
//		mtr[i] = make([]float64, m)
//	}
//	fmt.Println("Введите матрицу:")
//	for i := 0; i < m; i++ {
//		if i == 0 || i == m-1 {
//			if i == 0 {
//				fmt.Scan(&t)
//				mtr[i][i] = t
//				fmt.Scan(&t)
//				mtr[i][i+1] = t
//			} else {
//				fmt.Scan(&t)
//				mtr[i][i-1] = t
//				fmt.Scan(&t)
//				mtr[i][i] = t
//			}
//			continue
//		}
//		for j := i - 1; j < i+2; j++ {
//			fmt.Scan(&t)
//			mtr[i][j] = t
//		}
//	}
//	fmt.Println("Введите правые части:")
//	for i := 0; i < m; i++ {
//		fmt.Scan(&t)
//		val[i] = t
//	}
//	fmt.Println("Исходная матрица:")
//	prnt(mtr, m)
//	fmt.Println()
//	X, P, Q := prog(mtr, val, m)
//	fmt.Println("Результат:")
//	for _, val := range X {
//		fmt.Print(val, " ") //fmt.Print(math.Round(val), " ")
//	}
//	fmt.Println()
//	fmt.Println("P:")
//	fmt.Println(P)
//	fmt.Println("Q:")
//	fmt.Println(Q)
//}
