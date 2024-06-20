// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: cmd.proto

package cmd

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 批量发送产出消耗日志
type BatchFinLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string    `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	FinLogs []*FinLog `protobuf:"bytes,2,rep,name=FinLogs,proto3" json:"FinLogs,omitempty"`
}

func (x *BatchFinLog) Reset() {
	*x = BatchFinLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchFinLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchFinLog) ProtoMessage() {}

func (x *BatchFinLog) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchFinLog.ProtoReflect.Descriptor instead.
func (*BatchFinLog) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{0}
}

func (x *BatchFinLog) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *BatchFinLog) GetFinLogs() []*FinLog {
	if x != nil {
		return x.FinLogs
	}
	return nil
}

type FinLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UID              int32         `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`                            // 用户ID
	SpecialUser      int32         `protobuf:"varint,2,opt,name=SpecialUser,proto3" json:"SpecialUser,omitempty"`            // 特殊用户标记
	PaidType         int32         `protobuf:"varint,3,opt,name=PaidType,proto3" json:"PaidType,omitempty"`                  // 用户付费类型
	TotalPay         float64       `protobuf:"fixed64,4,opt,name=TotalPay,proto3" json:"TotalPay,omitempty"`                 // 总付费
	SeriesId         int64         `protobuf:"varint,5,opt,name=SeriesId,proto3" json:"SeriesId,omitempty"`                  // 机器系列ID
	SlotId           string        `protobuf:"bytes,6,opt,name=SlotId,proto3" json:"SlotId,omitempty"`                       // 机器ID
	Balance          uint64        `protobuf:"varint,7,opt,name=Balance,proto3" json:"Balance,omitempty"`                    // 账户筹码余额
	Win              uint64        `protobuf:"varint,8,opt,name=Win,proto3" json:"Win,omitempty"`                            // 当前赢钱
	Bet              uint64        `protobuf:"varint,9,opt,name=Bet,proto3" json:"Bet,omitempty"`                            // 当前下注
	BalanceUSD       float64       `protobuf:"fixed64,10,opt,name=BalanceUSD,proto3" json:"BalanceUSD,omitempty"`            // 账户美元值
	WinUSD           float64       `protobuf:"fixed64,11,opt,name=WinUSD,proto3" json:"WinUSD,omitempty"`                    // 赢钱美元值
	BetUSD           float64       `protobuf:"fixed64,12,opt,name=BetUSD,proto3" json:"BetUSD,omitempty"`                    // 下注美元值
	SpinNum          int32         `protobuf:"varint,13,opt,name=SpinNum,proto3" json:"SpinNum,omitempty"`                   // SPIN总次数
	IsBroke          int32         `protobuf:"varint,14,opt,name=IsBroke,proto3" json:"IsBroke,omitempty"`                   // 破产标记：1是，0否
	Reason           uint64        `protobuf:"varint,15,opt,name=Reason,proto3" json:"Reason,omitempty"`                     // 来源标记
	GamePlay         int64         `protobuf:"varint,16,opt,name=GamePlay,proto3" json:"GamePlay,omitempty"`                 // 玩法
	OneDollar        uint64        `protobuf:"varint,17,opt,name=OneDollar,proto3" json:"OneDollar,omitempty"`               // 当前1美元兑筹码
	RegisterTime     int64         `protobuf:"varint,18,opt,name=RegisterTime,proto3" json:"RegisterTime,omitempty"`         // 注册时间
	AvgBet           float64       `protobuf:"fixed64,19,opt,name=AvgBet,proto3" json:"AvgBet,omitempty"`                    // 千次平均BET
	CreateTime       int64         `protobuf:"varint,20,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`             // 创建时间
	WorkScope        int32         `protobuf:"varint,21,opt,name=WorkScope,proto3" json:"WorkScope,omitempty"`               // 0外网和内网正式；1是内网测试
	Platform         int32         `protobuf:"varint,22,opt,name=Platform,proto3" json:"Platform,omitempty"`                 // 1:android 2:ios
	OutRTP           float64       `protobuf:"fixed64,23,opt,name=OutRTP,proto3" json:"OutRTP,omitempty"`                    // 输出RTP
	RTPState         int32         `protobuf:"varint,24,opt,name=RTPState,proto3" json:"RTPState,omitempty"`                 // 0：正常状态 1：持金状态 2：水池状态
	MinBet           uint64        `protobuf:"varint,25,opt,name=MinBet,proto3" json:"MinBet,omitempty"`                     // 固定档MinBet
	TriggerLucky     int32         `protobuf:"varint,26,opt,name=TriggerLucky,proto3" json:"TriggerLucky,omitempty"`         // 机器是否触发了幸运值和假数据：0否1是
	Lv               int32         `protobuf:"varint,27,opt,name=Lv,proto3" json:"Lv,omitempty"`                             // 用户当前等级
	VipLv            int32         `protobuf:"varint,28,opt,name=VipLv,proto3" json:"VipLv,omitempty"`                       // VIP等级
	LastPayTime      int64         `protobuf:"varint,29,opt,name=LastPayTime,proto3" json:"LastPayTime,omitempty"`           // 最近付费时间
	Spend            float64       `protobuf:"fixed64,30,opt,name=Spend,proto3" json:"Spend,omitempty"`                      // 此条产出关联的当前付费价格，比如：商城购买物品消费100$
	RemarkUser       int32         `protobuf:"varint,31,opt,name=RemarkUser,proto3" json:"RemarkUser,omitempty"`             // 备注用户：1后台备注用户
	RemarkFilter     int32         `protobuf:"varint,32,opt,name=RemarkFilter,proto3" json:"RemarkFilter,omitempty"`         // 是否过滤统计：1过滤；0不过滤
	JackpotLv        int32         `protobuf:"varint,33,opt,name=JackpotLv,proto3" json:"JackpotLv,omitempty"`               // 解锁Jackpot的Bet档位，机器内spin、领钱、玩法都需要传递，档位：1-4档
	UnlockPlay       int32         `protobuf:"varint,34,opt,name=UnlockPlay,proto3" json:"UnlockPlay,omitempty"`             // 解锁玩法类型 0：未解锁， 1：解锁长线玩法 （如果有新的玩法，可以用2、4、8）
	UnlockMaxJackpot bool          `protobuf:"varint,35,opt,name=UnlockMaxJackpot,proto3" json:"UnlockMaxJackpot,omitempty"` // 是否解锁最高jackpot档位
	ScheduleID       int32         `protobuf:"varint,36,opt,name=ScheduleID,proto3" json:"ScheduleID,omitempty"`             // 目前生效档期
	IncrGems         uint64        `protobuf:"varint,37,opt,name=IncrGems,proto3" json:"IncrGems,omitempty"`                 // 增加的二级货币
	SubGems          uint64        `protobuf:"varint,38,opt,name=SubGems,proto3" json:"SubGems,omitempty"`                   // 消耗的二级货币
	IncrGemsUSD      float64       `protobuf:"fixed64,39,opt,name=IncrGemsUSD,proto3" json:"IncrGemsUSD,omitempty"`          // 增加的二级货币美元值
	SubGemsUSD       float64       `protobuf:"fixed64,40,opt,name=SubGemsUSD,proto3" json:"SubGemsUSD,omitempty"`            // 消耗的二级货币美元值
	Gems             uint64        `protobuf:"varint,41,opt,name=Gems,proto3" json:"Gems,omitempty"`                         // 二级货币余额
	GemsUSD          float64       `protobuf:"fixed64,42,opt,name=GemsUSD,proto3" json:"GemsUSD,omitempty"`                  // 二级货币余额美元
	GemsOneDollar    float64       `protobuf:"fixed64,43,opt,name=GemsOneDollar,proto3" json:"GemsOneDollar,omitempty"`      // 当前1美金价值
	WinJackpots      []*WinJackpot `protobuf:"bytes,44,rep,name=WinJackpots,proto3" json:"WinJackpots,omitempty"`            // 命中jackpot及数量 json格式 eg: [{"Mini":2,"Major":1}] 如果没有中jackpot 字段为空
	UserGroup        string        `protobuf:"bytes,45,opt,name=UserGroup,proto3" json:"UserGroup,omitempty"`                // 玩家所属分组
	Description      string        `protobuf:"bytes,46,opt,name=Description,proto3" json:"Description,omitempty"`            // 说明
	ResVer           string        `protobuf:"bytes,47,opt,name=ResVer,proto3" json:"ResVer,omitempty"`                      // 资源版本号
	MediaSource      string        `protobuf:"bytes,48,opt,name=MediaSource,proto3" json:"MediaSource,omitempty"`            // 玩家注册渠道
	LastMinBet       int64         `protobuf:"varint,49,opt,name=LastMinBet,proto3" json:"LastMinBet,omitempty"`             // 最近的minBet
	LastMaxBet       int64         `protobuf:"varint,50,opt,name=LastMaxBet,proto3" json:"LastMaxBet,omitempty"`             // 最近的maxBet
	RTP200           float64       `protobuf:"fixed64,51,opt,name=RTP200,proto3" json:"RTP200,omitempty"`                    // 最近200次平均RTP
	RTP1000          float64       `protobuf:"fixed64,52,opt,name=RTP1000,proto3" json:"RTP1000,omitempty"`                  // 最近1000次平均RTP
	RFMGroup         string        `protobuf:"bytes,53,opt,name=RFMGroup,proto3" json:"RFMGroup,omitempty"`                  // rfm分组
	BetUnlock        int64         `protobuf:"varint,54,opt,name=BetUnlock,proto3" json:"BetUnlock,omitempty"`               // 解锁BET
	BetType          int32         `protobuf:"varint,55,opt,name=BetType,proto3" json:"BetType,omitempty"`                   // BET类型
	BetParam         string        `protobuf:"bytes,56,opt,name=BetParam,proto3" json:"BetParam,omitempty"`                  // BET参数
	PayLeft          float64       `protobuf:"fixed64,57,opt,name=PayLeft,proto3" json:"PayLeft,omitempty"`                  // 充值剩余筹码
}

