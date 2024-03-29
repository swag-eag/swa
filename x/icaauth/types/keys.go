package types

const (
	// ModuleName defines the module name
	ModuleName = "icaauth"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_icaauth"

	// Version defines the current version the IBC module supports
	Version = "icaauth-1"
)

// prefix bytes for the icaauth persistent store
const (
	paramsKey = iota + 1
)

// KVStore key prefixes
var (
	// ParamsKey is the key for params.
	ParamsKey = []byte{paramsKey}
)
