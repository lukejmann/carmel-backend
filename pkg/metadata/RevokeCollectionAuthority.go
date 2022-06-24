// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RevokeCollectionAuthority is the `RevokeCollectionAuthority` instruction.
type RevokeCollectionAuthority struct {

	// [0] = [WRITE] collectionAuthorityRecord
	//
	// [1] = [SIGNER] updateAuthority
	//
	// [2] = [] metadata
	//
	// [3] = [] mint
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRevokeCollectionAuthorityInstructionBuilder creates a new `RevokeCollectionAuthority` instruction builder.
func NewRevokeCollectionAuthorityInstructionBuilder() *RevokeCollectionAuthority {
	nd := &RevokeCollectionAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetCollectionAuthorityRecordAccount sets the "collectionAuthorityRecord" account.
func (inst *RevokeCollectionAuthority) SetCollectionAuthorityRecordAccount(collectionAuthorityRecord ag_solanago.PublicKey) *RevokeCollectionAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(collectionAuthorityRecord).WRITE()
	return inst
}

// GetCollectionAuthorityRecordAccount gets the "collectionAuthorityRecord" account.
func (inst *RevokeCollectionAuthority) GetCollectionAuthorityRecordAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
func (inst *RevokeCollectionAuthority) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *RevokeCollectionAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
func (inst *RevokeCollectionAuthority) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *RevokeCollectionAuthority) SetMetadataAccount(metadata ag_solanago.PublicKey) *RevokeCollectionAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *RevokeCollectionAuthority) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAccount sets the "mint" account.
func (inst *RevokeCollectionAuthority) SetMintAccount(mint ag_solanago.PublicKey) *RevokeCollectionAuthority {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *RevokeCollectionAuthority) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst RevokeCollectionAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RevokeCollectionAuthority,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RevokeCollectionAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RevokeCollectionAuthority) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CollectionAuthorityRecord is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Mint is not set")
		}
	}
	return nil
}

func (inst *RevokeCollectionAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RevokeCollectionAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("collectionAuthorityRecord", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          updateAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                 metadata", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                     mint", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj RevokeCollectionAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RevokeCollectionAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRevokeCollectionAuthorityInstruction declares a new RevokeCollectionAuthority instruction with the provided parameters and accounts.
func NewRevokeCollectionAuthorityInstruction(
	// Accounts:
	collectionAuthorityRecord ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey) *RevokeCollectionAuthority {
	return NewRevokeCollectionAuthorityInstructionBuilder().
		SetCollectionAuthorityRecordAccount(collectionAuthorityRecord).
		SetUpdateAuthorityAccount(updateAuthority).
		SetMetadataAccount(metadata).
		SetMintAccount(mint)
}