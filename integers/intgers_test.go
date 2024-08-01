package integers

import (
	"fmt"
	"testing"
)

func TestIntegers(t *testing.T) {
	t.Run("adds two numbers", func(t *testing.T) {
		got := Add(2, 2)
		expected := 4

		if got != expected {
			t.Errorf("got %d, expected %d", got, expected)
		}
	})
}

// Adding this code will cause the example to appear in your godoc documentation, making your code even more accessible.
// If ever your code changes so that the example is no longer valid, your build will fail.
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
