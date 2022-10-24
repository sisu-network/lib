package chain

var (
	SolanaChains = map[string]bool{
		"solana-devnet":  true,
		"solana-testnet": true,
		"solana-mainnet": true,
	}
)

func IsSolanaChain(chain string) bool {
	return SolanaChains[chain]
}
