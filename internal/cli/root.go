package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const TimeLayout = "2006-01-02 15:04:05"

// Environment variables
var ASCII bool
var Verbose bool
var KeystorePassword string
var RootKeystorePath string
var RootKeystoreAlias string
var CloudKeystorePath string
var CloudKeystoreAlias string
var SysOpsKeystorePath string
var Truststore string
var ArrowheadTLS bool
var ArrowheadAuthorizationHost string
var ArrowheadAuthorizationPort int
var ArrowheadServiceRegistryHost string
var ArrowheadServiceRegistryPort int
var ArrowheadOrchestratorHost string
var ArrowheadOrchestratorPort int

// Flags
var SystemName string
var ServiceID int
var SystemID int
var AuthID int
var AuthInfo bool
var Filter string
var Address string
var Port int
var HTTPMethod string
var ServiceDefinition string
var ServiceURI string
var ConsumerName string
var ProviderName string
var KeystorePath string
var Compact bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(envCmd)

	parseEnv()
}

var rootCmd = &cobra.Command{
	Use:   "arrowhead",
	Short: "arrowhead",
	Long:  "CLI to interact with Arrowhead Core Systems",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version",
	Long:  "Version",
	Run: func(cmd *cobra.Command, args []string) {
		ASCII = false
		ASCIIStr := os.Getenv("ARROWHEAD_CLI_ASCII")
		if ASCIIStr == "true" {
			ASCII = true
		}

		printVersionTable()
	},
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Environment",
	Long:  "Environment",
	Run: func(cmd *cobra.Command, args []string) {
		ASCII = false
		ASCIIStr := os.Getenv("ARROWHEAD_CLI_ASCII")
		if ASCIIStr == "true" {
			ASCII = true
		}

		printEnvTable()
	},
}
