package shop1

import (
	"context"

	pb "github.com/Shemistan/pkg/shopV1"
	"github.com/Shemistan/uzum_shop/internal/models"
	"github.com/golang/protobuf/ptypes/empty"
)

func (serv *Shop) AddProductToBasket(ctx context.Context, req *pb.AddProductToBasket_Request) (*empty.Empty, error) {

	addProdToBask := &models.AddProductToBasketModel{
		ProductId: req.ProductId,
		Count:     req.Count,
	}

	err := serv.ShopService.AddProductToBasketService(ctx, addProdToBask)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}