package Command

import "testing"

func TestCommand_Execute(t *testing.T) {
	stock := Stock{}
	buystock:=BuyStock{stock: stock}
	sellstock:=SellStock{stock: stock}
	broker:=Broker{}
	broker.takeOrder(&buystock)
	broker.takeOrder(&sellstock)
	broker.placeOrders()
}
