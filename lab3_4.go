package main

func ldif1(x, y []float64, l, r int) float64 {
	return (y[r] - y[l]) / (x[r] - x[l])
}

func lldif1(x, y []float64, val float64, l, r int) float64 {
	return ldif1(x, y, l, r) + (ldif1(x, y, r, r+1)-ldif1(x, y, l, r))/(x[r+1]-x[l])*(2*val-x[l]-x[r])
}

func ldif2(x, y []float64, l, r int) float64 {
	return 2 * (ldif1(x, y, r+1, r) - ldif1(x, y, l, r)) / (x[r+1] - x[l])
}

//func main() {
//	x := []float64{0.2, 0.5, 0.8, 1.1, 1.4}
//	y := []float64{12.906, 5.5273, 3.8777, 3.2692, 3.0319}
//	fmt.Println("Левая производная")
//	fmt.Println(ldif1(x, y, 1, 2))
//	fmt.Println("Правая производная")
//	fmt.Println(ldif1(x, y, 2, 3))
//	fmt.Println("Второй порядок точности")
//	fmt.Println(lldif1(x, y, 0.8, 1, 2))
//	fmt.Println("Вторая производная")
//	fmt.Println(ldif2(x, y, 1, 2))
//}
