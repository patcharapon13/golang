package adapters

import (
	"clean_architecture/entities"
	"clean_architecture/usecases"
	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func (g *GormOrderRepository) Save(order entities.Order) error {
	return g.db.Create(&order).Error
}

func NewGormOrderRepository(db *gorm.DB) usecases.OrderRepository {
	return &GormOrderRepository{db: db}
}
