package testjson

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id int32 `json:"id"`
}

var userJson = `{"id":"1"}` // 强类型，不可转

func Unmarshal() error {
	var user User
	err := json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", user)
	return nil
}
