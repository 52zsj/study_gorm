package example

import (
	"github.com/gookit/goutil/dump"
	"gorm.io/gorm"
	"gorm/entity"
)

// HasManyFind 关联查询
func HasManyFind(db *gorm.DB, showDump bool) entity.User {
	if showDump {
		dump.P("------------ 接下来进入的是 HasMany 介绍 ------------")
	}
	//关联查询
	var user, user1, user2 entity.User

	db.Model(&user).Preload("UserRole.Role").Last(&user1)

	db.Model(&user).Preload("UserRole").Last(&user2)

	dump.P(user1)

	dump.P(user2)

	return user1
}

// HasManyUpdate 关联更新
func HasManyUpdate(db *gorm.DB) {
	user := HasManyFind(db, false)
	//更新
	user.UserRole = []entity.UserRole{
		{
			Id:     user.UserRole[0].Id,
			UserId: user.Id,
			RoleId: 4,
		},
	}
	//
	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&user)
}

// HasManyCreate 关联创建
func HasManyCreate(db *gorm.DB) {
	createUserData := entity.User{
		Username: "炎帝",
		Nickname: "炎帝",
		Password: "123456",
		UserOther: entity.UserOther{
			OtherInfo: "我是炎帝的扩展信息",
		},
		UserRole: []entity.UserRole{
			{
				RoleId: 2,
			},
			{
				RoleId: 1,
			},
		},
	}
	db.Create(&createUserData)
}
