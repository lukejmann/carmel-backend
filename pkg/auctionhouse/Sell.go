package auction_house

import (
	"errors"

	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Sell is the `sell` instruction.
type Sell struct {
	TradeStateBump      *uint8
	FreeTradeStateBump  *uint8
	ProgramAsSignerBump *uint8
	BuyerPrice          *uint64
	TokenSize           *uint64
	// Expiry              *uint8
	// ExpiryTwo           *uint8

	// [0] = [] wallet
	//
	// [1] = [WRITE] tokenAccount
	//
	// [2] = [] metadata
	//
	// [3] = [] authority
	//
	// [4] = [] auctionHouse
	//
	// [5] = [WRITE] auctionHouseFeeAccount
	//
	// [6] = [WRITE] sellerTradeState
	//
	// [7] = [WRITE] freeSellerTradeState
	//
	// [8] = [] tokenProgram
	//
	// [9] = [] systemProgram
	//
	// [10] = [] programAsSigner
	//
	// [11] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSellInstructionBuilder creates a new `Sell` instruction builder.
func NewSellInstructionBuilder() *Sell {
	nd := &Sell{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
	return nd
}

// SetTradeStateBump sets the "tradeStateBump" parameter.
func (inst *Sell) SetTradeStateBump(tradeStateBump uint8) *Sell {
	inst.TradeStateBump = &tradeStateBump
	return inst
}

// SetFreeTradeStateBump sets the "freeTradeStateBump" parameter.
func (inst *Sell) SetFreeTradeStateBump(freeTradeStateBump uint8) *Sell {
	inst.FreeTradeStateBump = &freeTradeStateBump
	return inst
}

// SetProgramAsSignerBump sets the "programAsSignerBump" parameter.
func (inst *Sell) SetProgramAsSignerBump(programAsSignerBump uint8) *Sell {
	inst.ProgramAsSignerBump = &programAsSignerBump
	return inst
}

// SetBuyerPrice sets the "buyerPrice" parameter.
func (inst *Sell) SetBuyerPrice(buyerPrice uint64) *Sell {
	inst.BuyerPrice = &buyerPrice
	return inst
}

// SetTokenSize sets the "tokenSize" parameter.
func (inst *Sell) SetTokenSize(tokenSize uint64) *Sell {
	inst.TokenSize = &tokenSize
	return inst
}

// SetWalletAccount sets the "wallet" account.
func (inst *Sell) SetWalletAccount(wallet ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(wallet)
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *Sell) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenAccountAccount sets the "tokenAccount" account.
func (inst *Sell) SetTokenAccountAccount(tokenAccount ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenAccount).WRITE()
	return inst
}

// GetTokenAccountAccount gets the "tokenAccount" account.
func (inst *Sell) GetTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *Sell) SetMetadataAccount(metadata ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *Sell) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Sell) SetAuthorityAccount(authority ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Sell) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAuctionHouseAccount sets the "auctionHouse" account.
func (inst *Sell) SetAuctionHouseAccount(auctionHouse ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(auctionHouse)
	return inst
}

// GetAuctionHouseAccount gets the "auctionHouse" account.
func (inst *Sell) GetAuctionHouseAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAuctionHouseFeeAccountAccount sets the "auctionHouseFeeAccount" account.
func (inst *Sell) SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(auctionHouseFeeAccount).WRITE()
	return inst
}

// GetAuctionHouseFeeAccountAccount gets the "auctionHouseFeeAccount" account.
func (inst *Sell) GetAuctionHouseFeeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSellerTradeStateAccount sets the "sellerTradeState" account.
func (inst *Sell) SetSellerTradeStateAccount(sellerTradeState ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(sellerTradeState).WRITE()
	return inst
}

// GetSellerTradeStateAccount gets the "sellerTradeState" account.
func (inst *Sell) GetSellerTradeStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetFreeSellerTradeStateAccount sets the "freeSellerTradeState" account.
func (inst *Sell) SetFreeSellerTradeStateAccount(freeSellerTradeState ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(freeSellerTradeState).WRITE()
	return inst
}

// GetFreeSellerTradeStateAccount gets the "freeSellerTradeState" account.
func (inst *Sell) GetFreeSellerTradeStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Sell) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Sell) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Sell) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Sell) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetProgramAsSignerAccount sets the "programAsSigner" account.
func (inst *Sell) SetProgramAsSignerAccount(programAsSigner ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(programAsSigner)
	return inst
}

