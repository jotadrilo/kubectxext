package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/juju/errors"
	configapi "k8s.io/client-go/tools/clientcmd/api/v1"
	"sigs.k8s.io/yaml"
)

// Load reads the provided `f` kubernetes config file and returns a new instance
// of it.
func Load(f string) (*configapi.Config, error) {
	r, err := os.Open(f)
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, errors.Trace(err)
	}

	config := &configapi.Config{}
	return config, yaml.Unmarshal(data, config)
}

// Print prints the `c` config object into the `f` file.
func Print(c *configapi.Config) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return errors.Trace(err)
	}
	fmt.Printf("%s", string(data))
	return nil
}
