package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	config "github.com/jotadrilo/kubectxext/internal/configutil"
	"github.com/jotadrilo/kubectxext/internal/splitter"
	"github.com/jotadrilo/kubectxext/internal/tui"
	"github.com/juju/errors"
)

var (
	kubeconfig = flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "kube config file to extract from.")
	context    = flag.String("context", "", "context to extract.")
)

func validateFlags() error {
	if *kubeconfig == "" {
		return errors.Errorf("invalid input file: %q", *kubeconfig)
	}
	return nil
}

func main() {
	flag.Parse()
	if err := validateFlags(); err != nil {
		flag.Usage()
		log.Fatal(err)
	}

	c, err := config.Load(*kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Prompt select table if the context is not provided
	if *context == "" {
		contexts := []tui.Context{}
		for _, ctx := range c.Contexts {
			contexts = append(contexts, tui.Context{
				Name:    ctx.Name,
				Cluster: ctx.Context.Cluster,
				Auth:    ctx.Context.AuthInfo,
			})
		}
		prompt := tui.ContextSelect(contexts)
		i, _, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}
		*context = contexts[i].Name
	}

	if err := config.Print(splitter.ByContexts(c, *context)); err != nil {
		log.Fatal(err)
	}
}
