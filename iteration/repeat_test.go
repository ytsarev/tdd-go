package iteration

import (
	"testing"
  "fmt"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i:=0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 7)
	fmt.Println(repeat)
	// Output: aaaaaaa
}
