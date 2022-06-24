// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// PrintBidReceipt is the `printBidReceipt` instruction.
type PrintBidReceipt struct {
	ReceiptBump *uint8

	// [0] = [WRITE] receipt
	//
	// [1] = [WRITE, SIGNER] bookkeeper
	//
	// [2] = [] systemProgram
	//
	// [3] = [] rent
	//
	// [4] = [] instruction
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewPrintBidReceiptInstructionBuilder creates a new `PrintBidReceipt` instruction builder.
func NewPrintBidReceiptInstructionBuilder() *PrintBidReceipt {
	nd := &PrintBidReceipt{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetReceiptBump sets the "receiptBump" parameter.
func (inst *PrintBidReceipt) SetReceiptBump(receiptBump uint8) *PrintBidReceipt {
	inst.ReceiptBump = &receiptBump
	return inst
}

// SetReceiptAccount sets the "receipt" account.
func (inst *PrintBidReceipt) SetReceiptAccount(receipt ag_solanago.PublicKey) *PrintBidReceipt {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(receipt).WRITE()
	return inst
}

// GetReceiptAccount gets the "receipt" account.
func (inst *PrintBidReceipt) GetReceiptAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBookkeeperAccount sets the "bookkeeper" account.
func (inst *PrintBidReceipt) SetBookkeeperAccount(bookkeeper ag_solanago.PublicKey) *PrintBidReceipt {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(bookkeeper).WRITE().SIGNER()
	return inst
}

// GetBookkeeperAccount gets the "bookkeeper" account.
func (inst *PrintBidReceipt) GetBookkeeperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *PrintBidReceipt) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *PrintBidReceipt {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *PrintBidReceipt) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetRentAccount sets the "rent" account.
func (inst *PrintBidReceipt) SetRentAccount(rent ag_solanago.PublicKey) *PrintBidReceipt {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *PrintBidReceipt) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetInstructionAccount sets the "instruction" account.
func (inst *PrintBidReceipt) SetInstructionAccount(instruction ag_solanago.PublicKey) *PrintBidReceipt {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(instruction)
	return inst
}

// GetInstructionAccount gets the "instruction" account.
func (inst *PrintBidReceipt) GetInstructionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst PrintBidReceipt) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_PrintBidReceipt,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PrintBidReceipt) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PrintBidReceipt) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ReceiptBump == nil {
			return errors.New("ReceiptBump parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Receipt is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Bookkeeper is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Instruction is not set")
		}
	}
	return nil
}

func (inst *PrintBidReceipt) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PrintBidReceipt")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ReceiptBump", *inst.ReceiptBump))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      receipt", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   bookkeeper", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         rent", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("  instruction", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj PrintBidReceipt) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ReceiptBump` param:
	err = encoder.Encode(obj.ReceiptBump)
	if err != nil {
		return err
	}
	return nil
}
func (obj *PrintBidReceipt) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ReceiptBump`:
	err = decoder.Decode(&obj.ReceiptBump)
	if err != nil {
		return err
	}
	return nil
}

// NewPrintBidReceiptInstruction declares a new PrintBidReceipt instruction with the provided parameters and accounts.
func NewPrintBidReceiptInstruction(
	// Parameters:
	receiptBump uint8,
	// Accounts:
	receipt ag_solanago.PublicKey,
	bookkeeper ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	instruction ag_solanago.PublicKey) *PrintBidReceipt {
	return NewPrintBidReceiptInstructionBuilder().
		SetReceiptBump(receiptBump).
		SetReceiptAccount(receipt).
		SetBookkeeperAccount(bookkeeper).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetInstructionAccount(instruction)
}