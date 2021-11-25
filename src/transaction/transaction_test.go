package transaction

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	timeNow := time.Now()

	transactionMock := Transaction{Date: timeNow, Amount: 2.00, Balance: 2.00, UserId: 1}
	transaction := CreateDepositTransaction(timeNow, 2.00, 1, nil)

	assert.Equal(t, transactionMock, transaction, "The two Transactions should be the same.")
}

func TestCreateTwoDepositTransactionBalance(t *testing.T) {

	timeNow := time.Now()
	transactionOne := CreateDepositTransaction(timeNow, 2.00, 1, nil)
	transactionTwo := CreateDepositTransaction(timeNow, 3.00, 1, &transactionOne)

	InsertTransanction(transactionOne)
	InsertTransanction(transactionTwo)

	assert.EqualValues(t, 5.00, ArrayTransactions[1].Balance, "The two Balances should be the same.")
}

func TestInsertTwoTransaction(t *testing.T) {

	timeNow := time.Now()

	transactionOne := CreateDepositTransaction(timeNow, 2.00, 1, nil)
	transactionTwo := CreateDepositTransaction(timeNow, 3.00, 1, &transactionOne)

	InsertTransanction(transactionOne)
	InsertTransanction(transactionTwo)

	assert.Len(t, ArrayTransactions, 2)
}

func TestWithdrawalTransaction(t *testing.T) {
	timeNow := time.Now()
	transactionOne := CreateDepositTransaction(timeNow, 3.00, 1, nil)
	transactionTwo := CreateWithdrawTransaction(timeNow, 3.00, 1, &transactionOne)

	InsertTransanction(transactionOne)
	InsertTransanction(transactionTwo)

	assert.EqualValues(t, 0, ArrayTransactions[1].Balance, "The two Balances should be the same.")
}

func TestWithdrawalBalanceminorZeroTransaction(t *testing.T) {
	timeNow := time.Now()
	transactionOne := CreateDepositTransaction(timeNow, 3.00, 1, nil)
	transactionTwo := CreateWithdrawTransaction(timeNow, 4.00, 1, &transactionOne)

	InsertTransanction(transactionOne)
	InsertTransanction(transactionTwo)

	assert.EqualValues(t, -1, ArrayTransactions[1].Balance, "The two Balances should be the same.")
}

func TestGetTransactions(t *testing.T) {
	timeNow := time.Now()

	transactionOne := CreateDepositTransaction(timeNow, 2.00, 1, nil)
	transactionTwo := CreateDepositTransaction(timeNow, 3.00, 1, &transactionOne)

	InsertTransanction(transactionOne)
	InsertTransanction(transactionTwo)

	totalTransactions := GetTransactions()

	assert.Len(t, totalTransactions, 2)
}

func TestGetTransactionsZero(t *testing.T) {

	totalTransactions := GetTransactions()

	assert.Len(t, totalTransactions, 0)
}
