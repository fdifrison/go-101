package numbers

import (
	"fmt"
	"testing"
)

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers#testable-examples
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAddition(t *testing.T) {

	t.Run("Should add", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})
}
