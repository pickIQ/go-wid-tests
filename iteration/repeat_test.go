package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	exptected := "aaaaa"

	if repeated != exptected {
		t.Errorf("expected %q but got %q", exptected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("s", 3)
	fmt.Println(repeated)
	// Output: sss
} 

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}