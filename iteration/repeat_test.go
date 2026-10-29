package iteration

import (
	"fmt"
	"go-wid-tests/helpers"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("default repeatTime", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "aaaaa"

		helpers.AssertCorrectMessage(t, got, want)

	})

	t.Run("custom repeatTime", func(t *testing.T) {
		got := Repeat("a", 3)
		want := "aaa"

		helpers.AssertCorrectMessage(t, got, want)

	})
}

func ExampleRepeat() {
	got := Repeat("s", 3)
	fmt.Println(got)
	// Output: sss
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}
