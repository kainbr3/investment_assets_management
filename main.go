package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kainbr3/investment_assets_management/database"
	"github.com/kainbr3/investment_assets_management/database/schemas"

	"github.com/joho/godotenv"
)

func main() {
	// load envs from root .env file
	godotenv.Load()

	// executes the mgiration to create table if not exists
	db := database.Start(context.Background())

	// if err := db.AutoMigrate(&schemas.WorldCurrencySymbol{}); err != nil {
	// 	log.Println(err)
	// }

	// if err := db.AutoMigrate(&schemas.IbanCurrency{}); err != nil {
	// 	log.Println(err)
	// }

	// if err := db.AutoMigrate(&schemas.IbanCountry{}); err != nil {
	// 	log.Println(err)
	// }

	// if err := db.AutoMigrate(&schemas.IbanCountryCurrency{}); err != nil {
	// 	log.Println(err)
	// }

	// sql := `
	// INSERT INTO world_currencies_symbols (name, code, symbol, is_active)
	// VALUES
	// ('Albania Lek', 'ALL', 'Lek', true),
	// ('Afghanistan Afghani', 'AFN', '؋', true),
	// ('Argentina Peso', 'ARS', '$', true),
	// ('Aruba Guilder', 'AWG', 'ƒ', true),
	// ('Australia Dollar', 'AUD', '$', true),
	// ('Azerbaijan Manat', 'AZN', '₼', true),
	// ('Bahamas Dollar', 'BSD', '$', true),
	// ('Barbados Dollar', 'BBD', '$', true),
	// ('Belarus Ruble', 'BYN', 'Br', true),
	// ('Belize Dollar', 'BZD', 'BZ$', true),
	// ('Bermuda Dollar', 'BMD', '$', true),
	// ('Bolivia Bolíviano', 'BOB', '$b', true),
	// ('Bosnia and Herzegovina Convertible Mark', 'BAM', 'KM', true),
	// ('Botswana Pula', 'BWP', 'P', true),
	// ('Bulgaria Lev', 'BGN', 'лв', true),
	// ('Brazil Real', 'BRL', 'R$', true),
	// ('Brunei Darussalam Dollar', 'BND', '$', true),
	// ('Cambodia Riel', 'KHR', '៛', true),
	// ('Canada Dollar', 'CAD', '$', true),
	// ('Cayman Islands Dollar', 'KYD', '$', true),
	// ('Chile Peso', 'CLP', '$', true),
	// ('Colombia Peso', 'COP', '$', true),
	// ('Costa Rica Colon', 'CRC', '₡', true),
	// ('Croatia Kuna', 'HRK', 'kn', true),
	// ('Cuba Peso', 'CUP', '₱', true),
	// ('Czech Republic Koruna', 'CZK', 'Kč', true),
	// ('Denmark Krone', 'DKK', 'kr', true),
	// ('Dominican Republic Peso', 'DOP', 'RD$', true),
	// ('East Caribbean Dollar', 'XCD', '$', true),
	// ('Egypt Pound', 'EGP', '£', true),
	// ('El Salvador Colon', 'SVC', '$', true),
	// ('Euro Member Countries', 'EUR', '€', true),
	// ('Falkland Islands (Malvinas) Pound', 'FKP', '£', true),
	// ('Fiji Dollar', 'FJD', '$', true),
	// ('Ghana Cedi', 'GHS', '¢', true),
	// ('Gibraltar Pound', 'GIP', '£', true),
	// ('Guatemala Quetzal', 'GTQ', 'Q', true),
	// ('Guernsey Pound', 'GGP', '£', true),
	// ('Guyana Dollar', 'GYD', '$', true),
	// ('Honduras Lempira', 'HNL', 'L', true),
	// ('Hong Kong Dollar', 'HKD', '$', true),
	// ('Hungary Forint', 'HUF', 'Ft', true),
	// ('Iceland Krona', 'ISK', 'kr', true),
	// ('India Rupee', 'INR', '₨', true),
	// ('Indonesia Rupiah', 'IDR', 'Rp', true),
	// ('Iran Rial', 'IRR', '﷼', true),
	// ('Isle of Man Pound', 'IMP', '£', true),
	// ('Israel Shekel', 'ILS', '₪', true),
	// ('Jamaica Dollar', 'JMD', 'J$', true),
	// ('Japan Yen', 'JPY', '¥', true),
	// ('Jersey Pound', 'JEP', '£', true),
	// ('Kazakhstan Tenge', 'KZT', 'лв', true),
	// ('Korea (North) Won', 'KPW', '₩', true),
	// ('Korea (South) Won', 'KRW', '₩', true),
	// ('Kyrgyzstan Som', 'KGS', 'лв', true),
	// ('Laos Kip', 'LAK', '₭', true),
	// ('Lebanon Pound', 'LBP', '£', true),
	// ('Liberia Dollar', 'LRD', '$', true),
	// ('Macedonia Denar', 'MKD', 'ден', true),
	// ('Malaysia Ringgit', 'MYR', 'RM', true),
	// ('Mauritius Rupee', 'MUR', '₨', true),
	// ('Mexico Peso', 'MXN', '$', true),
	// ('Mongolia Tughrik', 'MNT', '₮', true),
	// ('Mozambique Metical', 'MZN', 'MT', true),
	// ('Namibia Dollar', 'NAD', '$', true),
	// ('Nepal Rupee', 'NPR', '₨', true),
	// ('Netherlands Antilles Guilder', 'ANG', 'ƒ', true),
	// ('New Zealand Dollar', 'NZD', '$', true),
	// ('Nicaragua Cordoba', 'NIO', 'C$', true),
	// ('Nigeria Naira', 'NGN', '₦', true),
	// ('Norway Krone', 'NOK', 'kr', true),
	// ('Oman Rial', 'OMR', '﷼', true),
	// ('Pakistan Rupee', 'PKR', '₨', true),
	// ('Panama Balboa', 'PAB', 'B/.', true),
	// ('Paraguay Guarani', 'PYG', 'Gs', true),
	// ('Peru Sol', 'PEN', 'S/.', true),
	// ('Philippines Peso', 'PHP', '₱', true),
	// ('Poland Zloty', 'PLN', 'zł', true),
	// ('Qatar Riyal', 'QAR', '﷼', true),
	// ('Romania Leu', 'RON', 'lei', true),
	// ('Russia Ruble', 'RUB', '₽', true),
	// ('Saint Helena Pound', 'SHP', '£', true),
	// ('Saudi Arabia Riyal', 'SAR', '﷼', true),
	// ('Serbia Dinar', 'RSD', 'Дин.', true),
	// ('Seychelles Rupee', 'SCR', '₨', true),
	// ('Singapore Dollar', 'SGD', '$', true),
	// ('Solomon Islands Dollar', 'SBD', '$', true),
	// ('Somalia Shilling', 'SOS', 'S', true),
	// ('South Africa Rand', 'ZAR', 'R', true),
	// ('Sri Lanka Rupee', 'LKR', '₨', true),
	// ('Sweden Krona', 'SEK', 'kr', true),
	// ('Switzerland Franc', 'CHF', 'CHF', true),
	// ('Suriname Dollar', 'SRD', '$', true),
	// ('Syria Pound', 'SYP', '£', true),
	// ('Taiwan New Dollar', 'TWD', 'NT$', true),
	// ('Thailand Baht', 'THB', '฿', true),
	// ('Trinidad and Tobago Dollar', 'TTD', 'TT$', true),
	// ('Turkey Lira', 'TRY', '₺', true),
	// ('Tuvalu Dollar', 'TVD', '$', true),
	// ('Ukraine Hryvnia', 'UAH', '₴', true),
	// ('United Kingdom Pound', 'GBP', '£', true),
	// ('United States Dollar', 'USD', '$', true),
	// ('Uruguay Peso', 'UYU', '$U', true),
	// ('Uzbekistan Som', 'UZS', 'лв', true),
	// ('Venezuela Bolívar', 'VEF', 'Bs', true),
	// ('Viet Nam Dong', 'VND', '₫', true),
	// ('Yemen Rial', 'YER', '﷼', true),
	// ('Zimbabwe Dollar', 'ZWD', 'Z$', true);
	// `

	// tx := db.Exec(sql)
	// fmt.Println(tx.RowsAffected)
	// fmt.Println(tx.Error)
	// fmt.Println(tx)
	// if tx.Error != nil {
	// 	log.Println(tx.Error)
	// }

	//lista := []*schemas.WorldCurrencySymbol{}
	//lista := []*schemas.IbanCurrency{}
	//lista := []*schemas.IbanCountry{}
	lista := []*schemas.IbanCountryCurrency{}
	db.Order("name asc").Find(&lista, "1=1").Order("name asc")

	jsonString, err := json.Marshal(lista)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	fmt.Println(string(jsonString))
	//fmt.Println(lista)
}
