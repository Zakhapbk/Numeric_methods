package main

import (
	"math"
)

var X = 1.0

var kn1 = []float64{0, math.Pi / 6, math.Pi / 3, math.Pi / 2}

//var kn1 = []float64{0.0, 1.0, 2.0, 3.0} // 2nd variant of the net
var ValInKn [4]float64

func vals() {
	for i := range ValInKn {
		ValInKn[i] = inp(kn1[i])
	}
}

func inp(x float64) float64 {
	return math.Cos(x) + x
}

func acc(A, B func(x float64) float64, x float64) float64 {
	return math.Abs(math.Abs(A(x)) - math.Abs(B(x)))
}

func omega(x float64) float64 {
	z := 1.0
	for i := range kn1 {
		z *= x - kn1[i]
	}
	return z
}

func nDifOmega(n int) float64 {
	z := 1.0
	for i := range kn1 {
		if i == n {
			continue
		}
		z *= kn1[n] - kn1[i]
	}
	return z
}

func Lagrange(x float64) float64 {
	res := 0.0
	for i := range kn1 {
		res += ValInKn[i] * omega(x) / ((x - kn1[i]) * nDifOmega(i))
	}
	return res
}

func fdel(a, b, c int) float64 {
	if c == 2 {
		return (ValInKn[a] - ValInKn[b]) / (kn1[a] - kn1[b])
	} else {
		return (fdel(a, b-1, c-1) - fdel(a+1, b, c-1)) / (kn1[a] - kn1[b])
	}

}

func Newton(x float64) float64 {
	sum := ValInKn[0]
	mult := x - kn1[0]
	sum += mult * fdel(1, 0, 2)
	for i := 1; i < 3; i++ {
		mult *= x - kn1[i]
		sum += mult * fdel(0, i+1, i+2)
	}
	return sum
}

//
//func main() {
//	vals()
//	fmt.Println("Лагранж")
//	fmt.Println("Значение в точке X*:", Lagrange(X))
//	fmt.Println("Погрешность: ", acc(Lagrange, inp, X))
//	for i := 0; i < len(kn1); i++ {
//		fmt.Printf("+ (%.5f)", ValInKn[i]/nDifOmega(i))
//		for j := 0; j < len(kn1); j++ {
//			if i == j {
//				continue
//			}
//			fmt.Printf("(x-%.5f)", kn1[j])
//		}
//		fmt.Print(" ")
//
//	}
//	fmt.Println()
//	fmt.Println()
//	fmt.Println("Ньютон")
//	fmt.Println("Значение в точке X*: ", Newton(X))
//	fmt.Println("Погрешность: ", acc(Newton, inp, X))
//	for i := 0; i < len(kn1); i++ {
//		if i == 0 {
//			fmt.Print(ValInKn[0], " ")
//			continue
//		}
//		if i == 1 {
//			fmt.Printf("+ (%.5f)", fdel(1, 0, 2))
//		} else {
//			fmt.Printf("+ (%.5f)", fdel(0, i, i+1))
//		}
//		for j := 0; j < i; j++ {
//
//			fmt.Printf("(x-%.5f)", kn1[j])
//		}
//		fmt.Print(" ")
//
//	}
//
//}
