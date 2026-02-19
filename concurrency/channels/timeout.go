/* Ödev 3: Timeout (Zaman Aşımı - Select ile)
Bir fonksiyon yaz, 2 saniye uyusun sonra kanala "Cevap Geldi" yazsın.

Main fonksiyonunda select kur.

Eğer cevap 1 saniye içinde gelmezse, "Çok geç kaldın, işlem iptal" yazıp programı bitirsin.

İpucu: time.After(1 * time.Second) kullanman gerekecek.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	chan1 := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(chan1 chan string) {
		time.Sleep(time.Second * 2)
		chan1 <- "Siparis Hazır."
	}(chan1)

	select {
	case food := <-chan1:
		fmt.Println(food)
	case <-time.After(time.Second * 3):
		fmt.Println("İptal")
	}

}
