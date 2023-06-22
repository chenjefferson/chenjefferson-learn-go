package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		wallet.Deposit(10)

		got = wallet.Balance()
		want = Bitcoin(20)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.run("withdraw", func(t *testing.T) {
		wallet := Wallet{}

		got, err := wallet.Withdraw(10)

	})
}
