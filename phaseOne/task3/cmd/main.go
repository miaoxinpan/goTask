package main

import (
	"fmt"
	"gotask/phaseOne/task3/code"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gormproject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&code.Student{})
	db.AutoMigrate(&code.Transaction{})
	db.AutoMigrate(&code.Account{})
	// fmt.Println("添加了%d条学生信息\n", code.CreateStuInfo(db))
	var students []code.Student = code.ConditionalQueryStuInfo(db)
	fmt.Println(students)
	var student code.Student = code.UpdateStuInfo(db, "Bob", 6)
	fmt.Println(student)
	// fmt.Println("添加了%d条账户信息\n", code.CreateAccountInfo(db))
	fmt.Println(code.AccounTrading(db, 1, 2, 100))
}
