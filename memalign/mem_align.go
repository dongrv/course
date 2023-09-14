package memalign

import "unsafe"

// BasicTypeAlign 基本类型对齐
func BasicTypeAlign() {
	println("###Go基本类型内存对齐测试###\n")
	var byteVar byte = 'a'
	var runeVar rune = '1'

	var int8Var int8 = 0
	var int16Var int16 = 0
	var int32Var int32 = 0
	var int64Var int64 = 0
	var uint8Var uint8 = 0

	var uint16Var uint16 = 0
	var uint32Var uint32 = 0
	var uint64Var uint64 = 0

	var intVar = 0
	var uintVar = 0

	var float32Var float32 = 0.0
	var float64Var float64 = 0.0

	var stringVar = ""
	var stringsVar []string
	var int32sVar []int32
	var int64sVar []int64

	println("byte最大对齐长度\t", unsafe.Alignof(byteVar), "总字节数：", unsafe.Sizeof(byteVar))
	println("rune最大对齐长度\t", unsafe.Alignof(runeVar), "总字节数：", unsafe.Sizeof(runeVar))
	println()
	println("int8最大对齐长度\t", unsafe.Alignof(int8Var), "总字节数：", unsafe.Sizeof(int8Var))
	println("int16最大对齐长度\t", unsafe.Alignof(int16Var), "总字节数：", unsafe.Sizeof(int16Var))
	println("int32最大对齐长度\t", unsafe.Alignof(int32Var), "总字节数：", unsafe.Sizeof(int32Var))
	println("int64最大对齐长度\t", unsafe.Alignof(int64Var), "总字节数：", unsafe.Sizeof(int64Var))
	println()
	println("uint8最大对齐长度\t", unsafe.Alignof(uint8Var), "总字节数：", unsafe.Sizeof(uint16Var))
	println("uint16最大对齐长度\t", unsafe.Alignof(uint16Var), "总字节数：", unsafe.Sizeof(uint16Var))
	println("uint32最大对齐长度\t", unsafe.Alignof(uint32Var), "总字节数：", unsafe.Sizeof(uint32Var))
	println("uint64最大对齐长度\t", unsafe.Alignof(uint64Var), "总字节数：", unsafe.Sizeof(uint64Var))
	println()
	println("int最大对齐长度\t", unsafe.Alignof(intVar), "总字节数：", unsafe.Sizeof(intVar))
	println("uint最大对齐长度\t", unsafe.Alignof(uintVar), "总字节数：", unsafe.Sizeof(uintVar))
	println()
	println("float32最大对齐长度\t", unsafe.Alignof(float32Var), "总字节数：", unsafe.Sizeof(float32Var))
	println("float64最大对齐长度\t", unsafe.Alignof(float64Var), "总字节数：", unsafe.Sizeof(float64Var))
	println()
	println("string最大对齐长度\t", unsafe.Alignof(stringVar), "总字节数：", unsafe.Sizeof(stringVar))
	println("string切片最大对齐长度\t", unsafe.Alignof(stringsVar), "总字节数：", unsafe.Sizeof(stringsVar))
	println("int32切片最大对齐长度\t", unsafe.Alignof(int32sVar), "总字节数：", unsafe.Sizeof(int32sVar))
	println("int64切片最大对齐长度\t", unsafe.Alignof(int64sVar), "总字节数：", unsafe.Sizeof(int64sVar))
}

// StructTypeAlign 结构体对齐
func StructTypeAlign() {
	t1 := StructTest1{}
	t2 := StructTest2{}
	t3 := StructTest3{}
	t4 := StructTest4{}
	t5 := StructTest5{}
	println("StructTest1最大对齐长度：", unsafe.Alignof(t1), "总字节数：", unsafe.Sizeof(t1))
	println("StructTest1最大对齐长度：", unsafe.Alignof(t2), "总字节数：", unsafe.Sizeof(t2))
	println("StructTest1最大对齐长度：", unsafe.Alignof(t3), "总字节数：", unsafe.Sizeof(t3))
	println("StructTest1最大对齐长度：", unsafe.Alignof(t4), "总字节数：", unsafe.Sizeof(t4))
	println("StructTest1最大对齐长度：", unsafe.Alignof(t5), "总字节数：", unsafe.Sizeof(t5))
}

type StructTest1 struct {
	a int8     // 对齐字节长度：1
	b int32    // 对齐字节长度：4
	c []string // 对齐字节长度：8
}

type StructTest2 struct {
	c []string // 对齐字节长度：8
	b int64    // 对齐字节长度：8
	a int32    // 对齐字节长度：4
}

type StructTest3 struct {
	a int8
	b int64
	c struct{}
}

type StructTest4 struct {
	a struct{}
	b int8
	c int32
}

