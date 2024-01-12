package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func connectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func createTable(db *gorm.DB) {
	err := db.AutoMigrate(&TableTemplate{})
	if err != nil {
		panic("fail to migrate table: " + err.Error())
	}
}

func insertDataInTable(db *gorm.DB) {
	fmt.Println("inserting data...")
	for i := 0; i < 1000; i++ {
		db.Create(&TableTemplate{
			Column1: i,
			Column2: i,
			Column3: i,
			Column4: i,
			Column5: i,
			Column6: i,
			Column7: i,
			Column8: i,
			Column9: i,
		})
	}
}

func queryColumns(db *gorm.DB, columnName1, columnName2 string) {
	var result []TableTemplate
	avgElapsedTime := 0.0
	repeatTimes := 1000
	for i := 0; i < repeatTimes; i++ {
		start := time.Now()
		db.Select("column1, column8").Find(&result)
		end := time.Now()
		avgElapsedTime += end.Sub(start).Seconds()
	}
	fmt.Println("query "+columnName1+" and "+columnName2+" average elapsed time: ",
		avgElapsedTime/float64(repeatTimes))
}

type TableTemplate struct {
	ID       int `gorm:"primary_key"`
	Column1  int
	Column2  int
	Column3  int
	Column4  int
	Column5  int
	Column6  int
	Column7  int
	Column8  int
	Column9  int
	Column10 int
}

func main() {
	port1 := "3307"
	dsn1 := "root:1234@tcp(127.0.0.1:" + port1 + ")/testlocality?charset=utf8mb4&parseTime=True&loc=Local"
	db1 := connectDB(dsn1)
	createTable(db1)
	insertDataInTable(db1)
	queryColumns(db1, "column1", "column8")

	port2 := "3308"
	dsn2 := "root:1234@tcp(127.0.0.1:" + port2 + ")/testlocality?charset=utf8mb4&parseTime=True&loc=Local"
	db2 := connectDB(dsn2)
	createTable(db2)
	insertDataInTable(db2)
	queryColumns(db2, "column1", "column9")
}
