package robot

import (
	"casino/protocol/cmd"
	"math/rand"
	"time"
)

// Meta 访问单元
type Meta struct {
	Seq        int
	Repeat     Repeat        // 重复访问
	Weight     int           // 权重
	ExpectCode cmd.ErrorCode // 应该获取的错误码
}

type Repeat struct {
	Daemon  bool          // 守护进程
	Tick    time.Duration // 发送间隔
	TickNum int           // 发送次数
}

type Plan struct {
	step   int    // 访问位置
	flag   bool   // 是否已访问必须序列
	Must   []Meta // 必须请求序列号
	Random []Meta // 随机请求序列号
}

func NewPlan() *Plan {
	return new(Plan)
}

type INext interface {
	Next() Meta
}

var Empty = Meta{}

// Next 下一个访问序号
func (p *Plan) Next() Meta {
	if !p.flag {
		if p.step >= len(p.Must) {
			p.flag = true
			goto end
		}
		c := p.step
		p.step++
		//logrus.Debugf("Plan.Must: %d", p.Must[c].Seq)
		return p.Must[c]
	}

end:
	if len(p.Random) == 0 {
		return Empty
	}
	total := sum(p.Random)
	if total == 0 {
		return p.Random[rand.Intn(len(p.Random))] // 纯随机
	}

	// 权重控制
	random := rand.Intn(total)
	for k, v := range p.Random {
		random -= v.Weight
		if random < 0 {
			//logrus.Debugf("Plan.Random: %d", p.Random[k].Seq)
			return p.Random[k]
		}
	}
	return Empty // empty meta
}

func sum(ws []Meta) int {
	var total int
	for _, w := range ws {
		total += w.Weight
	}
	return total
}
