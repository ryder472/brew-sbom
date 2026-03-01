package model

// Package represents a Homebrew-installed package.
type Package struct {
	Name        string
	Version     string
	License     string
	Homepage    string
	Description string
}
