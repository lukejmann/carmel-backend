// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction_house

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CancelBidReceipt is the `cancelBidReceipt` instruction.
type CancelBidReceipt struct {

	// [0] = [WRITE] receipt
	//
	// [1] = [] systemProgram
	//
	// [2] = [] instruction
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCancelBidReceiptInstructionBuilder creates a new `CancelBidReceipt` instruction builder.
func NewCancelBidReceiptInstructionBuilder() *CancelBidReceipt {
	nd := &CancelBidReceipt{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetReceiptAccount sets the "receipt" account.
func (inst *CancelBidReceipt) SetReceiptAccount(receipt ag_solanago.PublicKey) *CancelBidReceipt {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(receipt).WRITE()
	return inst
}

// GetReceiptAccount gets the "receipt" account.
func (inst *CancelBidReceipt) GetReceiptAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CancelBidReceipt) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CancelBidReceipt {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CancelBidReceipt) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetInstructionAccount sets the "instruction" account.
func (inst *CancelBidReceipt) SetInstructionAccount(instruction ag_solanago.PublicKey) *CancelBidReceipt {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(instruction)
	return inst
}

// GetInstructionAccount gets the "instruction" account.
func (inst *CancelBidReceipt) GetInstructionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst CancelBidReceipt) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CancelBidReceipt,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CancelBidReceipt) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CancelBidReceipt) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Receipt is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Instruction is not set")
		}
	}
	return nil
}

func (inst *CancelBidReceipt) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CancelBidReceipt")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      receipt", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  instruction", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj CancelBidReceipt) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CancelBidReceipt) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCancelBidReceiptInstruction declares a new CancelBidReceipt instruction with the provided parameters and accounts.
func NewCancelBidReceiptInstruction(
	// Accounts:
	receipt ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	instruction ag_solanago.PublicKey) *CancelBidReceipt {
	return NewCancelBidReceiptInstructionBuilder().
		SetReceiptAccount(receipt).
		SetSystemProgramAccount(systemProgram).
		SetInstructionAccount(instruction)
}