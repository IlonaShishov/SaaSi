package config

import (
	"errors"
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"
)

// ----------------------
// ----Deployer Config----
// ----------------------

type DeployerConfig struct {
	Deployer ComponentConfig `yaml:"deployer"`
}

type ComponentConfig struct {
	ClusterConfig     ClusterConfig     `yaml:"cluster"`
	ApplicationConfig ApplicationConfig `yaml:"application"`
}

// ----------------------
// ----Cluster Config----
// ----------------------

type ClusterConfig struct {
	Server        string        `yaml:"server"`
	User          string        `yaml:"user"`
	Token         string        `yaml:"token"`
	FromClusterId string        `yaml:"fromClusterId"`
	ClusterId     string        `yaml:"clusterId"`
	Aws           AwsSettings   `yaml:"aws"`
	Params        ClusterParams `yaml:"params"`
}

type AwsSettings struct {
	AwsPublicDomain    string `yaml:"aws_public_domain"`
	AwsAccountName     string `yaml:"aws_account_name"`
	AwsAccessKeyId     string `yaml:"aws_access_key_id"`
	AwsSecretAccessKey string `yaml:"aws_secret_access_key"`
}

type ClusterParams struct {
	ClusterName           string `yaml:"CLUSTER_NAME",omitempty`
	ClusterBaseDomain     string `yaml:"CLUSTER_BASE_DOMAIN",omitempty`
	WorkerCount           string `yaml:"WORKER_COUNT",omitempty`
	ClusterVersion        string `yaml:"CLUSTER_VERSION",omitempty`
	ClusterNetwork        string `yaml:"CLUSTER_NETWORK",omitempty`
	HostPrefix            string `yaml:"HOST_PREFIX",omitempty`
	ServiceNetwork        string `yaml:"SERVICE_NETWORK",omitempty`
	NetworkType           string `yaml:"NETWORK_TYPE",omitempty`
	RegistryRouteHostname string `yaml:"REGISTRY_ROUTE_HOSTNAME",omitempty`
	RegistryIsExposed     string `yaml:"REGISTRY_IS_EXPOSED",omitempty`
	ProvCloudProvider     string `yaml:"PROV_CLOUD_PROVIDER",omitempty`
	ProvCloudRegion       string `yaml:"PROV_CLOUD_REGION",omitempty`
	MasterCount           string `yaml:"MASTER_COUNT",omitempty`
}

// ----------------------
// ------App Config------
// ----------------------

type ApplicationConfig struct {
	Name                   string       `yaml:"name"`
	NamespaceMappingFormat string       `yaml:"namespaceMappingFormat"`
	Namespaces             []Namespaces `yaml:"namespaces"`
}

type Namespaces struct {
	Name       string       `yaml:"name"`
	Target     string       `yaml:"target"`
	ConfigMaps []ConfigMaps `yaml:"params"`
	Secrets    []Secrets    `yaml:"secrets"`
}

type ConfigMaps struct {
	ConfigMap string              `yaml:"configMap"`
	Params    []ApplicationParams `yaml:"params"`
}

type Secrets struct {
	Secret string              `yaml:"secret"`
	Params []ApplicationParams `yaml:"params"`
}

type ApplicationParams struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func InitDeployerConfig(configFile string) *ComponentConfig {

	// read deployer config file
	yfile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal deployer config
	config := DeployerConfig{}
	err = yaml.Unmarshal(yfile, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config.Deployer
}

func (e *ComponentConfig) Validate() error {
	if reflect.ValueOf(e.ClusterConfig).IsZero() {
		return errors.New("missing cluster configuration")
	} else {
		if e.ClusterConfig.ClusterId == "" {
			return errors.New("missing clusterId configuration")
		}
		if e.ClusterConfig.Server == "" {
			return errors.New("missing server configuration")
		}
		if e.ClusterConfig.Token == "" {
			return errors.New("missing token configuration")
		}
	}
	if reflect.ValueOf(e.ApplicationConfig).IsZero() {
		return errors.New("missing application configuration")
	} else {
		if e.ApplicationConfig.Name == "" {
			return errors.New("missing application name configuration")
		}
	}
	return nil
}
