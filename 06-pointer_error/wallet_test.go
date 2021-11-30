package main

import (
  "testing"
)

func TestWallet(t *testing.T) {

  t.Run("Deposit", func(t *testing.T) {
    wallet := Wallet{}
    wallet.Deposit(Bitcoin(10))
    assertBalance(t, wallet, Bitcoin(10))
  })

  t.Run("Withdraw", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(20)}
    err := wallet.Withdraw(Bitcoin(10))

    assertNoError(t, err, errInsufficientFunds)
    assertBalance(t, wallet, Bitcoin(10))
  })

  t.Run("Withdraw insufficient funds", func(t *testing.T) {
    startingBalance := Bitcoin(10)
    wallet := Wallet{startingBalance}
    err := wallet.Withdraw(Bitcoin(20))

    assertError(t, err, errInsufficientFunds)
    assertBalance(t, wallet, startingBalance)
  })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin){
  t.Helper()
  got := wallet.Balance()

  if got != want {
    t.Errorf("got %s but want %s", got, want)
  }
}

func assertError(t testing.TB, got error, want error){
  t.Helper()
  if got == nil {
    t.Fatal("wanted an error but didn't get one")
  }

  if got.Error() != want.Error() {
    t.Errorf("got %q but want %q", got, want)
  }
}

func assertNoError(t testing.TB, got error, want error){
  t.Helper()
  if got != nil {
    t.Fatal("got an error but didn't want one")
  }
}
