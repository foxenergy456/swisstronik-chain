package types

const (
	// ModuleName defines the module name
	ModuleName = "did"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_did"

	DIDMethod 			= "swtr"
	LatestDocumentVersionKey = "did-latest:"
	DocumentVersionKey       = "did-version:"
	DocumentCountKey         = "did-count:"
	UpdatedPostfix			 = "-updated"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetLatestDocumentVersionKey(did string) []byte {
	return []byte(LatestDocumentVersionKey + did)
}

func GetDocumentVersionKey(did string, version string) []byte {
	return []byte(DocumentVersionKey + did + ":" + version)
}

func GetLatestDocumentVersionPrefix() []byte {
	return []byte(LatestDocumentVersionKey)
}

func GetDocumentVersionsPrefix(did string) []byte {
	return []byte(DocumentVersionKey + did + ":")
}
