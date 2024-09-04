package example1

import (
	"aaagame/tests/course/protobuf/example2/protocol"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"math"
	"os"
)

var file = "resp.raw"

func MarshalBinary() {
	resp := protocol.HelleResp{
		UID:  10001,
		Name: "Tony",
		Exp:  math.MaxInt32 / 2,
		//Rank: 1000,
	}
	f, _ := os.Create(file)
	defer f.Close()
	bs, _ := proto.Marshal(&resp)
	f.Write(bs)
}

func UnmarshalBinary() {
	resp := &protocol.HelleResp{}
	bs, _ := os.ReadFile(file)
	err := proto.Unmarshal(bs, resp)
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(resp)
	fmt.Printf("resp-> %s\n", marshal)
}

// string类型改整数，字段值失效，反之亦然
