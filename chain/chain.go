// Package chain 责任链
package chain

type BoardingProcessor interface {
	SetNextProcessor(processor BoardingProcessor)
	ProcessFor(passenger *Passenger)
}

type Passenger struct {
	name                string // 姓名
	hasBoardingPass     bool   // 是否已办理登机牌
	hasLuggage          bool   // 是否有行李托运
	isPassIdentifyCheck bool   // 是否通过身份校验
	isPassSecurityCheck bool   // 是否通过安检
	isCompleteBoard     bool   // 是否完成等级
}

// baseBoardingProcessor 登机流程基类
type baseBoardingProcessor struct {
	nextProcessor BoardingProcessor // 下一个登机流程
}

func (b *baseBoardingProcessor) SetNextProcessor(processor BoardingProcessor) {
	b.nextProcessor = processor
}

func (b *baseBoardingProcessor) ProcessFor(passenger *Passenger) {
	if b.nextProcessor != nil {
		b.nextProcessor.ProcessFor(passenger)
	}
}

// boardingPassProcessor 办理登机牌处理器
type boardingPassProcessor struct {
	baseBoardingProcessor
}

func (b *boardingPassProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		println("办理登机牌")
		passenger.hasBoardingPass = true
	}
	b.baseBoardingProcessor.ProcessFor(passenger)
}

type luggageCheckInProcessor struct {
	baseBoardingProcessor
}

func (l *luggageCheckInProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		println("未办理登机牌，不能托运行李")
		passenger.hasBoardingPass = true
	}
	l.baseBoardingProcessor.ProcessFor(passenger)
}

type identityCheckProcessor struct {
	baseBoardingProcessor
}

func (i *identityCheckProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		println("没有办理登机牌，不能进行身份核验")
		return
	}
	if !passenger.isPassIdentifyCheck {
		println("核验身份")
		passenger.isPassIdentifyCheck = true
	}
	i.baseBoardingProcessor.ProcessFor(passenger)
}

type securityCheckProcessor struct {
	baseBoardingProcessor
}

func (s *securityCheckProcessor) ProcessFor(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		println("没有办理登机牌，不能进行身份核验")
		return
	}
	if !passenger.isPassSecurityCheck {
		println("为乘客安检")
		passenger.isPassSecurityCheck = true
	}
	s.baseBoardingProcessor.ProcessFor(passenger)
}

type completeBoardingProcessor struct {
	baseBoardingProcessor
}

func (c *completeBoardingProcessor) ProcessFor(passenger *Passenger) {
	if passenger.hasBoardingPass && passenger.isPassIdentifyCheck && passenger.isPassSecurityCheck {
		println("乘客可以登机")
		passenger.isCompleteBoard = true
		return
	}
	println("登机检查流程未完成，不能登机")
}
