package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	tm := TypeMeta{
		Name:       "zhang",
		Department: "jf",
	}

	jb, err := json.Marshal(tm)
	if err != nil {
		fmt.Printf("json.Marshal error: %v \n", err)
	}
	fmt.Printf("json.Marshal success, json tm = %s \n", string(jb))

	PrintTag()
}

type TypeMeta struct {
	Name       string `json:"name" excel:"A, 姓名"`
	Department string `json:"department" excel:"B, 部门"`
}

func PrintTag() {
	tm := TypeMeta{
		Name:       "zhang",
		Department: "jf",
	}

	ty := reflect.TypeOf(tm)

	for i := 0; i < ty.NumField(); i++ {
		fmt.Printf("Field: %s, Tag: %s \n", ty.Field(i).Name, ty.Field(i).Tag.Get("excel"))
	}
}
