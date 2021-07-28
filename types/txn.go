package types

// Txn is the current transaction.
type Txn struct {
	// 32 byte address
	Sender []byte
	// micro-Algos
	Fee uint64
	// round number
	FirstValid uint64
	// Causes program to fail; reserved for future use
	FirstValidTime uint64
	// round number
	LastValid uint64
	// Any data up to 1024 bytes
	Note []byte
	// 32 byte lease value
	Lease []byte
	// 32 byte address
	Receiver []byte
	// micro-Algos
	Amount uint64
	// 32 byte address
	CloseRemainderTo []byte
	// 32 byte address
	VotePK []byte
	// 32 byte address
	SelectionPK []byte
	// The first round that the participation key is valid.
	VoteFirst uint64
	// The last round that the participation key is valid.
	VoteLast uint64
	// Dilution for the 2-level participation key
	VoteKeyDilution uint64
	// Transaction type as bytes
	Type []byte
	// See table below
	TypeEnum uint64
	// Asset ID
	XferAsset uint64
	// value in Asset's units
	AssetAmount uint64
	// 32 byte address. Causes clawback of all value of asset from AssetSender if Sender is the Clawback address of the asset.
	AssetSender []byte
	// 32 byte address
	AssetReceiver []byte
	// 32 byte address
	AssetCloseTo []byte
	// Position of this transaction within an atomic transaction group. A stand-alone transaction is implicitly element 0 in a group of 1
	GroupIndex uint64
	// The computed ID for this transaction. 32 bytes.
	TxID []byte
	// ApplicationID from ApplicationCall transaction
	ApplicationID uint64
	// ApplicationCall transaction on completion action
	OnCompletion uint64
	// Arguments passed to the application in the ApplicationCall transaction
	ApplicationArgs []byte
	// Number of ApplicationArgs
	NumAppArgs uint64
	// Accounts listed in the ApplicationCall transaction
	Accounts []byte
	// Number of Accounts
	NumAccounts uint64
	// Approval program
	ApprovalProgram []byte
	// Clear state program
	ClearStateProgram []byte
	// 32 byte Sender's new AuthAddr
	RekeyTo []byte
	// Asset ID in asset config transaction
	ConfigAsset uint64
	// Total number of units of this asset created
	ConfigAssetTotal uint64
	// Number of digits to display after the decimal place when displaying the asset
	ConfigAssetDecimals uint64
	// Whether the asset's slots are frozen by default or not, 0 or 1
	ConfigAssetDefaultFrozen uint64
	// Unit name of the asset
	ConfigAssetUnitName []byte
	// The asset name
	ConfigAssetName []byte
	// URL
	ConfigAssetURL []byte
	// 32 byte commitment to some unspecified asset metadata
	ConfigAssetMetadataHash []byte
	// 32 byte address
	ConfigAssetManager []byte
	// 32 byte address
	ConfigAssetReserve []byte
	// 32 byte address
	ConfigAssetFreeze []byte
	// 32 byte address
	ConfigAssetClawback []byte
	// Asset ID being frozen or un-frozen
	FreezeAsset uint64
	// 32 byte address of the account whose asset slot is being frozen or un-frozen
	FreezeAssetAccount []byte
	// The new frozen value, 0 or 1
	FreezeAssetFrozen uint64
	// Foreign Assets listed in the ApplicationCall transaction
	Assets uint64
	// Number of Assets
	NumAssets uint64
	// Foreign Apps listed in the ApplicationCall transaction
	Applications uint64
	// Number of Applications
	NumApplications uint64
	// Number of global state integers in ApplicationCall
	GlobalNumUint uint64
	// Number of global state byteslices in ApplicationCall
	GlobalNumByteSlice uint64
	// Number of local state integers in ApplicationCall
	LocalNumUint uint64
	// Number of local state byteslices in ApplicationCall
	LocalNumByteSlice uint64
	// Number of additional pages for each of the application's approval and clear state programs. An ExtraProgramPages of 1 means 2048 more total bytes, or 1024 for each program.
	ExtraProgramPages uint64
}
