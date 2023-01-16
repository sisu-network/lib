package chain

var (
	LiskChains = map[string]bool{
		"lisk-testnet": true,
		"lisk-mainnet": true,
	}
)

func IsLiskChain(chain string) bool {
	return LiskChains[chain]
}
