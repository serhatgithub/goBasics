package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	Balance int
	lock    sync.Mutex
}

func main() {

	bank := &BankAccount{}
	bank.Balance = 0
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 50; i++ {
		go Deposit(bank, &wg)
	}

	for i := 0; i < 50; i++ {
		go Withdraw(bank, &wg)
	}

	wg.Wait()

	fmt.Println(bank.Balance)
}

func Deposit(bank *BankAccount, wg *sync.WaitGroup) {
	bank.lock.Lock()
	bank.Balance += 10
	bank.lock.Unlock()
	wg.Done()
}

func Withdraw(bank *BankAccount, wg *sync.WaitGroup) {
	bank.lock.Lock()
	bank.Balance -= 10
	bank.lock.Unlock()
	wg.Done()
}