type StructTest5 struct {
	a struct{}
}

func StructTypeAlign2() {
	oldStruct := UserOld{}
	newStruct := UserNew{}
	println("UserOld最大对齐长度：", unsafe.Alignof(oldStruct), "总字节数：", unsafe.Sizeof(oldStruct))
	println("UserNew最大对齐长度：", unsafe.Alignof(newStruct), "总字节数：", unsafe.Sizeof(newStruct))
}

type UserOld struct {
	ScheduleID   int32   // 目前生效档期数
	RoundNum     int32   // 用户轮次
	UnlockBet    uint64  // 当前轮次对应的活动解锁bet
	LV           int32   // 本轮用户等级
	VipLV        int32   // 当前用户等级
	BuyNum       int32   // 当天矿镐购买次数
	RefreshTime  int64   // 刷新时间
	PickNum      int32   // 矿镐数量
	IsGetPick    int32   // 是否获取首次矿镐
	SilverNum    int32   // 银矿数量
	BarrierSeq   int32   // 当前关卡号
	StoneArr     []int32 // 已挖石头序号 0、1、2、3
	DigRewardArr []int32 // 已挖奖励序号
	WheelArr     []int32 // 转轮转的序号
	WheelState   int32   // 转轮状态 0：正常 1：转轮待进入 2：转轮转动
	MissionArr   []int32 // 任务编号
}

type UserNew struct {
	ScheduleID   int32   // 目前生效档期数
	RoundNum     int32   // 用户轮次
	LV           int32   // 本轮用户等级
	VipLV        int32   // 当前用户等级
	BuyNum       int32   // 当天矿镐购买次数
	PickNum      int32   // 矿镐数量
	IsGetPick    int32   // 是否获取首次矿镐
	SilverNum    int32   // 银矿数量
	BarrierSeq   int32   // 当前关卡号
	WheelState   int32   // 转轮状态 0：正常 1：转轮待进入 2：转轮转动
	UnlockBet    uint64  // 当前轮次对应的活动解锁bet
	RefreshTime  int64   // 刷新时间
	StoneArr     []int32 // 已挖石头序号 0、1、2、3
	DigRewardArr []int32 // 已挖奖励序号
	WheelArr     []int32 // 转轮转的序号
	MissionArr   []int32 // 任务编号
}

func StructTypeAlign3() {
	oldStruct := FreeSpinDataOld{}
	newStruct := FreeSpinDataNew{}
	println("UserOld最大对齐长度：", unsafe.Alignof(oldStruct), "总字节数：", unsafe.Sizeof(oldStruct))
	println("UserNew最大对齐长度：", unsafe.Alignof(newStruct), "总字节数：", unsafe.Sizeof(newStruct))
}

type FreeSpinDataOld struct {
	HadSpinCnt      int    `json:"HadSpinCnt,omitempty"`      // 已转了几次FreeSpin
	LastSpinLeftCnt int    `json:"LastSpinLeftCnt,omitempty"` // 最后一次spin剩余次数
	TotalSpinCnt    int    `json:"TotalSpinCnt,omitempty"`    // 总FreeSpin次数
	HadFinish       bool   `json:"HadFinish,omitempty"`       // 已经完成了， 假如总共6次，第5次就强制结束了，此标记直接改为True
	TotalWin        uint64 `json:"TotalWin,omitempty"`        // 累计奖金
	TotalBonusWin   uint64 `json:"TotalBonusWin,omitempty"`   // 额外奖励奖金 包括jackpot
	Bet             uint64 `json:"Bet,omitempty"`             // 触发FreeSpin时的Bet
	RowCnt          int    `json:"RowCnt,omitempty"`          // 转轮行数 6002机器用到这个参数
	HadReSpin       bool   `json:"HadReSpin,omitempty"`       // 是否已经进行过ReSpin

	BonusValue       uint64  `json:"BonusValue,omitempty"`       // 6007 价值字段 6001中scatter_low总价值
	PayWindowCnt     int     `json:"PayWindowCnt,omitempty"`     // 屏幕数量
	AddPayWindow     int     `json:"AddPayWindow,omitempty"`     // 增加屏幕数量
	LockCfgId        []int32 `json:"LockCfgId,omitempty"`        // 6010锁列 NODE: 已上线机器 json字段后面的空格暂不能删除
	FreeSpinWheelCnt int     `json:"FreeSpinWheelCnt,omitempty"` // FreeSpinWheel 剩余次数 NODE: 已上线机器 json字段后面的空格暂不能删除
	FreeSpinSlotID   string  `json:"FreeSpinSlotID,omitempty"`   // FreeSpin使用的SlotID
	CenterMulti      int32   `json:"CenterMulti,omitempty"`      // 6008用 中间Wild的倍数

	ReSpinTriggerCnt int     `json:"ReSpinTriggerCnt,omitempty"` //触发respin的次数
	LockSymbolFlag   int32   `json:"LockSymbolFlag,omitempty"`   // 6018 记录fs锁wild的位置 6044 记录fs使用Wild权重Idx
	LockSymbols      []int32 `json:"LockSymbols,omitempty"`      // 6018 锁定标签
	AddSpinTimes     int32   `json:"AddSpinTimes,omitempty"`     // 增加转动次数
	SpinReel         string  `json:"SpinReel,omitempty"`         // fs转动表
	AddStickyCnt     int32   `json:"AddStickyCnt,omitempty"`     // 进入freeSpin增加的Sticky标签数--6004

	IsMinimum   bool    `json:"IsMinimum,omitempty"`   // 是否已保底转动过
	Multi       float64 `json:"Multi,omitempty"`       // 赢钱翻倍 6032
	MinimumSpin int32   `json:"MinimumSpin,omitempty"` // 触发保底的转数
	PlaySign    string  `json:"PlaySign,omitempty"`    // 不同玩法标记 --6024

	Rates []float64 `json:"Rates,omitempty"` // 6025 概率列表
}

