package model

type Dependencies map[string]PackageItem

type PackageItem struct {
	Version      string            `json:"version"`
	Resolved     string            `json:"resolved"`
	Integrity    string            `json:"integrity"`
	Requires     map[string]string `json:"requires"`
	Dependencies Dependencies      `json:"dependencies"`
	Dev          bool              `json:"dev"`
}

type PackageLockJson struct {
	Version         string       `json:"version"`
	LockfileVersion int          `json:"lockfileVersion"`
	Requires        bool         `json:"requires"`
	Dependencies    Dependencies `json:"dependencies"`
}
