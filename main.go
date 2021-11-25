package main

import (
	"fmt"
	"time"

	"peya-kata/src/transaction"
)

/*
	US1: Client into Date, Amount, balance
*/

func main() {
	transactions := transaction.ArrayTransactions

	timeNow := time.Now()
	transactionOne := transaction.CreateDepositTransaction(timeNow, 2.00, 1, nil)
	transactionTwo := transaction.CreateDepositTransaction(timeNow, 3.00, 1, &transactionOne)

	transactions = append(transactions, transactionOne, transactionTwo)
	fmt.Println(transactions)
}

func IsTrue(value bool) bool {
	return value
}

// Hexagonal | DDD
