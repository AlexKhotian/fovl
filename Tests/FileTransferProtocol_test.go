package Tests

import (
	"encoding/json"
	"fovl/FileSession"
	"testing"
)

func TestMarshallJson(t *testing.T) {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	testStruct := &FileSession.TransferCommand{
		CommandType: FileSession.Start,
		DataChunk:   byt}
	data, err := json.Marshal(testStruct)
	if err != nil || data == nil {
		t.Log("Failed to marshall json")
		t.Fail()
	}
}

func TestUnMarshallJson(t *testing.T) {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	inStruct := &FileSession.TransferCommand{
		CommandType: FileSession.Start,
		DataChunk:   byt}
	data, err := json.Marshal(inStruct)
	if err != nil {
		t.Log("Failed to marshall json")
		t.Fail()
	}

	outStruct := &FileSession.TransferCommand{}
	err = json.Unmarshal(data, outStruct)
	if err != nil {
		t.Log("Failed to unmarshall json")
		t.Fail()
	}

	if inStruct.CommandType != outStruct.CommandType ||
		outStruct.DataChunk == nil {
		t.Log("Wrong data unmarshalled")
		t.Fail()
	}
}
