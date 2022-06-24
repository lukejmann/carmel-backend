// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package mpl_token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateMetadataAccountV2 is the `CreateMetadataAccountV2` instruction.
type CreateMetadataAccountV2 struct {
	CreateMetadataAccountArgsV2 *CreateMetadataAccountArgsV2

	// [0] = [WRITE] metadata
	//
	// [1] = [] mint
	//
	// [2] = [SIGNER] mintAuthority
	//
	// [3] = [SIGNER] payer
	//
	// [4] = [] updateAuthority
	//
	// [5] = [] systemProgram
	//
	// [6] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateMetadataAccountV2InstructionBuilder creates a new `CreateMetadataAccountV2` instruction builder.
func NewCreateMetadataAccountV2InstructionBuilder() *CreateMetadataAccountV2 {
	nd := &CreateMetadataAccountV2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetCreateMetadataAccountArgsV2 sets the "createMetadataAccountArgsV2" parameter.
func (inst *CreateMetadataAccountV2) SetCreateMetadataAccountArgsV2(createMetadataAccountArgsV2 CreateMetadataAccountArgsV2) *CreateMetadataAccountV2 {
	inst.CreateMetadataAccountArgsV2 = &createMetadataAccountArgsV2
	return inst
}

// SetMetadataAccount sets the "metadata" account.
func (inst *CreateMetadataAccountV2) SetMetadataAccount(metadata ag_solanago.PublicKey) *CreateMetadataAccountV2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *CreateMetadataAccountV2) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
func (inst *CreateMetadataAccountV2) SetMintAccount(mint ag_solanago.PublicKey) *CreateMetadataAccountV2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *CreateMetadataAccountV2) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
func (inst *CreateMetadataAccountV2) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *CreateMetadataAccountV2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mintAuthority).SIGNER()
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
func (inst *CreateMetadataAccountV2) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *CreateMetadataAccountV2) SetPayerAccount(payer ag_solanago.PublicKey) *CreateMetadataAccountV2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *CreateMetadataAccountV2) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
func (inst *CreateMetadataAccountV2) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *CreateMetadataAccountV2 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(updateAuthority)
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
func (inst *CreateMetadataAccountV2) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateMetadataAccountV2) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateMetadataAccountV2 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateMetadataAccountV2) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetRentAccount sets the "rent" account.
func (inst *CreateMetadataAccountV2) SetRentAccount(rent ag_solanago.PublicKey) *CreateMetadataAccountV2 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CreateMetadataAccountV2) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst CreateMetadataAccountV2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateMetadataAccountV2,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateMetadataAccountV2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateMetadataAccountV2) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CreateMetadataAccountArgsV2 == nil {
			return errors.New("CreateMetadataAccountArgsV2 parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *CreateMetadataAccountV2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateMetadataAccountV2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("CreateMetadataAccountArgsV2", *inst.CreateMetadataAccountArgsV2))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       metadata", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("           mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("  mintAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("updateAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           rent", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj CreateMetadataAccountV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CreateMetadataAccountArgsV2` param:
	err = encoder.Encode(obj.CreateMetadataAccountArgsV2)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateMetadataAccountV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CreateMetadataAccountArgsV2`:
	err = decoder.Decode(&obj.CreateMetadataAccountArgsV2)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateMetadataAccountV2Instruction declares a new CreateMetadataAccountV2 instruction with the provided parameters and accounts.
func NewCreateMetadataAccountV2Instruction(
	// Parameters:
	createMetadataAccountArgsV2 CreateMetadataAccountArgsV2,
	// Accounts:
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *CreateMetadataAccountV2 {
	return NewCreateMetadataAccountV2InstructionBuilder().
		SetCreateMetadataAccountArgsV2(createMetadataAccountArgsV2).
		SetMetadataAccount(metadata).
		SetMintAccount(mint).
		SetMintAuthorityAccount(mintAuthority).
		SetPayerAccount(payer).
		SetUpdateAuthorityAccount(updateAuthority).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}