func (x *FinLog) Reset() {
	*x = FinLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinLog) ProtoMessage() {}

func (x *FinLog) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinLog.ProtoReflect.Descriptor instead.
func (*FinLog) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{1}
}

func (x *FinLog) GetUID() int32 {
	if x != nil {
		return x.UID
	}
	return 0
}

func (x *FinLog) GetSpecialUser() int32 {
	if x != nil {
		return x.SpecialUser
	}
	return 0
}

func (x *FinLog) GetPaidType() int32 {
	if x != nil {
		return x.PaidType
	}
	return 0
}

func (x *FinLog) GetTotalPay() float64 {
	if x != nil {
		return x.TotalPay
	}
	return 0
}

func (x *FinLog) GetSeriesId() int64 {
	if x != nil {
		return x.SeriesId
	}
	return 0
}

func (x *FinLog) GetSlotId() string {
	if x != nil {
		return x.SlotId
	}
	return ""
}

func (x *FinLog) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *FinLog) GetWin() uint64 {
	if x != nil {
		return x.Win
	}
	return 0
}

func (x *FinLog) GetBet() uint64 {
	if x != nil {
		return x.Bet
	}
	return 0
}

func (x *FinLog) GetBalanceUSD() float64 {
	if x != nil {
		return x.BalanceUSD
	}
	return 0
}

