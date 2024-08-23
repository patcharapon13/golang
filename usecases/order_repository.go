package usecases

import "clean_architecture/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
