package main

import (
	"github.com/kalifun/aco-2022/logic"
	"github.com/kalifun/aco-2022/repo/decrypt"
)

func main() {
	handles := decrypt.NewDecyptHandle(
		logic.NewCalorie(),
		logic.NewMorraGame(),
		logic.NewRuclSack(),
		logic.NewCampCleanup(),
		logic.NewSupplyStacks(),
		logic.NewTuningTrouble(),
		logic.NewDevice(),
	)
	handles.Decypt()
}
