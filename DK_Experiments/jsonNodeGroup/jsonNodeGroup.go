package jsonNodeGroup

import (
	"encoding/json"
	"fmt"
)

type NodeGroup struct {
	Name       string
	IsSelected bool
	SubGroup   []NodeGroup
}

func (ng NodeGroup) ToJson() (string, error) {
	b, err := json.Marshal(ng)
	if err != nil {
		return "", fmt.Errorf("json.Marshal(ng) error: %v", err)
	}
	return string(b), nil
}
