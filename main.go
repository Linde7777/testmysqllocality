package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"unsafe"
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
	var i int
	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(i))
	dbName := "testlocality1"
	dsn := "root:1234@tcp(127.0.0.1:9999)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db := connectDB(dsn)
	createTable(db)
	insertDataInTable(db)
	query1stAnd8thColumn(db)
}
