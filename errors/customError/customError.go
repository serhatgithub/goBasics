package main

import (
	"errors"
	"fmt"
)

type User struct {
	Age      int
	Username string
}

type AgeError struct {
	Age int
}

func (a *AgeError) Error() string {
	return fmt.Sprintf("Yasiniz: %d, en az 18 yasinda olmalisiniz.", a.Age)
}

func ValidateUser(u User) error {
	if u.Username == "" {
		return errors.New("Kullanici bos olamaz")
	}

	if u.Age < 18 {
		return &AgeError{Age: u.Age}
	}

	return nil
}

func main() {
	user1 := User{Username: "", Age: 24}
	err1 := ValidateUser(user1)
	fmt.Println("Deneme 1 Sonuç:", err1)

	user2 := User{Username: "Serhat", Age: 16}
	err2 := ValidateUser(user2)
	fmt.Println("Deneme 2 Sonuç:", err2)

	user3 := User{Username: "Ahmet", Age: 25}
	err3 := ValidateUser(user3)
	fmt.Println("Deneme 3 Sonuç:", err3)
}
