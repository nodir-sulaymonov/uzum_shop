package models

// import (
// 	"context"
// 	"fmt"

// 	"github.com/Shemistan/uzum_shop/internal/service/shop1"
// 	// "github.com/Shemistan/uzum_shop/internal/models"
// )

// func (s *shop1.IShopSystemService) AddProductToBasketService(ctx context.Context, req *models.AddProductToBasketModel) error {
// 	userId, err := s.GetUserIdFromLoginServ(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	count, err := s.storage.GetProductCountStorage(ctx, req.ProductId)
// 	if err != nil {
// 		return err
// 	}

// 	if count < int(req.Count) {
// 		return fmt.Errorf("sorry, not enough inventory. Only %v left in stock", count)
// 	}

// 	req.UserId = userId

// 	err = s.storage.AddProductToBasketStorage(ctx, req)
// 	if err != nil {
// 		return err
// 	}

// 	count = count - int(req.Count)

// 	res, err := s.storage.CalculateProductCountStorage(ctx, int(req.ProductId), count)
// 	if err != nil {
// 		return err
// 	}
// 	if res == 0 {
// 		return fmt.Errorf("%v rows affected", res)
// 	}

// 	return nil
// }

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
