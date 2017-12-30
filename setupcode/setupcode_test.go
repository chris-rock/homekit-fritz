package setupcode

import "testing"

func TestGenxhmuri(t *testing.T) {

	expected := "X-HM://0023ISYWYFRIT"

	result := GenXhmURI(2, 0, "03145154", "FRIT")
	if result != expected {
		t.Errorf("genxhmuri was incorrect, got: %s, want: %s.", result, expected)
	}
}
