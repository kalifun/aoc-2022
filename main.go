package main

import (
	"github.com/kalifun/aco-2022/aoc_go/logic"
	"github.com/kalifun/aco-2022/aoc_go/repo/decypt"
)

func main() {
	handles := decypt.NewDecyptHandle(
		logic.NewCalorie(),
	)
	handles.Decypt()
}
