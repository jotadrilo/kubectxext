package splitter

import (
	configapi "k8s.io/client-go/tools/clientcmd/api/v1"
)

// ByContext returns a new Config object containing the resources
// attached to the provided `name` context.
func ByContext(cfg *configapi.Config, name string) *configapi.Config {
	config := &configapi.Config{}

	auths := []configapi.NamedAuthInfo{}
	clusters := []configapi.NamedCluster{}
	contexts := []configapi.NamedContext{}
	for _, ctx := range cfg.Contexts {
		if ctx.Name == name {
			contexts = append(contexts, ctx)
			auths = append(auths, GetAuthInfoByName(cfg, ctx.Context.AuthInfo))
			clusters = append(clusters, GetClusterByName(cfg, ctx.Context.Cluster))
		}
	}

	config.AuthInfos = auths
	config.Clusters = clusters
	config.Contexts = contexts
	config.CurrentContext = name

	return config
}

// ByContexts returns a new Config object containing the resources
// attached to the provided `names` contexts.
func ByContexts(cfg *configapi.Config, names ...string) *configapi.Config {
	config := &configapi.Config{}
	for _, name := range names {
		c := ByContext(cfg, name)
		config.AuthInfos = append(config.AuthInfos, c.AuthInfos...)
		config.Clusters = append(config.Clusters, c.Clusters...)
		config.Contexts = append(config.Contexts, c.Contexts...)
		config.CurrentContext = name
	}
	return config
}

// GetAuthInfoByName returns the AuthInfo that matches the `name`
func GetAuthInfoByName(cfg *configapi.Config, name string) configapi.NamedAuthInfo {
	for _, auth := range cfg.AuthInfos {
		if auth.Name == name {
			return auth
		}
	}
	return configapi.NamedAuthInfo{}
}

// GetClusterByName returns the Cluster that matches the `name`
func GetClusterByName(cfg *configapi.Config, name string) configapi.NamedCluster {
	for _, cluster := range cfg.Clusters {
		if cluster.Name == name {
			return cluster
		}
	}
	return configapi.NamedCluster{}
}
