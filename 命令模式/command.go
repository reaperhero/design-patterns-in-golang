package Command

import "fmt"

type Order interface {
	execute()
}

type Stock struct {
	name     string
	quantity int
}

func (s *Stock) buy() {
	fmt.Printf("%s buy %d", s.name, s.quantity)
}

func (s *Stock) sell() {
	fmt.Printf("%s sell %d", s.name, s.quantity)
}

type BuyStock struct {
	stock Stock
}

func (b *BuyStock) execute() {
	b.stock.buy()
}

type SellStock struct {
	stock Stock
}

func (b *SellStock) execute() {
	b.stock.sell()
}

type Broker struct {
	orderList []Order
}

func (b *Broker) takeOrder(order Order) {
	b.orderList = append(b.orderList, order)
}

func (b *Broker) placeOrders() {
	for _, order := range b.orderList {
		order.execute()
	}
}
