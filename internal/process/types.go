package process

import (
	"fmt"

	ah "github.com/lukejmann/carmel-backend/pkg/auctionhouse"
	"github.com/near/borsh-go"
)

func AuctionHouseDeserialize(data []byte) (ah.AuctionHouse, error) {
	var auctionHouse ah.AuctionHouse
	err := borsh.Deserialize(&auctionHouse, data)
	if err != nil {
		return ah.AuctionHouse{}, fmt.Errorf("failed to deserialize data, err: %v", err)
	}

	return auctionHouse, nil
}

type MeLog struct {
	Price        uint64 `json:"price"`
	SellerExpiry int64  `json:"seller_expiry"`
}

type MESellLog struct {
	Price        uint64 `json:"price"`
	SellerExpiry int64  `json:"seller_expiry"`
}

type MEBuyLog struct {
	Price       uint64 `json:"price"`
	BuyerExpiry int64  `json:"buyer_expiry"`
}

type MEExecuteSaleLog struct {
	Price        uint64 `json:"price"`
	BuyerExpiry  int64  `json:"buyer_expiry"`
	SellerExpiry int64  `json:"seller_expiry"`
}
