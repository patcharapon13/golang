package usecases

import "clean_architecture/entities"

type OrderUseCase interface {
	CreateOrder(order entities.Order) error
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repository OrderRepository) OrderUseCase {
	return &OrderService{repo: repository}
}

func (o *OrderService) CreateOrder(order entities.Order) error {
	return o.repo.Save(order)
}