// GetProgramAsSignerAccount gets the "programAsSigner" account.
func (inst *Sell) GetProgramAsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetRentAccount sets the "rent" account.
func (inst *Sell) SetRentAccount(rent ag_solanago.PublicKey) *Sell {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *Sell) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

func (inst Sell) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Sell,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Sell) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Sell) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.TradeStateBump == nil {
			return errors.New("TradeStateBump parameter is not set")
		}
		if inst.FreeTradeStateBump == nil {
			return errors.New("FreeTradeStateBump parameter is not set")
		}
		if inst.ProgramAsSignerBump == nil {
			return errors.New("ProgramAsSignerBump parameter is not set")
		}
		if inst.BuyerPrice == nil {
			return errors.New("BuyerPrice parameter is not set")
		}
		if inst.TokenSize == nil {
			return errors.New("TokenSize parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AuctionHouse is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AuctionHouseFeeAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SellerTradeState is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.FreeSellerTradeState is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.ProgramAsSigner is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *Sell) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Sell")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("     TradeStateBump", *inst.TradeStateBump))
						paramsBranch.Child(ag_format.Param(" FreeTradeStateBump", *inst.FreeTradeStateBump))
						paramsBranch.Child(ag_format.Param("ProgramAsSignerBump", *inst.ProgramAsSignerBump))
						paramsBranch.Child(ag_format.Param("         BuyerPrice", *inst.BuyerPrice))
						paramsBranch.Child(ag_format.Param("          TokenSize", *inst.TokenSize))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=12]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("              wallet", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("               token", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            metadata", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        auctionHouse", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("     auctionHouseFee", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("    sellerTradeState", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("freeSellerTradeState", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("        tokenProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("     programAsSigner", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                rent", inst.AccountMetaSlice.Get(11)))
					})
				})
		})
}

func (obj Sell) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TradeStateBump` param:
	err = encoder.Encode(obj.TradeStateBump)
	if err != nil {
		return err
	}
	// Serialize `FreeTradeStateBump` param:
	err = encoder.Encode(obj.FreeTradeStateBump)
	if err != nil {
		return err
	}
	// Serialize `ProgramAsSignerBump` param:
	err = encoder.Encode(obj.ProgramAsSignerBump)
	if err != nil {
		return err
	}
	// Serialize `BuyerPrice` param:
	err = encoder.Encode(obj.BuyerPrice)
	if err != nil {
		return err
	}
	// Serialize `TokenSize` param:
	err = encoder.Encode(obj.TokenSize)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Sell) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TradeStateBump`:
	err = decoder.Decode(&obj.TradeStateBump)
	if err != nil {
		return err
	}
	// Deserialize `FreeTradeStateBump`:
	err = decoder.Decode(&obj.FreeTradeStateBump)
	if err != nil {
		return err
	}
	// Deserialize `ProgramAsSignerBump`:
	err = decoder.Decode(&obj.ProgramAsSignerBump)
	if err != nil {
		return err
	}
	// Deserialize `BuyerPrice`:
	err = decoder.Decode(&obj.BuyerPrice)
	if err != nil {
		return err
	}
	// Deserialize `TokenSize`:
	err = decoder.Decode(&obj.TokenSize)
	if err != nil {
		return err
	}
	return nil
}

// NewSellInstruction declares a new Sell instruction with the provided parameters and accounts.
func NewSellInstruction(
	// Parameters:
	tradeStateBump uint8,
	freeTradeStateBump uint8,
	programAsSignerBump uint8,
	buyerPrice uint64,
	tokenSize uint64,
	// Accounts:
	wallet ag_solanago.PublicKey,
	tokenAccount ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	auctionHouse ag_solanago.PublicKey,
	auctionHouseFeeAccount ag_solanago.PublicKey,
	sellerTradeState ag_solanago.PublicKey,
	freeSellerTradeState ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	programAsSigner ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *Sell {
	return NewSellInstructionBuilder().
		SetTradeStateBump(tradeStateBump).
		SetFreeTradeStateBump(freeTradeStateBump).
		SetProgramAsSignerBump(programAsSignerBump).
		SetBuyerPrice(buyerPrice).
		SetTokenSize(tokenSize).
		SetWalletAccount(wallet).
		SetTokenAccountAccount(tokenAccount).
		SetMetadataAccount(metadata).
		SetAuthorityAccount(authority).
		SetAuctionHouseAccount(auctionHouse).
		SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount).
		SetSellerTradeStateAccount(sellerTradeState).
		SetFreeSellerTradeStateAccount(freeSellerTradeState).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetProgramAsSignerAccount(programAsSigner).
		SetRentAccount(rent)
}
