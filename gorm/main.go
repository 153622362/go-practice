package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
	Disscount uint
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(192.168.199.101:3306)/go_practice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 5) // 查询id为1的product
	//db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	//// 更新 - 更新product的price为2000
	//db.Model(&product).Update("Price", 2000)
	//fmt.Println(product.Price)
	// 删除 - 删除product
	if product.ID > 0 {
		db.Delete(&product)
	}

}