func (x *FinLog) GetWinUSD() float64 {
	if x != nil {
		return x.WinUSD
	}
	return 0
}

func (x *FinLog) GetBetUSD() float64 {
	if x != nil {
		return x.BetUSD
	}
	return 0
}

func (x *FinLog) GetSpinNum() int32 {
	if x != nil {
		return x.SpinNum
	}
	return 0
}

func (x *FinLog) GetIsBroke() int32 {
	if x != nil {
		return x.IsBroke
	}
	return 0
}

func (x *FinLog) GetReason() uint64 {
	if x != nil {
		return x.Reason
	}
	return 0
}

func (x *FinLog) GetGamePlay() int64 {
	if x != nil {
		return x.GamePlay
	}
	return 0
}

func (x *FinLog) GetOneDollar() uint64 {
	if x != nil {
		return x.OneDollar
	}
	return 0
}

func (x *FinLog) GetRegisterTime() int64 {
	if x != nil {
		return x.RegisterTime
	}
	return 0
}

func (x *FinLog) GetAvgBet() float64 {
	if x != nil {
		return x.AvgBet
	}
	return 0
}

func (x *FinLog) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *FinLog) GetWorkScope() int32 {
	if x != nil {
		return x.WorkScope
	}
	return 0
}

func (x *FinLog) GetPlatform() int32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *FinLog) GetOutRTP() float64 {
	if x != nil {
		return x.OutRTP
	}
	return 0
}

