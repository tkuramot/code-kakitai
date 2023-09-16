package order

import (
	"context"
	"time"

	"github.com/code-kakitai/go-pkg/errors"

	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	orderDomain "github/code-kakitai/code-kakitai/domain/order"
)

type SaveOrderUseCase struct {
	orderDomainService orderDomain.OrderDomainService
	cartRepo           cartDomain.CartRepository
}

func NewSaveOrderUseCase(
	orderDomainService orderDomain.OrderDomainService,
	cartRepo cartDomain.CartRepository,
) *SaveOrderUseCase {
	return &SaveOrderUseCase{
		orderDomainService: orderDomainService,
		cartRepo:           cartRepo,
	}
}

type SaveOrderUseCaseInputDto struct {
	ProductID string
	Quantity  int
}

func (uc *SaveOrderUseCase) Run(ctx context.Context, userID string, dtos []SaveOrderUseCaseInputDto, now time.Time) (string, error) {
	// カートから商品情報を取得
	cart, err := uc.getValidCart(ctx, userID, dtos)
	if err != nil {
		return "", err
	}
	// 購入処理
	orderID, err := uc.orderDomainService.OrderProducts(ctx, cart, now)
	if err != nil {
		return "", err
	}
	// 管理者とユーザーにメール送信
	return orderID, nil
}

// カートの中身の整合性をチェック
func (uc *SaveOrderUseCase) getValidCart(ctx context.Context, userID string, dtos []SaveOrderUseCaseInputDto) (*cartDomain.Cart, error) {
	// カートから商品情報を取得
	cart, err := uc.cartRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	for _, dto := range dtos {
		pq, err := cart.QuantityByProductID(dto.ProductID)
		if err != nil {
			return nil, err
		}
		// DTOで渡ってきた数量とカートの数量が一致しない場合はエラー
		if pq != dto.Quantity {
			return nil, errors.NewError("カートの商品数が変更されています。")
		}
	}
	return cart, nil
}