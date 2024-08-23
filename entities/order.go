package entities

type Order struct {
	Id   uint
	Name string
}

func (Order) TableName() string {
	return "saving.order"
}
