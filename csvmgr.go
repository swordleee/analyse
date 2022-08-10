package main

import (
	"strconv"
	"strings"
	"utils"
)

type CsvMgr struct {
	HistoryBusiness map[int][]*HistoryBusiness
}

var csvmgrsingleton *CsvMgr = nil

func GetCsvMgr() *CsvMgr {
	if csvmgrsingleton == nil {
		csvmgrsingleton = new(CsvMgr)
		csvmgrsingleton.InitStruct()
	}

	return csvmgrsingleton
}

func (cm *CsvMgr) InitStruct() {
	cm.HistoryBusiness = make(map[int][]*HistoryBusiness)
}

func (cm *CsvMgr) LoadCsv() {
	cm.LoadHistoryBusiness()
}

func (cm *CsvMgr) LoadHistoryBusiness() {
	historyBusiness := make([]*HistoryBusiness, 0)
	utils.GetCsvUtilMgr().LoadCsv("0603986", &historyBusiness)
	for _, v := range historyBusiness {
		//str to num
		codeStr := strings.Replace(v.Code, "'", "", -1)
		codeNo, err := strconv.Atoi(codeStr)
		if err != nil {
			continue
		}

		_, ok := cm.HistoryBusiness[codeNo]
		if !ok {
			cm.HistoryBusiness[codeNo] = make([]*HistoryBusiness, 0)
		}
		cm.HistoryBusiness[codeNo] = append(cm.HistoryBusiness[codeNo], v)
	}
}
