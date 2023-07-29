package toolkit

import (
	"testing"
)

func TestTools_RandomString(t *testing.T) {
	var tool Tools

	s1 := tool.RandomStringL(10)
	if len(s1) != 10 {
		t.Error("Random string length error")
	}

	s2 := tool.RandomString(10)
	if len(s2) != 10 {
		t.Error("Random string length error")
	}
}

func BenchmarkTools_RandomStringL(b *testing.B) {
	var tool Tools

	for i := 0; i < b.N; i++ {
		tool.RandomStringL(10)
	}
}

func BenchmarkTools_RandomString(b *testing.B) {
	var tool Tools

	for i := 0; i < b.N; i++ {
		tool.RandomString(10)
	}
}
