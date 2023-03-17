package integers

import "fmt"
import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}

func TestSub(t *testing.T) {
	result := Sub(4, 2)
	expected := 2

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
