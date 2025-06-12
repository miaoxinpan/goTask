package code

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

/*
SQL语句练习
题目1：基本CRUD操作

		假设有一个名为 students 的表，包含字段 id （主键，自增）、
		name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
	要求 ：
		编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
		编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
		编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
		编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

题目2：事务语句

		假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	要求 ：
		编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
		在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
		向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
		如果余额不足，则回滚事务。
*/
type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

var gradeMap = map[int]string{
	1: "一年级",
	2: "二年级",
	3: "三年级",
	4: "四年级",
	5: "五年级",
	6: "六年级",
}
var nameMap = map[int]string{
	0: "Alice",
	1: "Bob",
	2: "Charlie",
	3: "David",
	4: "Eve",
	5: "Frank",
	6: "Grace",
	7: "Helen",
	8: "Ivy",
	9: "Jack",
}

// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 随机插入5条 不同年纪不同名字不同年级的数据
func CreateStuInfo(db *gorm.DB) int64 {
	var students []Student
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		var Student Student = Student{Name: nameMap[rand.Intn(9)], Age: rand.Intn(99) + 1, Grade: gradeMap[rand.Intn(5)+1]}
		students = append(students, Student)
	}
	result := db.Create(students)

	if result.Error != nil {
		log.Printf("数据库操作失败: %v", result.Error)
		return -1
	}
	return result.RowsAffected // returns inserted records count
}

// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
func ConditionalQueryStuInfo(db *gorm.DB) []Student {
	var students []Student
	db.Where("age > ?", 18).Find(&students)
	return students
}

// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
func UpdateStuInfo(db *gorm.DB, name string, grade int) Student {
	var student Student
	db.Where("name = ?", name).First(&student)
	student.Grade = gradeMap[grade]
	db.Save(&student)
	return student
}

// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。 age < 15
func ConditionalDeleteStuInfo(db *gorm.DB, field string, condition string, param int) int64 {
	result := db.Where(field+condition+"?", param).Delete(&Student{})
	if result.Error != nil {
		log.Printf("数据库操作失败: %v", result.Error)
		return -1
	}
	return result.RowsAffected
}

type Transaction struct {
	ID              int
	From_account_id int
	To_account_id   int
	Amount          int
}
type Account struct {
	ID      int
	Balance int
}

func CreateAccountInfo(db *gorm.DB) int64 {
	// 先创建账户
	var accounts []Account
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		account := Account{Balance: rand.Intn(1000) + 1} // 随机余额 1~1000
		accounts = append(accounts, account)
	}
	result := db.Create(&accounts)
	if result.Error != nil {
		log.Printf("数据库操作失败: %v", result.Error)
		return -1
	}
	return result.RowsAffected // 返回插入的记录数
}

/*
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
如果余额不足，则回滚事务。
*/
func AccounTrading(db *gorm.DB, accountA int, accountB int, balance int) string {

	err := db.Transaction(func(tx *gorm.DB) error {
		var a Account
		result := tx.Where("id = ? AND balance >= ?", accountA, balance).First(&a)
		if result.Error != nil {
			return result.Error // 没查到直接返回，事务回滚
		}
		//查询b账户
		var b Account
		tx.Where("id = ? ", accountB).First(&b)
		b.Balance = b.Balance + balance
		if err := tx.Save(&b).Error; err != nil {
			return err
		}
		a.Balance = a.Balance - balance
		if err := tx.Save(&a).Error; err != nil {
			return err
		}
		//记录交易
		trans := Transaction{From_account_id: accountA, To_account_id: accountB, Amount: balance}
		if err := tx.Create(&trans).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Sprintf("转账失败: %v", err)
	} else {
		return "转账成功"
	}

}
