package utils

import (
	"fmt"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func PrintJsonShort(v interface{}) (err error) {
	b, err := json.Marshal(v)
	if err == nil {
		fmt.Println(string(b))
	}
	return
}