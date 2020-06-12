package user

import (
	"fmt"
	"testing"
)

func TestToJsonString(t *testing.T) {
	user1 := UserInfo{
		UserID:         0,
		UserName:       "u1",
		ChineseName:    "c1",
		Department:     "department1",
		DepartmentCode: "d1",
		Section:        "s1",
		Mobile:         "123",
		Role:           0,
		Status:         0,
		Remarks:        "",
	}
	user2 := UserInfo{
		UserID:         0,
		UserName:       "u2",
		ChineseName:    "c2",
		Department:     "department2",
		DepartmentCode: "d2",
		Section:        "s2",
		Mobile:         "456",
		Role:           0,
		Status:         0,
		Remarks:        "",
	}

	t.Logf("%s \n", ToJsonString([]UserInfo{user1, user2}))
	t.Logf("%s \n", user1.ToJsonString())
}

func TestSetUserName(t *testing.T) {
	var name, dc, sname string
	dc = "jf"

	name = "zhang"
	sname = StandardizedUserName(name, dc)
	if sname != fmt.Sprintf("%s-%s", dc, name) {
		t.Errorf("test failed: got %s", sname)
	}

	name = "jf-zh"
	sname = StandardizedUserName(name, dc)
	if sname != name {
		t.Errorf("test failed: got %s", sname)
	}

	name = "z"
	sname = StandardizedUserName(name, dc)
	if sname != fmt.Sprintf("%s-%s", dc, name) {
		t.Errorf("test failed: got %s", sname)
	}
}
