package lesson1

import (
	"fmt"
)
// За многие годы заточения узник замка Иф проделал в стене прямоугольное
// отверстие размером D × E. Замок Иф сложен из кирпичей, размером A × B × C.
// Определите, сможет ли узник выбрасывать кирпичи в море через это отверстие, если
// стороны кирпича должны быть параллельны сторонам отверстия.

func PrisonerCastle() {
	in, err := readArgs(5)
	if err != nil {
		panic(err.Error())
	}
	pass := isPassBlockThrowHole(in[:3], in[3:])
	if pass {
		fmt.Printf("%s", "YES")
	} else {
		fmt.Printf("%s", "NO")
	}
}

func isPassBlockThrowHole(block []int, hole []int) bool {
	for i, value := range block {
		for j := 0; j < len(hole); j++ {
			if value <= 0 || hole[j] <= 0 {
				continue
			}
			if hole[j] >= value {
				block = remove(block, i)
				hole = remove(hole, j)
			}
		}
	}
	return len(hole) == 0
}

func remove(array []int, i int) []int {
	if i+1 >= len(array) {
		return array[:i]
	}
	return append(array[:i], array[i+1:]...)
}

func readArgs(argsCnt int) ([]int, error) {
	in := make([]int, argsCnt)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			return in[:1], err
		}
	}
	return in, nil
}
