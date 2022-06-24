// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Buy is the `buy` instruction.
type Buy struct {
	TradeStateBump    *uint8
	EscrowPaymentBump *uint8
	BuyerPrice        *uint64
	TokenSize         *uint64

	// [0] = [SIGNER] wallet
	//
	// [1] = [WRITE] paymentAccount
	//
	// [2] = [] transferAuthority
	//
	// [3] = [] treasuryMint
	//
	// [4] = [] tokenAccount
	//
	// [5] = [] metadata
	//
	// [6] = [WRITE] escrowPaymentAccount
	//
	// [7] = [] authority
	//
	// [8] = [] auctionHouse
	//
	// [9] = [WRITE] auctionHouseFeeAccount
	//
	// [10] = [WRITE] buyerTradeState
	//
	// [11] = [] tokenProgram
	//
	// [12] = [] systemProgram
	//
	// [13] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewBuyInstructionBuilder creates a new `Buy` instruction builder.
func NewBuyInstructionBuilder() *Buy {
	nd := &Buy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	return nd
}

// SetTradeStateBump sets the "tradeStateBump" parameter.
func (inst *Buy) SetTradeStateBump(tradeStateBump uint8) *Buy {
	inst.TradeStateBump = &tradeStateBump
	return inst
}

// SetEscrowPaymentBump sets the "escrowPaymentBump" parameter.
func (inst *Buy) SetEscrowPaymentBump(escrowPaymentBump uint8) *Buy {
	inst.EscrowPaymentBump = &escrowPaymentBump
	return inst
}

// SetBuyerPrice sets the "buyerPrice" parameter.
func (inst *Buy) SetBuyerPrice(buyerPrice uint64) *Buy {
	inst.BuyerPrice = &buyerPrice
	return inst
}

// SetTokenSize sets the "tokenSize" parameter.
func (inst *Buy) SetTokenSize(tokenSize uint64) *Buy {
	inst.TokenSize = &tokenSize
	return inst
}

// SetWalletAccount sets the "wallet" account.
func (inst *Buy) SetWalletAccount(wallet ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(wallet).SIGNER()
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *Buy) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPaymentAccountAccount sets the "paymentAccount" account.
func (inst *Buy) SetPaymentAccountAccount(paymentAccount ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(paymentAccount).WRITE()
	return inst
}

// GetPaymentAccountAccount gets the "paymentAccount" account.
func (inst *Buy) GetPaymentAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTransferAuthorityAccount sets the "transferAuthority" account.
func (inst *Buy) SetTransferAuthorityAccount(transferAuthority ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(transferAuthority)
	return inst
}

// GetTransferAuthorityAccount gets the "transferAuthority" account.
func (inst *Buy) GetTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTreasuryMintAccount sets the "treasuryMint" account.
func (inst *Buy) SetTreasuryMintAccount(treasuryMint ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(treasuryMint)
	return inst
}

// GetTreasuryMintAccount gets the "treasuryMint" account.
func (inst *Buy) GetTreasuryMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenAccountAccount sets the "tokenAccount" account.
func (inst *Buy) SetTokenAccountAccount(tokenAccount ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenAccount)
	return inst
}

// GetTokenAccountAccount gets the "tokenAccount" account.
func (inst *Buy) GetTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *Buy) SetMetadataAccount(metadata ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *Buy) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetEscrowPaymentAccountAccount sets the "escrowPaymentAccount" account.
func (inst *Buy) SetEscrowPaymentAccountAccount(escrowPaymentAccount ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(escrowPaymentAccount).WRITE()
	return inst
}

// GetEscrowPaymentAccountAccount gets the "escrowPaymentAccount" account.
func (inst *Buy) GetEscrowPaymentAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *Buy) SetAuthorityAccount(authority ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *Buy) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetAuctionHouseAccount sets the "auctionHouse" account.
func (inst *Buy) SetAuctionHouseAccount(auctionHouse ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(auctionHouse)
	return inst
}

