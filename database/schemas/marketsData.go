package schemas

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type MarketData struct {
	AssetID     uuid.UUID       `json:"asset_id" gorm:"column:asset_id;type:uuid;not null"`
	Asset       *Asset          `json:"asset" gorm:"foreignKey:AssetID;references:ID"`
	MarketPrice decimal.Decimal `json:"market_price" gorm:"column:market_price;type:decimal(18,2);default:0.00"`
}

func (i *MarketData) TableName() string {
	return "markets_data"
}
