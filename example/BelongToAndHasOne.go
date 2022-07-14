package example

import (
	"github.com/gookit/goutil/dump"
	"gorm.io/gorm"
	"gorm/entity"
)

// Belongs2HasOneFind 关联查询
func Belongs2HasOneFind(db *gorm.DB, showDump bool) entity.User {

	if showDump {
		dump.P("belongsTo 和hasOne 非常类似,从记录上来说都是指的1对1, hasOne是指正向联系比如:user表存在一个扩展表user_other这个就是1对1关系")
		dump.P("belongsTo是指反向联系还是上面的例子:user_other属于user,简单概括就是 我和你 1对1对应,你属于我,我也属于你.只是逻辑上的偏差")
		dump.P("------------ 关联查询 介绍 ------------")
	}
	var user, user1, user2 entity.User
	db.Model(&user).Joins("UserOther").Preload("UserRole.Role").Last(&user1) //使用 Preload + Joins 进行关联查询
	dump.P(user1)
	//已知 Joins 只适用于1对1关联的场景 可自行验证
	db.Model(&user).Joins("UserOther").Last(&user2) //使用 Joins 进行关联查询

	dump.P(user2)
	return user1

}

// Belong2HasoneCreate 关联创建
func Belong2HasoneCreate(db *gorm.DB) {

	dump.P("------------ 关联创建 介绍 ------------")
	insertUser := entity.User{
		Username: "张三",
		Nickname: "张三的昵称",
		Password: "1234567",
		UserOther: entity.UserOther{
			OtherInfo: "张三的扩展信息->关联创建",
		},
	}
	_ = insertUser
	db.Create(&insertUser)
}

// Belong2HasoneUpdate 关联更新
func Belong2HasoneUpdate(db *gorm.DB) {

	insertUser := Belongs2HasOneFind(db, false)

	dump.P("------------ 关联更新 介绍 ------------")
	dump.P("通常的操作是查找出数据后在原始的数据上做数据修改后在进行更新,因此此处踩了1个坑,\n就是userOther里面的ID还是必须要存在的!不然的话就是会一直插入数据")
	updateUser := entity.User{
		Id:       insertUser.Id,
		Username: "52zsj",
		Nickname: "我被修改了!!!酷哇伊",
		Password: "7654321",
		UserOther: entity.UserOther{
			Id:        insertUser.UserOther.Id,
			OtherInfo: "52zsj",
		},
	}
	_ = updateUser
	//只更新主表信息
	//db.Updates(&updateUser)

	//关联更新
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updateUser)
}
