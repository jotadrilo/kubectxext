# kubectxext

The `


-----

# kubectxext

kubectxext is a utility to extract a valid kubernetes config file from your kubernetes config for a given context.

```
USAGE:
  kubectxext                    : prompt a select list using the information in your kubernetes config file and extracts a kube config file for the selected context.
  kubectxext -context <NAME>    : extracts a kube config file for the provided context name.
  kubectxext -kubeconfig <FILE> : specify the kube config file to work with (`~/.kube/config` by default)
```


### Usage

```sh
$ kubectxext -context my-context
Context "test" set.
Active namespace is "kube-system".

$ kubens -
Context "test" set.
Active namespace is "default".
```

`kubens` also supports <kbd>Tab</kbd> completion on bash/zsh/fish shells.

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

    brew install kubectx

#### go

    go get -u github.com/jotadrilo/kubectxext/cmd
