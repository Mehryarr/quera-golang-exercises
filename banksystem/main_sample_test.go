package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepositCredit(t *testing.T) {
	var a Account = NewInvestmentAccount()
	a.Deposit(250)
	assert.Equal(t, 250, a.CheckBalance())
}

func TestWithdrawCredit(t *testing.T) {
	var a Account = NewSavingsAccount()
	a.Deposit(250)
	a.Withdraw(50)
	assert.Equal(t, 200, a.CheckBalance())
}

func TestTransferSuccess(t *testing.T) {
	var a1, a2 Account = NewCheckingAccount(), NewInvestmentAccount()
	a1.Deposit(300)
	s := a1.Transfer(a2, 250)
	assert.Equal(t, "Success", s)
}

func TestMonthlyInterest2(t *testing.T) {
	var a1 Account = NewSavingsAccount()
	a1.Deposit(2000)
	i := a1.MonthlyInterest()
	assert.Equal(t, 8, i)
}
