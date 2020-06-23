package myGob

import "testing"

func TestStoreLoad(t *testing.T) {

	str := "Hello World!"
	t.Logf("[]byte(str) = %v", []byte(str))

	b := Store(str)
	t.Logf("b = %v", b)

	strRead := ""
	Load(&strRead, b)
	t.Logf("strRead = %s", strRead)
}
