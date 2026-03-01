package cyclonedx

import (
	"bytes"
	"fmt"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/ryder472/brew-sbom/internal/model"
)

// Generate produces a CycloneDX BOM in JSON format for the given packages.
func Generate(packages []model.Package) ([]byte, error) {
	components := make([]cdx.Component, 0, len(packages))
	for _, pkg := range packages {
		c := cdx.Component{
			Type:        cdx.ComponentTypeLibrary,
			Name:        pkg.Name,
			Version:     pkg.Version,
			Description: pkg.Description,
			ExternalReferences: externalRefs(pkg.Homepage),
		}
		if pkg.License != "" {
			licenses := cdx.Licenses{
				cdx.LicenseChoice{Expression: pkg.License},
			}
			c.Licenses = &licenses
		}
		components = append(components, c)
	}

	bom := cdx.NewBOM()
	bom.Components = &components

	var buf bytes.Buffer
	enc := cdx.NewBOMEncoder(&buf, cdx.BOMFileFormatJSON)
	enc.SetPretty(true)
	if err := enc.Encode(bom); err != nil {
		return nil, fmt.Errorf("encoding CycloneDX BOM: %w", err)
	}
	return buf.Bytes(), nil
}

func externalRefs(homepage string) *[]cdx.ExternalReference {
	if homepage == "" {
		return nil
	}
	refs := []cdx.ExternalReference{
		{Type: cdx.ERTypeWebsite, URL: homepage},
	}
	return &refs
}
