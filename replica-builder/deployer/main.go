package main

import (
	"log"
	"os"
	"reflect"

	"github.com/RHEcosystemAppEng/SaaSi/replica-builder/deployer/pkg/config"
	"github.com/RHEcosystemAppEng/SaaSi/replica-builder/deployer/pkg/deployer"
	"github.com/RHEcosystemAppEng/SaaSi/replica-builder/deployer/pkg/packager"
	"github.com/kr/pretty"
)

var (
	err error
)

func main() {

	// get deployer config yaml as input
	if len(os.Args) != 2 {
		log.Fatal("Expected 1 argument, got ", len(os.Args)-1)
	}

	// Unmarshal deployer config and get cluster and application configs
	componentConfig := config.ReadDeployerConfig(os.Args[1])
	pretty.Printf("Deploying the following configuration: \n%# v", componentConfig)

	//
	// TODO - create and deploy cluster
	//

	// check if application deployment has been requested
	if !reflect.ValueOf(componentConfig.Application).IsZero() {

		// create application deployment package
		applicationPkg := packager.NewApplicationPkg(componentConfig.Application)

		// deploy application deployment package
		deployer.DeployApplication(applicationPkg)

	} else {
		log.Println("No application to deploy")
	}
}
