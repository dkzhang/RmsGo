package resGNodeTree

import (
	"fmt"
	"github.com/dkzhang/RmsGo/ResourceSM/model/resGNode"
	"io/ioutil"
)

func ToJson(t Tree) (string, error) {
	return resGNode.ToJson(t.Root)
}

func ToJsonIndent(t Tree) (string, error) {
	return resGNode.ToJsonIndent(t.Root)
}

func LoadFromJson(bJson []byte) (t Tree, err error) {
	t.Root, err = resGNode.LoadFromJson(bJson)
	if err != nil {
		return Tree{}, fmt.Errorf("LoadFromJson error since resGNode.LoadFromJson error: %v", err)
	}

	t.NodesNum = CountRO(&t)
	return t, nil
}

func LoadFromJsonFile(filename string) (t Tree, err error) {
	bJson, err := ioutil.ReadFile(filename)
	if err != nil {
		return Tree{}, fmt.Errorf("ioutil.ReadFile error: %v", err)
	}

	return LoadFromJson(bJson)
}
