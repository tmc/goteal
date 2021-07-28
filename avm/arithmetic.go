package avm

// Sha256 returns the SHA256 hash of value X.
func Sha256(X []byte) [32]byte {
	// TODO: implement
	return [32]byte{}
}

// Keccak256 returns the Keccak256 hash of value X.
func Keccak256(X []byte) [32]byte {
	// TODO: implement
	return [32]byte{}
}

// Ed25519verify computes for (data A, signature B, pubkey C) verify the signature of ("ProgData" || program_hash || data) against the pubkey => {0 or 1}
func Ed25519verify(data, signature, pubkey []byte) uint64 {
	// TODO: implement
	return 0
}
