package main

type Account interface {
	MonthlyInterest() int
	Transfer(receiver Account, amount int) string
	Deposit(amount int) string
	Withdraw(amount int) string
	CheckBalance() int
}
