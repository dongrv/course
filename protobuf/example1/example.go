package example1

import (
	"aaagame/tests/course/protobuf/example1/protocol"
	"encoding/binary"
	"fmt"
	"google.golang.org/protobuf/proto"
	"math"
	"os"
)

// InnerProto 内部二进制排列
func InnerProto() {
	v := uint32(1)
	send := protocol.Send{
		Fuint32:  v,
		Fuint322: v + 1,
		Fuint323: v + 2,
		Values:   []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Values2:  []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}
	d, _ := proto.Marshal(&send)

	fmt.Printf("%b\n", d)
	fmt.Printf("%08b\n", d)
	fmt.Printf("%d\n", len(d))

	var i uint32 = v
	fmt.Printf("%08b\n", i)
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	fmt.Printf("buf: %08b\n", buf)
}

var file = "resp.raw"

func MarshalBinary() {
	resp := protocol.HelleResp{
		UID:  10001,
		Name: "Tony",
		Exp:  math.MaxInt32 / 2,
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
	fmt.Printf("resp-> %+v\n", resp)
}

// 无符号整数同类型由小改大，兼容，	举例：uint32 -> uint64
// 无符号整数同类型由大改小，不兼容，	举例：uint64 -> uint32
// 有符号整数同类型由小改大，兼容，	举例：int32 -> int64
// 有符号整数同类型由大改小，不兼容，	举例：int64 -> int32

// 无符号和有符号类型互相转换时，如果通讯中存在负数值，不兼容，举例：int32 -> uint32
