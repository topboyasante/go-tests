package iterations

import (
	"fmt"
	"testing"
)

func TestIterations(t *testing.T) {
	t.Run("it should repeat the provided string", func(t *testing.T) {
		got := Repeat("a" , 5)
		expected := "aaaaa"

		if got != expected {
			t.Errorf("expected %q got %q", expected, got)
		}
	})
}

// When the benchmark code is executed, it runs b.N times and measures how long it takes.
// The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("No. ", b.N)
		Repeat("a", 5)
	}
}
