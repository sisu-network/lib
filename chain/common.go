package chain

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
