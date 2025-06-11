// 2. Реализуйте класс BankAccount, который имеет приватный атрибут balance. Добавьте методы для пополнения и снятия средств, а также метод для получения текущего баланса.

package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	balance float64
}

func NewBankAccount(initialBalance float64) *BankAccount {
	return &BankAccount{balance: initialBalance}
}

func (acc *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма для пополнения должна быть положительной")
	}
	acc.balance += amount
	return nil
}

func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма для снятия должна быть положительной")
	}
	if amount > acc.balance {
		return errors.New("недостаточно средств на счете")
	}
	acc.balance -= amount
	return nil
}

func (acc *BankAccount) GetBalance() float64 {
	return acc.balance
}

func main() {
	account := NewBankAccount(1000)

	fmt.Printf("Начальный баланс: %.2f\n", account.GetBalance())

	err := account.Deposit(500)
	if err != nil {
		fmt.Println("Ошибка при пополнении:", err)
	} else {
		fmt.Printf("После пополнения на 500: %.2f\n", account.GetBalance())
	}

	err = account.Withdraw(200)
	if err != nil {
		fmt.Println("Ошибка при снятии:", err)
	} else {
		fmt.Printf("После снятия 200: %.2f\n", account.GetBalance())
	}

	err = account.Withdraw(2000)
	if err != nil {
		fmt.Println("Ошибка при снятии:", err)
	} else {
		fmt.Printf("После снятия 2000: %.2f\n", account.GetBalance())
	}

	err = account.Deposit(-100)
	if err != nil {
		fmt.Println("Ошибка при пополнении:", err)
	} else {
		fmt.Printf("После пополнения на -100: %.2f\n", account.GetBalance())
	}

	fmt.Printf("Финальный баланс: %.2f\n", account.GetBalance())
	fmt.Printf("Финальный баланс: %.2f\n", account.GetBalance())
	fmt.Println("\nНажмите Enter для выхода...")
	fmt.Scanln()
}
