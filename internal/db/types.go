package db

import (
	"encoding/json"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/near/borsh-go"
)

type WatchedWallets struct {
	WatchedWallets []string `json:"watchedWallets" firestore:"watchedWallets"`
}

type OwnedTokens struct {
	Mint string `json:"mint" firestore:"mint"`
	Size int    `json:"size" firestore:"size"`
}

type Wallet struct {
	Pubkey             string `json:"pubkey" firestore:"pubkey"`
	OldestProcessedSig string `json:"oldestProcessSig" firestore:"oldestProcessSig"`
	LatestProcessedSig string `json:"latestProcessSig" firestore:"latestProcessSig"`
	SyncedToSlot       int    `json:"syncedToSlot" firestore:"syncedToSlot"`
	SyncToSlot         int    `json:"syncToSlot" firestore:"syncToSlot"`
}

type Collection struct {
	ID                   string   `json:"id" firestore:"id"`
	Type                 string   `json:"type" firestore:"type"`
	User                 string   `json:"user" firestore:"user"`
	Mints                []string `json:"mints" firestore:"mints"`
	Name                 string   `json:"name" firestore:"name"`
	Description          string   `json:"description" firestore:"description"`
	Image                string   `json:"image" firestore:"image"`
	Discord              string   `json:"discord" firestore:"discord"`
	Website              string   `json:"website" firestore:"website"`
	Twitter              string   `json:"twitter" firestore:"twitter"`
	CreatedDate          string   `json:"createdDate" firestore:"createdDate"`
	UpdatedDate          string   `json:"updatedDate" firestore:"updatedDate"`
	CandyMachine         string   `json:"candyMachine" firestore:"candyMachine"`
	OnChainCollectionKey string   `json:"onChainCollectionKey" firestore:"onChainCollectionKey"`
}

type ArrayValue struct {
	Values []Value `json:"values"`
}
type Value struct {
	MapValue     MapValue     `json:"mapValue,omitempty"`
	StringValue  StringValue  `json:"stringValue,omitempty"`
	IntegerValue IntegerValue `json:"integerValue,omitempty"`
	ArrayValue   ArrayValue   `json:"arrayValue,omitempty"`
}
type IntegerValue struct {
	IntegerValue string `json:"integerValue"`
}
type StringValue struct {
	StringValue string `json:"stringValue"`
}
type MapValue struct {
	Fields interface{} `json:"fields"`
}

type FirestoreTriggerCollection struct {
	ID    StringValue `json:"id" firestore:"id"`
	Mints ArrayValue  `json:"mints" firestore:"mints"`
}

const EDITION_MARKER_BIT_SIZE uint64 = 248

type Key borsh.Enum

const (
	KeyUninitialized Key = iota
	KeyEditionV1
	KeyMasterEditionV1
	KeyReservationListV1
	KeyMetadataV1
	KeyReservationListV2
	KeyMasterEditionV2
	KeyEditionMarker
	KeyUseAuthorityRecord
	KeyCollectionAuthorityRecord
)

type Creator struct {
	Address  solana.PublicKey
	Verified bool
	Share    uint8
}

type FirestoreCreator struct {
	Address  string `firestore:"address"`
	Verified bool   `firestore:"verified"`
	Share    uint8  `firestore:"share"`
}

type Data struct {
	Name                 string     `json:"name"`
	Symbol               string     `json:"symbol"`
	Uri                  string     `json:"uri"`
	SellerFeeBasisPoints uint16     `json:"sellerFeeBasisPoints"`
	Creators             *[]Creator `json:"creators"`
}

type DataV2 struct {
	Name                 string     `json:"name"`
	Symbol               string     `json:"symbol"`
	Uri                  string     `json:"uri"`
	SellerFeeBasisPoints uint16     `json:"sellerFeeBasisPoints"`
	Creators             *[]Creator `json:"creators"`
}

type Metadata struct {
	Key                 Key                 `json:"key"`
	UpdateAuthority     solana.PublicKey    `json:"updateAuthority"`
	Mint                solana.PublicKey    `json:"mint"`
	Data                DataV2              `json:"data"`
	PrimarySaleHappened bool                `json:"primarySaleHappened"`
	IsMutable           bool                `json:"isMutable"`
	EditionNonce        *uint8              `json:"editionNonce"`
	TokenStandard       *uint8              `json:"tokenStandard"`
	Collection          *MetadataCollection `json:"collection"`
	Uses                *Uses               `json:"uses"`
}

