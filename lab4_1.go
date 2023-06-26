package main

import (
	"fmt"
	"math"
)

var H = 0.1
var K = [4]float64{}
var L = [4]float64{}
var z = [12]float64{}
var y = make([]float64, 12)
var y1 = make([]float64, 12)
var y2 = make([]float64, 12)

func Adams(x []float64, h int, H float64) {
	kutt(y2, 3, H)
	for i := 3; i < h; i++ { //10
		y2[i+1] = y2[i] + H*(55*z[i]-59*z[i-1]+37*z[i-2]-9*z[i-3])/24
		z[i+1] = z[i] + +H*(55*g(z[i], y[i], x[i])-59*g(z[i-1], y[i-1], x[i-1])+37*g(z[i-2], y[i-2], x[i-2])-9*g(z[i-3], y[i-3], x[i-3]))/24
	}
}

func g(z, y, x float64) float64 {

	return (z*2*math.Tan(x) + 3*y)
}

func kutt(y1 []float64, k int, H float64) {
	x := 0.0
	dy := 0.0
	dz := 0.0
	yr := 0.0
	fmt.Printf("| k |     xk    |        yk        |        zk       |      dyk      |       dzk      |     y ист     |       e       | \n")
	for i := 0; x <= 1 && i < k; i++ {
		K[0] = H * z[i]
		//fmt.Println(K[0])
		L[0] = H * g(z[i], y1[i], x)
		//fmt.Println(L[0])
		K[1] = H * (z[i] + 0.5*L[0])
		//fmt.Println(K[1])
		L[1] = H * g(z[i]+0.5*L[0], y1[i]+0.5*K[0], x+0.5*H)
		//fmt.Println(L[1])
		K[2] = H * (z[i] + 0.5*L[1])
		//fmt.Println(K[2])
		L[2] = H * g(z[i]+0.5*L[1], y1[i]+0.5*K[1], x+0.5*H)
		//fmt.Println(L[2])
		K[3] = H * (z[i] + L[2])
		//fmt.Println(K[3])
		L[3] = H * g(z[i]+L[2], y1[i]+K[2], x+H)
		//fmt.Println(L[3])
		dy = (K[0] + 2*K[1] + 2*K[2] + K[3]) / 6
		dz = (L[0] + 2*L[1] + 2*L[2] + L[3]) / 6
		y1[i+1] = y1[i] + dy
		z[i+1] = z[i] + dz
		yr = math.Cos(x)*math.Cos(x)*math.Cos(x) + math.Sin(x)*(1+2*math.Cos(x)*math.Cos(x))

		fmt.Printf("| %d | %f  |     %.6f     |    %.6f     |    %.6f   |    %.6f    |   %.6f    |   %.6f    |\n", i, x, y1[i], z[i], dy, dz, yr, yr-y[i])
		x += H

	}
}

//func main() {
//	z[0] = 3
//	y[0] = 1
//	y1[0] = 1
//	y2[0] = 1
//	x := make([]float64, 12)
//	rv := make([]float64, 12)
//	fmt.Println("Рунге кутт")
//	kutt(y1, 12, 0.1)
//	fmt.Println()
//	kutt(y2, 12, 0.2)
//	fmt.Println("Уточнение")
//	for i := 0; i < 6; i++ {
//		if i == 0 {
//			fmt.Println(runge(y2[i], y1[i], 0.5, 4))
//		} else {
//			fmt.Println(runge(y2[i], y1[2*i], 0.5, 4))
//		}
//	}
//	fmt.Println()
//	fmt.Println("Эйлер")
//	for i := 0; i < 5; i++ { //x < 1
//		y2[i+1] = y2[i] + 0.2*z[i]
//		z[i+1] = z[i] + 0.2*(2*math.Tan(x[i])*z[i]+3*y[i])
//		rv[i] = 0.25 * math.Exp(-math.Sqrt(2)*x[i]) * ((2+3*math.Sqrt(2))*math.Exp(2*math.Sqrt(2)*x[i]) + 2 - 3*math.Sqrt(2)) / math.Cos(x[i]) // math.Pow(math.Cos(x[i]), 3.0) + math.Sin(x[i])*(1+2*math.Pow(math.Cos(x[i]), 2.0)) // :)))))))))
//		fmt.Printf("y[i]: %f, real: %f \n", y2[i], rv[i])
//		x[i+1] = x[i] + 0.2
//	}
//	fmt.Println()
//	for i := 0; i < 10; i++ { //x < 1
//		y[i+1] = y[i] + H*z[i]
//
//		z[i+1] = z[i] + H*(2*math.Tan(x[i])*z[i]+3*y[i])
//		rv[i] = 0.25 * math.Exp(-math.Sqrt(2)*x[i]) * ((2+3*math.Sqrt(2))*math.Exp(2*math.Sqrt(2)*x[i]) + 2 - 3*math.Sqrt(2)) / math.Cos(x[i]) //math.Pow(math.Cos(x[i]), 3.0) + math.Sin(x[i])*(1+2*math.Pow(math.Cos(x[i]), 2.0))
//		fmt.Printf("y[i]: %f, real: %f \n", y[i], rv[i])
//		x[i+1] = x[i] + 0.1
//	}
//	fmt.Println("Уточнение")
//	for i := 0; i < 6; i++ {
//		if i == 0 {
//			fmt.Println(runge(y2[i], y[i], 0.5, 4))
//		} else {
//			fmt.Println(runge(y2[i], y[2*i], 0.5, 4))
//		}
//	}
//	fmt.Println("Adams")
//
//	Adams(x, 11, 0.1)
//	fmt.Println(y2)
//	p := plot.New()
//	p.Title.Text = "графики"
//	p.X.Label.Text = "X"
//	p.Y.Label.Text = "F(x)"
//
//	err := plotutil.AddLinePoints(p,
//		//	"Tабличные данные,", makepoints(val, res, N+1),
//		"euler", makepoints(x, y, 10),
//		"real", makepoints(x, rv, 10),
//		"kutt", makepoints(x, y1, 10),
//		"adams", makepoints(x, y2, 10),
//	)
//
//	if err != nil {
//		panic(err)
//	}
//
//	if err := p.Save(4*vg.Inch, 4*vg.Inch, "p.png"); err != nil {
//		panic(err)
//	}
//}
