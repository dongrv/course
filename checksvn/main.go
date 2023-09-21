package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type CMD string

func (c CMD) String() string {
	return string(c)
}

var (
	/*
		格式：
			------------------------------------------------------------------------
			r13475 | developer | 2023-09-21 13:34:34 +0800 (周四, 21 9月 2023) | 1 line

			[m] 预选bet的代码改到触发活动逻辑下面
			------------------------------------------------------------------------
	*/
	SvnShowLog CMD = "svn log %s -l 1 --username %s --password %s" // 获取日志最近一行

	splitLine = `------------------------------------------------------------------------`
)

// svn 命令集合
var svnCMDs = map[CMD]func(*string) []Row{
	SvnShowLog: parseSvnLog,
}

func main() {
	println(1)
	svnURL := `svn://xxx.com`
	username := ""
	password := ""

	cmd := fmt.Sprintf(SvnShowLog.String(), svnURL, username, password)
	str, err := Command(cmd)
	if err != nil {
		panic(err)
	}
	str = strings.TrimSpace(strings.ReplaceAll(str, splitLine, ""))
	println(2, str)
	results := svnCMDs[SvnShowLog](&str)
	if len(results) == 0 {
		return
	}
	currentLog := results[0] // 取最近行
	if time.Now().Unix()-currentLog.Timestamp <= 5*60 {
		msg := fmt.Sprintf("### 业务服Pre-release Branch有更新！\n >**提交：** %s\n >**注释：** %s\n >**时间：** %s\n\n 请及时处理 <@101>",
			currentLog.Author, currentLog.Comment, currentLog.DateTime)
		println(3, msg)
		notifyGroup(msg)
	}
}

var notifyURL = `https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=8fa40bf3-2f0d-4823-9085-fd4a91813740`

type Markdown struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type Message struct {
	MsgType  string    `json:"msgtype"`
	Markdown *Markdown `json:"markdown"`
}

// 通知群
func notifyGroup(str string) {
	msg := &Message{
		MsgType:  `markdown`,
		Markdown: &Markdown{Content: str, MentionedList: []string{"@118"}},
	}
	client := &http.Client{Timeout: 3 * time.Second}
	defer client.CloseIdleConnections()
	bs, _ := json.Marshal(msg)
	resp, err := client.Post(notifyURL, `application/json`, bytes.NewBuffer(bs))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "推送错误：%s\n", err.Error())
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "推送错误：%s\n", err.Error())
		return
	}
	defer resp.Body.Close()
	_, _ = fmt.Fprintf(os.Stdout, "推送结果：%s\n", body)
}

type Row struct {
	Revision  string // 当前版本
	Author    string // 作者
	DateTime  string // 日期时间
	Timestamp int64  // 秒级时间戳
	Comment   string // 注释
}

func parseSvnLog(str *string) []Row {
	msg := *str
	logs := strings.Split(msg, "line")
	var rows []Row
	for i := 0; i < len(logs)/2; {
		logs[i] = strings.TrimSpace(logs[i])
		meta := TrimSlice(strings.Split(logs[i], "|"))
		date, _, _ := strings.Cut(meta[2], "+0800")
		timestamp, _ := time.ParseInLocation(time.DateTime, strings.TrimSpace(date), time.FixedZone("CST", 28800))
		seconds := timestamp.Unix()
		row := Row{
			Revision:  meta[0],
			Author:    meta[1],
			DateTime:  date,
			Timestamp: seconds,
			Comment:   strings.TrimSpace(logs[i+1]),
		}
		rows = append(rows, row)
		i += 2
	}
	return rows
}

func TrimSlice(strs []string) []string {
	var result []string
	for _, str := range strs {
		result = append(result, strings.TrimSpace(str))
	}
	return result
}

func Command(args ...string) (string, error) {
	name := "/bin/bash"
	c := "-c"
	// 命令兼容系统差异
	if runtime.GOOS == "windows" {
		name = "cmd"
		c = "/C"
	}
	// 组装可执行命令结构
	args = append([]string{c}, args...)
	cmd := exec.Command(name, args...)
	// 创建获取输出命令的管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err = cmd.Start(); err != nil {
		return "", err
	}
	// 读取管道
	bytes, newErr := io.ReadAll(stdout)
	if newErr != nil {
		return "", newErr
	}
	if err = cmd.Wait(); err != nil {
		return "", err
	}
	output, err := simplifiedchinese.GB18030.NewDecoder().String(string(bytes))
	if err != nil {
		return "", err
	}
	return output, nil
}
