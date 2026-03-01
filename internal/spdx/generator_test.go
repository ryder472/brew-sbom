package spdx

import (
	"encoding/json"
	"testing"

	"github.com/ryder472/brew-sbom/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	packages := []model.Package{
		{Name: "wget", Version: "1.21.4", License: "GPL-3.0-or-later"},
		{Name: "curl", Version: "8.4.0", License: "curl"},
	}

	out, err := Generate(packages)
	require.NoError(t, err)

	var result map[string]interface{}
	require.NoError(t, json.Unmarshal(out, &result))

	assert.Equal(t, "SPDX-2.3", result["spdxVersion"])
	assert.Equal(t, "CC0-1.0", result["dataLicense"])
	assert.NotEmpty(t, result["documentNamespace"])

	pkgs, ok := result["packages"].([]interface{})
	require.True(t, ok)
	assert.Len(t, pkgs, 2)

	first := pkgs[0].(map[string]interface{})
	assert.Equal(t, "wget", first["name"])
	assert.Equal(t, "1.21.4", first["versionInfo"])
}
