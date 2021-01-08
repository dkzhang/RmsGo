package security

import "testing"

func TestGenLoginSecurity(t *testing.T) {
	t.Log(GenLoginSecurity())
	t.Log(GenLoginSecurity())
}