func (x *FinLog) GetRTPState() int32 {
	if x != nil {
		return x.RTPState
	}
	return 0
}

func (x *FinLog) GetMinBet() uint64 {
	if x != nil {
		return x.MinBet
	}
	return 0
}

func (x *FinLog) GetTriggerLucky() int32 {
	if x != nil {
		return x.TriggerLucky
	}
	return 0
}

func (x *FinLog) GetLv() int32 {
	if x != nil {
		return x.Lv
	}
	return 0
}

func (x *FinLog) GetVipLv() int32 {
	if x != nil {
		return x.VipLv
	}
	return 0
}

func (x *FinLog) GetLastPayTime() int64 {
	if x != nil {
		return x.LastPayTime
	}
	return 0
}

func (x *FinLog) GetSpend() float64 {
	if x != nil {
		return x.Spend
	}
	return 0
}

func (x *FinLog) GetRemarkUser() int32 {
	if x != nil {
		return x.RemarkUser
	}
	return 0
}

func (x *FinLog) GetRemarkFilter() int32 {
	if x != nil {
		return x.RemarkFilter
	}
	return 0
}

func (x *FinLog) GetJackpotLv() int32 {
	if x != nil {
		return x.JackpotLv
	}
	return 0
}

func (x *FinLog) GetUnlockPlay() int32 {
	if x != nil {
		return x.UnlockPlay
	}
	return 0
}

func (x *FinLog) GetUnlockMaxJackpot() bool {
	if x != nil {
		return x.UnlockMaxJackpot
	}
	return false
}

func (x *FinLog) GetScheduleID() int32 {
	if x != nil {
		return x.ScheduleID
	}
	return 0
}

func (x *FinLog) GetIncrGems() uint64 {
	if x != nil {
		return x.IncrGems
	}
	return 0
}

func (x *FinLog) GetSubGems() uint64 {
	if x != nil {
		return x.SubGems
	}
	return 0
}

func (x *FinLog) GetIncrGemsUSD() float64 {
	if x != nil {
		return x.IncrGemsUSD
	}
	return 0
}

func (x *FinLog) GetSubGemsUSD() float64 {
	if x != nil {
		return x.SubGemsUSD
	}
	return 0
}

func (x *FinLog) GetGems() uint64 {
	if x != nil {
		return x.Gems
	}
	return 0
}

func (x *FinLog) GetGemsUSD() float64 {
	if x != nil {
		return x.GemsUSD
	}
	return 0
}

func (x *FinLog) GetGemsOneDollar() float64 {
	if x != nil {
		return x.GemsOneDollar
	}
	return 0
}

func (x *FinLog) GetWinJackpots() []*WinJackpot {
	if x != nil {
		return x.WinJackpots
	}
	return nil
}

func (x *FinLog) GetUserGroup() string {
	if x != nil {
		return x.UserGroup
	}
	return ""
}

func (x *FinLog) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FinLog) GetResVer() string {
	if x != nil {
		return x.ResVer
	}
	return ""
}

func (x *FinLog) GetMediaSource() string {
	if x != nil {
		return x.MediaSource
	}
	return ""
}

func (x *FinLog) GetLastMinBet() int64 {
	if x != nil {
		return x.LastMinBet
	}
	return 0
}

func (x *FinLog) GetLastMaxBet() int64 {
	if x != nil {
		return x.LastMaxBet
	}
	return 0
}

func (x *FinLog) GetRTP200() float64 {
	if x != nil {
		return x.RTP200
	}
	return 0
}

