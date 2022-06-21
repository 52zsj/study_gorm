package entity

type Role struct {
	Id   uint32 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

//
func (Role) TableName() string {
	return "role"
}

// 多对多关联
type M2MRole struct {
	Id       uint32 `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	UserList []User `json:"user_list" gorm:"many2many:user_role;joinForeignKey:RoleId"`
}

//
func (M2MRole) TableName() string {
	return "role"
}
