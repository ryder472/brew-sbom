package brew

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockCommandRunner struct {
	output []byte
	err    error
}

func (m *mockCommandRunner) Run(_ string, _ ...string) ([]byte, error) {
	return m.output, m.err
}

const sampleBrewJSON = `{
  "formulae": [
    {
      "name": "wget",
      "desc": "Internet file retriever",
      "homepage": "https://www.gnu.org/software/wget/",
      "license": "GPL-3.0-or-later",
      "installed": [{"version": "1.21.4"}]
    },
    {
      "name": "curl",
      "desc": "Get a file from an HTTP, HTTPS or FTP server",
      "homepage": "https://curl.se",
      "license": "curl",
      "installed": [{"version": "8.4.0"}]
    }
  ],
  "casks": []
}`

func TestInstalledPackages(t *testing.T) {
	mock := &mockCommandRunner{output: []byte(sampleBrewJSON)}
	runner := NewBrewRunner(mock)

	pkgs, err := runner.InstalledPackages()
	require.NoError(t, err)
	assert.Len(t, pkgs, 2)

	assert.Equal(t, "wget", pkgs[0].Name)
	assert.Equal(t, "1.21.4", pkgs[0].Version)
	assert.Equal(t, "GPL-3.0-or-later", pkgs[0].License)
	assert.Equal(t, "https://www.gnu.org/software/wget/", pkgs[0].Homepage)
	assert.Equal(t, "Internet file retriever", pkgs[0].Description)

	assert.Equal(t, "curl", pkgs[1].Name)
	assert.Equal(t, "8.4.0", pkgs[1].Version)
}
