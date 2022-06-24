// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ExecuteSale is the `executeSale` instruction.
type ExecuteSale struct {
	EscrowPaymentBump   *uint8
	FreeTradeStateBump  *uint8
	ProgramAsSignerBump *uint8
	BuyerPrice          *uint64
	TokenSize           *uint64

	// [0] = [WRITE] buyer
	//
	// [1] = [WRITE] seller
	//
	// [2] = [WRITE] tokenAccount
	//
	// [3] = [] tokenMint
	//
	// [4] = [] metadata
	//
	// [5] = [] treasuryMint
	//
	// [6] = [WRITE] escrowPaymentAccount
	//
	// [7] = [WRITE] sellerPaymentReceiptAccount
	//
	// [8] = [WRITE] buyerReceiptTokenAccount
	//
	// [9] = [] authority
	//
	// [10] = [] auctionHouse
	//
	// [11] = [WRITE] auctionHouseFeeAccount
	//
	// [12] = [WRITE] auctionHouseTreasury
	//
	// [13] = [WRITE] buyerTradeState
	//
	// [14] = [WRITE] sellerTradeState
	//
	// [15] = [WRITE] freeTradeState
	//
	// [16] = [] tokenProgram
	//
	// [17] = [] systemProgram
	//
	// [18] = [] ataProgram
	//
	// [19] = [] programAsSigner
	//
	// [20] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewExecuteSaleInstructionBuilder creates a new `ExecuteSale` instruction builder.
func NewExecuteSaleInstructionBuilder() *ExecuteSale {
	nd := &ExecuteSale{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 21),
	}
	return nd
}

// SetEscrowPaymentBump sets the "escrowPaymentBump" parameter.
func (inst *ExecuteSale) SetEscrowPaymentBump(escrowPaymentBump uint8) *ExecuteSale {
	inst.EscrowPaymentBump = &escrowPaymentBump
	return inst
}

// SetFreeTradeStateBump sets the "freeTradeStateBump" parameter.
func (inst *ExecuteSale) SetFreeTradeStateBump(freeTradeStateBump uint8) *ExecuteSale {
	inst.FreeTradeStateBump = &freeTradeStateBump
	return inst
}

// SetProgramAsSignerBump sets the "programAsSignerBump" parameter.
func (inst *ExecuteSale) SetProgramAsSignerBump(programAsSignerBump uint8) *ExecuteSale {
	inst.ProgramAsSignerBump = &programAsSignerBump
	return inst
}

// SetBuyerPrice sets the "buyerPrice" parameter.
func (inst *ExecuteSale) SetBuyerPrice(buyerPrice uint64) *ExecuteSale {
	inst.BuyerPrice = &buyerPrice
	return inst
}

// SetTokenSize sets the "tokenSize" parameter.
func (inst *ExecuteSale) SetTokenSize(tokenSize uint64) *ExecuteSale {
	inst.TokenSize = &tokenSize
	return inst
}

// SetBuyerAccount sets the "buyer" account.
func (inst *ExecuteSale) SetBuyerAccount(buyer ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(buyer).WRITE()
	return inst
}

// GetBuyerAccount gets the "buyer" account.
func (inst *ExecuteSale) GetBuyerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSellerAccount sets the "seller" account.
func (inst *ExecuteSale) SetSellerAccount(seller ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(seller).WRITE()
	return inst
}

// GetSellerAccount gets the "seller" account.
func (inst *ExecuteSale) GetSellerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenAccountAccount sets the "tokenAccount" account.
func (inst *ExecuteSale) SetTokenAccountAccount(tokenAccount ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenAccount).WRITE()
	return inst
}

// GetTokenAccountAccount gets the "tokenAccount" account.
func (inst *ExecuteSale) GetTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenMintAccount sets the "tokenMint" account.
func (inst *ExecuteSale) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintAccount gets the "tokenMint" account.
func (inst *ExecuteSale) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *ExecuteSale) SetMetadataAccount(metadata ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *ExecuteSale) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTreasuryMintAccount sets the "treasuryMint" account.
func (inst *ExecuteSale) SetTreasuryMintAccount(treasuryMint ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(treasuryMint)
	return inst
}

// GetTreasuryMintAccount gets the "treasuryMint" account.
func (inst *ExecuteSale) GetTreasuryMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetEscrowPaymentAccountAccount sets the "escrowPaymentAccount" account.
func (inst *ExecuteSale) SetEscrowPaymentAccountAccount(escrowPaymentAccount ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(escrowPaymentAccount).WRITE()
	return inst
}

