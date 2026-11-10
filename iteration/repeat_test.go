package iteration

import "testing"

func TestForLoop(t *testing.T) {

	t.Run("Should repeat string n times", func(t *testing.T) {
		repeat := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeat != expected {
			t.Errorf("expected %s got %s", expected, repeat)
		}

	})

	t.Run("Should repeat fast string n times", func(t *testing.T) {
		repeat := RepeatFast("a", 10)
		expected := "aaaaaaaaaa"

		if repeat != expected {
			t.Errorf("expected %s got %s", expected, repeat)
		}

	})

}

// https://pkg.go.dev/testing#hdr-Benchmarks
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}

// go test -bench=. -benchmem
// B/op: the number of bytes allocated per iteration
// allocs/op: the number of memory allocations per iteration
func BenchmarkRepeatFast(b *testing.B) {
	for b.Loop() {
		RepeatFast("a", 10)
	}
}
