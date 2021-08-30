package lesson1

import "fmt"

// Решите в целых числах уравнение: SQRT(ax + b) = c a, b, c – данные целые
// числа: найдите все решения или сообщите, что решений в целых числах нет.

func RootEquation() {
	in := make([]int, 3)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			panic(err.Error())
		}
	}
	res := findingSolution(in[0], in[1], in[2])
	fmt.Print(res)
}

func findingSolution(a, b, c int) string {
	if c < 0 {
		return "NO SOLUTION"
	}
	if a == 0 && b == c*c {
		return "MANY SOLUTIONS"
	}
	if a != 0 && (c*c - b) % a == 0  {
		res := (c*c - b) / a
		return fmt.Sprintf("%d", res)
	}
	return "NO SOLUTION"
}
