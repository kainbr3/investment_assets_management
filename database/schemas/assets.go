package schemas

import "github.com/google/uuid"

type Asset struct {
	Base
	Ticker               string     `json:"ticker" gorm:"column:ticker;unique;not null;type:varchar(25);default:null"`
	Name                 string     `json:"name" gorm:"column:name;not null;type:varchar(250);default:null"`
	IdentificationNumber string     `json:"identification_number" gorm:"column:ideidentification_numberntification_number;unique;not null;type:varchar(100);default:null"`
	AssetTypeID          uuid.UUID  `json:"asset_type_id" gorm:"column:asset_type_id;type:uuid;not null"`
	AssetType            *AssetType `json:"asset_type" gorm:"foreignKey:AssetTypeID;references:ID"`
	CountryID            uuid.UUID  `json:"country_id" gorm:"column:country_id;type:uuid;not null"`
	Country              *Country   `json:"country" gorm:"foreignKey:CountryID;references:ID"`
}

func (i *Asset) TableName() string {
	return "assets"
}
