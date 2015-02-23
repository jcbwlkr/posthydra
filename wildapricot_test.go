package posthydra_test

import (
	"testing"

	"github.com/jcbwlkr/posthydra"
)

func TestTokenAuthHeader(t *testing.T) {
	cfg := &posthydra.WildApricotConfig{Key: "Banana", AccountId: 7}

	token := cfg.TokenAuthHeader()
	expected := "Basic abcde"
	// QVBJS0VZOkJhbm#

	if expected != token {
		t.Errorf("WA Token Auth failed: Got %s expected %s", token, expected)
	}
}
