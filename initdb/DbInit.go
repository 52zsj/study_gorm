package initdb

import (
	"database/sql"
	"fmt"
	"github.com/gookit/goutil/dump"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

//初始化数据库标志
const lockDatabasePath = "./sql/init_database.lock"

//初始化数据标志
const lockTablePath = "./sql/init_table.lock"

const (
	databaseName = "study_gorm"
	host         = "127.0.0.1"
	username     = "root"
	password     = "root"
	port         = "3306"
	dsnStr       = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

// InitDatabase 首次调用使用
func InitDatabase() {
	if _, err := os.Stat(lockDatabasePath); err != nil && os.IsExist(err) == false {
		//自动创建数据库
		createDatabse := "CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;"
		createDatabse = fmt.Sprintf(createDatabse, databaseName)

		inidatabse, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port))
		if err = inidatabse.Ping(); err != nil {
			//panic(err.Error())
			dump.P(err.Error())
			os.Exit(0)
		}
		_, err = inidatabse.Exec(createDatabse)
		if err != nil {
			dump.P(err.Error())
			os.Exit(0)
		}
		//关闭创建的数据库
		defer func(inidatabse *sql.DB) {
			err := inidatabse.Close()
			if err != nil {
				dump.P(err.Error())
			}
		}(inidatabse)
		//创建lock文件
		var d1 = []byte("1")
		_ = ioutil.WriteFile(lockDatabasePath, d1, 0755) //写入文件(字节数组)
	}

}

// InitTableAndCreateDb 创建表格 且初始化db实例
func InitTableAndCreateDb(debug bool) *gorm.DB {
	//自动创建表
	var err error
	var db *gorm.DB
	dsn := fmt.Sprintf(dsnStr, username, password, host, port, databaseName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		dump.P(err.Error())
		return nil
	}

	if debug {
		db = db.Debug()
	}

	if _, err := os.Stat(lockTablePath); err != nil && os.IsExist(err) == false {
		dump.P("首次执行 正在初始化数据库...请稍后...")
		sqls := returnTableData()
		insertSql := insertDefaultData()
		for k, v := range sqls {
			dump.P("正在创建表:`" + k + "`")
			//直接删除
			db.Exec("DROP TABLE IF EXISTS `" + k + "`;")
			db.Exec(v)
		}
		dump.P("正在初始化基础数据...")

		db.Exec(insertSql)

		dump.P("很重要的一点,后续如果使用到多对多关联,注意将数据填充到user_role表中,谢谢")
		//创建lock文件
		var d1 = []byte("1")
		_ = ioutil.WriteFile(lockTablePath, d1, 0755) //写入文件(字节数组)
		dump.P("数据库初始化完毕,请继续操作.....")
	}
	return db
}

// returnTableData 表格创建语句
func returnTableData() map[string]string {
	var table = make(map[string]string, 0)
	table["role"] = "CREATE TABLE `role` (`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(32) DEFAULT NULL COMMENT '角色名称',`created_at` datetime DEFAULT CURRENT_TIMESTAMP,`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,PRIMARY KEY (`id`)) ENGINE = InnoDB  AUTO_INCREMENT = 1  DEFAULT CHARSET = utf8mb4 COMMENT ='角色表';"
	table["user"] = "CREATE TABLE `user`(`id` int(11) NOT NULL AUTO_INCREMENT,`username` varchar(255) DEFAULT NULL,`nickname` varchar(255) DEFAULT NULL,`password` varchar(255) DEFAULT NULL,PRIMARY KEY (`id`)) ENGINE = InnoDB  AUTO_INCREMENT = 1  DEFAULT CHARSET = utf8mb4;"
	table["user_other"] = "CREATE TABLE `user_other`(`id` int(11) NOT NULL AUTO_INCREMENT,`user_id`int(11)  DEFAULT NULL,`other_info` varchar(255) DEFAULT NULL,PRIMARY KEY (`id`),KEY `q` (`user_id`)) ENGINE = InnoDB  AUTO_INCREMENT = 1  DEFAULT CHARSET = utf8mb4;"
	table["user_role"] = "CREATE TABLE `user_role`(`id`  int(11) NOT NULL AUTO_INCREMENT,`role_id` int(11) DEFAULT NULL COMMENT 'role表主键',`user_id` int(11) DEFAULT NULL COMMENT 'user表主键',PRIMARY KEY (`id`)) ENGINE = InnoDB  AUTO_INCREMENT = 1  DEFAULT CHARSET = utf8mb4 COMMENT ='用户角色表 1对多关系';"
	return table
}

// insertDefaultData 初始化基础数据
func insertDefaultData() string {
	insertSql := "INSERT INTO `role` (`name`) VALUES ('超级管理员'), ('管理员'),('子管理员'),('子子管理员');"
	return insertSql
}
