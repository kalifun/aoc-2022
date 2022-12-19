package decrypt

import (
	"log"

	"github.com/kalifun/aco-2022/entity/proto"
)

// decryptHandle  TODO
type decryptHandle struct {
	handles []proto.CollectStars
}

// NewDecyptHandle
//
//	@param handles
//	@return decyptHandle
func NewDecyptHandle(handles ...proto.CollectStars) decryptHandle {
	return decryptHandle{
		handles: handles,
	}
}

// Decypt
//
//	@receiver d
func (d decryptHandle) Decypt() {
	for _, handle := range d.handles {
		err := handle.GetStar()
		if err != nil {
			log.Printf("decypt err: %s\n", err.Error())
		}
	}
}
