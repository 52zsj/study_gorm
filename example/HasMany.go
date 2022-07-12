package example

import (
	"github.com/gookit/goutil/dump"
	"gorm.io/gorm"
	"gorm/constant"
	"gorm/entity"
)

func HasManyExample(db *gorm.DB, method uint32) {
	switch method {
	case constant.MethodHasManyFind:
		hasManyFindExample(db, true)
	case constant.MethodHasManyUpdate:
		hasManyUpdate(db)
	case constant.MethodHasManyCreate:
		hasManyCreate(db)

	}
}

// hasManyFindExample 关联查询
func hasManyFindExample(db *gorm.DB, showDump bool) entity.User {
	if showDump {
		dump.P("------------ 接下来进入的是 HasMany 介绍 ------------")
	}

	//关联查询
	var user, user1, user2 entity.User

	db.Model(&user).Preload("UserRole.Role").First(&user1)

	db.Model(&user).Preload("UserRole").First(&user2)

	dump.P(user1)

	dump.P(user2)

	return user1
}

// hasManyUpdate 关联更新
func hasManyUpdate(db *gorm.DB) {
	user := hasManyFindExample(db, false)
	//更新
	user.UserRole = []entity.UserRole{
		{
			Id:     user.UserRole[0].Id,
			UserId: user.Id,
			RoleId: 3,
		},
	}
	//
	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&user)
}

// hasManyCreate 关联创建
func hasManyCreate(db *gorm.DB) {
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
		},
	}
	db.Create(&createUserData)
}
