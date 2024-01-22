package orchestration

import "testing"

func BenchmarkStart(b *testing.B) {
	Start(b.N, 10000000000)
}
