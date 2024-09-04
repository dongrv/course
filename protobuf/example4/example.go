package example1

import (
	"aaagame/tests/course/protobuf/example4/protocol"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"os"
)

var file = "resp.raw"

func MarshalBinary() {
	resp := protocol.HelleResp{
		UID:  10001,
		Name: "Tony",
		Profile: &protocol.Profile{
			//Addr:       "Shanghai/China",
			RegionCode: 101,
		},
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

// 只是修改结构体中的字段名称，不影响序列化和反序列化
