package orchestration

import (
	"github.com/vposloncec/go-ssip/base"
	"math/rand"
)

func GetRandomNeighbours(nodeId base.NodeID, minId int, maxId int, amount int) (n []int) {
	u := newUniqueRand()
	u.Exclude(int(nodeId))

	for i := 0; i < amount; i++ {
		n = append(n, u.Int(minId, maxId))
	}
	return
}

func newUniqueRand() *UniqueRand {
	return &UniqueRand{generated: make(map[int]bool)}
}

type UniqueRand struct {
	generated map[int]bool
}

func (u *UniqueRand) Int(min int, max int) int {
	for {
		i := rand.Intn(max-min+1) + min
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}

func (u *UniqueRand) Exclude(num int) {
	u.generated[num] = true
}
