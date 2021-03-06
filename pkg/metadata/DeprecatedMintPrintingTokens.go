// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// DeprecatedMintPrintingTokens is the `DeprecatedMintPrintingTokens` instruction.
type DeprecatedMintPrintingTokens struct {
	MintPrintingTokensViaTokenArgs *MintPrintingTokensViaTokenArgs

	// [0] = [WRITE] destination
	//
	// [1] = [WRITE] printingMint
	//
	// [2] = [SIGNER] updateAuthority
	//
	// [3] = [] metadata
	//
	// [4] = [] masterEdition
	//
	// [5] = [] tokenProgram
	//
	// [6] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDeprecatedMintPrintingTokensInstructionBuilder creates a new `DeprecatedMintPrintingTokens` instruction builder.
func NewDeprecatedMintPrintingTokensInstructionBuilder() *DeprecatedMintPrintingTokens {
	nd := &DeprecatedMintPrintingTokens{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetMintPrintingTokensViaTokenArgs sets the "mintPrintingTokensViaTokenArgs" parameter.
func (inst *DeprecatedMintPrintingTokens) SetMintPrintingTokensViaTokenArgs(mintPrintingTokensViaTokenArgs MintPrintingTokensViaTokenArgs) *DeprecatedMintPrintingTokens {
	inst.MintPrintingTokensViaTokenArgs = &mintPrintingTokensViaTokenArgs
	return inst
}

// SetDestinationAccount sets the "destination" account.
func (inst *DeprecatedMintPrintingTokens) SetDestinationAccount(destination ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(destination).WRITE()
	return inst
}

// GetDestinationAccount gets the "destination" account.
func (inst *DeprecatedMintPrintingTokens) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPrintingMintAccount sets the "printingMint" account.
func (inst *DeprecatedMintPrintingTokens) SetPrintingMintAccount(printingMint ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(printingMint).WRITE()
	return inst
}

// GetPrintingMintAccount gets the "printingMint" account.
func (inst *DeprecatedMintPrintingTokens) GetPrintingMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
func (inst *DeprecatedMintPrintingTokens) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
func (inst *DeprecatedMintPrintingTokens) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *DeprecatedMintPrintingTokens) SetMetadataAccount(metadata ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *DeprecatedMintPrintingTokens) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMasterEditionAccount sets the "masterEdition" account.
func (inst *DeprecatedMintPrintingTokens) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(masterEdition)
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
func (inst *DeprecatedMintPrintingTokens) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *DeprecatedMintPrintingTokens) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *DeprecatedMintPrintingTokens) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetRentAccount sets the "rent" account.
func (inst *DeprecatedMintPrintingTokens) SetRentAccount(rent ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *DeprecatedMintPrintingTokens) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst DeprecatedMintPrintingTokens) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_DeprecatedMintPrintingTokens,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeprecatedMintPrintingTokens) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeprecatedMintPrintingTokens) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MintPrintingTokensViaTokenArgs == nil {
			return errors.New("MintPrintingTokensViaTokenArgs parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Destination is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PrintingMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *DeprecatedMintPrintingTokens) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeprecatedMintPrintingTokens")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MintPrintingTokensViaTokenArgs", *inst.MintPrintingTokensViaTokenArgs))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    destination", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   printingMint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("updateAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       metadata", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("  masterEdition", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("   tokenProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           rent", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj DeprecatedMintPrintingTokens) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MintPrintingTokensViaTokenArgs` param:
	err = encoder.Encode(obj.MintPrintingTokensViaTokenArgs)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DeprecatedMintPrintingTokens) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MintPrintingTokensViaTokenArgs`:
	err = decoder.Decode(&obj.MintPrintingTokensViaTokenArgs)
	if err != nil {
		return err
	}
	return nil
}

// NewDeprecatedMintPrintingTokensInstruction declares a new DeprecatedMintPrintingTokens instruction with the provided parameters and accounts.
func NewDeprecatedMintPrintingTokensInstruction(
	// Parameters:
	mintPrintingTokensViaTokenArgs MintPrintingTokensViaTokenArgs,
	// Accounts:
	destination ag_solanago.PublicKey,
	printingMint ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *DeprecatedMintPrintingTokens {
	return NewDeprecatedMintPrintingTokensInstructionBuilder().
		SetMintPrintingTokensViaTokenArgs(mintPrintingTokensViaTokenArgs).
		SetDestinationAccount(destination).
		SetPrintingMintAccount(printingMint).
		SetUpdateAuthorityAccount(updateAuthority).
		SetMetadataAccount(metadata).
		SetMasterEditionAccount(masterEdition).
		SetTokenProgramAccount(tokenProgram).
		SetRentAccount(rent)
}
