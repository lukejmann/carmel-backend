// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateMasterEdition is the `CreateMasterEdition` instruction.
type CreateMasterEdition struct {
	CreateMasterEditionArgs *CreateMasterEditionArgs

	// [0] = [WRITE] edition
	//
	// [1] = [WRITE] mint
	//
	// [2] = [SIGNER] updateAuthority
	//
	// [3] = [SIGNER] mintAuthority
	//
	// [4] = [SIGNER] payer
	//
	// [5] = [] metadata
	//
	// [6] = [] tokenProgram
	//
	// [7] = [] systemProgram
	//
	// [8] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateMasterEditionInstructionBuilder creates a new `CreateMasterEdition` instruction builder.
func NewCreateMasterEditionInstructionBuilder() *CreateMasterEdition {
	nd := &CreateMasterEdition{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetCreateMasterEditionArgs sets the "createMasterEditionArgs" parameter.
func (inst *CreateMasterEdition) SetCreateMasterEditionArgs(createMasterEditionArgs CreateMasterEditionArgs) *CreateMasterEdition {
	inst.CreateMasterEditionArgs = &createMasterEditionArgs
	return inst
}

// SetEditionAccount sets the "edition" account.
func (inst *CreateMasterEdition) SetEditionAccount(edition ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(edition).WRITE()
	return inst
}

// GetEditionAccount gets the "edition" account.
func (inst *CreateMasterEdition) GetEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
func (inst *CreateMasterEdition) SetMintAccount(mint ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint).WRITE()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *CreateMasterEdition) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
func (inst *CreateMasterEdition) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
func (inst *CreateMasterEdition) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
func (inst *CreateMasterEdition) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mintAuthority).SIGNER()
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
func (inst *CreateMasterEdition) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPayerAccount sets the "payer" account.
func (inst *CreateMasterEdition) SetPayerAccount(payer ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *CreateMasterEdition) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *CreateMasterEdition) SetMetadataAccount(metadata ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *CreateMasterEdition) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CreateMasterEdition) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CreateMasterEdition) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateMasterEdition) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateMasterEdition) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetRentAccount sets the "rent" account.
func (inst *CreateMasterEdition) SetRentAccount(rent ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CreateMasterEdition) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst CreateMasterEdition) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateMasterEdition,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateMasterEdition) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateMasterEdition) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CreateMasterEditionArgs == nil {
			return errors.New("CreateMasterEditionArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Edition is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *CreateMasterEdition) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateMasterEdition")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("CreateMasterEditionArgs", *inst.CreateMasterEditionArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        edition", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("updateAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("  mintAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          payer", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("           rent", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj CreateMasterEdition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CreateMasterEditionArgs` param:
	err = encoder.Encode(obj.CreateMasterEditionArgs)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateMasterEdition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CreateMasterEditionArgs`:
	err = decoder.Decode(&obj.CreateMasterEditionArgs)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateMasterEditionInstruction declares a new CreateMasterEdition instruction with the provided parameters and accounts.
func NewCreateMasterEditionInstruction(
	// Parameters:
	createMasterEditionArgs CreateMasterEditionArgs,
	// Accounts:
	edition ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *CreateMasterEdition {
	return NewCreateMasterEditionInstructionBuilder().
		SetCreateMasterEditionArgs(createMasterEditionArgs).
		SetEditionAccount(edition).
		SetMintAccount(mint).
		SetUpdateAuthorityAccount(updateAuthority).
		SetMintAuthorityAccount(mintAuthority).
		SetPayerAccount(payer).
		SetMetadataAccount(metadata).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
