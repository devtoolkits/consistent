package rings

import (
	"testing"
)

var nodes = []string{"node-001", "node-002", "node-003", "node-004", "node-005", "node-006", "node-007", "node-008", "node-009", "node-010", "node-011", "node-012", "node-013", "node-014", "node-015", "node-016", "node-017", "node-018", "node-019", "node-020", "node-021", "node-022", "node-023", "node-024", "node-025", "node-026", "node-027", "node-028", "node-029", "node-030", "node-031", "node-032", "node-033", "node-034", "node-035", "node-036", "node-037", "node-038", "node-039", "node-040", "node-041", "node-042", "node-043", "node-044", "node-045", "node-046", "node-047", "node-048", "node-049", "node-050", "node-051", "node-052", "node-053", "node-054", "node-055", "node-056", "node-057", "node-058", "node-059", "node-060", "node-061", "node-062", "node-063", "node-064", "node-065", "node-066", "node-067", "node-068", "node-069", "node-070", "node-071", "node-072", "node-073", "node-074", "node-075", "node-076", "node-077", "node-078", "node-079", "node-080", "node-081", "node-082", "node-083", "node-084", "node-085", "node-086", "node-087", "node-088", "node-089", "node-090", "node-091", "node-092", "node-093", "node-094", "node-095", "node-096", "node-097", "node-098", "node-099", "node-100"}
var ring_500, ring_1000, ring_2000, ring_5000, ring_10000, ring_20000, ring_50000 *ConsistentHashNodeRing

func Benchmark_NewConsistentHashNodesRing_500(b *testing.B) {
	ring_500 = NewConsistentHashNodesRing(500, nodes)
}
func Benchmark_NewConsistentHashNodesRing_1000(b *testing.B) {
	ring_1000 = NewConsistentHashNodesRing(1000, nodes)
}
func Benchmark_NewConsistentHashNodesRing_2000(b *testing.B) {
	ring_2000 = NewConsistentHashNodesRing(2000, nodes)
}
func Benchmark_NewConsistentHashNodesRing_5000(b *testing.B) {
	ring_5000 = NewConsistentHashNodesRing(5000, nodes)
}
func Benchmark_NewConsistentHashNodesRing_10000(b *testing.B) {
	ring_10000 = NewConsistentHashNodesRing(10000, nodes)
}
func Benchmark_NewConsistentHashNodesRing_20000(b *testing.B) {
	ring_20000 = NewConsistentHashNodesRing(20000, nodes)
}
func Benchmark_NewConsistentHashNodesRing_50000(b *testing.B) {
	ring_50000 = NewConsistentHashNodesRing(50000, nodes)
}

func Benchmark_GetNodes_500(b *testing.B) {
	_, err := ring_500.GetNode("1")
	if err != nil {
		b.Error("unexpect")
	}
}
func Benchmark_GetNodes_1000(b *testing.B) {
	_, err := ring_1000.GetNode("1")
	if err != nil {
		b.Error("unexpect")
	}
}
func Benchmark_GetNodes_2000(b *testing.B) {
	_, err := ring_2000.GetNode("1")
	if err != nil {
		b.Error("unexpect")
	}
}
func Benchmark_GetNodes_5000(b *testing.B) {
	_, err := ring_5000.GetNode("1")
	if err != nil {
		b.Error("unexpect")
	}
}
func Benchmark_GetNodes_10000(b *testing.B) {
	_, err := ring_10000.GetNode("1")
	if err != nil {
		b.Error("unexpect")
	}
}
func Benchmark_GetNodes_20000(b *testing.B) {
	_, err := ring_20000.GetNode("1")
	if err != nil {
		b.Error("unexpect")
	}
}
func Benchmark_GetNodes_50000(b *testing.B) {
	_, err := ring_50000.GetNode("1")
	if err != nil {
		b.Error("unexpect")
	}
}
