package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	character := "a"
	repeatTime := 8
	repeated := Repeat(character, repeatTime)
	expected := "aaaaaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	//Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
