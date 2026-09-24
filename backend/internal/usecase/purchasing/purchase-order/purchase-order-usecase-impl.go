package purchase_order

import (
	"context"
	"errors"
	"fmt"
	"time"

	dto "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/purchasing/purchase-order"
	entity "github.com/nathanchristiawan02/icon-commission-system-backend/internal/entity/purchasing/purchase-order"
	repository "github.com/nathanchristiawan02/icon-commission-system-backend/internal/repository/purchasing/purchase-order"
)

type purchaseOrderUsecase struct {
	repository repository.PurchaseOrderRepository
}

func NewPurchaseOrderUsecase(
	repository repository.PurchaseOrderRepository,
) PurchaseOrderUsecase {
	return &purchaseOrderUsecase{
		repository: repository,
	}
}

// =========================
// GET ALL
// =========================

func (u *purchaseOrderUsecase) GetAll(
	ctx context.Context,
) ([]dto.PurchaseOrderResponse, error) {

	purchaseOrders, err := u.repository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.PurchaseOrderResponse, 0, len(purchaseOrders))

	for _, po := range purchaseOrders {

		items, err := u.repository.FindItems(uint(po.ID))
		if err != nil {
			return nil, err
		}

		response := mapPurchaseOrderResponse(po)
		response.Total = calculateTotal(items)

		responses = append(responses, response)
	}

	return responses, nil
}

// =========================
// GET BY ID
// =========================

func (u *purchaseOrderUsecase) GetByID(
	ctx context.Context,
	id int,
) (*dto.PurchaseOrderResponse, error) {

	if id <= 0 {
		return nil, errors.New("invalid purchase order id")
	}

	po, err := u.repository.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	items, err := u.repository.FindItems(uint(id))
	if err != nil {
		return nil, err
	}

	response := mapPurchaseOrderResponse(*po)
	response.Total = calculateTotal(items)

	return &response, nil
}

// =========================
// GET ITEMS
// =========================

func (u *purchaseOrderUsecase) GetItems(
	ctx context.Context,
	id int,
) ([]dto.PurchaseOrderItemResponse, error) {

	if id <= 0 {
		return nil, errors.New("invalid purchase order id")
	}

	items, err := u.repository.FindItems(uint(id))
	if err != nil {
		return nil, err
	}

	responses := make(
		[]dto.PurchaseOrderItemResponse,
		0,
		len(items),
	)

	for _, item := range items {
		responses = append(
			responses,
			dto.PurchaseOrderItemResponse{
				ID:              item.ID,
				PurchaseOrderID: item.PurchaseOrderID,
				ProductID:       item.ProductID,
				ProductName:     item.ProductName,
				Brand:           item.Brand,
				Quantity:        item.Quantity,
				UnitPrice:       item.UnitPrice,
				Subtotal:        item.Subtotal,
				Remark:          item.Remark,
			},
		)
	}

	return responses, nil
}

// =========================
// CREATE
// =========================

func (u *purchaseOrderUsecase) Create(
	ctx context.Context,
	req dto.CreatePurchaseOrderRequest,
) (*dto.PurchaseOrderResponse, error) {

	if req.VendorName <= "" {
		return nil, errors.New("vendor_name is required")
	}

	if len(req.Items) == 0 {
		return nil, errors.New(
			"purchase order must have at least one item",
		)
	}

	orderDate, err := time.Parse(
		"2006-01-02",
		req.OrderDate,
	)
	if err != nil {
		return nil, errors.New(
			"invalid order_date format, use YYYY-MM-DD",
		)
	}

	documentNo := req.DocumentNo

	if documentNo == "" {
		documentNo = fmt.Sprintf(
			"PO-%s",
			time.Now().Format("20060102-150405"),
		)
	}

	po := &entity.PurchaseOrder{
		DocumentName: "Purchase Order",
		DocumentNo:   documentNo,
		VendorName:   req.VendorName,
		OrderDate:    orderDate,
		Status:       entity.StatusDraft,
		Remark:       req.Remark,
	}

	items := make(
		[]entity.PurchaseOrderItem,
		0,
		len(req.Items),
	)

	for _, item := range req.Items {

		if item.ProductID <= 0 {
			return nil, errors.New(
				"product_id is required",
			)
		}

		if item.Quantity <= 0 {
			return nil, errors.New(
				"quantity must be greater than 0",
			)
		}

		if item.UnitPrice < 0 {
			return nil, errors.New(
				"unit_price cannot be negative",
			)
		}

		items = append(
			items,
			entity.PurchaseOrderItem{
				ProductID:   item.ProductID,
				ProductName: item.ProductName,
				Brand:       item.Brand,
				Quantity:    item.Quantity,
				UnitPrice:   item.UnitPrice,
				Subtotal:    item.Quantity * item.UnitPrice,
				Remark:      item.Remark,
			},
		)
	}

	err = u.repository.Create(po, items)
	if err != nil {
		return nil, err
	}

	// Hitung total berdasarkan item yang baru dibuat
	response := mapPurchaseOrderResponse(*po)
	response.Total = calculateTotal(items)

	return &response, nil
}

