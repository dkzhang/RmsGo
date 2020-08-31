package resNodeTree

import (
	"encoding/json"
	"fmt"
)

func GroupToJson(g Group) (string, error) {
	bj, err := json.Marshal(g)
	if err != nil {
		return "", fmt.Errorf("json Marshal Group error: %v", err)
	}
	return string(bj), nil
}

func GroupToJsonIndent(g Group) (string, error) {
	bj, err := json.MarshalIndent(g, "", "    ")
	if err != nil {
		return "", fmt.Errorf("json Marshal Group error: %v", err)
	}
	return string(bj), nil
}

func LoadGroupFromJson(str string) (g Group, err error) {
	err = json.Unmarshal([]byte(str), &g)
	if err != nil {
		return Group{}, fmt.Errorf("json UnMarshal Group error: %v", err)
	}
	return g, nil
}
