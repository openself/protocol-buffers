package main

import (
	"./src/complex"
	"./src/enum_example"
	"./src/simple"
	"bufio"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	msg := getSimpleMsg()
	writeAndReadMsgBin(msg)
	writeAndReadMsgJSON(msg)

	doEnum()

	doComplex()
}

func getSimpleMsg() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple message",
		SampleList: []int32{1, 4, 7, 8},
	}
	log.Println(sm)
	sm.Name = "I rename you"
	log.Println(sm)
	log.Println("The ID is:", sm.GetId())
	return &sm
}

func writeAndReadMsgBin(m *simplepb.SimpleMessage) {
	buf, err := proto.Marshal(m)
	if err != nil {
		log.Fatalln("Data marshal error:", err)
	}
	outputFile, err := ioutil.TempFile(".", "msg_*.bin")
	if err != nil {
		log.Fatalln("File write error:", err)
	}
	defer os.Remove(outputFile.Name())
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.Write(buf)
	outputFile.Close()

	buf, err = ioutil.ReadFile(outputFile.Name())
	if err != nil {
		log.Fatalln("File read error:", err)
	}
	sm := simplepb.SimpleMessage{}
	err = proto.Unmarshal(buf, &sm)
	if err != nil {
		log.Fatalln("Data marshal error:", err)
	}
	log.Println("Tranferred message:", sm)
}

func writeAndReadMsgJSON(m *simplepb.SimpleMessage) {
	marshaler := jsonpb.Marshaler{}
	outJSON, err := marshaler.MarshalToString(m)
	if err != nil {
		log.Fatalln("Data marshal to JSON error:", err)
	}
	log.Println("JSON message:", outJSON)

	sm := simplepb.SimpleMessage{}
	err = jsonpb.UnmarshalString(outJSON, &sm)
	if err != nil {
		log.Fatalln("Data unmarshal from JSON error:", err)
	}
	log.Println("Tranferred from JSON message:", sm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           777,
		DayOfTheWeek: enumpb.DayOfTheWeek_FRIDAY,
	}
	log.Println(em)

	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY
	log.Println(em)
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "One message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   2,
				Name: "Second message",
			},
			{
				Id:   3,
				Name: "Third message",
			},
		},
	}
	log.Println(cm)

	cm.OneDummy.Name = "First message"
	log.Println(cm)
}
