package main

import (
	"sync"

	"github.com/expertmaksud/go-clean-archi/controllers"
	repo "github.com/expertmaksud/go-clean-archi/repositories"
	service "github.com/expertmaksud/go-clean-archi/services"
	"github.com/jinzhu/gorm"
)

//IInjectionContainer ...
type IInjectionContainer interface {
	ProvideProductController() *controllers.ProducrController
}

//container struct ....
type container struct {
	db *gorm.DB
}

//ProvideProductController provide product controller
func (con *container) ProvideProductController() *controllers.ProducrController {
	productRepository := repo.NewProductRepository(con.db)
	productService := service.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	return productController
}

var (
	k             *container
	containerOnce sync.Once
)

//NewInjectionContainer retunrs implement of IInjectionContainer interface
func NewInjectionContainer(db *gorm.DB) IInjectionContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &container{db: db}
		})
	}
	return k

}
