package purchase_order

import (
	"context"

	purchase_order "github.com/nathanchristiawan02/icon-commission-system-backend/internal/dto/purchasing/purchase-order"
)

type PurchaseOrderUsecase interface {
	GetAll(ctx context.Context) ([]purchase_order.PurchaseOrderResponse, error)
	GetByID(ctx context.Context, id int) (*purchase_order.PurchaseOrderResponse, error)
	GetItems(ctx context.Context, id int) ([]purchase_order.PurchaseOrderItemResponse, error)

	Create(
		ctx context.Context,
		req purchase_order.CreatePurchaseOrderRequest,
	) (*purchase_order.PurchaseOrderResponse, error)

	Update(
		ctx context.Context,
		id int,
		req purchase_order.UpdatePurchaseOrderRequest,
	) (*purchase_order.PurchaseOrderResponse, error)

	Submit(ctx context.Context, id int) error

	Delete(ctx context.Context, id int) error
}