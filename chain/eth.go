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
		"avaxc-testnet":    true,
	}
	ethSigners map[string]etypes.Signer
)

func initSigners() {
	ethSigners = make(map[string]etypes.Signer)

	ethSigners["eth"] = etypes.NewLondonSigner(GetChainIntFromId("eth"))
	ethSigners["ropsten-testnet"] = etypes.NewLondonSigner(GetChainIntFromId("ropsten-testnet"))
	ethSigners["goerli-testnet"] = etypes.NewLondonSigner(GetChainIntFromId("goerli-testnet"))
	ethSigners["binance-testnet"] = etypes.NewLondonSigner(GetChainIntFromId("binance-testnet"))
	ethSigners["ganache1"] = etypes.NewLondonSigner(GetChainIntFromId("ganache1"))
	ethSigners["ganache2"] = etypes.NewLondonSigner(GetChainIntFromId("ganache2"))
	ethSigners["fantom-testnet"] = etypes.NewLondonSigner(GetChainIntFromId("fantom-testnet"))
	ethSigners["xdai"] = etypes.NewLondonSigner(GetChainIntFromId("xdai"))
	ethSigners["polygon-testnet"] = etypes.NewLondonSigner(GetChainIntFromId("polygon-testnet"))
	ethSigners["arbitrum-testnet"] = etypes.NewLondonSigner(GetChainIntFromId("arbitrum-testnet"))
	ethSigners["avaxc-testnet"] = etypes.NewLondonSigner(GetChainIntFromId("avaxc-testnet"))
}

func GetChainIntFromId(chain string) *big.Int {
	switch chain {
	case "eth":
		return big.NewInt(1)
	case "ropsten-testnet":
		return big.NewInt(3)
	case "goerli-testnet":
		return big.NewInt(5)
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
	case "avaxc-testnet":
		return big.NewInt(43113)

	// Non-evm
	case "cardano-testnet":
		return big.NewInt(98723843487)
	default:
		log.Error("unknown chain:", chain)
		return nil
	}
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
	}

	return ""
}

func IsETHBasedChain(chain string) bool {
	switch chain {
	case "eth", "ropsten-testnet", "goerli-testnet", "binance-testnet", "ganache1", "ganache2",
		"fantom-testnet", "polygon-testnet", "xdai", "arbitrum-testnet", "avaxc-testnet":
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
