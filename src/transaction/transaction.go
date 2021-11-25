package transaction

import (
	"errors"
	"time"
)

type Transaction struct {
	Date    time.Time
	Amount  float32
	Balance float32
	UserId  int
}

var ArrayTransactions []Transaction

func CreateDepositTransaction(date time.Time, amount float32, userId int, lastTransaction *Transaction) Transaction {

	var transaction Transaction

	transaction.Date = date
	transaction.Amount = amount
	transaction.UserId = userId

	transaction.Balance = transaction.Amount
	if lastTransaction != nil {
		transaction.Balance = transaction.Amount + lastTransaction.Balance
	}

	return transaction
}

func InsertTransanction(transaction Transaction) []Transaction {
	ArrayTransactions = append(ArrayTransactions, transaction)
	return ArrayTransactions
}

func GetTransactions() []Transaction {
	return ArrayTransactions
}

func CreateWithdrawTransaction(date time.Time, amount float32, userId int, lastTransaction *Transaction) Transaction {

	var transaction Transaction

	transaction.Date = date
	transaction.Amount = amount
	transaction.UserId = userId

	transaction.Balance = transaction.Amount
	if lastTransaction != nil {
		transaction.Balance = lastTransaction.Balance - transaction.Amount
	}

	if transaction.Balance < 0 {
		errors.New("string del error")
	}

	return transaction
}
