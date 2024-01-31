package shop1

import (
	"context"

	"github.com/Shemistan/uzum_shop/internal/models"
	repo "github.com/Shemistan/uzum_shop/internal/storage"
	loginPb "github.com/Shemistan/uzum_shop/pkg/login_v1"
)

type IShopSystemService interface {
	GetUserIdFromLoginServ(ctx context.Context) (int, error)
	GetProductService(ctx context.Context, prodId uint32) (*models.Product, error)
	GetAllProductsService(ctx context.Context) ([]*models.Product, error)
	AddProductToBasketService(ctx context.Context, req *models.AddProductToBasketModel) error
	UpdateBasketService(ctx context.Context, req *models.AddProductToBasketModel) error
	DeleteBasketService(ctx context.Context, req *models.DeleteFomBasked) error
	GetBasketService(ctx context.Context) ([]*models.BasketItem, error)
	CreateOrderService(ctx context.Context, req *models.Order) (int32, error)
	CancelOrderService(ctx context.Context, orderId uint32) error
}

type shopSystemService struct {
	storage     repo.IStorage
	loginClient loginPb.LoginV1Client
}

func NewShopSystemService(storage repo.IStorage, loginClient loginPb.LoginV1Client) IShopSystemService {
	return &shopSystemService{
		storage:     storage,
		loginClient: loginClient,
	}
}
