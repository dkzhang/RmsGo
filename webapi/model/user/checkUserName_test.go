package user

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCheckUserName(t *testing.T) {
	type NameTable struct {
		Name    string
		IsLegal bool
	}
	nameTable := []NameTable{
		{"jf-zhang", true},
		{"z12-xx-qq8", true},
		{"z11_xx-qq8", true},
		{"x", false},
		{"1x", false},
		{"abcdefghijklmnopqrstuvwxyz", false},
		{"zhang@jf", false},
		{"x=abs", false},
	}

	Convey("Test Name Table", t, func() {
		for _, nt := range nameTable {
			r := CheckUserName(nt.Name)
			So(r, ShouldEqual, nt.IsLegal)
		}
	})
}
