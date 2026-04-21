package iteration

import "testing"

func TestIteration(t *testing.T){
	repeated:= Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}


func BenchmarkRepeat(b *testing.B){
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 6)
	println(repeated)
	// Output: "aaaaaa"
}