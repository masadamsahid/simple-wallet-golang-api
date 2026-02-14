package repository

import (
	"wallet-app/model"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

type ITransactionRepository interface {
	CreateTransaction(db *gorm.DB, transaction *model.Transaction) (*model.Transaction, error)
	GetTransactionsByUserID(db *gorm.DB, userID uint) ([]model.Transaction, error)
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{db: db}
}

func (t *TransactionRepository) CreateTransaction(db *gorm.DB, transaction *model.Transaction) (*model.Transaction, error) {
	if db == nil {
		db = t.db
	}

	err := db.Create(transaction).Preload("User").First(&transaction, transaction.ID).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *TransactionRepository) GetTransactionsByUserID(db *gorm.DB, userID uint) ([]model.Transaction, error) {
	if db == nil {
		db = t.db
	}

	var transactions []model.Transaction
	err := db.Where("user_id = ?", userID).Order("created_at desc").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
