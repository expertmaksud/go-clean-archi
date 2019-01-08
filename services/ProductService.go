package services

import (
	"sync"

	"github.com/devfeel/mapper"
	"github.com/expertmaksud/go-clean-archi/dtos"
	repo "github.com/expertmaksud/go-clean-archi/repositories/interfaces"
	"github.com/expertmaksud/go-clean-archi/services/interfaces"

	_ "github.com/devfeel/mapper" //for automap..
)

//ProductService ..
type productService struct {
	productRepo repo.IProductRepository
}

//GetAllProductsByTenant return all products by tenant id
func (service *productService) GetAllProductsByTenant(tenentID int) ([]*dtos.ProductDto, error) {
	products, err := service.productRepo.GetAll(tenentID)
	if err != nil {
		return nil, err
	}

	productDtos := make([]*dtos.ProductDto, 0)

	maperr := mapper.MapperSlice(products, &productDtos)

	if maperr != nil {
		return nil, maperr
	}

	return productDtos, nil
}

var (
	prodService *productService
	syncOnce    sync.Once
)

//NewProductService factory...
func NewProductService(productRepo repo.IProductRepository) interfaces.IProductService {
	if prodService == nil {
		syncOnce.Do(func() {
			prodService = &productService{productRepo: productRepo}
		})
	}
	return prodService
}
