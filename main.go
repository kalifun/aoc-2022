package main

import (
	"github.com/kalifun/aco-2022/aoc_go/logic"
	"github.com/kalifun/aco-2022/aoc_go/repo/decrypt"
)

func main() {
	handles := decrypt.NewDecyptHandle(
		logic.NewCalorie(),
		logic.NewMorraGame(),
		logic.NewRuclSack(),
		logic.NewCampCleanup(),
	)
	handles.Decypt()
}
