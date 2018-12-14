package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type UseData struct {
	Length  int
	MaxSize int
}

func TestJason(t *testing.T) {
	js := `{"length":12,"maxsize":14}`
	w := &UseData{}
	error := json.Unmarshal([]byte(js), w)
	if error != nil {
		fmt.Println("json error", error)
		t.FailNow()
	}
	fmt.Println(w.Length)
	fmt.Println(w.MaxSize)
}
