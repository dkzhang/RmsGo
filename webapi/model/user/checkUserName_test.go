package user

import "testing"

func TestCheckUserName(t *testing.T) {
	nameList := []string{"zhang", "az1", "12345bb", "-xxxyx", "z12xx8"}
	for _, n := range nameList {
		r := CheckUserName(n)
		t.Logf("name <%s> result = <%v>", n, r)
	}

}
