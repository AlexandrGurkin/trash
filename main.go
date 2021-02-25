package main

import (
	"log"
	"trash/proto_out"

	"github.com/golang/protobuf/proto"
)

func main() {
	comp := &proto_out.Company{
		Name: "Example Corp",
		Type: proto_out.CompanyType_Public,
		Employees: []*proto_out.Employee{
			&proto_out.Employee{
				Name: "John",
			},
		},
	}

	data, err := proto.Marshal(comp)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	log.Print(data)

	serialized := &proto_out.Company{}
	err = proto.Unmarshal(data, serialized)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
}
