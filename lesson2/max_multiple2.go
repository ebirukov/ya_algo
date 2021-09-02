package lesson2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MaxMultiple2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	seq := numbers(scanner.Text())
	first, second := findMaxMultiple2(seq)
	fmt.Printf("%d %d", first, second)
}

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func kthStats(seq []int, left, right, k int) []int {
	for left < right {
		pivot := seq[(left+right)/2]
		ptr := left
		pivotPtr := left
		fmt.Printf("left=%v right=%v\n", left, right)
		for i := left; i < right+1; i++ {
			current := seq[i]
			if current < pivot {
				seq[i] = seq[ptr]
				seq[ptr] = seq[pivotPtr]
				seq[pivotPtr] = current
				ptr++
				pivotPtr++
			} else if current == pivot {
				seq[i], seq[ptr] = seq[ptr], current
				ptr++
			}
			fmt.Printf("%v\n", seq)
		}
		fmt.Printf("%v\n", ptr)
		if ptr <= k {
			left = ptr
		} else if pivotPtr > k {
			right = pivotPtr - 1
		} else {
			return seq
		}
	}
	return seq
}

func _findMaxMultiple2(seq []int) (int, int) {
	len := len(seq)
	seq = kthStats(seq, 0, len-1, len-1)
	seq = kthStats(seq, 0, len-2, len-2)
	seq = kthStats(seq, 0, len-3, 2)
	if seq[len-1]*seq[len-2] > seq[0]*seq[1] {
		return seq[len-2], seq[len-1]
	} else {
		return seq[0], seq[1]
	}
}

func findMaxMultiple2(seq []int) (int, int) {
	if len(seq) == 2 {
		if seq[0] > seq[1] {
			return seq[1], seq[0]
		}
		return seq[0], seq[1]
	}
	maxNegative := []int{0, 0}
	maxPositive := []int{0, 0}
	for i := 0; i < len(seq); i++ {
		if seq[i] > 0 {
			if seq[i] > maxPositive[0] {
				maxPositive[0], maxPositive[1] = seq[i], maxPositive[0]
			} else if seq[i] > maxPositive[1] {
				maxPositive[1] = seq[i]
			}
		} else if seq[i] < 0 {
			if seq[i] < maxNegative[0] {
				maxNegative[0], maxNegative[1] = seq[i], maxNegative[0]
			} else if seq[i] < maxNegative[1] {
				maxNegative[1] = seq[i]
			}
		}
	}
	//if -1 * maxNegative[0] > maxPositive[0] || -1 * maxNegative[1] > maxPositive[1] {
	if maxNegative[0]*maxNegative[1] > maxPositive[0]*maxPositive[1] {
		return maxNegative[0], maxNegative[1]
	}
	return maxPositive[1], maxPositive[0]
}
