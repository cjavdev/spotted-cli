# Spotted CLI

The official CLI for the [Spotted REST API](https://spotted.stldocs.com).

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Homebrew

```sh
brew tap cjavdev/spotted-cli
brew install spotted
```

### Installing with Go

<!-- x-release-please-start-version -->

```sh
go install 'github.com/cjavdev/spotted-cli/cmd/spotted@latest'
```

### Running Locally

<!-- x-release-please-start-version -->

```sh
go run cmd/spotted/main.go
```

<!-- x-release-please-end -->

## Usage

The CLI follows a resource-based command structure:

```sh
spotted [resource] [command] [flags]
```

```sh
spotted albums retrieve \
  --id 4aawyAB9vmqN3uQ7FjRGTy
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