// GetAuctionHouseAccount gets the "auctionHouse" account.
func (inst *Buy) GetAuctionHouseAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAuctionHouseFeeAccountAccount sets the "auctionHouseFeeAccount" account.
func (inst *Buy) SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(auctionHouseFeeAccount).WRITE()
	return inst
}

// GetAuctionHouseFeeAccountAccount gets the "auctionHouseFeeAccount" account.
func (inst *Buy) GetAuctionHouseFeeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetBuyerTradeStateAccount sets the "buyerTradeState" account.
func (inst *Buy) SetBuyerTradeStateAccount(buyerTradeState ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(buyerTradeState).WRITE()
	return inst
}

// GetBuyerTradeStateAccount gets the "buyerTradeState" account.
func (inst *Buy) GetBuyerTradeStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Buy) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Buy) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *Buy) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *Buy) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetRentAccount sets the "rent" account.
func (inst *Buy) SetRentAccount(rent ag_solanago.PublicKey) *Buy {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *Buy) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst Buy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Buy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Buy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Buy) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.TradeStateBump == nil {
			return errors.New("TradeStateBump parameter is not set")
		}
		if inst.EscrowPaymentBump == nil {
			return errors.New("EscrowPaymentBump parameter is not set")
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
			return errors.New("accounts.PaymentAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TransferAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TreasuryMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.EscrowPaymentAccount is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.AuctionHouse is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.AuctionHouseFeeAccount is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.BuyerTradeState is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *Buy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Buy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   TradeStateBump", *inst.TradeStateBump))
						paramsBranch.Child(ag_format.Param("EscrowPaymentBump", *inst.EscrowPaymentBump))
						paramsBranch.Child(ag_format.Param("       BuyerPrice", *inst.BuyerPrice))
						paramsBranch.Child(ag_format.Param("        TokenSize", *inst.TokenSize))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           wallet", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          payment", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("transferAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("     treasuryMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("            token", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("         metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("    escrowPayment", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("        authority", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("     auctionHouse", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("  auctionHouseFee", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("  buyerTradeState", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("     tokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("    systemProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("             rent", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj Buy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TradeStateBump` param:
	err = encoder.Encode(obj.TradeStateBump)
	if err != nil {
		return err
	}
	// Serialize `EscrowPaymentBump` param:
	err = encoder.Encode(obj.EscrowPaymentBump)
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
func (obj *Buy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TradeStateBump`:
	err = decoder.Decode(&obj.TradeStateBump)
	if err != nil {
		return err
	}
	// Deserialize `EscrowPaymentBump`:
	err = decoder.Decode(&obj.EscrowPaymentBump)
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

// NewBuyInstruction declares a new Buy instruction with the provided parameters and accounts.
func NewBuyInstruction(
	// Parameters:
	tradeStateBump uint8,
	escrowPaymentBump uint8,
	buyerPrice uint64,
	tokenSize uint64,
	// Accounts:
	wallet ag_solanago.PublicKey,
	paymentAccount ag_solanago.PublicKey,
	transferAuthority ag_solanago.PublicKey,
	treasuryMint ag_solanago.PublicKey,
	tokenAccount ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	escrowPaymentAccount ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	auctionHouse ag_solanago.PublicKey,
	auctionHouseFeeAccount ag_solanago.PublicKey,
	buyerTradeState ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *Buy {
	return NewBuyInstructionBuilder().
		SetTradeStateBump(tradeStateBump).
		SetEscrowPaymentBump(escrowPaymentBump).
		SetBuyerPrice(buyerPrice).
		SetTokenSize(tokenSize).
		SetWalletAccount(wallet).
		SetPaymentAccountAccount(paymentAccount).
		SetTransferAuthorityAccount(transferAuthority).
		SetTreasuryMintAccount(treasuryMint).
		SetTokenAccountAccount(tokenAccount).
		SetMetadataAccount(metadata).
		SetEscrowPaymentAccountAccount(escrowPaymentAccount).
		SetAuthorityAccount(authority).
		SetAuctionHouseAccount(auctionHouse).
		SetAuctionHouseFeeAccountAccount(auctionHouseFeeAccount).
		SetBuyerTradeStateAccount(buyerTradeState).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}