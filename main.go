package main

func main() {
	//日内
	GetDayStraMgr().Run()

	//历史
	GetHistoryStraMgr().RunStrategy()
}
