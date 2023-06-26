package main

var f = [5]float64{1.0, 1.5403, 1.5839, 2.01, 3.3464}
var x = [5]float64{0.0, 1.0, 2.0, 3.0, 4.0}

func h(i int) float64 {
	if i == 0 {
		return x[i]
	} else {
		return x[i] - x[i-1]
	}
}

func crsist(m int) []float64 { // m is number of knots-2
	mtr := make([][]float64, m)
	val := make([]float64, m)
	for i := range mtr {
		mtr[i] = make([]float64, m)
	}

	for i := 0; i < m; i++ {
		if i == 0 || i == m-1 {
			if i == 0 {
				mtr[i][i] = 2 * (h(1) + h(2))
				mtr[i][i+1] = h(2)
				val[i] = 3 * ((f[2]-f[1])/h(2) - (f[1]-f[0])/h(1))
			} else {
				mtr[i][i-1] = h(i + 1)
				mtr[i][i] = 2 * (h(i+1) + h(i+2))
				val[i] = 3 * ((f[m+1]-f[m])/h(m+1) - (f[m]-f[m-1])/h(m))
			}
			continue
		}
		for j := i - 1; j < i+2; j++ {
			if j == i-1 {
				mtr[i][j] = h(i + 1)
			} else if j == i {
				mtr[i][j] = 2 * (h(i+1) + h(i+2))
			} else {
				mtr[i][j] = h(i + 2)
			}
		}
		val[i] = 3 * ((f[i+2]-f[i+1])/h(i+2) - (f[i+1]-f[i])/h(i+1))
	}
	c := make([]float64, 1)
	tmp, _, _ := prog(mtr, val, m)
	c = append(c, tmp...)

	return c
}

func fsec(y float64) (int, int) {
	a, b := 0, 0
	for i := 0; i < len(x)-1; i++ {
		if y < x[i] {
			a, b = i-1, i
		}
	}
	return a, b
}

//func main() {
//	a := [len(x) - 1]float64{}
//	b := [len(x) - 1]float64{}
//	d := [len(x) - 1]float64{}
//	c := crsist(3)
//	// сoeifficent numeration starts with 1
//	for i := 0; i < len(x)-2; i++ {
//		a[i] = f[i]
//		b[i] = (f[i+1]-f[i])/h(i+1) - h(i+1)*(c[i+1]+2*c[i])/3
//		d[i] = (c[i+1] - c[i]) / (3 * h(i+1))
//	}
//	a[len(x)-2] = f[len(x)-2]
//	b[len(x)-2] = (f[len(x)-1]-f[len(x)-2])/h(len(x)-1) - 2*(h(len(x)-1)*c[len(x)-2])/3
//	d[len(x)-2] = -c[len(x)-2] / (3 * h(len(x)-1))
//	X := 1.5
//	l, r := fsec(X)
//	fmt.Println("Коэффициенты:")
//	fmt.Println("а: ", a)
//	fmt.Println("b: ", b)
//	fmt.Println("c: ", c)
//	fmt.Println("d: ", d)
//	fmt.Println("Многочлен")
//	fmt.Printf("%.5f + %.5f(X-%.5f) + %.5f(X-(%.5f))^2 + %.5f(X-%.5f)^3\n", a[r-2], b[r-2], x[l-1], c[r-2], x[l-1], d[r-2], x[l-1])
//	fmt.Print("f(X*): ")
//	fmt.Println(a[r-2] + b[r-2]*(X-x[l-1]) + c[r-2]*(X-x[l-1])*(X-x[l-1]) + d[r-2]*(X-x[l-1])*(X-x[l-1])*(X-x[l-1]))
//
//}
