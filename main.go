package main

import (
	"github.com/gookit/goutil/dump"
	"gorm.io/gorm"
	"gorm/constant"
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

	belongs2hasOne := make([]uint32, 0)
	belongs2hasOne = append(belongs2hasOne, constant.MethodBelongs2HasOneFind, constant.MethodBelongs2HasOneCreate, constant.MethodBelongs2HasOneUpdate)
	for _, v := range belongs2hasOne {
		example.Belongs2HasOneExample(db, v)
	}

	hasMany := make([]uint32, 0)
	hasMany = append(hasMany, constant.MethodHasManyFind, constant.MethodHasManyCreate, constant.MethodHasManyUpdate)
	for _, v := range hasMany {
		example.Belongs2HasOneExample(db, v)
	}

	example.Many2ManyExample(db)
	Thank()

}

// Thank 感谢
func Thank() {
	dump.P("------------ 谢谢聆听 ------------")
}
