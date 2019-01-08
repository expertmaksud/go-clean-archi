package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Product represent the entity in the DB
type Product struct {
	gorm.Model
	Batch              string
	Channel            string
	CompanyName        string
	CreatorUserID      int
	DeleterUserID      int
	IsChecked          bool
	IsDeleted          bool
	ProductCode        string
	ProductName        string
	ProductType        string
	TenantID           int
	Timestamp          time.Time
	TryCount           int
	UserID             int64
	Validity           time.Time
	LastModifierUserID int
}
