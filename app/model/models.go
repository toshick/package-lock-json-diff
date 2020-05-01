package model

type PackageItem struct {
	Version      string                     `json: version`
	Resolved     string                     `json: resolved`
	Integrity    string                     `json: integrity`
	Requires     map[string]string          `json: requires`
	Dependencies map[string]PackageLockJson `json: dependencies`
	dev          bool                       `json: dev`
}

type PackageLockJson struct {
	Version         string                 `json: version`
	LockfileVersion int                    `json: lockfileVersion`
	requires        bool                   `json: requires`
	Dependencies    map[string]PackageItem `json: requires`
}
