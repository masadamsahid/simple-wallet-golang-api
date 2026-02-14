package model

import "time"

type Transaction struct {
	ID              uint            `gorm:"column:id;primary_key" json:"id"`
	UserID          uint            `gorm:"column:user_id;not null" json:"user_id"`
	Amount          int64           `gorm:"column:amount;not null" json:"amount"`
	TransactionType TransactionType `gorm:"column:transaction_type;type:transaction_type_enum;not null" json:"transaction_type"`
	Description     string          `gorm:"column:description;type:text" json:"description"`

	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`

	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

func (Transaction) TableName() string {
	return "transactions"
}

type TransactionType string

const (
	TransactionType_WITHDRAW TransactionType = "withdrawal"
	TransactionType_DEPOSIT  TransactionType = "deposit"
)
