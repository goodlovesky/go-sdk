package number

import (
	"testing"
)

func TestMist(t *testing.T) {
	m := NewMist(10)
	for i := 0; i < 1000000; i++ {
		t.Log(m.GenerateUID())
	}
}

func BenchmarkTestMist(b *testing.B) {
	m := NewMist(1)
	for i := 0; i < b.N; i++ {
		b.Log(m.GenerateUID())
	}
}
