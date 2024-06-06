package schemas

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Trade struct {
	Base
	Date        time.Time       `json:"date" gorm:"column:date;not null;type:date;"`
	Price       decimal.Decimal `json:"price" gorm:"column:price;not null;type:decimal(18,2);default:0.00"`
	Shares      int32           `json:"shares" gorm:"column:shares;not null;type:int;default:0"`
	AssetID     uuid.UUID       `json:"asset_id" gorm:"column:asset_id;type:uuid;not null"`
	Asset       *Asset          `json:"asset" gorm:"foreignKey:AssetID;references:ID"`
	TradeTypeID uuid.UUID       `json:"trade_type_id" gorm:"column:asset_type_id;type:uuid;not null"`
	TradeType   *AssetType      `json:"trade_type" gorm:"foreignKey:TradeTypeID;references:ID"`
	CurrencyID  uuid.UUID       `json:"currency_id" gorm:"column:currency_id;type:uuid;not null"`
	Currency    *Currency       `json:"currency" gorm:"foreignKey:CurrencyID;references:ID"`
}

func (i *Trade) TableName() string {
	return "tardes"
}
