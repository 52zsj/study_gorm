package entity

type UserRole struct {
	Id     uint32 `json:"id" gorm:"primaryKey"`
	UserId uint32 `json:"user_id"`
	RoleId uint32 `json:"role_id"`
	// 可以用forignKey 来重新定义外键名称 gorm:"foreignKey:Name"
	Role Role `json:"role"`
}

func (UserRole) TableName() string {
	return "user_role"
}
