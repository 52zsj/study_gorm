package example

import (
	"github.com/gookit/goutil/dump"
	"gorm.io/gorm"
	"gorm/entity"
)

// Many2Many 多对多
func Many2ManyExample(db *gorm.DB) {
	var m2mUser entity.M2MUser
	var m2mRole entity.M2MRole

	dump.P("------------ 接下来进入的是 Many2Many 介绍 ------------")
	dump.P("------------ 通过User 找Role ------------")
	db.Model(&m2mUser).Where("id=2").Preload("RoleList").First(&m2mUser)
	dump.P(m2mUser)

	dump.P("------------ 通过Role 找User ------------")
	db.Model(&m2mRole).Where("id=3").Preload("UserList").First(&m2mRole)
	dump.P(m2mRole)
}

// Many2manyCreate 多对多创建
func Many2manyCreate(db *gorm.DB) {

	m2mUser := entity.M2MUser{
		Username: "李白乘舟将欲行",
		Nickname: "忽闻岸上踏歌声",
		Password: "12345678",
		RoleList: []entity.Role{
			{
				Name: "我是多对多关联创建的",
			},
		},
	}
	db.Create(&m2mUser)
}
