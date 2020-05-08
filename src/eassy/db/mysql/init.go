package mysql

import (
	"game_framework/src/eassy/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"strconv"
)

var dbRead *gorm.DB
var dbWrite *gorm.DB
var ConnStrW string
var ConnStrR string

func init() {

}

func Connect(dbName string) (*gorm.DB, *gorm.DB) {
	myConfW := conf.Mysql.Master["m1"]
	myConfR := conf.Mysql.Slaver["s1"]
	ConnStrW = myConfW.User + ":" + myConfW.Pass + "@tcp(" + myConfW.Host + ":" + strconv.Itoa(myConfW.Port) + ")"
	ConnStrR = myConfR.User + ":" + myConfR.Pass + "@tcp(" + myConfR.Host + ":" + strconv.Itoa(myConfR.Port) + ")"
	var err error
	dbRead, err = gorm.Open("mysql", ConnStrR+"/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	dbWrite, err = gorm.Open("mysql", ConnStrW+"/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	dbRead.DB().SetMaxIdleConns(10)
	dbRead.DB().SetMaxOpenConns(100)
	dbRead.SingularTable(true)
	dbRead.LogMode(true)

	dbWrite.DB().SetMaxIdleConns(10)
	dbWrite.DB().SetMaxOpenConns(100)
	dbWrite.SingularTable(true)
	dbWrite.LogMode(true)
	return dbRead, dbWrite
}
