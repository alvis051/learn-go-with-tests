package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 5)
	expect := "aaaaa"

	if repeat != expect {
		t.Errorf("expected: %q but got %q", expect, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	Repeat("a", b.N)
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}
