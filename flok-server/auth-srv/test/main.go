package main

import (
	proto "chappie-smarthome/chappie-cloud/auth-srv/authproto"
	"fmt"

	pro "github.com/golang/protobuf/proto"
)

func main() {
	u := &proto.User{
		Id:       "dsfdsf",
		Password: "dsfsdf",
	}

	p := &proto.User{}

	b, _ := pro.Marshal(u)
	pro.Unmarshal(b, p)

	fmt.Println(p)
}
