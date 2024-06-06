package schemas

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PortfolioAsset struct {
	AssetID      uuid.UUID       `json:"asset_id" gorm:"column:asset_id;type:uuid;not null"`
	Asset        *Asset          `json:"asset" gorm:"foreignKey:AssetID;references:ID"`
	Shares       int32           `json:"shares" gorm:"column:shares;not null;type:int;default:0"`
	AveragePrice decimal.Decimal `json:"average_price" gorm:"column:average_price;not null;type:decimal(18,2);default:0.00"`
	GainLoss     decimal.Decimal `json:"gain_loss" gorm:"column:gain_loss;type:decimal(18,2);default:0.00"`
	MarketData   *MarketData     `json:"market_data" gorm:"foreignKey:AssetID;references:AssetID"`
}

func (i *PortfolioAsset) TableName() string {
	return "portfolios_assets"
}
