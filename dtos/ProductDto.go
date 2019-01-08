package dtos

import "time"

//ProductDto ...
type ProductDto struct {
	ID                 uint `gorm:"primary_key"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
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
