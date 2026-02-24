package main

import (
	"errors"
	"fmt"
)

var BalanceErr = errors.New("Yetersiz bakiye!")

type Balance struct {
	Name    string
	balance int
}

type Product struct {
	price int
}

func checkBalance(b Balance, p Product) error {
	if b.balance < p.price {
		return fmt.Errorf("islem basarisiz (bakiye: %d, fiyat: %d): %w", b.balance, p.price, BalanceErr)
	}
	return nil
}

func odemeAl() error {
	if err := checkBalance(Balance{balance: 10}, Product{price: 20}); err != nil {
		return fmt.Errorf("Odeme Alinamadi %w", err)
	}
	return nil
}

func main() {
	err := odemeAl()
	if err != nil {
		if errors.Is(err, BalanceErr) {
			fmt.Println(BalanceErr)
		} else {
			fmt.Printf("Bilinmeyen bir hata olustu.", err)
		}
	} else {
		fmt.Println("Odeme Alindi!")
	}

}
