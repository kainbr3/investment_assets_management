package schemas

type Country struct {
	Base
	Name      string `json:"name" gorm:"column:name;not null;type:varchar(150);default:null"`
	Code      string `json:"code" gorm:"column:code;unique;not null;type:varchar(3);default:null"`
	ShortCode string `json:"short_code" gorm:"column:short_code;unique;not null;type:varchar(2);default:null"`
	IbanNumer string `json:"iban_number" gorm:"column:iban_number;unique;not null;type:varchar(3);default:null"`
}

func (i *Country) TableName() string {
	return "countries"
}
