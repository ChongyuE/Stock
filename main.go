package main

import (
	"fmt"
	// "os"
	// "runtime/pprof"
	"stock2.0/models"
	// "stock2.0/mylist"
	"stock2.0/utils"
	// "strconv"
	// "runtime"
	"strings"
	"sync"
	"time"
)

var l *utils.Mylist

//单秒数据写入数量
var s sync.Mutex
var count int

//单秒读入数据量
var s2 sync.Mutex
var count2 int

//

//
var states models.States

func init() {
	l = utils.NewMylist()
	count = 0
	count2 = 0
}
func main() {
	sinastockurl := `http://hq.sinajs.cn/list=`
	listSCode := models.GetEachStockCode()
	for _, c := range listSCode {
		go fmtRange(sinastockurl, c.Mar+c.StockCode)
	}
	for i := 0; i < 40; i++ {
		go InsertDB()
	}
	for {
		h := time.Now().Hour()
		if h < 16 && h >= 9 {
			s.Lock()
			s2.Lock()
			fmt.Printf(`当前队列长度: %d,单秒数据写入速度 %d 条/s,单秒读取数据 %d 条/s`, l.Len(), count, count2)
			fmt.Println()
			states.QueryLen = l.Len()
			states.Insert = count
			states.Read = count2
			models.InsertStates(states)
			count = 0
			s.Unlock()
			count2 = 0
			s2.Unlock()
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(10 * time.Minute)
		}
	}
}
func fmtRange(url string, s string) {
	var entity models.Sinastock
	var h string
	preTime := ""
	for {
		h = time.Now().Format("15:04:05")
		if (h >= "09:00:00" && (h <= "11:31:00")) || ((h >= "12:45:00") && (h <= "15:01:00")) {
			s = strings.ToLower(s)
			entity = utils.Convert(utils.HttpGet(url + s))
			if preTime == entity.Time {
				continue
			} else {
				l.PushBack(entity)
				s2.Lock()
				count2++
				s2.Unlock()
				time.Sleep(3 * time.Second)
			}
			preTime = entity.Time
		} else {
			time.Sleep(10 * time.Minute)
		}
	}
}

// Insert Use Sync Query
func InsertDB() {
	for {
		if l.Len() == 0 {
			continue
		}
		a := l.GetFrontRemove()
		if a != nil {
			models.InsertSinaStock(a.Value.(models.Sinastock))
			s.Lock()
			count++
			s.Unlock()
		}
	}
}
