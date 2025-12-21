package main

import "fmt"

type BankAccount struct {
	owner   string
	balance int
}

func debit(acc *BankAccount, amount int) {
	if acc == nil {
		return
	}
	if amount <= 0 {
		return
	}
	if amount > acc.balance {
		return
	}
	acc.balance -= amount
}
func main() {
	adarshAccount := BankAccount{"Adarsh Shukla", 5000}

	fmt.Println("Initial Balance:", adarshAccount.balance)

	debit(&adarshAccount, 2000)
	fmt.Println("Balance after attempting to debit 2000:", adarshAccount.balance)

	debit(&adarshAccount, 3500)
	fmt.Println("Balance after debit attempt for 3500:", adarshAccount.balance)

	debit(&adarshAccount, -500)
	fmt.Println("Balance after debit attempt for -500:", adarshAccount.balance)
}
