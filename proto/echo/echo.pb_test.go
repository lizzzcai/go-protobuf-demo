package echo

import (
	"testing"
	"github.com/golang/protobuf/proto"
)

func TestEcho(t *testing.T) {
	var cases = [] string{
		"lizzzcai",
		"echo",
		"",
		"123",
		"123.1234",
	}

	for _, c := range cases {

		req := &EchoRequest{Name: c}
		data, err := proto.Marshal(req)
		if err != nil {
			t.Errorf("Error while marshalling the object : %v", err)
		}
	
		res := &EchoRequest{}
		err = proto.Unmarshal(data, res)
		if err != nil {
			t.Errorf("Error while un-marshalling the object : %v", err)
		}

		if res.GetName() != c {
			t.Errorf("un-marshalled data = %v, want %v",res.GetName(), c)
		}

	}

}