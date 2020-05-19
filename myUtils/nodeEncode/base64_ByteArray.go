package nodeEncode

import "encoding/base64"

func ToBase64(nodesByteArray []byte) string {
	return base64.StdEncoding.EncodeToString(nodesByteArray)
}

func ToByteArray(nodesBase64 string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(nodesBase64)
	return decoded, err
}
