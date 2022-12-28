package pointers_errors

import (
	"testing"
)

func assertError(t testing.TB, got error, want error) {
	t.Helper();
	if got == nil {
		t.Fatal("Wanned an error, but didn't happend!")
	}

	if got != want {
		t.Errorf("Got: %q, Exptected: %q", got, want)
	}
}

func assertNotError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("An error occured")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Ether) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Got: %d, Exptected: %d", got, want)
	}
}
func TestWallet(t *testing.T) {


	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Ether(30))

		assertBalance(t, wallet, Ether(30))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Ether(35)}
		err := wallet.Withdraw(15)

		assertNotError(t, err)
		assertBalance(t, wallet, Ether(20))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Ether(30)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Ether(140))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("Should be an error, but didn't happen!")
		}
	})
}