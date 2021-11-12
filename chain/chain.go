package chain

import (
	"fmt"
	"math/big"

	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/sisu-network/lib/log"
)

var (
	ETH_CHAINS = map[string]bool{
		"eth":      true,
		"sisu-eth": true,
		"ganache1": true,
		"ganache2": true,
	}
	ethSigners map[string]etypes.Signer
)

func init() {
	initSigners()

	// Do Sanity checking to make sure that our functions do not miss any chain.
	for chain, _ := range ETH_CHAINS {
		id := GetChainIntFromId(chain)
		if id == nil {
			panic(fmt.Errorf("Cannot get chain id from %s", chain))
		}

		if !IsETHBasedChain(chain) {
			panic(fmt.Errorf("Chain %s is not an ETH based chain", chain))
		}

		if ethSigners[chain] == nil {
			panic(fmt.Errorf("There is no signer for chain %s", chain))
		}
	}
}

func initSigners() {
	ethSigners["eth"] = etypes.NewEIP2930Signer(GetChainIntFromId("eth"))
	ethSigners["sisu-eth"] = etypes.NewEIP2930Signer(GetChainIntFromId("sisu-eth"))
	ethSigners["ganache1"] = etypes.NewEIP2930Signer(GetChainIntFromId("ganache1"))
	ethSigners["ganache2"] = etypes.NewEIP2930Signer(GetChainIntFromId("ganache2"))
}

func GetChainIntFromId(chain string) *big.Int {
	switch chain {
	case "eth":
		return big.NewInt(1)
	case "sisu-eth":
		return big.NewInt(314767)
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
	case "sisu-eth", "eth", "ganache1", "ganache2":
		return true
	}

	return false
}

func GetEthChainSigner(chain string) etypes.Signer {
	return ethSigners[chain]
}