type MintData struct {
	Address string
	// Metadata *Metadata `json:"metadata"`
}

type MintDataExtended struct {
	Extended interface{} `json:"interface"`
}

type FlexString string

// UnmarshalJSON takes either a int or string and returns the string value
func (fi *FlexString) UnmarshalJSON(b []byte) error {
	var s string
	sErr := json.Unmarshal(b, &s)
	if sErr != nil {
		var i int
		iErr := json.Unmarshal(b, &i)
		if iErr != nil {
			return fmt.Errorf("unable to unmarshal string or int iErr: %v sErr: %v", iErr, sErr)
		}
		s = fmt.Sprintf("%d", i)
	}
	*fi = FlexString(s)
	return nil
}

type ExtendedAttribute struct {
	TraitType string     `json:"trait_type" firestore:"trait_type"`
	Value     FlexString `json:"value" firestore:"value"`
}

type MetadataExtended struct {
	Name        string                      `json:"name" firestore:"name"`
	Symbol      string                      `json:"symbol" firestore:"symbol"`
	Image       string                      `json:"image" firestore:"image"`
	Description string                      `json:"description" firestore:"description"`
	Attributes  []ExtendedAttribute         `json:"attributes" firestore:"attributes"`
	Collection  *MetadataExtendedCollection `json:"collection" firestore:"collection"`
}

type MetadataExtendedCollection struct {
	Name string `json:"name"`
}

// type MintDataEx

type MetadataCollection struct {
	Verified uint8            `json:"verified"`
	Key      solana.PublicKey `json:"key"`
}

type Uses struct {
	UseMethod UseMethod `json:"useMethod"`
	Remaining uint64    `json:"remaining"`
	Total     uint64    `json:"total"`
}

type UseMethod borsh.Enum

const (
	Burn UseMethod = iota
	Multiple
	Single
)

type MasterEditionV2 struct {
	Key       Key     `json:"key"`
	Supply    uint64  `json:"supply"`
	MaxSupply *uint64 `json:"maxSupply"`
}

type Order struct {
	DelegatePDA         string `json:"delegatePDA" firestore:"delegatePDA"`
	DelegatePDABumpSeed int64  `json:"delegatePDABumpSeed" firestore:"delegatePDABumpSeed"`
	Maker               string `json:"maker" firestore:"maker"`
	Taker               string `json:"taker" firestore:"taker"`
	Source              string `json:"source" firestore:"source"`
	AssetMint           string `json:"assetMint" firestore:"assetMint"`
	AssetSize           int64  `json:"assetSize" firestore:"assetSize"`
	PaymentMint         string `json:"paymentMint" firestore:"paymentMint"`
	PaymentBaseSize     int64  `json:"paymentBaseSize" firestore:"paymentBaseSize"`
	CreatorsFeeSize     int64  `json:"creatorsFeeSize" firestore:"creatorsFeeSize"`
	Curator             string `json:"curator" firestore:"curator"`
	CuratorBPS          int64  `json:"curatorBPS" firestore:"curatorBPS"`
	CuratorFeeSize      int64  `json:"curatorFeeSize" firestore:"curatorFeeSize"`
	ListingDate         int64  `json:"listingDate" firestore:"listingDate"`
	ExpirationDate      int64  `json:"expirationDate" firestore:"expirationDate"`
	Side                int64  `json:"side" firestore:"side"`
	Revoked             bool   `json:"revoked" firestore:"revoked"`
	RevokedDate         int64  `json:"revokedDate" firestore:"revokedDate"`
	Matched             bool   `json:"matched" firestore:"matched"`
	MatchedWith         string `json:"matchedWith" firestore:"matchedWith"`
	MatchedDate         int64  `json:"matchedDate" firestore:"matchedDate"`
	ProgramID           string `json:"programID" firestore:"programID"`
}

type WalletsProcessed struct {
	ProcessedSigs []string `json:"processedSigs"`
}
