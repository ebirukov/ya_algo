package lesson1

import (
	"errors"
	"fmt"
)

// На некоторых кросс-платформенных станциях метро (как, например,
// «Третьяковская») на разные стороны платформы приходят поезда разных
// направлений. Таня договорилась встретиться с подругой на такой станции, но
// поскольку подруга приехала из другого часового пояса, то из-за джетлага сильно
// проспала, и Тане пришлось долго её ждать. Поезда всегда ходят точно по
// расписанию, и Таня знает, что поезд стоит на платформе ровно одну минуту, а
// интервал между поездами (время, в течение которого поезда у платформы нет)
// составляет a минут для поездов на первом пути и b минут для поездов на втором
// пути. То есть на первый путь приезжает поезд и стоит одну минуту, затем в
// течение a минут поезда у платформы нет, затем в течение одной минуты у
// платформы стоит следующий поезд и т. д. Пока Таня стояла на платформе, она
// насчитала n поездов на первом пути и m поездов на втором пути. Определите
// минимальное и максимальное время, которое Таня могла провести на платформе,
// или сообщите, что она точно сбилась со счёта. Все поезда, которые видела Таня,
// она наблюдала в течение всей минуты, то есть Таня не приходит и не уходит с
// платформы посередине той минуты, когда поезд стоит на платформе.

func Metro() {
	in := make([]int, 4)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if in[i] > 1000 || in[i] < 1 {
			panic("number out of range")
		}
		if err != nil {
			panic(err.Error())
		}
	}
	res, err := timesOfWatch(in[0], in[1], in[2], in[3])
	if err != nil {
		fmt.Printf("%d", -1)
	} else {
		fmt.Printf("%d %d", res[0], res[1])
	}
}

func timesOfWatch(a, b, n, m int) ([]int, error) {
	leftTimeInterval := [2]int{(a + 1) * n - a, (a + 1) * n + a}
	rightTimeInterval := [2]int{(b + 1) * m - b, (b + 1) * m + b}
	if rightTimeInterval[0] < leftTimeInterval[0] {
		leftTimeInterval, rightTimeInterval = rightTimeInterval, leftTimeInterval
	}
	if leftTimeInterval[1] < rightTimeInterval[0] {
		return nil, errors.New("interval is not intersect")
	} else if leftTimeInterval[1] > rightTimeInterval[1] {
		return []int{rightTimeInterval[0], rightTimeInterval[1]}, nil
	} else {
		return []int{rightTimeInterval[0], leftTimeInterval[1]}, nil
	}
}
