// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_token_metadata

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type MintPrintingTokensViaTokenArgs struct {
	Supply uint64
}

func (obj MintPrintingTokensViaTokenArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Supply` param:
	err = encoder.Encode(obj.Supply)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MintPrintingTokensViaTokenArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Supply`:
	err = decoder.Decode(&obj.Supply)
	if err != nil {
		return err
	}
	return nil
}

type SetReservationListArgs struct {
	Reservations          []Reservation
	TotalReservationSpots *uint64 `bin:"optional"`
	Offset                uint64
	TotalSpotOffset       uint64
}

func (obj SetReservationListArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Reservations` param:
	err = encoder.Encode(obj.Reservations)
	if err != nil {
		return err
	}
	// Serialize `TotalReservationSpots` param (optional):
	{
		if obj.TotalReservationSpots == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.TotalReservationSpots)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Offset` param:
	err = encoder.Encode(obj.Offset)
	if err != nil {
		return err
	}
	// Serialize `TotalSpotOffset` param:
	err = encoder.Encode(obj.TotalSpotOffset)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SetReservationListArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Reservations`:
	err = decoder.Decode(&obj.Reservations)
	if err != nil {
		return err
	}
	// Deserialize `TotalReservationSpots` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.TotalReservationSpots)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Offset`:
	err = decoder.Decode(&obj.Offset)
	if err != nil {
		return err
	}
	// Deserialize `TotalSpotOffset`:
	err = decoder.Decode(&obj.TotalSpotOffset)
	if err != nil {
		return err
	}
	return nil
}

type UpdateMetadataAccountArgs struct {
	Data                *Data                  `bin:"optional"`
	UpdateAuthority     *ag_solanago.PublicKey `bin:"optional"`
	PrimarySaleHappened *bool                  `bin:"optional"`
}

func (obj UpdateMetadataAccountArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param (optional):
	{
		if obj.Data == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Data)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `UpdateAuthority` param (optional):
	{
		if obj.UpdateAuthority == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.UpdateAuthority)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `PrimarySaleHappened` param (optional):
	{
		if obj.PrimarySaleHappened == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.PrimarySaleHappened)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *UpdateMetadataAccountArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Data)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `UpdateAuthority` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.UpdateAuthority)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `PrimarySaleHappened` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.PrimarySaleHappened)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type UpdateMetadataAccountArgsV2 struct {
	Data                *DataV2                `bin:"optional"`
	UpdateAuthority     *ag_solanago.PublicKey `bin:"optional"`
	PrimarySaleHappened *bool                  `bin:"optional"`
	IsMutable           *bool                  `bin:"optional"`
}

func (obj UpdateMetadataAccountArgsV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param (optional):
	{
		if obj.Data == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Data)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `UpdateAuthority` param (optional):
	{
		if obj.UpdateAuthority == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.UpdateAuthority)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `PrimarySaleHappened` param (optional):
	{
		if obj.PrimarySaleHappened == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.PrimarySaleHappened)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `IsMutable` param (optional):
	{
		if obj.IsMutable == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.IsMutable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *UpdateMetadataAccountArgsV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Data)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `UpdateAuthority` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.UpdateAuthority)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `PrimarySaleHappened` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.PrimarySaleHappened)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `IsMutable` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.IsMutable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type CreateMetadataAccountArgs struct {
	Data      Data
	IsMutable bool
}

func (obj CreateMetadataAccountArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	// Serialize `IsMutable` param:
	err = encoder.Encode(obj.IsMutable)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CreateMetadataAccountArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	// Deserialize `IsMutable`:
	err = decoder.Decode(&obj.IsMutable)
	if err != nil {
		return err
	}
	return nil
}

type CreateMetadataAccountArgsV2 struct {
	Data      DataV2
	IsMutable bool
}

