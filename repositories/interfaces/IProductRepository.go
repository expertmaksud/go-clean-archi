package interfaces

import (
	"github.com/expertmaksud/go-clean-archi/entities"
)

//IProductRepository interface..
type IProductRepository interface {
	GetAll(tenantID int) ([]*entities.Product, error)
	GetByID(id int64) (*entities.Product, error)
	Create(product *entities.Product) (uint, error)
	Update(product *entities.Product) (*entities.Product, error)
	Delete(id uint) (bool, error)
}
