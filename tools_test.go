package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools

	s := testTools.RandomString(10)

	// Test that the random string is the correct length
	if len(s) != 10 {
		t.Errorf("Random string is not the correct length")
	}
}
