package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	gencyclonedx "github.com/ryder472/brew-sbom/internal/cyclonedx"
	genspdx "github.com/ryder472/brew-sbom/internal/spdx"

	"github.com/ryder472/brew-sbom/internal/brew"
	"github.com/ryder472/brew-sbom/internal/model"
)

const version = "0.1.0"

func main() {
	format := flag.String("format", "all", `Output format: "cyclonedx", "spdx", or "all"`)
	output := flag.String("output", ".", "Output directory")
	showVersion := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("brew-sbom version %s\n", version)
		os.Exit(0)
	}

	runner := brew.NewBrewRunner(brew.ExecRunner{})
	packages, err := runner.InstalledPackages()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching installed packages: %v\n", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(*output, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	switch *format {
	case "cyclonedx":
		if err := writeCycloneDX(packages, *output); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating CycloneDX SBOM: %v\n", err)
			os.Exit(1)
		}
	case "spdx":
		if err := writeSPDX(packages, *output); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating SPDX SBOM: %v\n", err)
			os.Exit(1)
		}
	case "all":
		if err := writeCycloneDX(packages, *output); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating CycloneDX SBOM: %v\n", err)
			os.Exit(1)
		}
		if err := writeSPDX(packages, *output); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating SPDX SBOM: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown format %q. Use cyclonedx, spdx, or all.\n", *format)
		os.Exit(1)
	}
}

func writeCycloneDX(packages []model.Package, dir string) error {
	data, err := gencyclonedx.Generate(packages)
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "sbom.cdx.json")
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	fmt.Printf("CycloneDX SBOM written to %s\n", path)
	return nil
}

func writeSPDX(packages []model.Package, dir string) error {
	data, err := genspdx.Generate(packages)
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "sbom.spdx.json")
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	fmt.Printf("SPDX SBOM written to %s\n", path)
	return nil
}
