package schemas

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Portfolio struct {
	Base
	Name          string           `json:"name" gorm:"column:name;not null;type:varchar(150);default:null"`
	UserID        uuid.UUID        `json:"user_id" gorm:"column:user_id;type:uuid;not null"`
	User          *User            `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Assets        []PortfolioAsset `json:"assets" gorm:"many2many:portfolio_assets"`
	TotalValue    decimal.Decimal  `json:"total_value" gorm:"column:total_value;type:decimal(18,2);default:0.00"`
	TotalGainLoss decimal.Decimal  `json:"total_gain_loss" gorm:"column:total_gain_loss;type:decimal(18,2);default:0.00"`
}

func (i *Portfolio) TableName() string {
	return "portfolios"
}
