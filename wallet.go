package main

import (
	"math"
	"sync"
	"time"
)

const (
	TotalCapital = 1000000
	HandNum      = 100 //一手的数量
)

type WalletMgr struct {
	Total        float64 `json:"total"`        //总金额
	Usable       float64 `json:"usable"`       //可用
	Profit       float64 `json:"profit"`       //盈亏
	HoldValue    float64 `json:"profit"`       //持有市值
	HoldInfo     string  `json:"holdinfo"`     //持仓信息
	BusinessInfo string  `json:"businessinfo"` //交易信息

	holdInfo     map[string]*HoldStock
	businessInfo []*BusinessStock
	locker       *sync.RWMutex
}

var walletmgrsingleton *WalletMgr = nil

func GetWalletMgr() *WalletMgr {
	if walletmgrsingleton == nil {
		walletmgrsingleton = new(WalletMgr)
		walletmgrsingleton.Total = TotalCapital
		walletmgrsingleton.Usable = TotalCapital
		walletmgrsingleton.Profit = 0
		walletmgrsingleton.HoldValue = 0
		walletmgrsingleton.holdInfo = make(map[string]*HoldStock)
		walletmgrsingleton.businessInfo = make([]*BusinessStock, 0)
		walletmgrsingleton.locker = new(sync.RWMutex)
	}
	return walletmgrsingleton
}

func (wal *WalletMgr) AddUsable(value float64) {
	wal.locker.Lock()
	defer wal.locker.Unlock()
	if value < 0 && math.Abs(value) > wal.Usable {
		return
	}
	wal.Usable += value
}

func (wal *WalletMgr) AddBusinessInfo(dayBusiness *DayBusiness, num int) {
	StockInfo := &BusinessStock{
		dayBusiness.Code,
		dayBusiness.Name,
		time.Now().Format("2006-01-02 15:04:05"),
		num,
		dayBusiness.Price}
	wal.businessInfo = append(wal.businessInfo, StockInfo)
}

func (wal *WalletMgr) AddHoldInfo(dayBusiness *DayBusiness, num int) {
	holdInfo, ok := wal.holdInfo[dayBusiness.Code]
	if !ok {
		StockInfo := &HoldStock{
			dayBusiness.Code,
			dayBusiness.Name,
			time.Now().Format("2006-01-02 15:04:05"),
			num,
			dayBusiness.Price,
			dayBusiness.Price}
		wal.holdInfo[dayBusiness.Code] = StockInfo
		return
	}

	holdInfo.Cost = (holdInfo.Cost*float64(holdInfo.TicketNum) + dayBusiness.Price*float64(num)) / float64(holdInfo.TicketNum+num)
	holdInfo.TicketNum += num
}

func (wal *WalletMgr) BuyStock(dayBusiness *DayBusiness, num int, price float64) {
	wal.locker.Lock()
	defer wal.locker.Unlock()
	if num < 100 {
		return
	}
	actNum := num / 100 * 100
	totalCost := float64(actNum) * price
	if totalCost > wal.Usable {
		return
	}
	//减去可用
	wal.Usable -= totalCost
	//记录交易信息
	wal.AddBusinessInfo(dayBusiness, actNum)
	//生成持仓信息
	wal.AddHoldInfo(dayBusiness, actNum)
}
