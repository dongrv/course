package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"httpproto/cmd"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	object := &cmd.FinLog{
		UID:              10001,
		SpecialUser:      0,
		PaidType:         1,
		TotalPay:         99999,
		SeriesId:         6001,
		SlotId:           "N6001A",
		Balance:          1000000000000,
		Win:              90000000,
		Bet:              10000000,
		BalanceUSD:       100000,
		WinUSD:           9,
		BetUSD:           1,
		SpinNum:          1000000,
		IsBroke:          0,
		Reason:           0,
		GamePlay:         1,
		OneDollar:        10000000,
		RegisterTime:     1718860332,
		AvgBet:           1.0,
		CreateTime:       1718870332,
		WorkScope:        0,
		Platform:         1,
		OutRTP:           9.9,
		RTPState:         1,
		MinBet:           1000000,
		TriggerLucky:     1,
		Lv:               100,
		VipLv:            2,
		LastPayTime:      1718870332,
		Spend:            1000,
		RemarkUser:       0,
		RemarkFilter:     0,
		JackpotLv:        50,
		UnlockPlay:       1,
		UnlockMaxJackpot: false,
		ScheduleID:       1,
		IncrGems:         100,
		SubGems:          10,
		IncrGemsUSD:      10,
		SubGemsUSD:       1,
		Gems:             123456789,
		GemsUSD:          100,
		GemsOneDollar:    10000000,
		WinJackpots:      []*cmd.WinJackpot{&cmd.WinJackpot{JackpotID: 1, Cnt: 100}},
		UserGroup:        "[1,2,3,4,5,6]",
		Description:      "",
		ResVer:           "1.0.0",
		MediaSource:      "ironsource",
		LastMinBet:       10000000,
		LastMaxBet:       900000000000,
		RTP200:           100000000,
		RTP1000:          20000000,
		RFMGroup:         "[1,2]",
		BetUnlock:        10000000,
		BetType:          1,
		BetParam:         "0",
		PayLeft:          99999999,
	}
	var count int
	finlogs := &cmd.BatchFinLog{
		FinLogs: make([]*cmd.FinLog, 0, 20),
	}
	for i := 0; i < 20; i++ {
		finlogs.FinLogs = append(finlogs.FinLogs, object)
	}
	bs, err := proto.Marshal(finlogs)
	if err != nil {
		println(err.Error())
		return
	}
	native, _ := json.Marshal(finlogs)
	println("实际大小：", len(bs), "JSON字符串大小：", len(native))
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			for k := 0; k < 60; k++ {
				for j := 0; j < 10; j++ {
					req, err := http.NewRequest(http.MethodPost, "http://192.168.8.114:223/", bytes.NewBuffer(bs))
					if err != nil {
						println(err.Error())
						return
					}
					client := &http.Client{Timeout: 3 * time.Second}
					resp, err := client.Do(req)
					if err != nil {
						println(err.Error())
						return
					}
					defer resp.Body.Close()
					result, err := io.ReadAll(resp.Body)
					if err != nil {
						println(err.Error())
						return
					}
					count++
					fmt.Printf("output: %s\n", result)
				}
				time.Sleep(time.Second)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	println("总消息：", count)
}
