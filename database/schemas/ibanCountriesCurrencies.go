package schemas

import "github.com/google/uuid"

type IbanCountryCurrency struct {
	Base
	Name           string        `json:"name" gorm:"column:name;not null;type:varchar(150);default:null"`
	IbanCountryID  uuid.UUID     `json:"iban_country_id" gorm:"column:iban_country_id;type:uuid"`
	IbanCountry    *IbanCountry  `json:"iban_country" gorm:"foreignKey:IbanCountryID;references:ID"`
	IbanCurrencyID uuid.UUID     `json:"iban_currency_id" gorm:"column:iban_currency_id;type:uuid"`
	IbanCurrency   *IbanCurrency `json:"iban_currency" gorm:"foreignKey:IbanCurrencyID;references:ID"`
}

func (i *IbanCountryCurrency) TableName() string {
	return "iban_countries_currencies"
}
