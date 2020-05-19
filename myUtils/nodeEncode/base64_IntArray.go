package nodeEncode

func IntArrayToBase64Str(intArray []int) string {
	return ToBase64(IntArrayToByteArray(intArray))
}
func IntArrayToByteArray(intArray []int) []byte {
	byteArray := make([]byte, max(intArray)/8+1)
	for _, v := range intArray {
		quotient := v / 8
		remainder := v % 8
		byteArray[quotient] = byteArray[quotient] | (1 << remainder)
	}
	return byteArray
}

func Base64StrToIntArray(base64Str string) ([]int, error) {
	byteArray, err := ToByteArray(base64Str)
	if err != nil {
		return nil, err
	}
	return ByteArrayToIntArray(byteArray), nil
}

func ByteArrayToIntArray(byteArray []byte) []int {
	intArray := make([]int, 0)
	for i := 0; i < len(byteArray); i++ {
		for j := 0; j < 8; j++ {
			if byteArray[i]&(1<<j) != 0 {
				intArray = append(intArray, j+8*i)
			}
		}
	}
	return intArray
}

func max(theArray []int) (theMax int) {
	theMax = theArray[0]
	for _, v := range theArray {
		if v > theMax {
			theMax = v
		}
	}
	return theMax
}
