# brew-sbom

SBOM (Software Bill of Materials) generator for macOS Homebrew packages.

Generates SBOMs in **CycloneDX** (JSON, spec v1.4+) and **SPDX** (JSON, v2.3) formats from your Homebrew-installed packages.

[![CI](https://github.com/ryder472/brew-sbom/actions/workflows/ci.yml/badge.svg)](https://github.com/ryder472/brew-sbom/actions/workflows/ci.yml)

## Requirements

- macOS with [Homebrew](https://brew.sh) installed
- Go 1.22+ (for building from source)

## Installation

### Download a pre-built binary

Download the latest release from the [Releases page](https://github.com/ryder472/brew-sbom/releases) and place the binary in your `$PATH`:

```bash
# Apple Silicon (M1/M2/M3)
curl -L https://github.com/ryder472/brew-sbom/releases/latest/download/brew-sbom-darwin-arm64 -o /usr/local/bin/brew-sbom
chmod +x /usr/local/bin/brew-sbom

# Intel Mac
curl -L https://github.com/ryder472/brew-sbom/releases/latest/download/brew-sbom-darwin-amd64 -o /usr/local/bin/brew-sbom
chmod +x /usr/local/bin/brew-sbom
```

### Build from source

```bash
git clone https://github.com/ryder472/brew-sbom.git
cd brew-sbom
make build
# binary is at bin/brew-sbom
```

## Usage

```bash
# Output both CycloneDX and SPDX (default)
brew-sbom --format all

# CycloneDX only
brew-sbom --format cyclonedx

# SPDX only
brew-sbom --format spdx

# Specify output directory
brew-sbom --format all --output /path/to/output

# Print version
brew-sbom --version

# Print help
brew-sbom --help
```

### Output files

| Format    | Filename           |
|-----------|--------------------|
| CycloneDX | `sbom.cdx.json`    |
| SPDX      | `sbom.spdx.json`   |

## Development

```bash
# Run tests
make test

# Run linter
make lint

# Build
make build

# Clean build artifacts
make clean
```

## License

[MIT License](LICENSE)