type FreeSpinDataNew struct {
	HadReSpin       bool `json:"HadReSpin,omitempty"`       // 是否已经进行过ReSpin
	HadFinish       bool `json:"HadFinish,omitempty"`       // 已经完成了， 假如总共6次，第5次就强制结束了，此标记直接改为True
	IsMinimum       bool `json:"IsMinimum,omitempty"`       // 是否已保底转动过
	HadSpinCnt      int  `json:"HadSpinCnt,omitempty"`      // 已转了几次FreeSpin
	LastSpinLeftCnt int  `json:"LastSpinLeftCnt,omitempty"` // 最后一次spin剩余次数
	TotalSpinCnt    int  `json:"TotalSpinCnt,omitempty"`    // 总FreeSpin次数

	TotalWin      uint64 `json:"TotalWin,omitempty"`      // 累计奖金
	TotalBonusWin uint64 `json:"TotalBonusWin,omitempty"` // 额外奖励奖金 包括jackpot
	Bet           uint64 `json:"Bet,omitempty"`           // 触发FreeSpin时的Bet
	RowCnt        int    `json:"RowCnt,omitempty"`        // 转轮行数 6002机器用到这个参数

	BonusValue       uint64  `json:"BonusValue,omitempty"`       // 6007 价值字段 6001中scatter_low总价值
	PayWindowCnt     int     `json:"PayWindowCnt,omitempty"`     // 屏幕数量
	AddPayWindow     int     `json:"AddPayWindow,omitempty"`     // 增加屏幕数量
	LockCfgId        []int32 `json:"LockCfgId,omitempty"`        // 6010锁列 NODE: 已上线机器 json字段后面的空格暂不能删除
	FreeSpinWheelCnt int     `json:"FreeSpinWheelCnt,omitempty"` // FreeSpinWheel 剩余次数 NODE: 已上线机器 json字段后面的空格暂不能删除
	FreeSpinSlotID   string  `json:"FreeSpinSlotID,omitempty"`   // FreeSpin使用的SlotID
	CenterMulti      int32   `json:"CenterMulti,omitempty"`      // 6008用 中间Wild的倍数

	ReSpinTriggerCnt int     `json:"ReSpinTriggerCnt,omitempty"` //触发respin的次数
	LockSymbolFlag   int32   `json:"LockSymbolFlag,omitempty"`   // 6018 记录fs锁wild的位置 6044 记录fs使用Wild权重Idx
	LockSymbols      []int32 `json:"LockSymbols,omitempty"`      // 6018 锁定标签
	AddSpinTimes     int32   `json:"AddSpinTimes,omitempty"`     // 增加转动次数
	SpinReel         string  `json:"SpinReel,omitempty"`         // fs转动表
	AddStickyCnt     int32   `json:"AddStickyCnt,omitempty"`     // 进入freeSpin增加的Sticky标签数--6004

	Multi       float64 `json:"Multi,omitempty"`       // 赢钱翻倍 6032
	MinimumSpin int32   `json:"MinimumSpin,omitempty"` // 触发保底的转数
	PlaySign    string  `json:"PlaySign,omitempty"`    // 不同玩法标记 --6024

	Rates []float64 `json:"Rates,omitempty"` // 6025 概率列表
}

func StructTypeAlign4() {
	println("最大对齐长度：", unsafe.Alignof(StructTypeChan{}), "总字节数：", unsafe.Sizeof(StructTypeChan{}))
}

type StructTypeChan struct {
	ch chan struct{} // 引用类型
}
