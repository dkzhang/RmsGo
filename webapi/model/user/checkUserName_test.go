package user

import "testing"

func TestCheckUserName(t *testing.T) {
	nameList := []string{"jf-zhang", "az-a1", "12345bb", "sadfa-xxxyx", "z12-xx-qq8"}
	for _, n := range nameList {
		r := CheckUserName(n)
		t.Logf("name <%s> result = <%v>", n, r)
	}

}
