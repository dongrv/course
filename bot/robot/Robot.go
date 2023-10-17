package robot

import (
	"bot/dialog"
	"bot/msg/cmd"
	"bot/msg/protocol"
	"context"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

// Center 机器人管理中心
type Center struct {
	wg       *sync.WaitGroup
	robots   []*Robot
	RobotNum int
	Prefix   string // 设备号前缀
}

func NewCenter(robotNum int, prefix string, addr string) *Center {
	r := &Center{
		wg:       new(sync.WaitGroup),
		robots:   make([]*Robot, 0, robotNum),
		RobotNum: robotNum,
		Prefix:   prefix,
	}
	// 构建机器人
	for i := 0; i < r.RobotNum; i++ {
		r.robots = append(r.robots, &Robot{DID: r.Prefix + `-` + strconv.Itoa(i), Addr: addr})
	}
	return r
}

func (m *Center) Activate(ctx context.Context) {
	if len(m.robots) == 0 {
		logrus.Debug("没有待激活的机器人")
		return
	}
	for _, robot := range m.robots {
		m.wg.Add(1)
		go robot.Run(ctx, m.wg)
		logrus.Debug("已激活机器人:" + robot.DID)
	}
}

func (m *Center) Wait() {
	logrus.Debug("Center.wait()")
	for _, r := range m.robots {
		if r.Conn == nil {
			continue
		}
		_ = r.Conn.Close()
	}
	m.wg.Wait()
}

// Robot 机器人基础信息
type Robot struct {
	Conn net.Conn
	DID  string
	Addr string // 目标服务器地址
}

func (r *Robot) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ln, err := net.Dial(`tcp`, r.Addr)
	if err != nil {
		logrus.Errorf("%s 拨号错误：%s", r.DID, err.Error())
		return
	}
	r.Conn = ln // 绑定连接

	user := &protocol.User{DID: r.DID}
	d := dialog.NewDialog()

	go createStrategy(ctx, d, user)   // 请求协程
	go d.OnProcess(ctx, wrapFn(user)) // 处理服务器返回值
	go d.OnNotify(ctx)                // 推送请求到服务器

	d.Start(ctx, ln) // 启动对话读写协程

}

func wrapFn(user *protocol.User) func(msgId int32, msg []byte) {
	return func(msgId int32, msg []byte) {
		if fn := protocol.Find(msgId); fn != nil {
			var err error
			if err = fn(user, msg); err != nil {
				logrus.Error(err.Error())
			}
			if err == cmd.ErrInterrupt {
				return
			}
		}
	}
}

// 请求计划和策略
func createStrategy(ctx context.Context, d *dialog.Dialog, user *protocol.User) {
	plan := PlanDoc()
	for {
		select {
		case <-ctx.Done():
			logrus.Debug("")
			return
		default:
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond) // 控制动作速度
			meta := plan.Next()
			if meta == Empty {
				continue
			}
			// 守护协程
			if meta.Repeat.Daemon {
				go pendAction(ctx, d, user, meta)
			}
			// 重复动作
			if meta.Repeat.TickNum > 0 {
				logrus.Debugf("%s send req->%s", user.DID, protocol.MsgName(int32(meta.Seq)))
				repeatAction(meta, d, user)
				continue
			}
			// 单次动作
			logrus.Debugf("%s send req->%s", user.DID, protocol.MsgName(int32(meta.Seq)))
			d.InputAction(wrapAction(meta, user))
		}
	}
}

// 包装动作
func wrapAction(meta Meta, user *protocol.User) dialog.IAction {
	return &dialog.Event{SeqNum: meta.Seq, Msg: protocol.FindWrap(int32(meta.Seq))(user)}
}

// pendAction 挂起协程
func pendAction(ctx context.Context, d *dialog.Dialog, user *protocol.User, meta Meta) {
	logrus.Debug("pendAction goroutine is on")
	for {
		select {
		case <-ctx.Done():
			logrus.Debug("Daemon goroutine exit")
			return
		case <-time.Tick(meta.Repeat.Tick):
			logrus.Debugf("%s send req->%s", user.DID, protocol.MsgName(int32(meta.Seq)))
			d.InputAction(wrapAction(meta, user))
		}
	}
}

// 重复动作
func repeatAction(meta Meta, d *dialog.Dialog, user *protocol.User) {
	for i := 0; i < meta.Repeat.TickNum; i++ {
		time.Sleep(meta.Repeat.Tick)
		d.InputAction(wrapAction(meta, user))
	}
}

// PlanDoc 计划表
func PlanDoc() *Plan {
	// TODO 从JSON文件加载
	return &Plan{
		Must: []Meta{
			{Seq: 3001}, {Seq: 3011}, {Seq: 3013}, {Seq: 3015}, {Seq: 3019}, {Seq: 3047},
			{Seq: 3071, Repeat: Repeat{Daemon: true, Tick: 10 * time.Second}}, // 心跳请求
		},
		Random: []Meta{
			{Seq: 3015, Weight: 20}, {Seq: 3019, Weight: 30}, {Seq: 3047, Weight: 50},
		},
	}
}
