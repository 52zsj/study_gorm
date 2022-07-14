package main

import (
	"github.com/gookit/goutil/dump"
	"gorm.io/gorm"
	"gorm/example"
	"gorm/initdb"
)

// gorm.Model 的定义
//type Model struct {
//  ID uint `gorm:"primaryKey"`
//  CreatedAt time.Time
//  UpdatedAt time.Time
//  DeletedAt gorm.DeletedAt `gorm:"index"`
//}

//GORM 倾向于约定，而不是配置。
//蛇形命名法是书写复合词或短语的一种形式，用下划线分隔每个单词，也称下划线命名法。
//默认情况下，GORM 使用 ID 作为主键，使用结构体名的 "蛇形复数" 作为表名，举个🌰 结构体:UserInfo -> user_infos
//字段名的 蛇形 作为列名，并使用 CreatedAt、UpdatedAt 字段追踪创建、更新时间
//模型 使用 gorm.DeletedAt 类型作为软删除的标志,字段名可随意

var (
	db *gorm.DB
)

func init() {
	initdb.InitDatabase()
	db = initdb.InitTableAndCreateDb(true)
}

func main() {

	//查找
	example.Belongs2HasOneFind(db, true)
	//创建
	//example.Belong2HasoneCreate(db)
	//更新
	//example.Belong2HasoneUpdate(db)

	//查找
	//example.HasManyFind(db, true)
	//创建
	//example.HasManyCreate(db)
	//更新
	//example.HasManyUpdate(db)

	//创建
	//example.Many2manyCreate(db)
	//查找
	//example.Many2ManyExample(db)

	//Thank()

}

// Thank 感谢
func Thank() {
	dump.P("------------ 谢谢聆听 ------------")
	dump.P("PS 多留意生成的SQL.")
}
