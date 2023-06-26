package main

import (
	"fmt"
	"math"
)

func difference(h float64, m int, x []float64) {
	val := make([]float64, m)
	mtr := make([][]float64, m)
	for i := range mtr {
		mtr[i] = make([]float64, m)
	}
	mtr[0][0], mtr[0][1] = -2+h*h*((-math.Exp(x[1]))/(math.Exp(x[1])+1)), 1+(-h/(math.Exp(x[1])+1))
	val[0] = -0.5 * (1 + h/(math.Exp(x[1])+1))
	mtr[m-1][m-2], mtr[m-1][m-1] = (1 - (-h / (math.Exp(x[m]) + 1))), -2+h*h*(-math.Exp(x[m]))/(math.Exp(x[m])+1)
	val[m-1] = -(math.E - 1 + (1 / (math.E + 1))) * (1 + (-h / (math.Exp(x[m]) + 1)))
	for i := 1; i < m-1; i++ {

		mtr[i][i-1] = 1 + (h / (math.Exp(x[i+1]) + 1))
		mtr[i][i] = -2 + h*h*(-math.Exp(x[i+1])/(math.Exp(x[i+1])+1))
		mtr[i][i+1] = 1 - (h / (math.Exp(x[i+1]) + 1))
		val[i] = 0

	}
	//fmt.Println(val)
	//prnt(mtr, m)
	res, _, _ := prog(mtr, val, m)
	res1 := make([]float64, len(val)+1)
	res1[0] = 0.5
	for i := range res {
		res1[i+1] = res[i]
	}
	res1 = append(res1, math.E-1+(1/(math.E+1)))
	for i := range res1 {
		fmt.Printf("%.4f ", res1[i])
	}
	fmt.Println()

}

func shooting() {
	fmt.Println("Метод стрельбы")
	y := make([]float64, 20)
	n0 := 1.0 //1.0
	n1 := 0.8
	nn := 0.0
	y[0] = 0.5 //1.0
	z[0] = math.E - 1 + 1/(math.E+1)

	y0, _ := mkutt(y, 26, 0.1)
	z[0] = n0
	y1, _ := mkutt(y, 26, 0.1)

	for {
		nn = n1 - (y1-(math.E-1+1/(math.E+1)))*(n1-n0)/(y1-y0)
		fmt.Println(nn)
		z[0] = nn
		y0 = y1
		y1, _ = mkutt(y, 26, 0.1)
		fmt.Println(math.E - 1 + 1/(math.E+1))
		if math.Abs(y1-(math.E-1+1/(math.E+1))) < 0.0001 {
			break
		}
		n0 = n1
		n1 = nn
	}
	y[0] = 0.5
	_, result := mkutt(y, 26, 0.1)
	fmt.Print("Приближение: ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%.7f ", result[i])
	}
	fmt.Println()
	nn = 0
	fmt.Print("x: ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%.7f ", nn)
		nn += 0.1
	}
}

func g1(z, y, x float64) float64 {

	return (2*z + math.Exp(x)*y) / (math.Exp(x) + 1)
}

func mkutt(y1 []float64, k int, H float64) (float64, []float64) {
	x := 0.0
	dy := 0.0
	dz := 0.0

	for i := 0; x <= 1 && i < k; i++ {
		K[0] = H * z[i]
		//fmt.Println(K[0])
		L[0] = H * g1(z[i], y1[i], x)
		//fmt.Println(L[0])
		K[1] = H * (z[i] + 0.5*L[0])
		//fmt.Println(K[1])
		L[1] = H * g1(z[i]+0.5*L[0], y1[i]+0.5*K[0], x+0.5*H)
		//fmt.Println(L[1])
		K[2] = H * (z[i] + 0.5*L[1])
		//fmt.Println(K[2])
		L[2] = H * g1(z[i]+0.5*L[1], y1[i]+0.5*K[1], x+0.5*H)
		//fmt.Println(L[2])
		K[3] = H * (z[i] + L[2])
		//fmt.Println(K[3])
		L[3] = H * g1(z[i]+L[2], y1[i]+K[2], x+H)
		//fmt.Println(L[3])
		dy = (K[0] + 2*K[1] + 2*K[2] + K[3]) / 6
		dz = (L[0] + 2*L[1] + 2*L[2] + L[3]) / 6
		y1[i+1] = y1[i] + dy
		z[i+1] = z[i] + dz

		x += H

	}

	return y1[10], y1
}

//func main() {
//	//x := make([]float64, 6)
//	//for i := 1; i < 6; i++ {
//	//	x[i] = x[i-1] + 0.2
//	//}
//	x := make([]float64, 11)
//	for i := 1; i < 11; i++ {
//		x[i] = x[i-1] + 0.1
//
//	}
//	fmt.Println("Разностный метод:")
//	fmt.Println("Приближение")
//	difference(0.1, 9, x)
//	fmt.Println("Действительные значения")
//	for i := range x {
//		fmt.Printf("%.4f ", math.Exp(x[i])-1+1/(math.Exp(x[i])+1))
//	}
//
//	z[0] = 0.8
//	y1[0] = 1.0
//	//fmt.Println(mkutt(y1, 30, 0.1))
//	fmt.Println("Метод стрельбы:")
//	fmt.Println()
//	shooting()
//	fmt.Println()
//	fmt.Print("Действительные: ")
//	for i := 0; i <= 10; i++ {
//		fmt.Printf("%.7f ", math.Exp(x[i])-1+1/(math.Exp(x[i])+1))
//
//	}
//
//}
