package checksvn

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// SvnAuth SVN授权信息
type SvnAuth struct {
	Host     string // svn://xxx.com/svn
	Username string // 用户名
	Password string // 访问密码
}

func NewSvnAuth(host string, username string, password string) *SvnAuth {
	return &SvnAuth{
		Host:     host,
		Username: username,
		Password: password,
	}
}

// 声明和注册SVN命令

type Cmd string

func (c Cmd) String() string {
	return string(c)
}

const splitLine = `------------------------------------------------------------------------` // 分割线

var (
	/*
		输出格式：
			------------------------------------------------------------------------
			r13475 | developer | 2023-09-21 13:34:34 +0800 (周四, 21 9月 2023) | 1 line

			[m] fixed bug
			------------------------------------------------------------------------
	*/
	ShowLog Cmd = "svn log %s -l 1 --username %s --password %s" // 获取日志最近一行

)

// SvnCMDMap 命令集合
var SvnCMDMap = map[Cmd]func(*string) []Row{
	ShowLog: ShowLogFunc,
}

// Row 行结构
type Row struct {
	Revision  string // 当前版本
	Author    string // 作者
	DateTime  string // 日期时间
	Timestamp int64  // 秒级时间戳
	Comment   string // 注释
}

func ShowLogFunc(str *string) []Row {
	var splitTag string // 切割标识
	msg := *str
	if strings.Contains(msg, "lines") {
		splitTag = "lines"
	} else {
		splitTag = "line"
	}
	revisions := strings.Split(msg, splitTag)
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