// =========================
// UPDATE
// =========================

func (u *purchaseOrderUsecase) Update(
	ctx context.Context,
	id int,
	req dto.UpdatePurchaseOrderRequest,
) (*dto.PurchaseOrderResponse, error) {

	if id <= 0 {
		return nil, errors.New(
			"invalid purchase order id",
		)
	}

	existing, err := u.repository.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	if existing.Status != entity.StatusDraft {
		return nil, errors.New(
			"only draft purchase order can be updated",
		)
	}

	if req.VendorName <= "" {
		return nil, errors.New(
			"vendor_name is required",
		)
	}

	if len(req.Items) == 0 {
		return nil, errors.New(
			"purchase order must have at least one item",
		)
	}

	orderDate, err := time.Parse(
		"2006-01-02",
		req.OrderDate,
	)
	if err != nil {
		return nil, errors.New(
			"invalid order_date format, use YYYY-MM-DD",
		)
	}

	documentNo := req.DocumentNo

	if documentNo == "" {
		documentNo = existing.DocumentNo
	}

	existing.DocumentNo = documentNo
	existing.VendorName = req.VendorName
	existing.OrderDate = orderDate
	existing.Remark = req.Remark

	items := make(
		[]entity.PurchaseOrderItem,
		0,
		len(req.Items),
	)

	for _, item := range req.Items {

		if item.ProductID <= 0 {
			return nil, errors.New(
				"product_id is required",
			)
		}

		if item.Quantity <= 0 {
			return nil, errors.New(
				"quantity must be greater than 0",
			)
		}

		if item.UnitPrice < 0 {
			return nil, errors.New(
				"unit_price cannot be negative",
			)
		}

		items = append(
			items,
			entity.PurchaseOrderItem{
				ProductID:   item.ProductID,
				ProductName: item.ProductName,
				Brand:       item.Brand,
				Quantity:    item.Quantity,
				UnitPrice:   item.UnitPrice,
				Subtotal:    item.Quantity * item.UnitPrice,
				Remark:      item.Remark,
			},
		)
	}

	err = u.repository.Update(existing, items)
	if err != nil {
		return nil, err
	}

	// Karena items sudah merupakan data terbaru,
	// total bisa langsung dihitung dari items.
	response := mapPurchaseOrderResponse(*existing)
	response.Total = calculateTotal(items)

	return &response, nil
}

// =========================
// SUBMIT
// =========================

// Submit changes the Purchase Order status from draft to ordered.
func (u *purchaseOrderUsecase) Submit(
	ctx context.Context,
	id int,
) error {

	if id <= 0 {
		return errors.New(
			"invalid purchase order id",
		)
	}

	purchaseOrder, err := u.repository.FindByID(uint(id))
	if err != nil {
		return err
	}

	if purchaseOrder.Status != entity.StatusDraft {
		return errors.New(
			"only draft purchase order can be submitted",
		)
	}

	err = u.repository.UpdateStatus(
		uint(id),
		entity.StatusOrdered,
	)

	if err != nil {
		return err
	}

	return nil
}

// =========================
// DELETE
// =========================

func (u *purchaseOrderUsecase) Delete(
	ctx context.Context,
	id int,
) error {

	if id <= 0 {
		return errors.New(
			"invalid purchase order id",
		)
	}

	purchaseOrder, err := u.repository.FindByID(uint(id))
	if err != nil {
		return err
	}

	if purchaseOrder.Status != entity.StatusDraft {
		return errors.New(
			"only draft purchase order can be deleted",
		)
	}

	return u.repository.Delete(uint(id))
}

// =========================
// CALCULATE TOTAL
// =========================

func calculateTotal(
	items []entity.PurchaseOrderItem,
) float64 {

	var total float64

	for _, item := range items {
		total += item.Subtotal
	}

	return total
}

// =========================
// MAPPER
// =========================

func mapPurchaseOrderResponse(
	po entity.PurchaseOrder,
) dto.PurchaseOrderResponse {

	return dto.PurchaseOrderResponse{
		ID:           po.ID,
		DocumentName: po.DocumentName,
		DocumentNo:   po.DocumentNo,
		VendorName:   po.VendorName,
		OrderDate:    po.OrderDate.Format("2006-01-02"),
		Status:       po.Status,
		Remark:       po.Remark,
		CreatedAt:    po.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    po.UpdatedAt.Format(time.RFC3339),
	}
}