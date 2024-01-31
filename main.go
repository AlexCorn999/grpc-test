package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id      int
	Name    string
	IsHuman bool
}

func (u *User) ProtoMessage()  {}
func (u *User) Reset()         { *u = User{} }
func (u *User) String() string { return fmt.Sprintf("%#v", u) }

func main() {
	user := User{
		Id:      1,
		Name:    "Alex",
		IsHuman: true,
	}

	dataJson, _ := json.Marshal(user)
	fmt.Printf("\n\nlen json - %d  |  %+v\n", len(dataJson), user)
	//dataPb, _ := proto.Marshal(&user)
	fmt.Printf("\n\nlen protobuf - %d  |  %+v\n", len(dataPb), user)
}
