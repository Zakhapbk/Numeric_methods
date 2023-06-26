package main

import (
	"fmt"
	"gonum.org/v1/plot/plotter"
	"math"
)

var val = []float64{-1.0, 0.0, 1.0, 2.0, 3.0, 4.0}
var res = []float64{-0.4597, 1.0, 1.5403, 1.5839, 2.010, 3.3464}

func makesist(x, y []float64, k, N int) []float64 { // к is power of polynomial + 1
	A := make([][]float64, k)
	for i := range A {
		A[i] = make([]float64, k)
	}
	B := make([]float64, k)
	F := make([]float64, N+1)
	A[0][0] = float64(N + 1)
	p := 0

	for i := 0; i < k; i++ {
		if i == 0 {
			p = i + 1
		} else {
			p = i * 2
		}
		for j := i; j < k; j++ {
			if (j == 0) && (j == i) {
				continue
			}
			for t := 0; t <= N; t++ {
				A[i][j] += math.Pow(x[t], float64(p))
				if j == k-1 {
					B[i] += math.Pow(x[t], float64(i)) * y[t]
				}
			}
			if i != j {
				A[j][i] = A[i][j]
			}
			p++
		}
	}

	_, _, _, res := LU(A, B, k)
	f := polinom(k, res)
	for i := 0; i < N+1; i++ {
		F[i] = f(x[i])
	}
	fmt.Println("Сумма квадратов ошибок:", accuracy(F, y, N))
	return F
}

func accuracy(a, b []float64, N int) float64 {
	res := 0.0
	for i := 0; i < N+1; i++ {
		res += (a[i] - b[i]) * (a[i] - b[i])
	}
	return res
}

func polinom(k int, arr []float64) func(x float64) float64 {
	f := func(x float64) float64 {
		sum := 0.0
		for i := 0; i < k; i++ {
			sum += arr[i] * math.Pow(x, float64(i))
		}
		return sum
	}
	return f
}
func makepoints(x, y []float64, N int) plotter.XYs {
	pts := make(plotter.XYs, N)
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}

	return pts
}

//
//func main() {
//	N := 5
//	F1 := makesist(val, res, 2, N)
//	F2 := makesist(val, res, 3, N)
//	p := plot.New()
//
//	fmt.Printf("Mногочлен 1 степени: %.5f + %.5fx\n", F1[0], F1[1])
//	fmt.Printf("Mногочлен 2 степени: %.5f + %.5fx + %.5fx^2 \n", F2[0], F2[1], F2[2])
//
//	p.Title.Text = "MHK"
//	p.X.Label.Text = "X"
//	p.Y.Label.Text = "F(x)"
//
//	err := plotutil.AddLinePoints(p,
//		//	"Tабличные данные,", makepoints(val, res, N+1),
//		"Mногочлен 1 степени", makepoints(val, F1, N+1),
//		"Mногочлен 2 степени", makepoints(val, F2, N+1),
//	)
//
//	if err != nil {
//		panic(err)
//	}
//	s, err := plotter.NewScatter(makepoints(val, res, N+1))
//	s.Shape = draw.CircleGlyph{}
//	p.Add(s)
//	p.Legend.Add("Tабличные данные", s)
//	if err != nil {
//		panic(err)
//	}
//
//	// Save the plot to a PNG file.
//	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
//		panic(err)
//	}
//
//}
