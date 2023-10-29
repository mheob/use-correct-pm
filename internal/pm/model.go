// Package pm holds the package manager related functions
package pm

// PackageManager structure of possible package manager
type PackageManager struct {
	Name        string
	Description string
	ShortFlag   string
	LockFile    string
}

// AvailablePackageManagers lists all available and supported package managers
var AvailablePackageManagers = []PackageManager{BUN, NPM, PNPM, YARN}

// BUN Package Manager
var BUN = PackageManager{
	Name:        "bun",
	Description: "Use BUN as package manager",
	ShortFlag:   "b",
	LockFile:    "bun.lockb",
}

// NPM Package Manager
var NPM = PackageManager{
	Name:        "npm",
	Description: "Use NPM as package manager",
	ShortFlag:   "n",
	LockFile:    "package-lock.json",
}

// PNPM Package Manager
var PNPM = PackageManager{
	Name:        "pnpm",
	Description: "Use PNPM as package manager",
	ShortFlag:   "p",
	LockFile:    "pnpm-lock.yaml",
}

// YARN Package Manager
var YARN = PackageManager{
	Name:        "yarn",
	Description: "Use YARN as package manager",
	ShortFlag:   "y",
	LockFile:    "yarn.lock",
}
