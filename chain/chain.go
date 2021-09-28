package chain

import (
	"fmt"
	"math/big"

	"github.com/sisu-network/lib/log"
)

var (
	ETH_CHAINS = map[string]bool{
		"eth":      true,
		"sisu-eth": true,
		"ganache1": true,
		"ganache2": true,
	}
)

func init() {
	// Do Sanity checking to make sure that our functions do not miss any chain.
	for chain, _ := range ETH_CHAINS {
		id := GetChainIntFromId(chain)
		if id == nil {
			panic(fmt.Errorf("Cannot get chain id from %s", chain))
		}

		if !IsETHBasedChain(chain) {
			panic(fmt.Errorf("Chain %s is not an ETH based chain", chain))
		}
	}
}

func GetChainIntFromId(chain string) *big.Int {
	switch chain {
	case "ganache1":
		return big.NewInt(189985)
	case "ganache2":
		return big.NewInt(189986)
	case "eth":
		return big.NewInt(1)
	case "sisu-eth":
		return big.NewInt(314767)
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
