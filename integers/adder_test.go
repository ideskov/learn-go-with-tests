package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Sum(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Sum(2, 4)
	fmt.Println(sum)
	// Output: 6
}
