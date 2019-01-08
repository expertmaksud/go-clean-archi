package repositories

import (
	"fmt"
	"sync"

	"github.com/expertmaksud/go-clean-archi/entities"
	"github.com/expertmaksud/go-clean-archi/repositories/interfaces"
	"github.com/jinzhu/gorm"
)

type productRepository struct {
	db *gorm.DB
}

//GetAll return all product by Tenant ID
func (repo *productRepository) GetAll(tenantID int) ([]*entities.Product, error) {
	products := make([]*entities.Product, 0)
	err := repo.db.Table("products").Where("tenant_id = ?", tenantID).Find(&products).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return products, nil
}

func (repo *productRepository) GetByID(id int64) (*entities.Product, error) {

	product := new(entities.Product)

	err := repo.db.First(&product, id).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return product, nil
}

func (repo *productRepository) Create(product *entities.Product) (uint, error) {
	res := repo.db.Create(&product)

	if res.Error != nil {
		return 0, res.Error
	}

	return product.ID, nil
}

func (repo *productRepository) Update(product *entities.Product) (*entities.Product, error) {
	res := repo.db.Save(&product)

	if res.Error != nil {
		return nil, res.Error
	}

	return product, nil
}

func (repo *productRepository) Delete(id uint) (bool, error) {
	var product *entities.Product

	err := repo.db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

var (
	productRepo *productRepository
	syncOnce    sync.Once
)

//NewProductRepository return the interface of ProductRepository
func NewProductRepository(database *gorm.DB) interfaces.IProductRepository {
	if productRepo == nil {
		syncOnce.Do(func() {
			productRepo = &productRepository{db: database}
		})
	}
	return productRepo
}
