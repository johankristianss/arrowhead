package cli

import (
	"github.com/johankristianss/arrowhead/pkg/rpc"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(orchestrateCmd)

	orchestrateCmd.Flags().StringVarP(&SystemName, "system", "s", "", "System name")
	orchestrateCmd.MarkFlagRequired("system")
	orchestrateCmd.Flags().StringVarP(&Address, "address", "a", "", "Address")
	orchestrateCmd.MarkFlagRequired("address")
	orchestrateCmd.Flags().IntVarP(&Port, "port", "p", 0, "Port")
	orchestrateCmd.MarkFlagRequired("port")
	orchestrateCmd.Flags().StringVarP(&ServiceDefinition, "service", "d", "", "Service definition")
	orchestrateCmd.MarkFlagRequired("service")
	orchestrateCmd.Flags().StringVarP(&KeystorePath, "keystore", "k", "", "Keystore path")
	orchestrateCmd.MarkFlagRequired("keystore")
	orchestrateCmd.Flags().StringVarP(&KeystorePassword, "password", "w", "", "Keystore password")
	orchestrateCmd.MarkFlagRequired("password")
	orchestrateCmd.Flags().BoolVarP(&Compact, "compact", "c", false, "Compact output")
}

var orchestrateCmd = &cobra.Command{
	Use:   "orchestrate",
	Short: "Orchestrate a service",
	Long:  "Orchestrate a service",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()
		config.KeystorePath = KeystorePath
		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)

		orchestrationRequest := rpc.BuildOrchestrationRequest(SystemName, Address, Port, ServiceDefinition)
		orchestrationResponse, err := arrowhead.Client.Orchestrate(orchestrationRequest)
		CheckError(err)

		if Compact {
			printOrchestrationResponseCompact(orchestrationResponse)
			return
		}
		printOrchestrationResponse(orchestrationResponse)
	},
}
