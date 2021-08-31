package lesson1

import (
	"fmt"
	"math"
)

// Даны числа a, b, c, d, e, f. Решите систему линейных уравнений:
// ab + by = e
// cx + dy = f
// Вывод программы зависит от вида решения этой системы. Если система не
// имеет решений, то программа должна вывести единственное число 0. Если система
// имеет бесконечно много решений, каждое из которых имеет вид y=kx+b, то
// программа должна вывести число 1, а затем значения k и b. Если система имеет
// единственное решение (x0,y0), то программа должна вывести число 2, а затем
// значения x0 и y0. Если система имеет бесконечно много решений вида x=x0, y —
// любое, то программа должна вывести число 3, а затем значение x0. Если система
// имеет бесконечно много решений вида y=y0, x — любое, то программа должна
// вывести число 4, а затем значение y0. Если любая пара чисел (x,y) является
// решением, то программа должна вывести число 5. Числа x0 и y0 будут проверяться
// с точностью до пяти знаков после точки.

func SysLinearEquations2() {
	in := make([]float32, 6)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			panic(err.Error())
		}
	}
	res := solveSysEquations(in[0], in[1], in[2], in[3], in[4], in[5])
	fmt.Printf("%v ", res[0])
	if len(res) > 1 {
		fmt.Printf("%v ", float32(math.Round(float64(res[1]*100000))/100000))
	}
	if len(res) > 2 {
		fmt.Printf("%v", float32(math.Round(float64(res[2]*100000))/100000))
	}
}

func solveSysEquations(a, b, c, d, e, f float32) (res []float32) {
	// x,y - any
	if a == 0 && b == 0 && c == 0 && d == 0 {
		return []float32{5}
	}
	// x - any, y = e / b or f / d
	if a == 0 && c == 0 {
		if b == 0 {
			// exchange equation
			a, b, c, d, e, f = c, d, a, b, f, e
		}
		if e/b == f/d || d == 0 {
			return []float32{4, e / b}
		}
		return []float32{0}
	}
	// y - any, x = e / a
	if b == 0 && d == 0 {
		if a == 0 {
			// exchange equation
			a, b, c, d, e, f = c, d, a, b, f, e
		}
		if e/a == f/c || c == 0 {
			return []float32{3, e / a}
		}
		return []float32{0}
	}
	// exchange equation for select main
	if a == 0 && c != 0 {
		// exchange equation
		a, b, c, d, e, f = c, d, a, b, f, e
	}
	// gauss solve
	a, b, e = 1, b/a, e/a
	if c != 0 {
		c, d, f = 1, d/c, f/c
		c, d, f = c-a, d-b, f-e
	}
	if c == 0 && d == 0 {
		if f == 0 {
			return []float32{1, -a / b, e / b}
		}
		return []float32{0}
	}
	y := f / d
	x := e - b*d/f
	return []float32{2, x, y}
}
