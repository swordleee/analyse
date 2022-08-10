package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//api接口
const (
	//日内实时盘口（JSON）
	DayAPI = "http://api.money.126.net/data/feed/%smoney.api"

	//历史成交数据（CSV）：
	HistoryAPI = "http://quotes.money.163.com/service/chddata.html" //?code=0601398&start=20000720&end=20150508

	//财务指标（CSV）：
	FinancialIndex = "http://quotes.money.163.com/service/zycwzb_601398.html?type=report"

	//资产负债表（CSV）：
	Balance = "http://quotes.money.163.com/service/zcfzb_601398.html"

	//利润表（CSV）：
	profit = "http://quotes.money.163.com/service/lrb_601398.html"

	//现金流表（CSV）：
	Cash = "http://quotes.money.163.com/service/xjllb_601398.html"

	//杜邦分析（HTML）：
	Dupont = "http://quotes.money.163.com/f10/dbfx_601398.html"
)

type NetCase struct {
}

var netcase *NetCase = nil

func GetNetCaseMgr() *NetCase {
	if netcase == nil {
		netcase = new(NetCase)
	}

	return netcase
}

func (nc *NetCase) GetDayBusiness(codes []string) (strData string) {
	if len(codes) <= 0 {
		return
	}
	strCodes := ""
	for i := 0; i < len(codes); i++ {
		strCodes += codes[i]
		strCodes += ","
	}

	dayUrl := fmt.Sprintf(DayAPI, strCodes)

	res, err := http.Get(dayUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//log.Println("result=", string(result))
	strResult, err := strconv.Unquote(strings.Replace(strconv.Quote(string(result)), `\\u`, `\u`, -1))
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(strResult)
	//取出json
	start := strings.Index(strResult, "(")
	end := strings.Index(strResult, ")")
	strData = strResult[start+1 : end]
	return
}

func (nc *NetCase) GetHistoryBusiness(code, start, end string) {

	postValue := url.Values{
		"code":  {code},
		"start": {start},
		"end":   {end},
	}

	res, err := http.PostForm(HistoryAPI, postValue)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//gbk to utf8
	data, err := GbkToUtf8(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	//log.Println("result=", string(data))
	//写入文件
	file, error := os.Create(fmt.Sprintf("./csv/%s.csv", code))
	if error != nil {
		fmt.Println(error)
	}
	file.Write(data)
	file.Close()
}
