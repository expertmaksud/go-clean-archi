package interfaces

import (
	"github.com/expertmaksud/go-clean-archi/dtos"
)

//IProductService interface..
type IProductService interface {
	GetAllProductsByTenant(tenentID int) ([]*dtos.ProductDto, error)
}
