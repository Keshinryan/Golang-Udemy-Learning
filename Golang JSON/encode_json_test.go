package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJSON(data interface{}) {
	bytess, err:=json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytess))
}

func TestEncode(t *testing.T) {
	logJSON("Jason")
	logJSON(1)
	logJSON(true)
	logJSON([]string{"Jason","Patrick"})
}