// GetEscrowPaymentAccountAccount gets the "escrowPaymentAccount" account.
func (inst *ExecuteSale) GetEscrowPaymentAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSellerPaymentReceiptAccountAccount sets the "sellerPaymentReceiptAccount" account.
func (inst *ExecuteSale) SetSellerPaymentReceiptAccountAccount(sellerPaymentReceiptAccount ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(sellerPaymentReceiptAccount).WRITE()
	return inst
}

// GetSellerPaymentReceiptAccountAccount gets the "sellerPaymentReceiptAccount" account.
func (inst *ExecuteSale) GetSellerPaymentReceiptAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetBuyerReceiptTokenAccountAccount sets the "buyerReceiptTokenAccount" account.
func (inst *ExecuteSale) SetBuyerReceiptTokenAccountAccount(buyerReceiptTokenAccount ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(buyerReceiptTokenAccount).WRITE()
	return inst
}

// GetBuyerReceiptTokenAccountAccount gets the "buyerReceiptTokenAccount" account.
func (inst *ExecuteSale) GetBuyerReceiptTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *ExecuteSale) SetAuthorityAccount(authority ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *ExecuteSale) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetAuctionHouseAccount sets the "auctionHouse" account.
func (inst *ExecuteSale) SetAuctionHouseAccount(auctionHouse ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(auctionHouse)
	return inst
}

// GetAuctionHouseAccount gets the "auctionHouse" account.
func (inst *ExecuteSale) GetAuctionHouseAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetAuctionHouseFeeAccountAccount sets the "auctionHouseFeeAccount" account.
func (inst *ExecuteSale) SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(auctionHouseFeeAccount).WRITE()
	return inst
}

// GetAuctionHouseFeeAccountAccount gets the "auctionHouseFeeAccount" account.
func (inst *ExecuteSale) GetAuctionHouseFeeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetAuctionHouseTreasuryAccount sets the "auctionHouseTreasury" account.
func (inst *ExecuteSale) SetAuctionHouseTreasuryAccount(auctionHouseTreasury ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(auctionHouseTreasury).WRITE()
	return inst
}

// GetAuctionHouseTreasuryAccount gets the "auctionHouseTreasury" account.
func (inst *ExecuteSale) GetAuctionHouseTreasuryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetBuyerTradeStateAccount sets the "buyerTradeState" account.
func (inst *ExecuteSale) SetBuyerTradeStateAccount(buyerTradeState ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(buyerTradeState).WRITE()
	return inst
}

// GetBuyerTradeStateAccount gets the "buyerTradeState" account.
func (inst *ExecuteSale) GetBuyerTradeStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetSellerTradeStateAccount sets the "sellerTradeState" account.
func (inst *ExecuteSale) SetSellerTradeStateAccount(sellerTradeState ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(sellerTradeState).WRITE()
	return inst
}

// GetSellerTradeStateAccount gets the "sellerTradeState" account.
func (inst *ExecuteSale) GetSellerTradeStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetFreeTradeStateAccount sets the "freeTradeState" account.
func (inst *ExecuteSale) SetFreeTradeStateAccount(freeTradeState ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(freeTradeState).WRITE()
	return inst
}

// GetFreeTradeStateAccount gets the "freeTradeState" account.
func (inst *ExecuteSale) GetFreeTradeStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *ExecuteSale) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *ExecuteSale) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *ExecuteSale) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *ExecuteSale) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetAtaProgramAccount sets the "ataProgram" account.
func (inst *ExecuteSale) SetAtaProgramAccount(ataProgram ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(ataProgram)
	return inst
}

// GetAtaProgramAccount gets the "ataProgram" account.
func (inst *ExecuteSale) GetAtaProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetProgramAsSignerAccount sets the "programAsSigner" account.
func (inst *ExecuteSale) SetProgramAsSignerAccount(programAsSigner ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(programAsSigner)
	return inst
}

// GetProgramAsSignerAccount gets the "programAsSigner" account.
func (inst *ExecuteSale) GetProgramAsSignerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

// SetRentAccount sets the "rent" account.
func (inst *ExecuteSale) SetRentAccount(rent ag_solanago.PublicKey) *ExecuteSale {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *ExecuteSale) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(20)
}

