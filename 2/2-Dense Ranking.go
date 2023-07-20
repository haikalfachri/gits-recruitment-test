package main

import (
	"fmt"
	"sort"
)
func denseRanking(scores []int, score int) int {
	var ranks []int
	var rank int = 1
	var result int

	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	for idx := range scores {
		if idx != len(scores)-1 {
			if scores[idx] != scores[idx+1] {
				ranks = append(ranks, rank)
				rank++
			} else {
				ranks = append(ranks, rank)
			}
		} else {
			ranks = append(ranks, rank)
		}

	}

	for idx, s := range scores{
		if s == score {
			result = ranks[idx]
		}
	}
	return result
}

func main() {
	var n_player int
	var scores []int
	var score int
	var n_GITS int
	var rank_result []int

	fmt.Scanln(&n_player)
	for i := 0; i < n_player; i++ {
		fmt.Scan(&score)
		scores = append(scores, score)
	}
	fmt.Scanln()

	fmt.Scanln(&n_GITS)
	for i := 0; i < n_GITS; i++ {
		fmt.Scan(&score)
		scores = append(scores, score)
		rank_result = append(rank_result, denseRanking(scores, score))
	}
	fmt.Scanln()
	
	fmt.Println("----Output----")
	for _, rank := range(rank_result){
		fmt.Print(rank, " ")
	}
}
