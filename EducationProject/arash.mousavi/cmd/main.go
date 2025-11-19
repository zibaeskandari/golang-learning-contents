package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, l int
	if _, err := fmt.Fscan(in, &n, &l); err != nil {
		return
	}

	weights := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &weights[i])
	}

	capacities := make([]int64, l)
	for i := 0; i < l; i++ {
		fmt.Fscan(in, &capacities[i])
	}

	sort.Slice(weights, func(i, j int) bool {
		return weights[i] > weights[j]
	})

	var score int64

	for _, w := range weights {

		for level := l - 1; level >= 0; level-- {
			if capacities[level] >= w {
				capacities[level] -= w
				levelIndex := int64(level + 1)
				score += w * levelIndex
				break
			}
		}

	}

	fmt.Println(score)
}
