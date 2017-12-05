package services

import (
	"time"

	"github.com/zanjs/y-mugg-v2/app/models"
	"github.com/zanjs/y-mugg-v2/db"
)

type (
	// InventoryServices is
	InventoryServices struct{}
)

// GetAll is
func (sev InventoryServices) GetAll(q models.QueryParams) ([]models.Inventory, models.PageModel, error) {
	var (
		datas []models.Inventory
		page  models.PageModel
		err   error
	)

	page.Limit = q.Limit
	page.Offset = q.Offset

	tx := gorm.MysqlConn().Begin()

	if page.Offset == 0 {
		err = tx.Preload("Wareroom").Preload("Product").Order("id desc").Limit(page.Limit).Find(&datas).Count(&page.Count).Error
	} else {

		err = tx.Preload("Wareroom").Preload("Product").Order("id desc").Offset(page.Offset * page.Limit).Limit(page.Limit).Find(&datas).Error
	}

	if err != nil {
		tx.Rollback()
		return datas, page, err
	}

	tx.Commit()

	return datas, page, err
}

// Create is
func (sev InventoryServices) Create(m models.Inventory) error {
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

// Update is
func (sev InventoryServices) Update(m models.Inventory) error {
	var err error

	m.UpdatedAt = time.Now()

	tx := gorm.MysqlConn().Begin()
	if err = tx.Save(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return err
}

// GetByPId is
func (sev InventoryServices) GetByPId(m models.Inventory) (models.Inventory, error) {
	var (
		inventory models.Inventory
		err       error
	)

	tx := gorm.MysqlConn().Begin()
	if err = tx.Where("wareroom_id = ? AND product_id = ?", m.WareroomID, m.ProductID).Find(&inventory).Error; err != nil {
		tx.Rollback()
		return inventory, err
	}
	tx.Commit()

	return inventory, err
}