func (x *FinLog) GetRTP1000() float64 {
	if x != nil {
		return x.RTP1000
	}
	return 0
}

func (x *FinLog) GetRFMGroup() string {
	if x != nil {
		return x.RFMGroup
	}
	return ""
}

func (x *FinLog) GetBetUnlock() int64 {
	if x != nil {
		return x.BetUnlock
	}
	return 0
}

func (x *FinLog) GetBetType() int32 {
	if x != nil {
		return x.BetType
	}
	return 0
}

func (x *FinLog) GetBetParam() string {
	if x != nil {
		return x.BetParam
	}
	return ""
}

func (x *FinLog) GetPayLeft() float64 {
	if x != nil {
		return x.PayLeft
	}
	return 0
}

type WinJackpot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JackpotID int32 `protobuf:"varint,1,opt,name=JackpotID,proto3" json:"JackpotID,omitempty"`
	Cnt       int32 `protobuf:"varint,2,opt,name=Cnt,proto3" json:"Cnt,omitempty"`
}

func (x *WinJackpot) Reset() {
	*x = WinJackpot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinJackpot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinJackpot) ProtoMessage() {}

func (x *WinJackpot) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinJackpot.ProtoReflect.Descriptor instead.
func (*WinJackpot) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{2}
}

func (x *WinJackpot) GetJackpotID() int32 {
	if x != nil {
		return x.JackpotID
	}
	return 0
}

func (x *WinJackpot) GetCnt() int32 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

var File_cmd_proto protoreflect.FileDescriptor

