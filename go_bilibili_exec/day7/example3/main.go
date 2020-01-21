package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"un"`
	NickName string `json:"nn"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

func main() {
	user1 := &User{
		UserName: "user1",
		NickName: "nick",
		Age:      18,
		Birthday: "2020/01/21",
		Sex:      "男",
		Email:    "test@qq.com",
		Phone:    "111111",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("Json.marshal failed, err:", err)
		return
	}
	fmt.Println("data: ", string(data))
}
