package serializer

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, []byte(data), 0644)
	return err
}

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)

	return err
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(bytes, message)
	return err
}