var file_cmd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x6d, 0x64,
	0x22, 0x48, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x46, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x25, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6d, 0x64, 0x2e, 0x46, 0x69, 0x6e, 0x4c, 0x6f,
	0x67, 0x52, 0x07, 0x46, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0xe3, 0x0c, 0x0a, 0x06, 0x46,
	0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x69,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x61, 0x69,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x57, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x57, 0x69,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x42, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x53,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x55, 0x53, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x69, 0x6e, 0x55, 0x53, 0x44, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x57, 0x69, 0x6e, 0x55, 0x53, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x42,
	0x65, 0x74, 0x55, 0x53, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x42, 0x65, 0x74,
	0x55, 0x53, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x70, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x70, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x49, 0x73, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x49, 0x73, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x4f,
	0x6e, 0x65, 0x44, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x4f, 0x6e, 0x65, 0x44, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x76, 0x67, 0x42, 0x65, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41,
	0x76, 0x67, 0x42, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x57, 0x6f, 0x72, 0x6b, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x16, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x52, 0x54, 0x50, 0x18, 0x17, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x4f, 0x75, 0x74, 0x52, 0x54, 0x50, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x54, 0x50, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x52, 0x54, 0x50, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x69, 0x6e, 0x42, 0x65, 0x74, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x4d, 0x69, 0x6e, 0x42, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4c, 0x75, 0x63, 0x6b, 0x79, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x4c, 0x75, 0x63, 0x6b, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x4c, 0x76, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x4c, 0x76, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x69, 0x70, 0x4c, 0x76, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x56, 0x69, 0x70, 0x4c, 0x76, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x50, 0x61, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x4c, 0x61, 0x73, 0x74,
	0x50, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x6e, 0x64,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x0c, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x4c, 0x76, 0x18, 0x21,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x4c, 0x76, 0x12,
	0x1e, 0x0a, 0x0a, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x18, 0x22, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x6c, 0x61, 0x79, 0x12,
	0x2a, 0x0a, 0x10, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x61, 0x78, 0x4a, 0x61, 0x63, 0x6b,
	0x70, 0x6f, 0x74, 0x18, 0x23, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x55, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x4d, 0x61, 0x78, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x24, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x49,
	0x6e, 0x63, 0x72, 0x47, 0x65, 0x6d, 0x73, 0x18, 0x25, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x49,
	0x6e, 0x63, 0x72, 0x47, 0x65, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x47, 0x65,
	0x6d, 0x73, 0x18, 0x26, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x53, 0x75, 0x62, 0x47, 0x65, 0x6d,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x6e, 0x63, 0x72, 0x47, 0x65, 0x6d, 0x73, 0x55, 0x53, 0x44,
	0x18, 0x27, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x49, 0x6e, 0x63, 0x72, 0x47, 0x65, 0x6d, 0x73,
	0x55, 0x53, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x47, 0x65, 0x6d, 0x73, 0x55, 0x53,
	0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x53, 0x75, 0x62, 0x47, 0x65, 0x6d, 0x73,
	0x55, 0x53, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x47, 0x65, 0x6d, 0x73, 0x18, 0x29, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x47, 0x65, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x65, 0x6d, 0x73, 0x55,
	0x53, 0x44, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x47, 0x65, 0x6d, 0x73, 0x55, 0x53,
	0x44, 0x12, 0x24, 0x0a, 0x0d, 0x47, 0x65, 0x6d, 0x73, 0x4f, 0x6e, 0x65, 0x44, 0x6f, 0x6c, 0x6c,
	0x61, 0x72, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x47, 0x65, 0x6d, 0x73, 0x4f, 0x6e,
	0x65, 0x44, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x12, 0x31, 0x0a, 0x0b, 0x57, 0x69, 0x6e, 0x4a, 0x61,
	0x63, 0x6b, 0x70, 0x6f, 0x74, 0x73, 0x18, 0x2c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x6d, 0x64, 0x2e, 0x57, 0x69, 0x6e, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x52, 0x0b, 0x57,
	0x69, 0x6e, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x56, 0x65, 0x72, 0x18, 0x2f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x56,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x30, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x69, 0x6e, 0x42,
	0x65, 0x74, 0x18, 0x31, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x69,
	0x6e, 0x42, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x61, 0x78, 0x42,
	0x65, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x61,
	0x78, 0x42, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x54, 0x50, 0x32, 0x30, 0x30, 0x18, 0x33,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x52, 0x54, 0x50, 0x32, 0x30, 0x30, 0x12, 0x18, 0x0a, 0x07,
	0x52, 0x54, 0x50, 0x31, 0x30, 0x30, 0x30, 0x18, 0x34, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x52,
	0x54, 0x50, 0x31, 0x30, 0x30, 0x30, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x46, 0x4d, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x35, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x46, 0x4d, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x65, 0x74, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x36, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x42, 0x65, 0x74, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x18, 0x0a, 0x07, 0x42, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x37, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x42, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x65,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x38, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x65,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x66,
	0x74, 0x18, 0x39, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x50, 0x61, 0x79, 0x4c, 0x65, 0x66, 0x74,
	0x22, 0x3c, 0x0a, 0x0a, 0x57, 0x69, 0x6e, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x4a, 0x61, 0x63, 0x6b, 0x70, 0x6f, 0x74, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03,
	0x43, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x43, 0x6e, 0x74, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_proto_rawDescOnce sync.Once
	file_cmd_proto_rawDescData = file_cmd_proto_rawDesc
)

func file_cmd_proto_rawDescGZIP() []byte {
	file_cmd_proto_rawDescOnce.Do(func() {
		file_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_proto_rawDescData)
	})
	return file_cmd_proto_rawDescData
}

var file_cmd_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmd_proto_goTypes = []interface{}{
	(*BatchFinLog)(nil), // 0: cmd.BatchFinLog
	(*FinLog)(nil),      // 1: cmd.FinLog
	(*WinJackpot)(nil),  // 2: cmd.WinJackpot
}
var file_cmd_proto_depIdxs = []int32{
	1, // 0: cmd.BatchFinLog.FinLogs:type_name -> cmd.FinLog
	2, // 1: cmd.FinLog.WinJackpots:type_name -> cmd.WinJackpot
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cmd_proto_init() }
func file_cmd_proto_init() {
	if File_cmd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchFinLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinJackpot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_proto_goTypes,
		DependencyIndexes: file_cmd_proto_depIdxs,
		MessageInfos:      file_cmd_proto_msgTypes,
	}.Build()
	File_cmd_proto = out.File
	file_cmd_proto_rawDesc = nil
	file_cmd_proto_goTypes = nil
	file_cmd_proto_depIdxs = nil
}
