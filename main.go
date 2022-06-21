package main

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm/entity"
	"io/ioutil"
	"os"
)

// gorm.Model 的定义
//type Model struct {
//  ID        uint           `gorm:"primaryKey"`
//  CreatedAt time.Time
//  UpdatedAt time.Time
//  DeletedAt gorm.DeletedAt `gorm:"index"`
//}
const lockPath = "./sql/init_sql.lock"

var (
	db  *gorm.DB
	err error
)

func init() {
	//初始化 gorm实例
	initGorm()
}
func main() {
	initDatabase()
	dump.P("------------ 漂亮的打印出来 ------------")
	belongs2HasOneExample()
	//hasManyExample()
	//many2Many()
	dump.P("------------ 谢谢聆听 ------------")

}

// belongs2HasOneExample 关联查询 预加载
func belongs2HasOneExample() {
	// belongs to  和 has one 非常类似 从记录上来说都是指的1对1
	// has one 是指正向联系 比如 user表存在一个扩展表 user_other 这个就是1对1关系
	// belongs to 是指反向联系 还是上面的例子 就是 user_other是属于user的
	// 简单来说的话就是 我和你1对1  你属于我  我也属于你这种.只是逻辑是的偏差
	dump.P("------------ 接下来进入的是 Belongs To 介绍 ------------")
	var users []entity.User
	var user entity.User
	_ = users
	_ = user

	dump.P("------------ 关联查询 介绍 ------------")

	//db.Model(&user).Preload("UserOther").Find(&user) //使用Preload进行关联查询
	db.Model(&user).Joins("UserOther").Find(&user) //使用Joins进行关联查询
	dump.P(user)
	dump.P("Join Preload 适用于一对一的关系，例如： has one, belongs to")

	dump.P("------------ 关联创建/更新 介绍 ------------")
	//HasOne 创建
	insertUser := entity.User{
		Username: "李白",
		Nickname: "李白",
		Password: "123456",
		UserOther: entity.UserOther{
			OtherInfo: "用户扩展信息",
		},
	}
	_ = insertUser
	db.Create(&insertUser)

	//此处踩了1个坑
	//就是userOther里面的ID还是必须要存在的!不然的话就是会一直插入数据
	//一般关联更新的时候都是先查询出这个数据,前提是用模型查询,或者自己组装数据 然后修改这个在进行更新
	updateUser := entity.User{
		Id:       insertUser.Id,
		Username: "李白被修改成杜甫",
		Nickname: "李白被修改成杜甫",
		Password: "7654321",
		UserOther: entity.UserOther{
			Id:        insertUser.UserOther.Id,
			OtherInfo: "用户扩展信息,杜甫很忙",
		},
	}
	_ = updateUser
	//db.Updates(&updateUser)
	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updateUser)
}

// hasManyExample 1 对多数据
func hasManyExample() {
	dump.P("------------ 接下来进入的是 HasMany 介绍 ------------")
	var users []entity.User
	var user entity.User
	_ = users
	_ = user
	//关联查询
	db.Model(&user).Preload("UserRole.Role").First(&user)
	dump.P(user)

	//关联创建
	/*	createUserData := entity.User{
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
		db.Create(&createUserData)*/
	//更新
	user.UserRole = []entity.UserRole{
		{
			Id:     user.UserRole[0].Id,
			UserId: user.Id,
			RoleId: 3,
		},
	}
	db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&user)

}

func many2Many() {
	var m2mUser entity.M2MUser
	var m2mRole entity.M2MRole
	_ = m2mUser
	_ = m2mRole
	dump.P("------------ 接下来进入的是 Many2Many 介绍 ------------")
	dump.P("------------ 通过User 找Role ------------")
	db.Model(&m2mUser).Where("id=1").Preload("RoleList").First(&m2mUser)

	dump.P(m2mUser)
	//
	dump.P("------------ 通过Role 找User ------------")
	db.Model(&m2mRole).Where("id=3").Preload("UserList").First(&m2mRole)
	dump.P(m2mRole)

}

