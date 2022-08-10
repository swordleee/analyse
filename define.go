package main

// 历史成交数据
type DayBusiness struct {
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Percent   float64 `json:"percent"` //涨跌幅
	High      float64 `json:"high"`    //最高价
	Low       float64 `json:"low"`     //最低价
	Price     float64 `json:"price"`   //当前价格
	Open      float64 `json:"open"`    //开盘价
	UpDown    float64 `json:"updown"`  //涨跌金额
	Type      string  `json:"type"`
	Symbol    string  `json:"symbol"`
	Status    int     `json:"status"`
	Update    string  `json:"update"`
	YestClose float64 `json:"yestclose"` //昨日收盘价
	Volume    float64 `json:"volume"`    //成交量
	Arrow     string  `json:"arrow"`
	Time      string  `json:"time"`
	Turnover  float64 `json:"turnover"` //成交额
	AskVol1   float64 `json:"askvol1"`
	AskVol2   float64 `json:"askvol2"`
	AskVol3   float64 `json:"askvol3"`
	AskVol4   float64 `json:"askvol4"`
	AskVol5   float64 `json:"askvol5"` //卖5数量
	Ask1      float64 `json:"ask1"`
	Ask2      float64 `json:"ask2"`
	Ask3      float64 `json:"ask3"`
	Ask4      float64 `json:"ask4"`
	Ask5      float64 `json:"ask5"` //卖5价格
	BidVol1   float64 `json:"bidvol1"`
	BidVol2   float64 `json:"bidvol2"`
	BidVol3   float64 `json:"bidvol3"`
	BidVol4   float64 `json:"bidvol4"`
	BidVol5   float64 `json:"bidvol5"` //买5数量
	Bid1      float64 `json:"bid1"`
	Bid2      float64 `json:"bid2"`
	Bid3      float64 `json:"bid3"`
	Bid4      float64 `json:"bid4"`
	Bid5      float64 `json:"bid5"` //买5价格
}

// 历史成交数据
type HistoryBusiness struct {
	Code             string  `json:"股票代码"`
	Date             string  `json:"日期"`
	Name             string  `json:"名称"`
	ClosePrice       float64 `json:"收盘价"`
	MaxPrice         float64 `json:"最高价"`
	MinPrice         float64 `json:"最低价"`
	OpenPrice        float64 `json:"开盘价"`
	Before           float64 `json:"前收盘"`
	ChangeAmount     float64 `json:"涨跌额"`
	ChangeRate       float64 `json:"涨跌幅"`
	TurnoverRate     float64 `json:"换手率"`
	Turnover         float64 `json:"成交量"`
	TurnoverAmount   float64 `json:"成交金额"`
	TotalMarketValue float64 `json:"总市值"`
	CirMarketValue   float64 `json:"流通市值"`
	DealNum          string  `json:"成交笔数"`
}

type HoldStock struct {
	Code      string  `json:"股票代码"`
	Name      string  `json:"名称"`
	Time      string  `json:"建仓时间"`
	TicketNum int     `json:"持仓数量"`
	Cost      float64 `json:"持仓成本"`
	CurPrice  float64 `json:"现价"`
}

type BusinessStock struct {
	Code        string  `json:"股票代码"`
	Name        string  `json:"名称"`
	Time        string  `json:"交易时间"`
	TicketNum   int     `json:"交易数量"`
	TicketPrice float64 `json:"交易价格"`
}
