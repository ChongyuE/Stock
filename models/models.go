package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)

type Stockcode struct {
	Id        int    `orm:"column(Id);auto"`
	StockCode string `orm:"column(StockCode);null"`
	StockName string `orm:"column(StockName);null"`
	Mar       string `orm:"column(Mar);null"`
}
type Sinastock struct {
	Id             int    `orm:"column(Id);auto"`
	StockCode      string `orm:"column(StockCode);null"`
	TodayBegin     string `orm:"column(TodayBegin);null"`
	YesterDayEnd   string `orm:"column(YesterDayEnd);null"`
	Now            string `orm:"column(Now);null"`
	TodayHeight    string `orm:"column(TodayHeight);null"`
	TodayLow       string `orm:"column(TodayLow);null"`
	BiddingFrist   string `orm:"column(BiddingFrist);null"`
	BiddingSec     string `orm:"column(BiddingSec);null"`
	TurnoverCount  string `orm:"column(TurnoverCount);null"`
	TurnoverMoney  string `orm:"column(TurnoverMoney);null"`
	Buy1Price      string `orm:"column(Buy1Price);null"`
	Buy1Count      string `orm:"column(Buy1Count);null"`
	Buy2Price      string `orm:"column(Buy2Price);null"`
	Buy2Count      string `orm:"column(Buy2Count);null"`
	Buy3Price      string `orm:"column(Buy3Price);null"`
	Buy3Count      string `orm:"column(Buy3Count);null"`
	Buy4Price      string `orm:"column(Buy4Price);null"`
	Buy4Count      string `orm:"column(Buy4Count);null"`
	Buy5Price      string `orm:"column(Buy5Price);null"`
	Buy5Count      string `orm:"column(Buy5Count);null"`
	Sell1Price     string `orm:"column(Sell1Price);null"`
	Sell1Count     string `orm:"column(Sell1Count);null"`
	Sell2Price     string `orm:"column(Sell2Price);null"`
	Sell2Count     string `orm:"column(Sell2Count);null"`
	Sell3Price     string `orm:"column(Sell3Price);null"`
	Sell3Count     string `orm:"column(Sell3Count);null"`
	Sell4Price     string `orm:"column(Sell4Price);null"`
	Sell4Count     string `orm:"column(Sell4Count);null"`
	Sell5Price     string `orm:"column(Sell5Price);null"`
	Sell5Count     string `orm:"column(Sell5Count);null"`
	Date           string `orm:"column(Date);null"`
	Time           string `orm:"column(Time);null"`
	SystemDateTime string `orm:"column(SystemDateTime);null"`
}
type States struct {
	Id       int       `orm:"column(Id);auto"`
	QueryLen int       `orm:"column(QueryLen);null"`
	Insert   int       `orm:"column(Insert);null"`
	Read     int       `orm:"column(Read);null"`
	SysTime  time.Time `orm:"column(SysTime);null"`
}

func init() {
	url := "root_cy:asdfasdf@tcp(rdsnh073m7tt922r6a8ypublic.mysql.rds.aliyuncs.com:3306)/stock?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", url)
	orm.RegisterModel(new(Stockcode), new(Sinastock), new(States))
	orm.Debug = false
}

func InsertStates(s States) bool {
	s.SysTime = time.Now()
	o := orm.NewOrm()
	_, err := o.Insert(&s)
	if err == nil {
		return true
	} else {
		return false
	}
}
func InsertStockCode(sc Stockcode) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&sc)
	if err == nil {
		return true
	} else {
		return false
	}
}
func GetEachStockCode() (sc []Stockcode) {
	o := orm.NewOrm()
	_, err := o.Raw("select * from stockcode").QueryRows(&sc)
	if err == nil {
		//fmt.Println("user nums: ", num)
	} else {
		fmt.Println("GetEachStockCode Error")
	}
	return
}
func InsertSinaStock(sc Sinastock) bool {
	o := orm.NewOrm()
	_, err := o.Insert(&sc)
	if err == nil {
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
