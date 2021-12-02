package chain

import (
	"math/big"

	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sisu-network/lib/log"
)

var (
	ETH_CHAINS = map[string]bool{
		"eth":              true,
		"eth-ropsten":      true,
		"eth-sisu-local":   true,
		"eth-sisu-testnet": true,
		"ganache1":         true,
		"ganache2":         true,
	}
	ethSigners map[string]etypes.Signer
)

func initSigners() {
	ethSigners = make(map[string]etypes.Signer)

	ethSigners["eth"] = etypes.NewEIP2930Signer(GetChainIntFromId("eth"))
	ethSigners["eth-ropsten"] = etypes.NewEIP2930Signer(GetChainIntFromId("eth-ropsten"))
	ethSigners["eth-sisu-local"] = etypes.NewEIP2930Signer(GetChainIntFromId("eth-sisu-local"))
	ethSigners["eth-sisu-testnet"] = etypes.NewEIP2930Signer(GetChainIntFromId("eth-sisu-testnet"))
	ethSigners["ganache1"] = etypes.NewEIP2930Signer(GetChainIntFromId("ganache1"))
	ethSigners["ganache2"] = etypes.NewEIP2930Signer(GetChainIntFromId("ganache2"))
}

func GetChainIntFromId(chain string) *big.Int {
	switch chain {
	case "eth":
		return big.NewInt(1)
	case "eth-ropsten":
		return big.NewInt(3)
	case "eth-sisu-local":
		return big.NewInt(314767)
	case "eth-sisu-testnet":
		return big.NewInt(314777)
	case "ganache1":
		return big.NewInt(189985)
	case "ganache2":
		return big.NewInt(189986)
	default:
		log.Error("unknown chain:", chain)
		return nil
	}
}

func IsETHBasedChain(chain string) bool {
	switch chain {
	case "eth", "eth-ropsten", "eth-sisu-local", "eth-sisu-testnet", "ganache1", "ganache2":
		return true
	}

	return false
}

func GetEthChainSigner(chain string) etypes.Signer {
	return ethSigners[chain]
}

func GetSupportedEthChains() map[string]bool {
	newMap := make(map[string]bool)
	for k, v := range ETH_CHAINS {
		newMap[k] = v
	}

	return newMap
}

func GetKeygenType(chain string) string {
	if IsETHBasedChain(chain) {
		return KEY_TYPE_ECDSA
	}

	return ""
}

func GetKeyTypeForChain(chain string) string {
	if IsETHBasedChain(chain) {
		return KEY_TYPE_ECDSA
	}

	return ""
}
