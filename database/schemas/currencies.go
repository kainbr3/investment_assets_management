package schemas

type Currency struct {
	Base
	Name   string `json:"name" gorm:"column:name;not null;type:varchar(150);default:null"`
	Code   string `json:"code" gorm:"column:code;unique;not null;type:varchar(3);default:null"`
	Symbol string `json:"symbol" gorm:"column:symbol;not null;type:varchar(3);default:null"`
}

func (i *Currency) TableName() string {
	return "currencies"
}
