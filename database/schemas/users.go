package schemas

import "github.com/google/uuid"

type User struct {
	Base
	FirstName string     `json:"first_name" gorm:"column:first_name;not null;type:varchar(150);default:null"`
	LastName  string     `json:"last_name" gorm:"column:first_name;not null;type:varchar(150);default:null"`
	Email     string     `json:"email" gorm:"column:email;not null;type:varchar(150);default:null"`
	Password  string     `json:"password" gorm:"column:password;not null;type:varchar(150);default:null"`
	RoleID    uuid.UUID  `json:"role_id" gorm:"column:asset_type_id;type:uuid;not null"`
	Role      *AssetType `json:"role" gorm:"foreignKey:RoleID;references:ID"`
}

func (i *User) TableName() string {
	return "users"
}
