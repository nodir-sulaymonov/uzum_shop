package shop1

import (
	shop_system "github.com/Shemistan/uzum_shop/internal/service/shop1"
)

type Shop struct {
	shop1.UnimplementedShopServer
	ShopService shop_system.IShopSystemService
}