func (inst ExecuteSale) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ExecuteSale,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ExecuteSale) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ExecuteSale) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.EscrowPaymentBump == nil {
			return errors.New("EscrowPaymentBump parameter is not set")
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
			return errors.New("accounts.Buyer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Seller is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TreasuryMint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.EscrowPaymentAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SellerPaymentReceiptAccount is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.BuyerReceiptTokenAccount is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.AuctionHouse is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.AuctionHouseFeeAccount is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.AuctionHouseTreasury is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.BuyerTradeState is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.SellerTradeState is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.FreeTradeState is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.AtaProgram is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.ProgramAsSigner is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *ExecuteSale) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ExecuteSale")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=5]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  EscrowPaymentBump", *inst.EscrowPaymentBump))
						paramsBranch.Child(ag_format.Param(" FreeTradeStateBump", *inst.FreeTradeStateBump))
						paramsBranch.Child(ag_format.Param("ProgramAsSignerBump", *inst.ProgramAsSignerBump))
						paramsBranch.Child(ag_format.Param("         BuyerPrice", *inst.BuyerPrice))
						paramsBranch.Child(ag_format.Param("          TokenSize", *inst.TokenSize))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=21]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               buyer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("              seller", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               token", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           tokenMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("            metadata", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("        treasuryMint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       escrowPayment", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("sellerPaymentReceipt", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("   buyerReceiptToken", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("           authority", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("        auctionHouse", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("     auctionHouseFee", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("auctionHouseTreasury", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("     buyerTradeState", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("    sellerTradeState", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("      freeTradeState", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("        tokenProgram", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("          ataProgram", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta("     programAsSigner", inst.AccountMetaSlice.Get(19)))
						accountsBranch.Child(ag_format.Meta("                rent", inst.AccountMetaSlice.Get(20)))
					})
				})
		})
}

func (obj ExecuteSale) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `EscrowPaymentBump` param:
	err = encoder.Encode(obj.EscrowPaymentBump)
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
func (obj *ExecuteSale) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `EscrowPaymentBump`:
	err = decoder.Decode(&obj.EscrowPaymentBump)
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

// NewExecuteSaleInstruction declares a new ExecuteSale instruction with the provided parameters and accounts.
func NewExecuteSaleInstruction(
	// Parameters:
	escrowPaymentBump uint8,
	freeTradeStateBump uint8,
	programAsSignerBump uint8,
	buyerPrice uint64,
	tokenSize uint64,
	// Accounts:
	buyer ag_solanago.PublicKey,
	seller ag_solanago.PublicKey,
	tokenAccount ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	treasuryMint ag_solanago.PublicKey,
	escrowPaymentAccount ag_solanago.PublicKey,
	sellerPaymentReceiptAccount ag_solanago.PublicKey,
	buyerReceiptTokenAccount ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	auctionHouse ag_solanago.PublicKey,
	auctionHouseFeeAccount ag_solanago.PublicKey,
	auctionHouseTreasury ag_solanago.PublicKey,
	buyerTradeState ag_solanago.PublicKey,
	sellerTradeState ag_solanago.PublicKey,
	freeTradeState ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	ataProgram ag_solanago.PublicKey,
	programAsSigner ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *ExecuteSale {
	return NewExecuteSaleInstructionBuilder().
		SetEscrowPaymentBump(escrowPaymentBump).
		SetFreeTradeStateBump(freeTradeStateBump).
		SetProgramAsSignerBump(programAsSignerBump).
		SetBuyerPrice(buyerPrice).
		SetTokenSize(tokenSize).
		SetBuyerAccount(buyer).
		SetSellerAccount(seller).
		SetTokenAccountAccount(tokenAccount).
		SetTokenMintAccount(tokenMint).
		SetMetadataAccount(metadata).
		SetTreasuryMintAccount(treasuryMint).
		SetEscrowPaymentAccountAccount(escrowPaymentAccount).
		SetSellerPaymentReceiptAccountAccount(sellerPaymentReceiptAccount).
		SetBuyerReceiptTokenAccountAccount(buyerReceiptTokenAccount).
		SetAuthorityAccount(authority).
		SetAuctionHouseAccount(auctionHouse).
		SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount).
		SetAuctionHouseTreasuryAccount(auctionHouseTreasury).
		SetBuyerTradeStateAccount(buyerTradeState).
		SetSellerTradeStateAccount(sellerTradeState).
		SetFreeTradeStateAccount(freeTradeState).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAtaProgramAccount(ataProgram).
		SetProgramAsSignerAccount(programAsSigner).
		SetRentAccount(rent)
}
