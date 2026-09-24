package purchase_order

import (
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/purchasing/purchase-order"
	"gorm.io/gorm"
)

type PurchaseOrderRepository interface {
	FindAll() ([]entity.PurchaseOrder, error)
	FindByID(id uint) (*entity.PurchaseOrder, error)
	FindItems(purchaseOrderID uint) ([]entity.PurchaseOrderItem, error)

	Create(po *entity.PurchaseOrder, items []entity.PurchaseOrderItem) error
	Update(po *entity.PurchaseOrder, items []entity.PurchaseOrderItem) error
	UpdateStatus(id uint, status string) error
	Delete(id uint) error
}

type purchaseOrderRepository struct {
	db *gorm.DB
}

func NewPurchaseOrderRepository(db *gorm.DB) PurchaseOrderRepository {
	return &purchaseOrderRepository{
		db: db,
	}
}

func (r *purchaseOrderRepository) FindAll() ([]entity.PurchaseOrder, error) {
	var purchaseOrders []entity.PurchaseOrder

	err := r.db.
		Find(&purchaseOrders).Error

	return purchaseOrders, err
}

func (r *purchaseOrderRepository) FindByID(id uint) (*entity.PurchaseOrder, error) {
	var purchaseOrder entity.PurchaseOrder

	err := r.db.
		First(&purchaseOrder, id).Error

	if err != nil {
		return nil, err
	}

	return &purchaseOrder, nil
}

func (r *purchaseOrderRepository) FindItems(purchaseOrderID uint) ([]entity.PurchaseOrderItem, error) {
	var items []entity.PurchaseOrderItem

	err := r.db.
		Where("purchase_order_id = ?", purchaseOrderID).
		Find(&items).Error

	return items, err
}

func (r *purchaseOrderRepository) Create(
	po *entity.PurchaseOrder,
	items []entity.PurchaseOrderItem,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(po).Error; err != nil {
			return err
		}

		for i := range items {
			items[i].PurchaseOrderID = po.ID

			if err := tx.Create(&items[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *purchaseOrderRepository) Update(
	po *entity.PurchaseOrder,
	items []entity.PurchaseOrderItem,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(&entity.PurchaseOrder{}).
			Where("id = ?", po.ID).
			Updates(map[string]interface{}{
				"document_name": po.DocumentName,
				"document_no":   po.DocumentNo,
				"vendor_id":     po.VendorID,
				"order_date":    po.OrderDate,
				"status":        po.Status,
				"remark":        po.Remark,
			}).Error; err != nil {
			return err
		}

		if err := tx.
			Where("purchase_order_id = ?", po.ID).
			Delete(&entity.PurchaseOrderItem{}).Error; err != nil {
			return err
		}

		for i := range items {
			items[i].PurchaseOrderID = po.ID

			if err := tx.Create(&items[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *purchaseOrderRepository) UpdateStatus(id uint, status string) error {
	return r.db.
		Model(&entity.PurchaseOrder{}).
		Where("id = ?", id).
		Update("status", status).
		Error
}

func (r *purchaseOrderRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.
			Where("purchase_order_id = ?", id).
			Delete(&entity.PurchaseOrderItem{}).Error; err != nil {
			return err
		}

		if err := tx.
			Delete(&entity.PurchaseOrder{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}