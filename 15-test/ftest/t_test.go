package test

import "testing"

func TestAdd(t *testing.T) {

	var (
		a1, a2 = 10, 10
		res    = 20
	)
	if Add(a1, a2) == res {
		t.Log("pass")
	} else {
		t.Error("failed")
	}
}

func BenchmarkAdd(b *testing.B) {

	var (
		a1, a2 = 10, 10
		res    = 20
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if Add(a1, a2) == res {
			b.Log("pass")
		} else {
			b.Error("failed")
		}
	}
}
