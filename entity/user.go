package entity

// 一个结构体就是一个模型 可以嵌入gorm.Model 来减少字段编写
type User struct {
	Id        uint32     `json:"id" gorm:"primaryKey"`
	Username  string     `json:"username"`
	Nickname  string     `json:"nickname"`
	Password  string     `json:"password"`
	UserOther UserOther  `json:"user_other" gorm:"foreignKey:UserId"` //使用ID作为外键,不过这一般是默认的可以指定 1对1关联
	UserRole  []UserRole `json:"user_role" gorm:"foreignKey:UserId"`  // 一个用户可能有多个角色
}

//
func (User) TableName() string {
	return "user"
}

//姑且叫扩展表吧
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
	//验证一下
	RoleList []Role `json:"role_list" gorm:"many2many:user_role;joinForeignKey:UserId"`
}

// TableName
func (M2MUser) TableName() string {
	return "user"
}
