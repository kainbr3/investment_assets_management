package schemas

type IbanCountry struct {
	Base
	Name        string `json:"name" gorm:"column:name;not null;type:varchar(150);default:null"`
	Alpha2Code  string `json:"alpha2_code" gorm:"column:alpha2_code;unique;not null;type:varchar(2);default:null"`
	Alpha3Code  string `json:"alpha3_code" gorm:"column:alpha3_code;unique;not null;type:varchar(3);default:null"`
	NumericCode int32  `json:"numeric_code" gorm:"column:numeric_code;unique;not null;type:int;default:0"`
}

func (i *IbanCountry) TableName() string {
	return "iban_countries"
}
