// Package cmd 系统协议管理
package cmd

import (
	"bot/msg/protocol"
	"casino/protocol/cmd"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var ErrInterrupt = errors.New("中断程序")

func pbName(msg proto.Message) string {
	return string(msg.ProtoReflect().Descriptor().FullName())
}

func RegisterFunc() {
	protocol.WrapFunc(3001, WrapLogin) // 包装请求体
	protocol.BindFunc(3002, LoginResp) // 返回值处理
	protocol.WrapFunc(3011, WrapInitUser)
	protocol.BindFunc(3012, InitUserResp)
	protocol.WrapFunc(3013, WrapInitSlots)
	protocol.BindFunc(3014, InitSlotsResp)
	protocol.WrapFunc(3015, WrapInitSys)
	protocol.BindFunc(3016, InitSysResp)
	protocol.BindFunc(3018, InitSceneResp)
	protocol.WrapFunc(3019, WrapSceneSwitch)
	protocol.BindFunc(3020, SceneSwitchResp)
	protocol.WrapFunc(3047, WrapShopPanel)
	protocol.BindFunc(3048, ShopPanelResp)
	protocol.WrapFunc(3071, WrapHeartbeat)
	protocol.BindFunc(3074, HeartBeatS2C)
	protocol.WrapFunc(3075, WrapReconnectLogin)
	protocol.BindFunc(3076, ReconnectLoginResp)

}

func WrapLogin(u *protocol.User) proto.Message {
	return &cmd.LoginReq{DID: u.DID, Platform: 1, LoginType: 1}
}

func LoginResp(u *protocol.User, msg []byte) error {
	r := &cmd.LoginResp{}
	err := proto.Unmarshal(msg, r)
	if err != nil {
		return fmt.Errorf("登录失败：%s", err.Error())
	}
	u.Summary, u.Token, u.NowMs = r.User, r.Token, r.TimestampMs
	u.Location = protocol.Lobby // 进入大厅
	if ValidateCode(r.Err, 0) != nil {
		return ErrInterrupt
	}
	logrus.Debugf("%s receive resp->%s", u.DID, pbName(r))
	return nil
}

func WrapInitUser(u *protocol.User) proto.Message {
	return &cmd.InitUserReq{UID: u.Summary.UID}
}

func InitUserResp(u *protocol.User, msg []byte) error {
	r := &cmd.InitUserResp{}
	_ = proto.Unmarshal(msg, r)
	u.GameInfo = &cmd.GameInfo{}
	u.GameInfo.Account = &cmd.Account{
		Balance: r.Balance,
		Level:   r.Lv,
		Exp:     r.Exp,
		VIP:     r.VIP,
	}
	logrus.Debugf("%s receive resp->%s", u.DID, pbName(r))
	return ValidateCode(r.Err, 0)
}

func WrapInitSlots(u *protocol.User) proto.Message {
	return &cmd.InitSlotsReq{UID: u.Summary.UID}
}

func InitSlotsResp(u *protocol.User, msg []byte) error {
	r := &cmd.InitSlotsResp{}
	_ = proto.Unmarshal(msg, r)
	u.Slots = r.Slots
	logrus.Debugf("%s receive resp->%s", u.DID, pbName(r))
	return ValidateCode(r.Err, 0)
}

func WrapInitSys(u *protocol.User) proto.Message {
	return &cmd.InitSysReq{UID: u.Summary.UID}
}

func InitSysResp(u *protocol.User, msg []byte) error {
	r := &cmd.InitSysResp{}
	_ = proto.Unmarshal(msg, r)
	u.Activities = r.Activities
	logrus.Debugf("%s receive resp->%s", u.DID, pbName(r))
	return ValidateCode(r.Err, 0)
}

func WrapShopPanel(u *protocol.User) proto.Message {
	return &cmd.ShopPanelReq{UID: u.Summary.UID}
}

// ShopPanelResp 商城面板
func ShopPanelResp(u *protocol.User, msg []byte) error {
	r := &cmd.ShopPanelResp{}
	_ = proto.Unmarshal(msg, r)
	u.Shop = &protocol.Shop{
		Bonus: r.ShopBonus,
		Items: r.ShopMalls,
	}
	logrus.Debugf("%s receive resp->%s", u.DID, pbName(r))
	return ValidateCode(r.Err, 0)
}

func WrapHeartbeat(_ *protocol.User) proto.Message {
	return &cmd.HeartBeatReq{}
}

func HeartBeatS2C(u *protocol.User, msg []byte) error {
	r := &cmd.HeartBeatS2C{}
	err := proto.Unmarshal(msg, r)
	if err != nil {
		return nil
	}
	u.NowMs = r.Ts // 校准时间
	logrus.Debugf("%s receive resp->%s", u.DID, pbName(r))
	return nil
}

func WrapReconnectLogin(u *protocol.User) proto.Message {
	return &cmd.ReconnectLoginReq{UID: u.Summary.UID, Token: u.Token}
}

func ReconnectLoginResp(u *protocol.User, msg []byte) error {
	r := &cmd.ReconnectLoginResp{}
	_ = proto.Unmarshal(msg, r)
	u.Summary = r.User
	u.Token = r.Token
	u.NowMs = r.TimestampMs
	logrus.Debugf("%s receive resp->%s", u.DID, pbName(r))
	return ValidateCode(r.Err, 0)
}

func InitSceneResp(u *protocol.User, msg []byte) error {
	r := &cmd.InitSceneResp{}
	_ = proto.Unmarshal(msg, r)
	logrus.Debugf("%s receive resp->%s", u.DID, pbName(r))
	return ValidateCode(r.Err, 0)
}

func WrapSceneSwitch(u *protocol.User) proto.Message {
	return &cmd.SceneSwitchReq{UID: u.Summary.UID}
}

func SceneSwitchResp(u *protocol.User, msg []byte) error {
	r := &cmd.SceneSwitchResp{}
	_ = proto.Unmarshal(msg, r)
	u.GameInfo.Account = r.Account
	u.Slots = r.Slots
	u.Activities = r.Activities
	logrus.Debugf("%s receive resp->%s", u.DID, pbName(r))
	return ValidateCode(r.Err, 0)
}

func ValidateCode(code cmd.ErrorCode, target cmd.ErrorCode) error {
	if code == target {
		return nil
	}
	return fmt.Errorf("状态码异常：%d %s 目标状态：%d %s",
		code, code.String(), target, target.String())
}
