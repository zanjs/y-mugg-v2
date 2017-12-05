package services

import (
	"fmt"
	"time"

	"github.com/zanjs/y-mugg-v2/app/middleware"
	"github.com/zanjs/y-mugg-v2/app/models"
	"github.com/zanjs/y-mugg-v2/db"
)

type (
	// SaleServices is
	SaleServices struct{}
)

// GetAll is
func (sev SaleServices) GetAll(q models.QueryParams) ([]models.Sale, models.PageModel, error) {
	var (
		sales []models.Sale
		page  models.PageModel
		err   error
	)

	page.Limit = q.Limit
	page.Offset = q.Offset

	tx := gorm.MysqlConn().Begin()

	if page.Offset == 0 {
		err = tx.Preload("Wareroom").Preload("Product").Order("id desc").Limit(page.Limit).Find(&sales).Count(&page.Count).Error
	} else {

		err = tx.Preload("Wareroom").Preload("Product").Order("id desc").Offset(page.Offset * page.Limit).Limit(page.Limit).Find(&sales).Error
	}

	if err != nil {
		tx.Rollback()
		return sales, page, err
	}

	tx.Commit()

	return sales, page, err
}

// Create is
func (sev SaleServices) Create(m models.Sale) error {
	var err error

	m.CreatedAt = time.Now()
	tx := gorm.MysqlConn().Begin()
	if err = tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// Delete is
func (sev SaleServices) Delete(m models.Sale) error {
	var err error
	tx := gorm.MysqlConn().Begin()
	if err = tx.Delete(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// WhereTime is
func (sev SaleServices) WhereTime(q models.QueryParams) ([]models.Sale, error) {
	var (
		sales []models.Sale
		err   error
	)

	queryTime := middleware.QueryStartEndTime(q)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Order("id desc").Where("created_at BETWEEN ? AND ?", queryTime.StartTime, queryTime.EndTime).Where("wareroom_id = ? AND product_id = ?", q.WareroomID, q.ProductID).Find(&sales).Error; err != nil {
		tx.Rollback()
		return sales, err
	}
	tx.Commit()

	fmt.Println("sales:", sales)

	return sales, err
}
