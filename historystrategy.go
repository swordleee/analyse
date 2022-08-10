package main

type HistoryStraMgr struct {
}

var hisstramgrsingleton *HistoryStraMgr = nil

func GetHistoryStraMgr() *HistoryStraMgr {
	if hisstramgrsingleton == nil {
		hisstramgrsingleton = new(HistoryStraMgr)
	}
	return hisstramgrsingleton
}

func (dsm *HistoryStraMgr) RunStrategy() {
	GetNetCaseMgr().GetHistoryBusiness("0603986", "20220801", "")
	GetCsvMgr().LoadCsv()
}
