package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/snluu/uuid"
)

func main() {
	m := map[string]interface{}{
		"a": 3,
		"b": 1,
		"c": "zhougang",
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, _ := json.MarshalToString(m)
	fmt.Println(b)

	var id = uuid.Rand()
	fmt.Println(id.Hex())
}
