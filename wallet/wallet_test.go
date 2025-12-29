package main

import "testing"

func TestWallet(t* testing.T){
	t.Run("Deposit Bitcoin", func (t* testing.T){
		wallet := NewWallet(Bitcoin(0))

		wallet.deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw Bitcoin", func (t* testing.T){
		wallet := NewWallet(Bitcoin(20))

		err := wallet.withdraw(Bitcoin(10))
		
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw more Bitcoin than stored", func (t* testing.T){
		wallet := NewWallet(Bitcoin(5))

		err := wallet.withdraw(Bitcoin(10))
		
		assertError(t, err, operationErrors["nofunds"])
		assertBalance(t, wallet, Bitcoin(5))
	})
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Wanted an error and didn't get one")
	}

	if got != want {
		t.Errorf("got error %s, wanted error %s", got, want)
	}
}
func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Got unexpected error: %s", err)
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.balance()

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
