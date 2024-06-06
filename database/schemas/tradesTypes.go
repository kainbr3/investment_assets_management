package schemas

type TradeType struct {
	Base
	Name string `json:"name" gorm:"column:name;not null;type:varchar(150);default:null"`
}

func (i *TradeType) TableName() string {
	return "trades_types"
}
