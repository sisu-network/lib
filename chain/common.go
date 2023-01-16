package chain

import "math/big"

func GetKeygenType(chain string) string {
	if IsETHBasedChain(chain) {
		return KEY_TYPE_ECDSA
	}

	if IsCardanoChain(chain) {
		return KEY_TYPE_EDDSA
	}

	return ""
}

func GetKeyTypeForChain(chain string) string {
	if IsETHBasedChain(chain) {
		return KEY_TYPE_ECDSA
	}

	if IsCardanoChain(chain) {
		return KEY_TYPE_EDDSA
	}

	return ""
}

func GetChainNameFromInt(bigI *big.Int) string {
	id := bigI.Int64()
	switch id {
	case 1:
		return "eth"
	case 3:
		return "ropsten-testnet"
	case 5:
		return "goerli-testnet"
	case 100:
		return "xdai"
	case 97:
		return "binance-testnet"
	case 4002:
		return "fantom-testnet"
	case 80001:
		return "polygon-testnet"
	case 421611:
		return "arbitrum-testnet"
	case 43113:
		return "avaxc-testnet"
	case 189985:
		return "ganache1"
	case 189986:
		return "ganache2"
	case 98723843487:
		return "cardano-testnet"
	case 2382734923:
		return "solana-devnet"
	case 9872347238347:
		return "lisk-testnet"
	case 9872347238348:
		return "lisk-mainnet"
	}

	return ""
}