func (obj CreateMetadataAccountArgsV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Data` param:
	err = encoder.Encode(obj.Data)
	if err != nil {
		return err
	}
	// Serialize `IsMutable` param:
	err = encoder.Encode(obj.IsMutable)
	if err != nil {
		return err
	}
	return nil
}

func (obj *CreateMetadataAccountArgsV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Data`:
	err = decoder.Decode(&obj.Data)
	if err != nil {
		return err
	}
	// Deserialize `IsMutable`:
	err = decoder.Decode(&obj.IsMutable)
	if err != nil {
		return err
	}
	return nil
}

type CreateMasterEditionArgs struct {
	MaxSupply *uint64 `bin:"optional"`
}

func (obj CreateMasterEditionArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MaxSupply` param (optional):
	{
		if obj.MaxSupply == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.MaxSupply)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *CreateMasterEditionArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MaxSupply` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.MaxSupply)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type MintNewEditionFromMasterEditionViaTokenArgs struct {
	Edition uint64
}

func (obj MintNewEditionFromMasterEditionViaTokenArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Edition` param:
	err = encoder.Encode(obj.Edition)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MintNewEditionFromMasterEditionViaTokenArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Edition`:
	err = decoder.Decode(&obj.Edition)
	if err != nil {
		return err
	}
	return nil
}

type ApproveUseAuthorityArgs struct {
	NumberOfUses uint64
}

func (obj ApproveUseAuthorityArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NumberOfUses` param:
	err = encoder.Encode(obj.NumberOfUses)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ApproveUseAuthorityArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NumberOfUses`:
	err = decoder.Decode(&obj.NumberOfUses)
	if err != nil {
		return err
	}
	return nil
}

type UtilizeArgs struct {
	NumberOfUses uint64
}

func (obj UtilizeArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NumberOfUses` param:
	err = encoder.Encode(obj.NumberOfUses)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UtilizeArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NumberOfUses`:
	err = decoder.Decode(&obj.NumberOfUses)
	if err != nil {
		return err
	}
	return nil
}

type Data struct {
	Name                 string
	Symbol               string
	Uri                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator `bin:"optional"`
}

func (obj Data) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `Uri` param:
	err = encoder.Encode(obj.Uri)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `Creators` param (optional):
	{
		if obj.Creators == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *Data) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `Uri`:
	err = decoder.Decode(&obj.Uri)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `Creators` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type DataV2 struct {
	Name                 string
	Symbol               string
	Uri                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator  `bin:"optional"`
	Collection           *Collection `bin:"optional"`
	Uses                 *Uses       `bin:"optional"`
}

func (obj DataV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `Uri` param:
	err = encoder.Encode(obj.Uri)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `Creators` param (optional):
	{
		if obj.Creators == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Collection` param (optional):
	{
		if obj.Collection == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Uses` param (optional):
	{
		if obj.Uses == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Uses)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *DataV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `Uri`:
	err = decoder.Decode(&obj.Uri)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `Creators` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Collection` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Collection)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Uses` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Uses)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Uses struct {
	UseMethod UseMethod
	Remaining uint64
	Total     uint64
}

func (obj Uses) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UseMethod` param:
	err = encoder.Encode(obj.UseMethod)
	if err != nil {
		return err
	}
	// Serialize `Remaining` param:
	err = encoder.Encode(obj.Remaining)
	if err != nil {
		return err
	}
	// Serialize `Total` param:
	err = encoder.Encode(obj.Total)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Uses) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UseMethod`:
	err = decoder.Decode(&obj.UseMethod)
	if err != nil {
		return err
	}
	// Deserialize `Remaining`:
	err = decoder.Decode(&obj.Remaining)
	if err != nil {
		return err
	}
	// Deserialize `Total`:
	err = decoder.Decode(&obj.Total)
	if err != nil {
		return err
	}
	return nil
}

type Collection struct {
	Verified bool
	Key      ag_solanago.PublicKey
}

