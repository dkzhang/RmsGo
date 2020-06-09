package userTempDM

import "github.com/dkzhang/RmsGo/myUtils/genPasswd"

type LoginSecurity struct {
	TheTokenSecurity TokenSecurity
}

type TokenSecurity struct {
	Key string
}

func LoadLoginSecurity() (sec LoginSecurity, err error) {
	pwLen := 12
	pwType := genPasswd.FlagLowerChar
	genPasswd.RandomSeed()
	return LoginSecurity{
		TheTokenSecurity: TokenSecurity{
			Key: genPasswd.GeneratePasswd(pwLen, pwType)}}, nil
}
