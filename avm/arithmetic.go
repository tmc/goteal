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

/* TODO:
"Arithmetic":           {"shl", "shr", "sqrt", "bitlen", "exp", "==", "!=", "!", "len", "itob", "btoi", "%", "|", "&", "^", "~", "mulw", "addw", "divmodw", "expw", "getbit", "setbit", "getbyte", "setbyte", "concat"},
"Byte Array Slicing":   {"substring", "substring3", "extract", "extract3", "extract16bits", "extract32bits", "extract64bits"},
"Byteslice Arithmetic": {"b+", "b-", "b/", "b*", "b<", "b>", "b<=", "b>=", "b==", "b!=", "b%"},
"Byteslice Logic":      {"b|", "b&", "b^", "b~"},
"Loading Values":       {"intcblock", "intc", "intc_0", "intc_1", "intc_2", "intc_3", "pushint", "bytecblock", "bytec", "bytec_0", "bytec_1", "bytec_2", "bytec_3", "pushbytes", "bzero", "arg", "arg_0", "arg_1", "arg_2", "arg_3", "txn", "gtxn", "txna", "gtxna", "gtxns", "gtxnsa", "global", "load", "store", "gload", "gloads", "gaid", "gaids"},
"Flow Control":         {"err", "bnz", "bz", "b", "return", "pop", "dup", "dup2", "dig", "swap", "select", "assert", "callsub", "retsub"},
"State Access":         {"balance", "min_balance", "app_opted_in", "app_local_get", "app_local_get_ex", "app_global_get", "app_global_get_ex", "app_local_put", "app_global_put", "app_local_del", "app_global_del", "asset_holding_get", "asset_params_get", "app_params_get"},

*/
