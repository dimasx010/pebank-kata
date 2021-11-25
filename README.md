# PeBank Kata Backend Golang

## Information
The project is a Kata for check knowledge in __Golang__, __TDD__ and __Backend Logic__.

## Kata Requirements
Think of your personal bank account experience When in doubt, go for the simplest solution.
Deposit and Withdrawal Account statement (date, amount, balance) Statement printing

User Stories
- US 1:
In order to save money As a bank client I want to make a deposit in my account

- US 2:
In order to retrieve some or all of my savings As a bank client I want to make a withdrawal from my account

- US 3:
In order to check my operations As a bank client I want to see the history (operation, date, amount, balance) of my operations

## Technologies
- [GoLang](https://golang.org/doc/code)
- [Golang/TestyFy](https://github.com/stretchr/testify)

## Instructions
- Clone project in your **Go directory**
  - `git clone git@github.com:dimasx010/pebank-kata.git`
- Run __go mod tidy__ for Go Modules
- Run __go run .__ for up project
- Run __go tests__ for tests

## Done
- Create new transaction
- Store Transaction (Deposit)
- Store Transaction (Withdraw)
- List of Transactions

## Tests
- Test for new transaction
- Test for create Two deposit and Check Balance is equal
- Test length of Transactions after intert Two Transactions
- Test Withdrawal Transaction
- Test Withdrawal Transaction and balance is minor to Zero (Is possible)
- Test for length of transactions
- Test for length of transactions is Zero

## Comments
It was expected to use different __Entities|Repositories__ with__DDD__ architecture, but for now we have used something simple following the TDD principles of make it work and then refactoring
