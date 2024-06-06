package schemas

import "github.com/google/uuid"

type IbanCountryCurrency struct {
	Base
	IbanCountryID  uuid.UUID     `json:"iban_country_id" gorm:"column:iban_country_id;type:uuid;not null"`
	IbanCountry    *IbanCountry  `json:"iban_country" gorm:"foreignKey:IbanCountryID;references:ID"`
	IbanCurrencyID uuid.UUID     `json:"iban_currency_id" gorm:"column:iban_currency_id;type:uuid;not null"`
	IbanCurrency   *IbanCurrency `json:"iban_currency" gorm:"foreignKey:IbanCurrencyID;references:ID"`
}
