package main

import (
	"math"
)

func rectangle(h, l, r float64) float64 {
	res := 0.0
	x0, xn := l, l+h
	for {
		res += und_int((x0 + xn) / 2)
		x0 = xn
		xn += h
		if xn > r {
			break
		}
	}
	return res * h
}

func trapezoid(h, l, r float64) float64 {
	res := 0.0
	xn := l
	for {
		if (xn == l) || (xn == r) {
			res += und_int(xn) / 2
		} else {
			res += und_int(xn)
		}
		xn += h
		if xn > r {
			break
		}
	}
	return res * h
}

func simpson(h, l, r float64) float64 {
	res := 0.0
	xn := l
	for i := 0; ; i++ {
		if (xn == l) || (xn == r) {
			res += und_int(xn)
		} else if i%2 == 1 {
			res += 4 * und_int(xn)
		} else {
			res += 2 * und_int(xn)
		}

		xn += h
		if xn > r {
			break
		}
	}
	return res * h / 3
}

func runge(F1, F2 float64, k, p float64) float64 { // k - the enlargment of the net. р - the order of accuracy
	return F1 + (F1-F2)/(math.Pow(float64(k), float64(p))-1)
}

func und_int(x float64) float64 {
	return x * x / (x*x*x - 27)
}

//func main() {
//	fmt.Println("Метод прямоугольников:", rectangle(1.0, -2, 2))
//	fmt.Println("Метод трапеций:", trapezoid(1.0, -2, 2))
//	fmt.Println("Метод Симпсона:", simpson(1.0, -2, 2))
//	fmt.Println("Шаг h2")
//	fmt.Println("Метод прямоугольников:", rectangle(0.5, -2, 2))
//	fmt.Println("Метод трапеций:", trapezoid(0.5, -2, 2))
//	fmt.Println("Метод Симпсона:", simpson(0.5, -2, 2))
//
//	a, b, c := runge(rectangle(1.0, -2, 2), rectangle(0.5, -2, 2), 0.5, 2), runge(trapezoid(1.0, -2, 2), trapezoid(0.5, -2, 2), 0.5, 2), runge(simpson(1.0, -2, 2), simpson(0.5, -2, 2), 0.5, 4)
//	fmt.Printf("Уточнение прямоугольников: %f Погрешность: %f \n", a, math.Abs(math.Abs(a)-math.Abs(-0.16474)))
//	fmt.Printf("Уточнение трапеций: %f Погрешность: %f \n", b, math.Abs(math.Abs(b)-math.Abs(-0.16474)))
//	fmt.Printf("Уточнение Симпсона: %f Погрешность: %f \n", c, math.Abs(math.Abs(c)-math.Abs(-0.16474)))
//}
