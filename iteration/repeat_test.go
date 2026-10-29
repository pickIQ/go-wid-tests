package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("default repeatTime", func(t *testing.T) {
	  repeated := Repeat("a", 0)
	  expected := "aaaaa"

	  if repeated != expected {
  		t.Errorf("expected %q but got %q", expected, repeated)
	  }
	})

	t.Run("custom repeatTime", func(t *testing.T) {
	  repeated := Repeat("a", 3)
	  expected := "aaa"

	  if repeated != expected {
  		t.Errorf("expected %q but got %q", expected, repeated)
	  }
	})
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