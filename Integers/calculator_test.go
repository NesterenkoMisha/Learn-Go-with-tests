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

func ExampleAdd() {
	result := Add(1, 5)
	fmt.Println(result)
	// Output: 6
}

func TestSub(t *testing.T) {
	result := Sub(4, 2)
	expected := 2

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}

func ExampleSub() {
	result := Sub(8, 3)
	fmt.Println(result)
	// Output: 5
}

func TestMul(t *testing.T) {
	result := Mul(4, 3)
	expected := 12

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}

func ExampleMul() {
	result := Mul(5, 4)
	fmt.Println(result)
	// Output: 20
}

func TestDiv(t *testing.T) {
	t.Run("Divide 12 by 4 should result in 3.0", func(t *testing.T) {
		result := Div(12, 4)
		expected := 3.0

		if result != expected {
			t.Errorf("expected '%f' but got '%f'", expected, result)
		}
	})
}

func ExampleDiv() {
	result := Div(12, 4)
	fmt.Println(result)
	// Output: 3
}
