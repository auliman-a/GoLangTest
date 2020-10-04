package model

var (
	InventoryData map[string]Item
	CartItems     map[string]int
	CartData      Cart
)

type CartDataModel struct {
	TotalPrice float64
	CartItems  []*CartOutputDataModel
}

type CartOutputDataModel struct {
	Sku      string
	Quantity int
	Name     string
	Amount   float64
}

func init() {
	CartItems = make(map[string]int)

	item1 := Item{
		Sku:      "120P90",
		Name:     "Google Home",
		Quantity: 10,
		Price:    49.99,
	}

	item2 := Item{
		Sku:      "43N23P",
		Name:     "MacBook Pro",
		Quantity: 5,
		Price:    5399.99,
	}

	item3 := Item{
		Sku:      "A304SD",
		Name:     "Alexa Speaker",
		Quantity: 10,
		Price:    109.50,
	}

	item4 := Item{
		Sku:      "234234",
		Name:     "Raspberry Pi B",
		Quantity: 2,
		Price:    30.00,
	}

	InventoryData = map[string]Item{
		item1.Sku: item1,
		item2.Sku: item2,
		item3.Sku: item3,
		item4.Sku: item4,
	}

	// InventoryData = append(InventoryData, item1)
	// InventoryData = append(InventoryData, item2)
	// InventoryData = append(InventoryData, item3)
	// InventoryData = append(InventoryData, item4)
}
