package brew

import (
	"encoding/json"
	"os/exec"

	"github.com/ryder472/brew-sbom/internal/model"
)

// CommandRunner is an interface for running external commands.
type CommandRunner interface {
	Run(name string, args ...string) ([]byte, error)
}

// ExecRunner is the real CommandRunner that uses os/exec.
type ExecRunner struct{}

func (e ExecRunner) Run(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).Output()
}

// Runner is the interface for fetching installed packages.
type Runner interface {
	InstalledPackages() ([]model.Package, error)
}

// BrewRunner implements Runner using the brew CLI.
type BrewRunner struct {
	cmd CommandRunner
}

// NewBrewRunner creates a new BrewRunner with the given CommandRunner.
func NewBrewRunner(cmd CommandRunner) *BrewRunner {
	return &BrewRunner{cmd: cmd}
}

type brewOutput struct {
	Formulae []brewFormula `json:"formulae"`
}

type brewFormula struct {
	Name      string          `json:"name"`
	Desc      string          `json:"desc"`
	Homepage  string          `json:"homepage"`
	License   string          `json:"license"`
	Installed []brewInstalled `json:"installed"`
}

type brewInstalled struct {
	Version string `json:"version"`
}

// InstalledPackages returns all installed Homebrew packages.
func (b *BrewRunner) InstalledPackages() ([]model.Package, error) {
	out, err := b.cmd.Run("brew", "info", "--json=v2", "--installed")
	if err != nil {
		return nil, err
	}

	var brewOut brewOutput
	if err := json.Unmarshal(out, &brewOut); err != nil {
		return nil, err
	}

	packages := make([]model.Package, 0, len(brewOut.Formulae))
	for _, f := range brewOut.Formulae {
		version := ""
		if len(f.Installed) > 0 {
			version = f.Installed[0].Version
		}
		packages = append(packages, model.Package{
			Name:        f.Name,
			Version:     version,
			License:     f.License,
			Homepage:    f.Homepage,
			Description: f.Desc,
		})
	}
	return packages, nil
}
