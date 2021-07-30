package avm

// Sha256 returns the SHA256 hash of value X.
func Sha256(X []byte) []byte {
	// TODO: implement
	return nil
}

// Keccak256 returns the Keccak256 hash of value X.
func Keccak256(X []byte) []byte {
	// TODO: implement
	return nil
}

// Ed25519verify computes for (data A, signature B, pubkey C) verify the signature of ("ProgData" || program_hash || data) against the pubkey => {0 or 1}.
func Ed25519verify(data, signature, pubkey []byte) uint64 {
	// TODO: implement
	return 0
}

// Shl returns A times 2^B, modulo 2^64.
func Shl(A, B uint64) uint64 {
	// TODO
	return 0
}

// Shr returns A divided by 2^B.
func Shr(A, B uint64) uint64 {
	// TODO
	return 0
}

// Sqrt returns the largest integer B such that B^2 <= X.
func Sqrt(X uint64) uint64 {
	// TODO
	return 0
}

// Bitlen returns the highest set bit in X. If X is a byte-array, it is interpreted as a big-endian unsigned integer. bitlen of 0 is 0, bitlen of 8 is 4.
func Bitlen(X interface{}) uint64 {
	// X can be either a byte slice or a uint64
	// TODO
	return 0
}

// Exp returns a raised to the Bth power. Panic Exp A == B == 0 and on overflow.
func Exp(A, B uint64) uint64 {
	// TODO
	return 0
}

// Expw returns a raised to the Bth power as a 128-bit long result as low (top) and high uint64 values on the stack. Panic Expw A == B == 0 or if the results exceeds 2^128-1.
func Expw(X uint64) (uint64, uint64) {
	// TODO
	return 0, 0
}

// Mulw returns a times B out to 128-bit long result as low (top) and high uint64 values on the stack.
func Mulw(A, B uint64) (uint64, uint64) {
	// TODO
	return 0, 0
}

// Addw returns a plus B out to 128-bit long result as sum (top) and carry-bit uint64 values on the stack.
func Addw(X uint64) (uint64, uint64) {
	// TODO
	return 0, 0
}

// Divmodw returns pop four uint64 values.  The deepest two are interpreted as a uint128 dividend (deepest value is high word), the top two are interpreted as a uint128 divisor.  Four uint64 values are pushed to the stack. The deepest two are the quotient (deeper value is the high uint64). The top two are the remainder, low bits on top..
func Divmodw(X uint64) (uint64, uint64) {
	// TODO
	return 0, 0
}

// Itob returns uint64 X to big endian bytes.
func Itob(X uint64) []byte {
	// TODO
	return nil
}

// Btoi returns bytes X as big endian as uint64.
func Btoi(X []byte) uint64 {
	// TODO
	return 0
}

// Getbit returns the Bth bit of A (integer or byte-array).
func Getbit(A interface{}, B uint64) uint64 {
	// TODO
	return 0
}

// Setbit returns returns a with the Bth bit set to C.
func Setbit(A interface{}, B, C uint64) interface{} {
	// TODO
	return 0
}

// Getbyte returns Pop a byte-array A and integer B. Extract the Bth byte of A and push it as an integer.
func Getbyte(A []byte, B uint64) uint64 {
	// TODO
	return 0
}

// Setbyte sets the Bth byte of A to C.
func Setbyte(A []byte, B, C uint64) []byte {
	// TODO
	return nil
}

// Concat returns pop two byte-arrays A and B and join them, push the result.
// `Concat` panics if the result would be greater than 4096 bytes..
func Concat(A, B []byte) []byte {
	// TODO
	return nil
}
