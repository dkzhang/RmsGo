package myGob

import (
	"bytes"
	"encoding/gob"
)

func Store(data interface{}) []byte {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
	//err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	//if err != nil{
	//	panic(err)
	//}
}

func Load(data interface{}, raw []byte) {
	//raw, err := ioutil.ReadFile(filename)
	//if err != nil{
	//	panic(err)
	//}

	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err := dec.Decode(data)
	if err != nil {
		panic(err)
	}
}
