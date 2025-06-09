package geeznum

import (
	"fmt"
	"testing"
)

func TestNumToGeez(t *testing.T) {
	got := NumToGeez(1)
	want := "፩"
	assertCorrectMessage(t, got, want)
}

func ExampleNumToGeez() {
	result := NumToGeez(2025)
	fmt.Println(result)

	// Output: ፳፻፳፭

}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
