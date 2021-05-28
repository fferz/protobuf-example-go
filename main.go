package main

import (
	"fmt"
	"io/ioutil"
	"log"

	homeworkpb "github.com/fferz/protobuf-example-go/src/homework"
	simplepb "github.com/fferz/protobuf-example-go/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"

	enumspb "github.com/fferz/protobuf-example-go/src/enums"

	complexpb "github.com/fferz/protobuf-example-go/src/complex"
)

func main() {

	sm := doSimple()

	writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content:", sm2)

	smAsString := toJSON(sm)
	fmt.Println("smAString", smAsString)

	sm3 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm3)
	fmt.Println("Succesfully created proto struct:", sm3)

	doEnum()

	doComplex()

	doHomework()
}

func doHomework() {
	hm := homeworkpb.AddressBook{
		People: []*homeworkpb.Person{
			{
				Name:  "Maria",
				Id:    1,
				Email: "maria@gmail.com",
				Phones: []*homeworkpb.Person_PhoneNumber{
					{
						Number: "3624789456",
						Type:   homeworkpb.Person_MOBILE,
					},
					{
						Number: "3624123456",
						Type:   homeworkpb.Person_WORK,
					},
				},
			},
			{
				Name:  "Juan",
				Id:    2,
				Email: "juan@gmail.com",
				Phones: []*homeworkpb.Person_PhoneNumber{
					{
						Number: "3624123111",
						Type:   homeworkpb.Person_HOME,
					},
					{
						Number: "3624123222",
						Type:   homeworkpb.Person_MOBILE,
					},
				},
			},
		},
	}
	fmt.Println("homework", hm)
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First dummy msg",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   2,
				Name: "Second dummy msg",
			},
			{
				Id:   2,
				Name: "Second dummy msg",
			},
		},
	}

	fmt.Println("complex message:", cm)
}

func doEnum() {
	em := enumspb.EnumsMessage{
		Id:        43,
		DayOfWeek: enumspb.DayOfWeek_MONDAY,
	}

	em.DayOfWeek = enumspb.DayOfWeek_FRIDAY
	fmt.Println("em msg", em)
}

func toJSON(m protoiface.MessageV1) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(m)
	if err != nil {
		log.Fatalln("Can't convert to json", err)
		return ""
	}
	return out
}

func fromJSON(in string, m protoiface.MessageV1) {
	err := jsonpb.UnmarshalString(in, m)
	if err != nil {
		log.Fatalln("Couldn't unmarshall the JSON into the pb struct", err)

	}
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can´t serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can´t write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Couldn't put the bytes into the protocol buffers struct", err2)
		return err2
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "my message",
		SampleList: []int32{1, 4, 7, 9},
	}

	fmt.Println(sm)

	sm.Name = "New name"

	fmt.Println(sm)
	fmt.Println("the id is:", sm.GetId())

	return &sm
}
