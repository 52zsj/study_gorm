package entity

import (
	"gorm.io/gorm"
)

// 1个用户有1条扩展信息 用户1对1扩展 hasOne  扩展属于用户belongsTo (当然也是1对1的关系)
// 1个用户有多重身份 用户1对多角色 hasMany
// 角色和权限是多对多 ManytoMany 1个角色对应多个权限  1个权限也可以属于多个角色

// 一个结构体就是一个模型 可以嵌入gorm.Model 来减少字段编写
type TestGromModel struct {
	gorm.Model
	Test1 string
	Test2 string
}

type User struct {
	Id        uint32     `json:"id" gorm:"primaryKey"`
	Username  string     `json:"username"`
	Nickname  string     `json:"nickname"`
	Password  string     `json:"password"`
	UserOther UserOther  `json:"user_other" gorm:"foreignKey:UserId"` // 使用ID作为外键,不过这一般是默认的可以指定 1对1关联
	UserRole  []UserRole `json:"user_role" gorm:"foreignKey:UserId"`  // 一个用户可能有多个角色
}

// 接收器User
func (User) TableName() string {
	return "user"
}

// 姑且叫扩展表吧
type UserOther struct {
	Id        uint32 `json:"id" gorm:"primarykey"`
	UserId    uint32 `json:"user_id"`
	OtherInfo string `json:"other_info"`
}

func (UserOther) TableName() string {
	return "user_other"
}

type M2MUser struct {
	Id       uint32 `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	//many2many 指定连接的表名(即中间表) joinForeignKey:用到了中间表的外键是哪个;
	//比如我现在用User查询role 那在user_role表中用到的外键就是user_id
	RoleList []Role `json:"role_list" gorm:"many2many:user_role;joinForeignKey:UserId"`
}

// TableName
func (M2MUser) TableName() string {
	return "user"
}
