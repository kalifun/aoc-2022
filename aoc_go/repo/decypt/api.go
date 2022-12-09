package decypt

import (
	"log"

	"github.com/kalifun/aco-2022/aoc_go/entity/proto"
)

// decyptHandle
type decyptHandle struct {
	handles []proto.CollectStars
}

// NewDecyptHandle
//  @param handles
//  @return decyptHandle
func NewDecyptHandle(handles ...proto.CollectStars) decyptHandle {
	return decyptHandle{
		handles: handles,
	}
}

// Decypt
//  @receiver d
func (d decyptHandle) Decypt() {
	for _, handle := range d.handles {
		err := handle.GetStar()
		if err != nil {
			log.Printf("decypt err: %s\n", err.Error())
		}
	}
}
