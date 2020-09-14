package resGNode

import (
	"encoding/json"
	"fmt"
)

func ToJson(gn ResGNode) (string, error) {
	bj, err := json.Marshal(gn)
	if err != nil {
		return "", fmt.Errorf("json Marshal ResGNode error: %v", err)
	}
	return string(bj), nil
}

func ToJsonIndent(gn ResGNode) (string, error) {
	bj, err := json.MarshalIndent(gn, "", "    ")
	if err != nil {
		return "", fmt.Errorf("json Marshal ResGNode error: %v", err)
	}
	return string(bj), nil
}

func LoadFromJson(bJson []byte) (gn ResGNode, err error) {
	err = json.Unmarshal(bJson, &gn)
	if err != nil {
		return ResGNode{}, fmt.Errorf("json UnMarshal ResGNode error: %v", err)
	}
	return gn, nil
}
