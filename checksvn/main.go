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
	"strconv"
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

	svnURL := ``
	username := ""
	password := ""

	cmd := fmt.Sprintf(SvnShowLog.String(), svnURL, username, password)
	str, err := Command(cmd)
	if err != nil {
		panic(err)
	}
	str = strings.TrimSpace(strings.ReplaceAll(str, splitLine, ""))
	results := svnCMDs[SvnShowLog](&str)
	if len(results) == 0 {
		return
	}
	currentLog := results[0] // 取最近行
	if time.Now().Unix()-currentLog.Timestamp <= 5*60 {
		msg := fmt.Sprintf("### 业务服Pre-release Branch有更新！\n >**提交：** %s\n >**注释：** %s\n >**时间：** %s\n\n 请及时处理 <@118>",
			currentLog.Author, currentLog.Comment, currentLog.DateTime)
		notifyGroup(msg)
	}
}

var notifyURL = `` // 机器人通知

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
	revisions := strings.Split(msg, "line")
	var (
		rows        []Row
		offsetHours int64
	)
	for i := 0; i < len(revisions)/2; {
		revisions[i] = strings.TrimSpace(revisions[i])
		meta := TrimSlice(strings.Split(revisions[i], "|"))
		times := strings.Split(meta[2], " ")
		timeZone := times[2] // +0800
		v := timeZone[1:]
		iv, _ := strconv.Atoi(v)
		iv /= 100
		if timeZone[0] == '-' {
			offsetHours = (int64(iv) + 8) * 3600 // +8对标东八区时间
		} else {
			var diff int64
			if int64(iv)-8 >= 0 {
				diff = -(int64(iv) - 8) // 东八区之前的时间减
			} else {
				diff = 8 - int64(iv) // 东八区之后的时间加
			}
			offsetHours = diff * 3600 // 和东八区时间相差时间
		}

		var cutSep string // 剪裁标识
		if iv <= 9 {
			cutSep = fmt.Sprintf("%c0%d", timeZone[0], iv*100)
		} else {
			cutSep = fmt.Sprintf("%c%d", timeZone[0], iv*100)
		}

		date, _, _ := strings.Cut(meta[2], cutSep)
		timestamp, err := time.ParseInLocation(time.DateTime, strings.TrimSpace(date), time.FixedZone("CST", 28800))
		if err != nil {
			fmt.Printf("err %v\n", err)
		}
		seconds := timestamp.Unix() + offsetHours
		date = time.Unix(seconds, 0).In(time.FixedZone("CST", 28800)).Format(time.DateTime) // 校准后的日期
		row := Row{
			Revision:  meta[0],
			Author:    meta[1],
			DateTime:  date,
			Timestamp: seconds,
			Comment:   strings.TrimSpace(revisions[i+1]),
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
	if runtime.GOOS == "windows" {
		output, err := simplifiedchinese.GB18030.NewDecoder().String(string(bytes))
		if err != nil {
			return "", err
		}
		return output, nil
	}
	return string(bytes), nil
}
