package proxy

import "github.com/gofiber/fiber/v2"

var wallets = map[string]int{}

type StoreProxy struct {
	key   string
	price int
}

func NewStoreProxy(key string, price int) StoreProxy {
	return StoreProxy{
		key:   key,
		price: price,
	}
}

func (p StoreProxy) Accept(key string) bool {
	return p.key == key
}

func (p StoreProxy) Proxy(c *fiber.Ctx) error {
	ip := c.IP()
	if 
}
