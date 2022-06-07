package chain

import (
	"math/big"

	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sisu-network/lib/log"
)

var (
	ethChains = map[string]bool{
		"eth":              true,
		"ropsten-testnet":  true,
		"goerli-testnet":   true,
		"binance-testnet":  true,
		"ganache1":         true,
		"ganache2":         true,
		"fantom-testnet":   true,
		"polygon-testnet":  true,
		"xdai":             true,
		"arbitrum-testnet": true,
	}
	ethSigners map[string]etypes.Signer
)

func initSigners() {
	ethSigners = make(map[string]etypes.Signer)

	ethSigners["eth"] = etypes.NewEIP2930Signer(GetChainIntFromId("eth"))
	ethSigners["ropsten-testnet"] = etypes.NewEIP2930Signer(GetChainIntFromId("ropsten-testnet"))
	ethSigners["goerli-testnet"] = etypes.NewEIP2930Signer(GetChainIntFromId("goerli-testnet"))
	ethSigners["binance-testnet"] = etypes.NewEIP2930Signer(GetChainIntFromId("binance-testnet"))
	ethSigners["ganache1"] = etypes.NewEIP2930Signer(GetChainIntFromId("ganache1"))
	ethSigners["ganache2"] = etypes.NewEIP2930Signer(GetChainIntFromId("ganache2"))
	ethSigners["fantom-testnet"] = etypes.NewEIP2930Signer(GetChainIntFromId("fantom-testnet"))
	ethSigners["xdai"] = etypes.NewEIP2930Signer(GetChainIntFromId("xdai"))
	ethSigners["polygon-testnet"] = etypes.NewEIP2930Signer(GetChainIntFromId("polygon-testnet"))
	ethSigners["arbitrum-testnet"] = etypes.NewEIP2930Signer(GetChainIntFromId("arbitrum-testnet"))
}

func GetChainIntFromId(chain string) *big.Int {
	switch chain {
	case "eth":
		return big.NewInt(1)
	case "ropsten-testnet":
		return big.NewInt(3)
	case "goerli-testnet":
		return big.NewInt(420)
	case "binance-testnet":
		return big.NewInt(97)
	case "xdai":
		return big.NewInt(100)
	case "ganache1":
		return big.NewInt(189985)
	case "ganache2":
		return big.NewInt(189986)
	case "fantom-testnet":
		return big.NewInt(4002)
	case "polygon-testnet":
		return big.NewInt(80001)
	case "arbitrum-testnet":
		return big.NewInt(421611)
	default:
		log.Error("unknown chain:", chain)
		return nil
	}
}

func IsETHBasedChain(chain string) bool {
	switch chain {
	case "eth", "ropsten-testnet", "goerli-testnet", "binance-testnet", "ganache1", "ganache2",
		"fantom-testnet", "polygon-testnet", "xdai", "arbitrum-testnet":
		return true
	}

	return false
}

func GetEthChainSigner(chain string) etypes.Signer {
	return ethSigners[chain]
}

func GetSupportedEthChains() map[string]bool {
	newMap := make(map[string]bool)
	for k, v := range ethChains {
		newMap[k] = v
	}

	return newMap
}