func (obj Collection) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Verified` param:
	err = encoder.Encode(obj.Verified)
	if err != nil {
		return err
	}
	// Serialize `Key` param:
	err = encoder.Encode(obj.Key)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Collection) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Verified`:
	err = decoder.Decode(&obj.Verified)
	if err != nil {
		return err
	}
	// Deserialize `Key`:
	err = decoder.Decode(&obj.Key)
	if err != nil {
		return err
	}
	return nil
}

type Creator struct {
	Address  ag_solanago.PublicKey
	Verified bool
	Share    uint8
}

func (obj Creator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `Verified` param:
	err = encoder.Encode(obj.Verified)
	if err != nil {
		return err
	}
	// Serialize `Share` param:
	err = encoder.Encode(obj.Share)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Creator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `Verified`:
	err = decoder.Decode(&obj.Verified)
	if err != nil {
		return err
	}
	// Deserialize `Share`:
	err = decoder.Decode(&obj.Share)
	if err != nil {
		return err
	}
	return nil
}

type Reservation struct {
	Address        ag_solanago.PublicKey
	SpotsRemaining uint64
	TotalSpots     uint64
}

func (obj Reservation) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `SpotsRemaining` param:
	err = encoder.Encode(obj.SpotsRemaining)
	if err != nil {
		return err
	}
	// Serialize `TotalSpots` param:
	err = encoder.Encode(obj.TotalSpots)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Reservation) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `SpotsRemaining`:
	err = decoder.Decode(&obj.SpotsRemaining)
	if err != nil {
		return err
	}
	// Deserialize `TotalSpots`:
	err = decoder.Decode(&obj.TotalSpots)
	if err != nil {
		return err
	}
	return nil
}

type ReservationV1 struct {
	Address        ag_solanago.PublicKey
	SpotsRemaining uint8
	TotalSpots     uint8
}

func (obj ReservationV1) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `SpotsRemaining` param:
	err = encoder.Encode(obj.SpotsRemaining)
	if err != nil {
		return err
	}
	// Serialize `TotalSpots` param:
	err = encoder.Encode(obj.TotalSpots)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ReservationV1) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `SpotsRemaining`:
	err = decoder.Decode(&obj.SpotsRemaining)
	if err != nil {
		return err
	}
	// Deserialize `TotalSpots`:
	err = decoder.Decode(&obj.TotalSpots)
	if err != nil {
		return err
	}
	return nil
}

type Key ag_binary.BorshEnum

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

func (value Key) String() string {
	switch value {
	case KeyUninitialized:
		return "Uninitialized"
	case KeyEditionV1:
		return "EditionV1"
	case KeyMasterEditionV1:
		return "MasterEditionV1"
	case KeyReservationListV1:
		return "ReservationListV1"
	case KeyMetadataV1:
		return "MetadataV1"
	case KeyReservationListV2:
		return "ReservationListV2"
	case KeyMasterEditionV2:
		return "MasterEditionV2"
	case KeyEditionMarker:
		return "EditionMarker"
	case KeyUseAuthorityRecord:
		return "UseAuthorityRecord"
	case KeyCollectionAuthorityRecord:
		return "CollectionAuthorityRecord"
	default:
		return ""
	}
}

type UseMethod ag_binary.BorshEnum

const (
	UseMethodBurn UseMethod = iota
	UseMethodMultiple
	UseMethodSingle
)

func (value UseMethod) String() string {
	switch value {
	case UseMethodBurn:
		return "Burn"
	case UseMethodMultiple:
		return "Multiple"
	case UseMethodSingle:
		return "Single"
	default:
		return ""
	}
}

type TokenStandard ag_binary.BorshEnum

const (
	TokenStandardNonFungible TokenStandard = iota
	TokenStandardFungibleAsset
	TokenStandardFungible
	TokenStandardNonFungibleEdition
)

func (value TokenStandard) String() string {
	switch value {
	case TokenStandardNonFungible:
		return "NonFungible"
	case TokenStandardFungibleAsset:
		return "FungibleAsset"
	case TokenStandardFungible:
		return "Fungible"
	case TokenStandardNonFungibleEdition:
		return "NonFungibleEdition"
	default:
		return ""
	}
}
