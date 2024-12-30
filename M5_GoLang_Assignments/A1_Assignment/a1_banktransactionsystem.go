package main

import (
	"errors"
	"fmt"
)

// Account struct to store account details
type Account struct {
	ID                 int
	Name               string
	Balance            float64
	TransactionHistory []string
}

// Constants for menu options
const (
	DepositOption  = 1
	WithdrawOption = 2
	ViewBalance    = 3
	ViewHistory    = 4
	ExitOption     = 5
)

// Slice to store all accounts
var accounts []Account

// Function to find an account by ID
func findAccountByID(id int) (*Account, error) {
	for i, acc := range accounts {
		if acc.ID == id {
			return &accounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

// Function to deposit money into an account
func deposit(accountID int, amount float64) error {
	// Validate deposit amount
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	// Find account
	account, err := findAccountByID(accountID)
	if err != nil {
		return err
	}

	// Perform deposit
	account.Balance += amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Deposited: $%.2f", amount))
	return nil
}

// Function to withdraw money from an account
func withdraw(accountID int, amount float64) error {
	// Validate withdrawal amount
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}

	// Find account
	account, err := findAccountByID(accountID)
	if err != nil {
		return err
	}

	// Ensure sufficient balance
	if account.Balance < amount {
		return errors.New("insufficient balance")
	}

	// Perform withdrawal
	account.Balance -= amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Withdrew: $%.2f", amount))
	return nil
}

// Function to view account balance
func viewBalance(accountID int) (float64, error) {
	account, err := findAccountByID(accountID)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}

// Function to view transaction history
func viewTransactionHistory(accountID int) ([]string, error) {
	account, err := findAccountByID(accountID)
	if err != nil {
		return nil, err
	}
	return account.TransactionHistory, nil
}

// Menu-driven program
func main() {
	// Add some test accounts
	accounts = append(accounts, Account{ID: 1, Name: "Alice", Balance: 500.0})
	accounts = append(accounts, Account{ID: 2, Name: "Bob", Balance: 300.0})

	var choice, accountID int
	var amount float64

	for {
		fmt.Println("\n--- Bank Transaction System ---")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. View Balance")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case DepositOption:
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&accountID)
			fmt.Print("Enter Amount to Deposit: ")
			fmt.Scan(&amount)
			err := deposit(accountID, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}
		case WithdrawOption:
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&accountID)
			fmt.Print("Enter Amount to Withdraw: ")
			fmt.Scan(&amount)
			err := withdraw(accountID, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful!")
			}
		case ViewBalance:
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&accountID)
			balance, err := viewBalance(accountID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Account Balance: $%.2f\n", balance)
			}
		case ViewHistory:
			fmt.Print("Enter Account ID: ")
			fmt.Scan(&accountID)
			history, err := viewTransactionHistory(accountID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Transaction History:")
				for _, transaction := range history {
					fmt.Println(transaction)
				}
			}
		case ExitOption:
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
