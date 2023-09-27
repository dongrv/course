package checksvn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// ChatGroupBot 微信群机器人
type ChatGroupBot struct {
	Url string
}

func NewChatGroupBot(url string) *ChatGroupBot {
	return &ChatGroupBot{Url: url}
}

// 推送通知

type Markdown struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type Message struct {
	MsgType  string    `json:"msgtype"`
	Markdown *Markdown `json:"markdown"`
}

// NotifyGroup 通知群
func NotifyGroup(str string, bot *ChatGroupBot) error {
	if strings.TrimSpace(str) == "" {
		return errors.New("推送消息地址错误")
	}
	msg := &Message{
		MsgType:  `markdown`,                                               // 默认推送 markdown 消息类型
		Markdown: &Markdown{Content: str, MentionedList: []string{"@118"}}, // 指定提示人
	}
	client := &http.Client{Timeout: 3 * time.Second}
	defer client.CloseIdleConnections()
	bs, _ := json.Marshal(msg)
	resp, err := client.Post(bot.Url, `application/json`, bytes.NewBuffer(bs))
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	StdOut(body, `NotifyGroup`)
	return nil
}

func WrapMessage(row Row, compileOk bool, note string) string {
	if compileOk {
		return fmt.Sprintf("### 业务服Pre-release Branch有更新！\n >**提交：** %s\n >**注释：** %s\n >**时间：** %s \n >**结果：** 编译通过 \n\n 请及时处理 <@118>",
			row.Author, row.Comment, row.DateTime)
	}
	return fmt.Sprintf("### 业务服Pre-release Branch有更新！\n >**编译失败：** %s\n\n 请及时处理 <@118>", note)
}
