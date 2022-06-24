package constants

import "github.com/gagliardetto/solana-go"

var (
	RPCEndPoint       = "https://weathered-delicate-feather.solana-mainnet.quiknode.pro/6f97c6fb63dbf6c02d6ed553bfd60584b5d04e62/"
	SolBase58         = "11111111111111111111111111111111"
	NameServicePubkey = solana.MustPublicKeyFromBase58("namesLPneVptA9Z5rqUDD9tMTWEJwofgaYwp8cawRkX")
	MagicEdenProgram  = solana.MustPublicKeyFromBase58("M2mx93ekt1fmXSVkTrUL9xVFHkmME8HTUi5Cyc5aF7K")
	HausProgram       = solana.MustPublicKeyFromBase58("hausS13jsjafwWwGqZTUQRmWyvyxn9EQpqMwV1PBBmk")
	MetaplexProgram   = solana.MustPublicKeyFromBase58("metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s")
)
