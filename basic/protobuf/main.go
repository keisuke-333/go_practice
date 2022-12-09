package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"protobuf-practice/pb"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Keisuke",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectTest": &pb.Company_Project{}},
		Profile:     &pb.Employee_Text{Text: "My name is Keisuke"},
		Birthday:    &pb.Date{Year: 2000, Month: 1, Day: 1},
	}

	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Can't marshal to json", err)
	}

	fmt.Println(out)

	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("Can't unmarshal from json", err)
	}
	fmt.Println(readEmployee)
}
