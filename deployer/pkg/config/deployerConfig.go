package config

import (
	"io/ioutil"
	"log"
	"net/http"

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
	ClusterName       string `yaml:"CLUSTER_NAME"`
	ClusterBaseDomain string `yaml:"CLUSTER_BASE_DOMAIN"`
	WorkerCount       string `yaml:"WORKER_COUNT"`
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

func InitDeployerConfig(byteConfig []byte) *ComponentConfig {

	// Unmarshal deployer config
	config := DeployerConfig{}
	err := yaml.Unmarshal(byteConfig, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config.Deployer
}

func ReadConfigFile(configFile string) {
	// read deployer config file
	yfile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	InitDeployerConfig(yfile)
}

func ReadRestRequest(w http.ResponseWriter, r *http.Request) {
	// read deployer config file
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err)
	}

	InitDeployerConfig(body)
}
