# kubectxext

The `kubectxext` utility is meant to be used for extracting specific kube config files from large kube config files.
kubectxext is a utility to extract a valid kubernetes config file from your kubernetes config for a given context.

## Usage

```
USAGE:
  kubectxext                    : prompt a select list using the information in your kubernetes config file and extracts a kube config file for the selected context.
  kubectxext -context <NAME>    : extracts a kube config file for the provided context name.
  kubectxext -kubeconfig <FILE> : specify the kube config file to work with (`~/.kube/config` by default)
```

## Installation

There are several installation options:

- macOS
  - Homebrew (recommended)
  - Go
- Linux
  - Go

### macOS

#### Homebrew

If you use [Homebrew](https://brew.sh/) you can install like this:

    brew tap jotadrilo/tap
    brew install jotadrilo/tap/kubectxext

#### go

    go get github.com/jotadrilo/kubectxext/cmd/...

### Linux

#### go

    go get github.com/jotadrilo/kubectxext/cmd/...
