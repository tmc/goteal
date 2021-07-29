package types

const (
	// UnknownTx signals an error
	UnknownTx = iota

	// PaymentTx indicates a payment transaction
	PaymentTx

	// KeyRegistrationTx indicates a transaction that registers participation keys
	KeyRegistrationTx

	// AssetConfigTx creates, re-configures, or destroys an asset
	AssetConfigTx

	// AssetTransferTx transfers assets between accounts (optionally closing)
	AssetTransferTx

	// AssetFreezeTx changes the freeze status of an asset
	AssetFreezeTx

	// ApplicationCallTx allows creating, deleting, and interacting with an application
	ApplicationCallTx

	// CompactCertTx records a compact certificate
	CompactCertTx
)
