package code

import (
	"log"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

/*
Sqlx入门
题目1：使用SQL扩展库进行查询

	假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。

要求 ：

	编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

题目2：实现类型安全映射

		假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	CREATE TABLE `books` (
	  `id` bigint(20) NOT NULL AUTO_INCREMENT,
	  `title` longtext,
	  `author` longtext,
	  `price` double DEFAULT NULL,
	  PRIMARY KEY (`id`)
	) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4

要求 ：

	定义一个 Book 结构体，包含与 books 表对应的字段。
	编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

var departmentMap = map[int]string{
	1:  "技术部",
	2:  "人事部",
	3:  "财务部",
	4:  "市场部",
	5:  "开发部",
	6:  "行政部",
	7:  "法务部",
	8:  "采购部",
	9:  "运营部",
	10: "客服部",
}

func CreateEmployInfo(db *gorm.DB) int64 {
	var employees []Employee
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		var employee Employee = Employee{Name: nameMap[rand.Intn(9)], Department: departmentMap[rand.Intn(9)+1], Salary: rand.Float64()*10000 + 5000}
		employees = append(employees, employee)
	}
	result := db.Create(employees)

	if result.Error != nil {
		log.Printf("数据库操作失败: %v", result.Error)
		return -1
	}
	return result.RowsAffected // returns inserted records count
}

func QueryDepartmentEmplInfo(dbsqlx *sqlx.DB, department int) []Employee {
	var employees []Employee
	dbsqlx.Select(&employees, "select * from employees where department = ?", departmentMap[department])
	return employees
}

func QueryMaxSalaryEmployee(dbsqlx *sqlx.DB) Employee {
	var employee Employee
	dbsqlx.Get(&employee, "select * from employees ORDER BY salary DESC LIMIT 1")
	return employee
}

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func QueryExpensiveBooks(dbsqlx *sqlx.DB, price float64) []Book {
	var books []Book
	dbsqlx.Select(&books, "SELECT * FROM books WHERE price > ?", price)
	return books
}
