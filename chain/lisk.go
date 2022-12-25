package chain

var (
	LiskChains = map[string]bool{
		"lisk-devnet":  true,
		"lisk-testnet": true,
		"lisk-mainnet": true,
	}
)

func IsLiskChain(chain string) bool {
	return LiskChains[chain]
}
