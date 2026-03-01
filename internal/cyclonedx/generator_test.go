package cyclonedx

import (
	"encoding/json"
	"testing"

	"github.com/ryder472/brew-sbom/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	packages := []model.Package{
		{Name: "wget", Version: "1.21.4", License: "GPL-3.0-or-later", Homepage: "https://www.gnu.org/software/wget/", Description: "Internet file retriever"},
		{Name: "curl", Version: "8.4.0", License: "curl", Homepage: "https://curl.se", Description: "Transfer data with URLs"},
	}

	out, err := Generate(packages)
	require.NoError(t, err)

	var result map[string]interface{}
	require.NoError(t, json.Unmarshal(out, &result))

	assert.Equal(t, "CycloneDX", result["bomFormat"])

	specVersion, ok := result["specVersion"].(string)
	require.True(t, ok)
	assert.GreaterOrEqual(t, specVersion, "1.4")

	components, ok := result["components"].([]interface{})
	require.True(t, ok)
	assert.Len(t, components, 2)

	first := components[0].(map[string]interface{})
	assert.Equal(t, "wget", first["name"])
	assert.Equal(t, "1.21.4", first["version"])
}
