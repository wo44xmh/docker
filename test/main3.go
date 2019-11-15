package main

import (
	"fmt"
	"sync"
	"time"
	"strconv"
)

func main() {
	hStr := time.Now().Local().Format("2006010215:04:05")
	fmt.Println(hStr[0:10])
	var str string = "N/A"
	val, _ := strconv.Atoi(str)
	fmt.Println(val)
	mutex := sync.RWMutex{}
	//mutex2 := sync.RWMutex{}
	var test_chan chan string = make(chan string)
	/*test_chan<-1*/
	go func() {
		for test := range test_chan {
			fmt.Println("test_chan:" + test)
		}
		fmt.Println("test_chan:" + <-test_chan)
		/*for test := range test_chan {
			fmt.Println(test)
		}*/
		mutex.Lock()
		/*mutex.RLock()
		mutex.RLock()*/
		/*mutex2.Lock()*/
		fmt.Println("test lock")
		mutex.Unlock()
	}()
	test_chan <- "1"
	close(test_chan)
	time.Sleep(2 * time.Second)
	currentTime := time.Now().Local()
	currentTimeFormat := currentTime.Format("2006-01-02 15:04:05")
	currentTimeFormat = currentTimeFormat[0:10]
	startDay := currentTimeFormat + " 00:00:00"
	fmt.Println(startDay)
	location, _ := time.LoadLocation("Local") //Asia/Chongqing"
	startDayUnixTime, _ := time.ParseInLocation("2006-01-02 15:04:05", startDay, location)
	fmt.Println(startDayUnixTime.Unix())
	endDay := currentTime.Format("2006-01-02") + " 23:59:59"
	endDayUnixTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endDay, location)
	fmt.Println(endDayUnixTime.Unix())
}
