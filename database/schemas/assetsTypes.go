package schemas

type AssetType struct {
	Base
	Name string `json:"name" gorm:"column:name;not null;type:varchar(150);default:null"`
}

func (i *AssetType) TableName() string {
	return "assets_types"
}
