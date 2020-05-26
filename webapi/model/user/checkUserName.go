package user

import "regexp"

func CheckUserName(name string) bool {
	uPattern := `^[a-zA-Z]{1}[a-zA-Z0-9_-]{3,16}$`

	r := regexp.MustCompile(uPattern)
	return r.MatchString(name)
}
