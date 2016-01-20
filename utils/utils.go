package utils

import (
	"fmt"
	// "strings"
	// "io"
	"io/ioutil"
	"net/http"
	// "os"
	"stock2.0/models"
	// "strconv"
	// "github.com/axgle/mahonia"
	// "container/list"
	"os"
	// "runtime"
	"strings"
	"time"
)

func HttpGet(url string) (body string) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	resp, err := client.Do(reqest)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	body = string(b)
	return
}
func Convert(str string) models.Sinastock {
	var s models.Sinastock
	strl := strings.Split(str, `"`)
	if strl[1] == "" {
		return s
	}
	strlsj := strings.Split(strl[1], `,`)
	stockcodel := strings.Split(strl[0], `_`)
	s.StockCode = stockcodel[2][:8]
	s.TodayBegin = strlsj[1]
	s.YesterDayEnd = strlsj[2]
	s.Now = strlsj[3]
	s.TodayHeight = strlsj[4]
	s.TodayLow = strlsj[5]
	s.BiddingFrist = strlsj[6]
	s.BiddingSec = strlsj[7]
	s.TurnoverCount = strlsj[8]
	s.TurnoverMoney = strlsj[9]
	s.Buy1Count = strlsj[10]
	s.Buy1Price = strlsj[11]
	s.Buy2Count = strlsj[12]
	s.Buy2Price = strlsj[13]
	s.Buy3Count = strlsj[14]
	s.Buy3Price = strlsj[15]
	s.Buy4Count = strlsj[16]
	s.Buy4Price = strlsj[17]
	s.Buy5Count = strlsj[18]
	s.Buy5Price = strlsj[19]
	s.Sell1Count = strlsj[20]
	s.Sell1Price = strlsj[21]
	s.Sell2Count = strlsj[22]
	s.Sell2Price = strlsj[23]
	s.Sell3Count = strlsj[24]
	s.Sell3Price = strlsj[25]
	s.Sell4Count = strlsj[26]
	s.Sell4Price = strlsj[27]
	s.Sell5Count = strlsj[28]
	s.Sell5Price = strlsj[29]
	s.Date = strlsj[30]
	s.Time = strlsj[31]
	s.SystemDateTime = time.Now().Format("2006-01-02T15:04:05.999999999Z07:00")
	return s
}
