package main

import (
	"log"
	"reflect"

	"github.com/RHEcosystemAppEng/SaaSi/deployer/pkg/config"
	"github.com/RHEcosystemAppEng/SaaSi/deployer/pkg/connect"
	"github.com/RHEcosystemAppEng/SaaSi/deployer/pkg/context"
	"github.com/RHEcosystemAppEng/SaaSi/deployer/pkg/deployer/app/deployer"
	"github.com/RHEcosystemAppEng/SaaSi/deployer/pkg/deployer/app/packager"
	"github.com/RHEcosystemAppEng/SaaSi/deployer/pkg/utils"
	"github.com/kr/pretty"
)

func main() {

	// parse command arguments from os into flags
	flagArgs := config.ParseFlagArgs()

	// expose REST POST service (placeholder)
	if !flagArgs.AsCli {
		log.Fatal("Note: start REST POST")
	}

	// validate that all required flags have been specified
	flagArgs.ValidateRequiredFlagArgs()

	// unmarshal deployer config and get cluster and application configs
	componentConfig := config.InitDeployerConfig(flagArgs.ConfigFile)
	pretty.Printf("Deploying the following configuration: \n%# v", componentConfig)

	// connect to cluster
	kubeConnection := connect.ConnectToCluster(componentConfig.ClusterConfig)

	// create deployer context to hold global variables
	deployerContext := context.InitDeployerContext(flagArgs, kubeConnection)

	// check if application deployment has been requested
	if !reflect.ValueOf(componentConfig.ApplicationConfig).IsZero() {

		// create application deployment package
		applicationPkg := packager.NewApplicationPkg(componentConfig.ApplicationConfig, deployerContext)

		// check if all mandatory variables have been set, else list unset vars and throw exception
		if len(applicationPkg.UnsetMandatoryParams) > 0 {
			log.Fatalf("ERROR: Please complete missing configuration for the following mandatory parameters (<FILEPATH>: <MANDATORY_PARAMETERS>.)\n%s", utils.StringifyMap(applicationPkg.UnsetMandatoryParams))
		}

		// deploy application deployment package
		deployer.DeployApplication(applicationPkg)

	} else {
		log.Println("No application to deploy")
	}
}
