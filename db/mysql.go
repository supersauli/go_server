package main

import (
	"database/sql"

	_ "github.com/mysql"
)

//dbinfo
type DBInfo struct {
	mDB string
}

type DBManage struct {
	mDBPrt *sql.DB
}

func (dbm *DBManage) QueryData() {
	//rows,err := dbm.
}

func main() {
	dbInfo := DBInfo{
		mDB: "root:password@tcp(127.0.0.1:3306)/test",
	}
	var dbm DBManage
	var err error
	dbm.mDBPrt, err = sql.Open("mysql", dbInfo.mDB)
	if err != nil {
		panic(err)
		return
	}
	defer dbm.mDBPrt.Close()

}
