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
	for i := 0; i < 10; i++ {
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

func query1stAnd8thColumn(db *gorm.DB) {
	var result []TableTemplate
	start := time.Now()
	db.Select("column1, column8").Find(&result)
	end := time.Now()
	fmt.Println("query 1st and 8th column elapsed: ", end.Sub(start))
}

func query1stAnd9thColumn(db *gorm.DB) {
	var result []TableTemplate
	start := time.Now()
	db.Select("column1, column9").Find(&result)
	end := time.Now()
	fmt.Println("query 1st and 9th column elapsed: ", end.Sub(start))
}

type TableTemplate struct {
	ID      int `gorm:"primary_key"`
	Column1 int
	Column2 int
	Column3 int
	Column4 int
	Column5 int
	Column6 int
	Column7 int
	Column8 int
	Column9 int
}

func main() {
	port1 := "3307"
	dsn1 := "root:1234@tcp(127.0.0.1:" + port1 + ")/testlocality?charset=utf8mb4&parseTime=True&loc=Local"
	db1 := connectDB(dsn1)
	createTable(db1)
	insertDataInTable(db1)
	query1stAnd8thColumn(db1)

	port2 := "3308"
	dsn2 := "root:1234@tcp(127.0.0.1:" + port2 + ")/testlocality?charset=utf8mb4&parseTime=True&loc=Local"
	db2 := connectDB(dsn2)
	createTable(db2)
	insertDataInTable(db2)
	query1stAnd9thColumn(db2)

}
