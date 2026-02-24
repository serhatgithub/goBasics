package main

import (
	"errors"
	"fmt"
)

var dbError = errors.New("DB baglantisi koptu")

func checkDB() error {
	return dbError
}

func getUser() error {
	if err := checkDB(); err != nil {
		return fmt.Errorf("Kullanici getirilemedi %w", err)
	}
	return nil
}

func main() {
	err := getUser()
	if err != nil {
		if errors.Is(err, dbError) {
			fmt.Println("Sorun veritabanından kaynaklanıyor!")
		} else {
			fmt.Println("Bilinmeyen bir hata oluştu:", err)
		}
	}

}
