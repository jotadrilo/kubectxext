# kubectxext

`kubectxext` is a utility to extract a valid kubernetes config file from your kubernetes config for a given context.

## Usage

![](img/kubectxext-demo.gif)

```
Usage of kubectxext:
  -context string
    	context to extract.
  -kubeconfig string
    	kube config file to extract from. (default "/Users/jotadrilo/.kube/config")
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
