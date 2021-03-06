# PeBank Kata Backend Golang :moneybag:

## Information
The project is a Kata for check knowledge in __Golang__, __TDD__ and __Backend Logic__.

## Kata Requirements
Think of your personal bank account experience When in doubt, go for the simplest solution.
Deposit and Withdrawal Account statement (date, amount, balance) Statement printing

User Stories
- :ok: US 1: 
In order to save money As a bank client I want to make a deposit in my account

- :ok: US 2: 
In order to retrieve some or all of my savings As a bank client I want to make a withdrawal from my account

- :ok: US 3: 
In order to check my operations As a bank client I want to see the history (operation, date, amount, balance) of my operations

## Technologies
- [GoLang](https://golang.org/doc/code)
- [Golang/Testyfy](https://github.com/stretchr/testify)

## Instructions
- Clone project in your **Go directory**
  - `git clone git@github.com:dimasx010/pebank-kata.git`
- Run __go mod tidy__ for Go Modules
- Run __go run .__ for up project
- Run __go tests__ for tests main
- Run __go test -v .\src\transaction__ for tests transaction (Windows)
- Run __go test -v /src/transaction__ for tests transaction (Others)

## Done
- Create new transaction
- Store Transaction (Deposit)
- Store Transaction (Withdraw)
- List of Transactions

## Tests
- Test for new transaction Desposit Operation
- Test for new transaction WithDraw Operation
- Test for create Two deposit and Check Balance is equal
- Test length of Transactions after intert Two Transactions
- Test Withdrawal Transaction
- Test Withdrawal Transaction and balance is minor to Zero (Is possible)
- Test for length of transactions
- Test for length of transactions is Zero

__Important__: Because __ArrayOfTransactions__ is Global Variable, is necessary restart value in each test

## Comments
It was expected to use different __Entities|Repositories__ with __DDD__ architecture, but for now we have used something simple following the __TDD__ principles of make it work and then refactoring

__Written with :green_heart: by Dimas__