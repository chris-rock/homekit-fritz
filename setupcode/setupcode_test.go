package setupcode

import "testing"

func TestGenxhmuri(t *testing.T) {

	expected := "X-HM://0023ISYWYFRIT"

	result := genxhmuri(2, 0, "03145154", "FRIT")
	if result != expected {
		t.Errorf("genxhmuri was incorrect, got: %d, want: %d.", result, expected)
	}
}
