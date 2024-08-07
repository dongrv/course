package abstract

// 第一阶段

// TODO 基本特征
// 犄角、四条腿、尾巴

type Feature interface {
	Horns() uint8  // 几个犄角
	Hooves() uint8 // 几个蹄子
	Tail() bool    // 有没有尾巴
}

// TODO 基本行为
// 吃草、迁徙

type Behavior interface {
	Eat()     // 吃东西
	Run()     // 跑
	Migrate() // 迁徙
}

type Bull interface {
	Feature
	Behavior
}

// AmericanBison 美洲野牛
type AmericanBison struct {
	Bull
}

type AfricanBison struct {
	Bull
}

func (ab *AfricanBison) Eat() {
	println("非洲野牛吃草...")
}

func (ab *AfricanBison) Run() {
	println("非洲野牛跑...")
}

func (ab *AfricanBison) Migrate() {
	println("非洲野牛迁徙")
}

func (ab *AfricanBison) Horns() uint8 {
	println("非洲野牛有2个犄角")
	return 2
}
func (ab *AfricanBison) Hooves() uint8 {
	println("非洲野牛有4个蹄子")
	return 4
}

func (ab *AfricanBison) Tail() bool {
	println("非洲野牛有1个尾巴")
	return true
}

// YellowCattle 大黄牛
type YellowCattle struct {
	Bull
}

func (yc *YellowCattle) Eat() {
	println("黄牛吃草...")
}

func (yc *YellowCattle) Run() {
	println("黄牛跑...")
}

func (yc *YellowCattle) Migrate() {
	println("黄牛不迁徙")
}

func (yc *YellowCattle) Horns() uint8 {
	println("黄牛有2个犄角")
	return 2
}
func (yc *YellowCattle) Hooves() uint8 {
	println("黄牛有4个蹄子")
	return 4
}

func (yc *YellowCattle) Tail() bool {
	println("黄牛有1个尾巴")
	return true
}

//

type Animal interface {
	Eat() // 吃东西
	Run() // 跑
}

type Tiger struct{}

func (t *Tiger) Eat() {
	println("老虎吃肉...")
}

func (t *Tiger) Run() {
	println("老虎跑...")
}

type Lion struct{}

func (l *Lion) Eat() {
	println("狮子吃肉...")
}

func (l *Lion) Run() {
	println("狮子跑...")
}

func BullDo(bull Bull) {
	bull.Horns()
	bull.Hooves()
	bull.Tail()
	bull.Eat()
	bull.Run()
	bull.Migrate()
}

func AnimalDo(animal Animal) {
	animal.Eat()
	animal.Run()
}

//

type Handler interface {
	Do()
}

type HandleFunc func()

func (h HandleFunc) Do() {
	h()
}

func Name() Handler {
	return HandleFunc(func() {
		println("老虎")
	})
}

func EatHandler(h Handler) Handler {
	return HandleFunc(func() {
		h.Do()
		println("吃东西")
	})
}

func RunHandler(h Handler) Handler {
	return HandleFunc(func() {
		h.Do()
		println("跑起来")
	})
}

func SayHandler(h Handler) Handler {
	return HandleFunc(func() {
		h.Do()
		println("咆哮")
	})
}

type Base struct {
	ID int32
}

func (b *Base) Call() {
	println("base call")
}

type Derived struct {
	Base
	ID int32
}

func (d *Derived) Call() {
	println("derived call")
}

func Call() {
	derived := Derived{
		Base: Base{ID: 10},
		ID:   100,
	}
	println(derived.Base.ID, derived.ID)
	derived.Base.Call()
	derived.Call()
}
