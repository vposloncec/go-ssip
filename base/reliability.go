package base

import (
	"math/rand"
)

//go:generate stringer -type=ReliabilityLevel
type ReliabilityLevel int

const (
	Super ReliabilityLevel = iota
	Reliable
	Common
	Occasional
	Erratic
)

var categories = map[ReliabilityLevel]float64{
	Super:      0.98,
	Reliable:   0.85,
	Common:     0.75,
	Occasional: 0.60,
	Erratic:    0.40,
}

// NewReliability returns a random ReliabilityLevel.
func NewReliability() ReliabilityLevel {
	return ReliabilityLevel(rand.Intn(len(categories)))
}

// ShouldDropPacket returns true if the packet should be dropped based on the category's chance.
func ShouldDropPacket(lvl ReliabilityLevel) bool {
	return rand.Float64() > categories[lvl]
}
