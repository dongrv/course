package dialog

import (
	"casino/protocol/cmd"
	"context"
	"github.com/sirupsen/logrus"
	"net"
	"testing"
	"time"
)

var protoFile = "./ProtoConfig.json"

func TestRobot_OnProcess(t *testing.T) {
	ln, err := net.Dial(`tcp`, `:2001`)
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	d := NewDialog()

	go d.OnProcess(nil, func(msgId int32, msg []byte) {
		logrus.Info(`receive->`, msgId)
	})

	go func() {
		time.AfterFunc(2*time.Second, func() {
			req := &cmd.LoginReq{DID: "Tony", Platform: 1, LoginType: 1}
			d.OnSend(req, 3001)
			logrus.Info(`send->`, req)
		})
	}()

	d.Start(context.Background(), ln)
}
