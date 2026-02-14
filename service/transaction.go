package service

import (
	"errors"
	"fmt"
	"wallet-app/model"
	"wallet-app/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionService struct {
	transactionRepo repository.ITransactionRepository
	userRepo        repository.IUserRepository
	db              *gorm.DB
}

type ITransactionService interface {
	Deposit(userID uint, amount int64) (*model.Transaction, error)
	Withdraw(userID uint, amount int64) (*model.Transaction, error)
	GetHistory(userID uint) ([]model.Transaction, error)
}

func NewTransactionService(transactionRepo repository.ITransactionRepository, userRepo repository.IUserRepository, db *gorm.DB) ITransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
		db:              db,
	}
}

func (s *TransactionService) Deposit(userID uint, amount int64) (*model.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	var newDeposit *model.Transaction

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var user model.User
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, userID).Error; err != nil {
			return err
		}

		user.Balance += amount
		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		newTransaction := &model.Transaction{
			UserID:          userID,
			Amount:          amount,
			TransactionType: model.TransactionType_DEPOSIT,
			Description:     fmt.Sprintf("Top Up Balance IDR %d by %s (uid: %d)", amount, user.Name, user.ID),
		}

		var err error
		newDeposit, err = s.transactionRepo.CreateTransaction(tx, newTransaction)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return newDeposit, nil
}

// Withdraw implements ITransactionService.
func (s *TransactionService) Withdraw(userID uint, amount int64) (*model.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	var newWithdrawal *model.Transaction

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var user model.User
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, userID).Error; err != nil {
			return err
		}

		if user.Balance < amount {
			return errors.New("insufficient balance")
		}

		user.Balance -= amount
		if err := tx.Save(&user).Error; err != nil {
			return err
		}

		newTransaction := &model.Transaction{
			UserID:          userID,
			Amount:          amount,
			TransactionType: model.TransactionType_WITHDRAW,
			Description:     fmt.Sprintf("Withdraw Balance IDR %d by %s (uid: %d)", amount, user.Name, user.ID),
		}

		var err error
		newWithdrawal, err = s.transactionRepo.CreateTransaction(tx, newTransaction)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return newWithdrawal, nil
}

func (s *TransactionService) GetHistory(userID uint) ([]model.Transaction, error) {
	return s.transactionRepo.GetTransactionsByUserID(nil, userID)
}
