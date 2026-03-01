package spdx

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ryder472/brew-sbom/internal/model"
	spdxcommon "github.com/spdx/tools-golang/spdx/v2/common"
	spdxv23 "github.com/spdx/tools-golang/spdx/v2/v2_3"
)

// Generate produces an SPDX 2.3 document in JSON format for the given packages.
func Generate(packages []model.Package) ([]byte, error) {
	pkgs := make([]*spdxv23.Package, 0, len(packages))
	for i, pkg := range packages {
		p := &spdxv23.Package{
			PackageName:             pkg.Name,
			PackageSPDXIdentifier:   spdxcommon.ElementID(fmt.Sprintf("Package-%d-%s", i, pkg.Name)),
			PackageVersion:          pkg.Version,
			PackageDownloadLocation: homepage(pkg.Homepage),
			PackageHomePage:         pkg.Homepage,
			FilesAnalyzed:           false,
			PackageLicenseConcluded: licenseValue(pkg.License),
			PackageLicenseDeclared:  licenseValue(pkg.License),
		}
		pkgs = append(pkgs, p)
	}

	doc := &spdxv23.Document{
		SPDXVersion:       spdxv23.Version,
		DataLicense:       spdxv23.DataLicense,
		SPDXIdentifier:    "DOCUMENT",
		DocumentName:      "brew-sbom",
		DocumentNamespace: fmt.Sprintf("https://github.com/ryder472/brew-sbom/%d", time.Now().UnixNano()),
		CreationInfo: &spdxv23.CreationInfo{
			Created: time.Now().UTC().Format(time.RFC3339),
			Creators: []spdxcommon.Creator{
				{CreatorType: "Tool", Creator: "brew-sbom"},
			},
		},
		Packages: pkgs,
	}

	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("encoding SPDX document: %w", err)
	}
	return data, nil
}

func homepage(h string) string {
	if h != "" {
		return h
	}
	return "NOASSERTION"
}

func licenseValue(l string) string {
	if l != "" {
		return l
	}
	return "NOASSERTION"
}
