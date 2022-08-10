package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var Codes = []string{
	"0603986",
	"1300709",
	"1000625"}

const (
	MaxDataNum        = 600 //数据最大条数（10分钟 一秒一条）
	RunStrategyMinNum = 300 //策略执行最小条数
)

type DayStraMgr struct {
	dayBusiness map[string][]*DayBusiness
}

var daystramgrsingleton *DayStraMgr = nil

func GetDayStraMgr() *DayStraMgr {
	if daystramgrsingleton == nil {
		daystramgrsingleton = new(DayStraMgr)
		daystramgrsingleton.dayBusiness = make(map[string][]*DayBusiness)
	}
	return daystramgrsingleton
}

func (dsm *DayStraMgr) AppendBusiness(dayBusiness *DayBusiness) {
	strBusiness, err := json.Marshal(dayBusiness)
	if err != nil {
		return
	}
	filePath := fmt.Sprintf("./daybusiness/%s", dayBusiness.Code)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(string(strBusiness) + "\n")
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func (dsm *DayStraMgr) LoadCodeData() {
	data := GetNetCaseMgr().GetDayBusiness(Codes)
	dayBusiness := make(map[string]*DayBusiness)
	err := json.Unmarshal([]byte(data), &dayBusiness)
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, val := range dayBusiness {
		if _, ok := dsm.dayBusiness[key]; !ok {
			dsm.dayBusiness[key] = make([]*DayBusiness, 0)
		}
		//判断刷新时间
		if len(dsm.dayBusiness[key]) > 0 && dsm.dayBusiness[key][len(dsm.dayBusiness[key])-1].Update == val.Update {
			continue
		}
		dsm.dayBusiness[key] = append(dsm.dayBusiness[key], val)

		if len(dsm.dayBusiness[key]) > MaxDataNum {
			dsm.dayBusiness[key] = dsm.dayBusiness[key][len(dsm.dayBusiness[key])-MaxDataNum:]
		}
		dsm.AppendBusiness(val)
	}

}

func (dsm *DayStraMgr) OpenHandicap() bool {
	now := time.Unix(time.Now().Unix(), 0)
	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday {
		return false
	}
	if now.Hour() < 9 || (now.Hour() == 9 && now.Minute() < 30) {
		return false
	}
	if (now.Hour() == 11 && now.Minute() > 30) || now.Hour() == 12 {
		return false
	}
	if now.Hour() >= 15 {
		return false
	}
	return true
}

func (dsm *DayStraMgr) RunStrategy() bool {
	if !dsm.OpenHandicap() {
		return false
	}

	time := time.Unix(time.Now().Unix(), 0)
	if time.Second()%1 == 0 {
		dsm.LoadCodeData()
	}
	return true
}

// 定时检查事件
func (dsm *DayStraMgr) Run() {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		if !dsm.RunStrategy() {
			break
		}
	}

	ticker.Stop()
}

func (dsm *DayStraMgr) GetUpDownNum(dataList []*DayBusiness) (upNum, downNum int) {

	for i := 0; i < len(dataList); i++ {
		if dataList[i].Arrow == "↑" {
			upNum++
		}
		if dataList[i].Arrow == "↓" {
			downNum++
		}
	}
	return
}

func (dsm *DayStraMgr) DayStrategy1Code(code string) {
	dataList, ok := dsm.dayBusiness[code]
	if !ok {
		return
	}
	if len(dataList) < RunStrategyMinNum {
		return
	}
	upNum, downNum := dsm.GetUpDownNum(dataList)
	//拉升量大于下跌量，且超过50%
	if upNum > downNum && (upNum-downNum)*100/downNum > 50 {

	}
}

//策略1
func (dsm *DayStraMgr) DayStrategy1() {
	for i := 0; i < len(Codes); i++ {
		dsm.DayStrategy1Code(Codes[1])
	}
}
