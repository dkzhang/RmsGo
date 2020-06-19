package user

import "regexp"

func CheckUserName(name string) bool {
	//uPattern := `^[a-zA-Z0-9]{2,4}[-]{1}[a-zA-Z0-9]{2,16}$`
	uPattern := `^[a-zA-Z_]{1}[a-zA-Z0-9_-]{1,15}$`

	r := regexp.MustCompile(uPattern)
	return r.MatchString(name)
}

func CheckDepartmentCode(dc string) bool {
	dcPattern := `^[a-zA-Z0-9]{2,4}$`

	r := regexp.MustCompile(dcPattern)
	return r.MatchString(dc)
}
