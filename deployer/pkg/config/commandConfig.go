package config

import (
	"log"
	"strings"

	"github.com/jessevdk/go-flags"
)

type FlagArgs struct {
	AsCli         bool   `long:"as-cli" description:"Indicator for the application to run from CLI"`
	ConfigFile    string `short:"f" long:"filename" description:"Application configuration file for deployemnt"`
	RootOutputDir string `short:"o" long:"output-dir" default:"output" description:"Root output folder"`
	RootSourceDir string `short:"s" long:"source-dir" description:"Root source folder"`
}

func ParseFlagArgs() *FlagArgs {

	// init flags for input arguments
	flagArgs := FlagArgs{}

	// parse input arguments from os into flags
	_, err := flags.Parse(&flagArgs)
	if err != nil {
		log.Fatalf("Failed to parse os input arguments: %s", err)
	}

	return &flagArgs
}

func (flagArgs *FlagArgs) ValidateRequiredFlagArgs() {

	names := make([]string, 0, 2)

	if flagArgs.ConfigFile == "" {
		names = append(names, "`-f, --filename`")
	}
	if flagArgs.RootSourceDir == "" {
		names = append(names, "`-s, --source-dir`")
	}

	if len(names) == 0 {
		return
	} else if len(names) == 1 {
		log.Fatalf("The required flag %s was not specified", names[0])
	} else {
		log.Fatalf("The required flags %s and %s were not specified",
			strings.Join(names[:len(names)-1], ", "), names[len(names)-1])
	}

}
