package example

import (
	"github.com/gookit/goutil/dump"
	"gorm.io/gorm"
	"gorm/constant"
	"gorm/entity"
	"time"
)

// Belongs2HasOneExample 关联查询 预加载
func Belongs2HasOneExample(db *gorm.DB, method uint32) {
	switch method {
	case constant.MethodBelongs2HasOneFind:
		belongs2HasOneFindExample(db, true)
	case constant.MethodBelongs2HasOneCreate:
		belong2HasoneCreate(db)
	case constant.MethodBelongs2HasOneUpdate:
		belong2HasoneUpdate(db)
	}

}

// belongs2HasOneFindExample 关联查询
func belongs2HasOneFindExample(db *gorm.DB, showDump bool) entity.User {
	if showDump {
		dump.P("belongs to  和 has one 非常类似 从记录上来说都是指的1对1,\n has one 是指正向联系 比如 user表存在一个扩展表 user_other 这个就是1对1关系\n belongs to 是指反向联系 还是上面的例子 就是 user_other是属于user的\n 简单来说的话就是 我和你1对1  你属于我  我也属于你这种.只是逻辑上的偏差")
		dump.P("------------ 关联查询 介绍 ------------")
	}
	var user, user1, user2 entity.User

	db.Model(&user).Preload("UserOther").Find(&user1) //使用 Preload 进行关联查询
	dump.P(user1)

	time.Sleep(time.Second * 3)
	db.Model(&user).Joins("UserOther").Find(&user2) //使用 Joins 进行关联查询

	dump.P(user2)
	if showDump {
		dump.P("Join Preload 适用于一对一的关系，例如： has one, belongs to")
	}
	return user1

}

// belong2HasoneCreate 关联创建
func belong2HasoneCreate(db *gorm.DB) {
	dump.P("------------ 关联创建 介绍 ------------")

	insertUser := entity.User{
		Username: "李白",
		Nickname: "李白",
		Password: "1234567",
		UserOther: entity.UserOther{
			OtherInfo: "用户扩展信息",
		},
	}
	_ = insertUser
	db.Create(&insertUser)

}

// belong2HasoneUpdate 关联更新
func belong2HasoneUpdate(db *gorm.DB) {
	insertUser := belongs2HasOneFindExample(db, false)
	dump.P("------------ 关联更新 介绍 ------------")
	dump.P("通常的操作是查找出数据后在原始的数据上做数据修改后在进行更新,因此此处踩了1个坑,\n就是userOther里面的ID还是必须要存在的!不然的话就是会一直插入数据")
	updateUser := entity.User{
		Id:       insertUser.Id,
		Username: "杜甫",
		Nickname: "我被修改了!!!酷哇伊",
		Password: "7654321",
		UserOther: entity.UserOther{
			Id:        insertUser.UserOther.Id,
			OtherInfo: "杜甫很忙",
		},
	}
	_ = updateUser
	//只更新主表信息
	//db.Updates(&updateUser)

	//关联更新
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updateUser)
}
