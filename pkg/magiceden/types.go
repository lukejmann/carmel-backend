// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import ag_binary "github.com/gagliardetto/binary"

type BidType ag_binary.BorshEnum

const (
	BidTypePublicSale BidType = iota
	BidTypePrivateSale
)

func (value BidType) String() string {
	switch value {
	case BidTypePublicSale:
		return "PublicSale"
	case BidTypePrivateSale:
		return "PrivateSale"
	default:
		return ""
	}
}