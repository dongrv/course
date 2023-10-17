package protocol

import (
	"casino/protocol/cmd"
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"os"
	"sync"
)

type Location int8 // 位置

const (
	Lobby Location = iota // 大厅
	Slot                  // 机器内
)

type User struct {
	DID        string
	Token      string
	NowMs      int64 // 毫秒级时间戳
	Summary    *cmd.UserSummary
	GameInfo   *cmd.GameInfo
	Slots      []*cmd.Slot
	Activities map[int32]*cmd.Activity
	Shop       *Shop
	Location   Location // 当前位置
}

type Shop struct {
	Bonus *cmd.ShopBonus
	Items []*cmd.ShopItem
}

type (
	HandleReqFunc  func(u *User) proto.Message     // 包装消息
	HandleRespFunc func(u *User, msg []byte) error // 函数格式
)

type MsgMap struct {
	mu         sync.RWMutex
	idName     map[int32]string         // 协议序号：协议名称
	nameId     map[string]int32         // 协议名称：协议序号
	reqIdFunc  map[int32]HandleReqFunc  // 协议序号：包装函数
	respIdFunc map[int32]HandleRespFunc // 协议序号：绑定函数
}

type Message struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

// 所有协议管理器
var msgMap *MsgMap

func MsgName(msgId int32) string {
	msgMap.mu.RLock()
	defer msgMap.mu.RUnlock()
	return msgMap.idName[msgId]
}

func MsgId(msgName string) int32 {
	msgMap.mu.RLock()
	defer msgMap.mu.RUnlock()
	return msgMap.nameId[msgName]
}

// BindFunc 绑定处理函数
func BindFunc(msgId int32, fn HandleRespFunc) {
	if _, ok := msgMap.respIdFunc[msgId]; ok {
		return
	}
	msgMap.respIdFunc[msgId] = fn
}

func Find(msgId int32) HandleRespFunc {
	msgMap.mu.RLock()
	defer msgMap.mu.RUnlock()
	v, ok := msgMap.respIdFunc[msgId]
	if !ok {
		return nil
	}
	return v
}

// WrapFunc 包装函数
func WrapFunc(msgId int32, fn HandleReqFunc) {
	if _, ok := msgMap.reqIdFunc[msgId]; ok {
		return
	}
	msgMap.reqIdFunc[msgId] = fn
}

func FindWrap(msgId int32) HandleReqFunc {
	msgMap.mu.RLock()
	defer msgMap.mu.RUnlock()
	v, ok := msgMap.reqIdFunc[msgId]
	if !ok {
		return nil
	}
	return v
}

// ReadProtocolFile 读取协议文件
func ReadProtocolFile(protoFile string) (map[string][]*Message, error) {
	var protocols = make(map[string][]*Message) // 协议序号：协议名称
	bytes, err := os.ReadFile(protoFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &protocols)
	if err != nil {
		return nil, err
	}
	return protocols, nil
}

// Register 注册所有协议
func Register(protoFile string) {
	protocols, err := ReadProtocolFile(protoFile)
	if err != nil {
		panic(err.Error())
	}
	msgMap = &MsgMap{
		idName:     make(map[int32]string),
		nameId:     make(map[string]int32),
		respIdFunc: make(map[int32]HandleRespFunc),
		reqIdFunc:  make(map[int32]HandleReqFunc),
	}
	for _, protos := range protocols {
		for _, v := range protos {
			if _, ok := msgMap.idName[v.Id]; ok {
				continue
			}
			msgMap.idName[v.Id] = v.Name
			msgMap.nameId[v.Name] = v.Id
		}
	}
}
