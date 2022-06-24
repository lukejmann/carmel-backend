// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateMetadataAccount is the `UpdateMetadataAccount` instruction.
type UpdateMetadataAccount struct {
	UpdateMetadataAccountArgs *UpdateMetadataAccountArgs

	// [0] = [WRITE] metadata
	//
	// [1] = [SIGNER] updateAuthority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateMetadataAccountInstructionBuilder creates a new `UpdateMetadataAccount` instruction builder.
func NewUpdateMetadataAccountInstructionBuilder() *UpdateMetadataAccount {
	nd := &UpdateMetadataAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetUpdateMetadataAccountArgs sets the "updateMetadataAccountArgs" parameter.
func (inst *UpdateMetadataAccount) SetUpdateMetadataAccountArgs(updateMetadataAccountArgs UpdateMetadataAccountArgs) *UpdateMetadataAccount {
	inst.UpdateMetadataAccountArgs = &updateMetadataAccountArgs
	return inst
}

// SetMetadataAccount sets the "metadata" account.
func (inst *UpdateMetadataAccount) SetMetadataAccount(metadata ag_solanago.PublicKey) *UpdateMetadataAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *UpdateMetadataAccount) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
func (inst *UpdateMetadataAccount) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *UpdateMetadataAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
func (inst *UpdateMetadataAccount) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst UpdateMetadataAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateMetadataAccount,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateMetadataAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateMetadataAccount) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.UpdateMetadataAccountArgs == nil {
			return errors.New("UpdateMetadataAccountArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
	}
	return nil
}

func (inst *UpdateMetadataAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateMetadataAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("UpdateMetadataAccountArgs", *inst.UpdateMetadataAccountArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       metadata", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("updateAuthority", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj UpdateMetadataAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UpdateMetadataAccountArgs` param:
	err = encoder.Encode(obj.UpdateMetadataAccountArgs)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateMetadataAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UpdateMetadataAccountArgs`:
	err = decoder.Decode(&obj.UpdateMetadataAccountArgs)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateMetadataAccountInstruction declares a new UpdateMetadataAccount instruction with the provided parameters and accounts.
func NewUpdateMetadataAccountInstruction(
	// Parameters:
	updateMetadataAccountArgs UpdateMetadataAccountArgs,
	// Accounts:
	metadata ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey) *UpdateMetadataAccount {
	return NewUpdateMetadataAccountInstructionBuilder().
		SetUpdateMetadataAccountArgs(updateMetadataAccountArgs).
		SetMetadataAccount(metadata).
		SetUpdateAuthorityAccount(updateAuthority)
}