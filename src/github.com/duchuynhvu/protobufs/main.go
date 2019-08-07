package main

import (
	"encoding/json"
	"fmt"

	pb "github.com/duchuynhvu/protofiles"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "Roger F",
		Email: "rf@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{
				Number: "555-4321",
				Type:   pb.Person_HOME},
		},
	}

	//EXAMPLE OF SERIALIZE...

	// p1 := &pb.Person{}
	// body, _ := proto.Marshal(p)
	// _ = proto.Unmarshal(body, p1)
	// fmt.Println("Original struct loaded from proto file:", p, "\n")
	// fmt.Println("Marshaled proto data: ", body, "\n")
	// fmt.Println("Unmarshaled struct: ", p1)

	//EXAMPLE OF GENERATE JSON...
	body, _ := json.Marshal(p)
	fmt.Println(string(body))

}
