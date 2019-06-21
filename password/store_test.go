package password

import (
	"testing"
)

func TestAdd(t *testing.T) {
	s := Init()

	s.Add("myFingerprint", "myTag", "myPassword")

	pws, ok := s.passwords["myFingerprint"]["myTag"]
	if !ok || pws[0] != "myPassword" {
		t.Error("Expected store to have passwort `myPassword` stored for fingerprint `myFingerprint` and tag `myTag` but found none")
	}
}
