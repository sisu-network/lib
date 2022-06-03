package chain

import "fmt"

func init() {
	initSigners()

	// Do Sanity checking to make sure that our functions do not miss any chain.
	for chain, _ := range ethChains {
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