func initGorm() {
	//TODO 自行创建数据库把
	_ = "CREATE DATABASE IF NOT EXISTS `study_gorm` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;"
	dsn := "root:root@tcp(127.0.0.1:3306)/study_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("数据库因为什么而没启动：", err)
		return
	}
	//开启SQL Debug
	db = db.Debug()
	//GORM 倾向于约定，而不是配置。
	//默认情况下，GORM 使用 ID 作为主键，使用结构体名的 "蛇形复数" 作为表名，
	//字段名的 蛇形 作为列名，并使用 CreatedAt、UpdatedAt 字段追踪创建、更新时间
	//模型 使用 gorm.DeletedAt 类型作为软删除的标志,字段名可随意
}

// 首次调用使用
func initDatabase() {

	if _, err := os.Stat(lockPath); err != nil && os.IsExist(err) == false {
		dump.P("首次执行 初始化数据库...请稍后...")
		sqls := ReturnTableData()
		insertSql := InsertDefaultData()
		for k, v := range sqls {
			dump.P("正在创建表:`" + k + "`")
			db.Exec("DROP TABLE IF EXISTS `" + k + "`;")
			db.Exec(v)
		}
		dump.P("正在初始化基础数据")
		db.Exec(insertSql)
		//创建lock文件
		var d1 = []byte("1")
		_ = ioutil.WriteFile(lockPath, d1, 0777) //写入文件(字节数组)
		dump.P("很重要的一点,后续如果使用到多对多关联,注意将数据填充到user_role表中,谢谢")
	}
}

func ReturnTableData() map[string]string {
	var table = make(map[string]string, 0)
	table["role"] = "CREATE TABLE `role`\n(\n    `id`         int(11) NOT NULL AUTO_INCREMENT,\n    `name`       varchar(32) DEFAULT NULL COMMENT '角色名称',\n    `created_at` datetime    DEFAULT CURRENT_TIMESTAMP,\n    `updated_at` datetime    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n    PRIMARY KEY (`id`)\n) ENGINE = InnoDB\n  AUTO_INCREMENT = 1\n  DEFAULT CHARSET = utf8mb4 COMMENT ='角色表';"
	table["user"] = "CREATE TABLE `user`\n(\n    `id`       int(11) NOT NULL AUTO_INCREMENT,\n    `username` varchar(255) DEFAULT NULL,\n    `nickname` varchar(255) DEFAULT NULL,\n    `password` varchar(255) DEFAULT NULL,\n    PRIMARY KEY (`id`)\n) ENGINE = InnoDB\n  AUTO_INCREMENT = 1\n  DEFAULT CHARSET = utf8mb4;"
	table["user_other"] = "CREATE TABLE `user_other`\n(\n    `id`         int(11) NOT NULL AUTO_INCREMENT,\n    `user_id`    int(11)      DEFAULT NULL,\n    `other_info` varchar(255) DEFAULT NULL,\n    PRIMARY KEY (`id`),\n    KEY `q` (`user_id`)\n) ENGINE = InnoDB\n  AUTO_INCREMENT = 1\n  DEFAULT CHARSET = utf8mb4;"
	table["user_role"] = "CREATE TABLE `user_role`\n(\n    `id`      int(11) NOT NULL AUTO_INCREMENT,\n    `role_id` int(11) DEFAULT NULL COMMENT 'role表主键',\n    `user_id` int(11) DEFAULT NULL COMMENT 'user表主键',\n    PRIMARY KEY (`id`)\n) ENGINE = InnoDB\n  AUTO_INCREMENT = 1\n  DEFAULT CHARSET = utf8mb4 COMMENT ='用户角色表 1对多关系';\n"
	return table
}
func InsertDefaultData() string {
	insertSql := "INSERT INTO `study_gorm`.`role` (`name`, `created_at`, `updated_at`) VALUES ('超级管理员', '2022-06-20 15:48:59', '2022-06-20 15:48:59'), ('管理员', '2022-06-20 15:49:04', '2022-06-20 15:49:04'),('子管理员', '2022-06-20 15:49:07', '2022-06-20 15:49:07');"
	return insertSql
}
