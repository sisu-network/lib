package chain

var (
	CardanoChains = map[string]bool{
		"cardano-testnet": true,
		"cardano-mainnet": true,
	}
)

func IsCardanoChain(chain string) bool {
	return CardanoChains[chain]
}
