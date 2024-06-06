package schemas

type Role struct {
	Base
	Name          string `json:"role" gorm:"column:role;not null;type:varchar(150);default:null"`
	HasPrivileges bool   `json:"has_privileges" gorm:"column:has_privileges;type:boolean;default:false"`
}

func (i *Role) TableName() string {
	return "roles"
}
