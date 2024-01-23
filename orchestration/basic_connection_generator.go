package orchestration

import (
	"fmt"
	"github.com/vposloncec/go-ssip/base"
	"math/rand"
	"slices"
)

func GenConnectionPairs(minId int, maxId int, totalNum int) (pairs []base.ConnectionPair) {
	// Max number of connections in N(N - 1) / 2
	nodesNum := maxId - minId + 1
	connMax := nodesNum * (nodesNum - 1) / 2
	if totalNum > connMax {
		fmt.Printf("Given number of connections (%v) is larger than maximum (%v) "+
			"for given network, using max instead\n", totalNum, connMax)
	}
	totalNum = min(totalNum, connMax)

	uniqueMap := make(map[base.ConnectionPair]struct{})
	uniqueMap2 := make(map[int]struct{})
	for len(uniqueMap) < totalNum {
		// Use UniqueRand to avoid self loops (where connection pair has 2 same nodeIDs)
		u := newUniqueRand()
		ids := []int{u.Int(minId, maxId), u.Int(minId, maxId)}

		// Pairs are sorted because (1,2) and (2,1) is the same thing
		slices.Sort(ids)
		pair := [2]int{ids[0], ids[1]}

		uniqueMap[pair] = struct{}{}
		uniqueMap2[rand.Intn(1000)] = struct{}{}
	}

	return mapToKeySlice(uniqueMap)
}

func mapToKeySlice(origMap map[base.ConnectionPair]struct{}) []base.ConnectionPair {
	keys := make([]base.ConnectionPair, len(origMap))

	i := 0
	for k := range origMap {
		keys[i] = k
		i++
	}

	return keys
}

func newUniqueRand() *uniqueRand {
	return &uniqueRand{generated: make(map[int]bool)}
}

type uniqueRand struct {
	generated map[int]bool
}

func (u *uniqueRand) Int(min int, max int) int {
	for {
		i := rand.Intn(max-min+1) + min
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}

func (u *uniqueRand) Exclude(num int) {
	u.generated[num] = true
}
