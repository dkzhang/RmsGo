package nodeEncode

import "testing"

//func TestIntArrayToByteArray(t *testing.T) {
//	intArray := []int{0,1,2,127}
//	byteArray := IntArrayToByteArray(intArray)
//	t.Logf("%v", byteArray)
//}

func TestIntArrayToBase64Str(t *testing.T) {
	int30_59 := make([]int, 30)
	for i := 0; i < 30; i++ {
		int30_59[i] = 30 + i
	}
	intArray := append([]int{0, 1, 2, 238, 239}, int30_59...)
	base64Str := IntArrayToBase64Str(intArray)
	t.Logf("base64Srt = \n %s\n", base64Str)
	t.Logf("len(base64Srt) = %d", len(base64Str))
	intArray1, err := Base64StrToIntArray(base64Str)
	if err != nil {
		t.Fatalf("Base64StrToIntArray error: %v", err)
	}
	t.Logf("intArray = %v", intArray1)
}
