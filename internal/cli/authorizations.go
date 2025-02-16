package cli

import (
	"errors"

	"github.com/johankristianss/arrowhead/pkg/rpc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authsCmd)
	authsCmd.AddCommand(lsAuthsCmd)
	authsCmd.AddCommand(removeAuthCmd)
	authsCmd.AddCommand(addAuthCmd)

	addAuthCmd.Flags().StringVarP(&ConsumerName, "consumer", "c", "", "Consumer name")
	addAuthCmd.MarkFlagRequired("consumer")
	addAuthCmd.Flags().StringVarP(&ProviderName, "provider", "p", "", "Provider name")
	addAuthCmd.MarkFlagRequired("provider")
	addAuthCmd.Flags().StringVarP(&ServiceDefinition, "service", "s", "", "Service definition")
	addAuthCmd.MarkFlagRequired("service")

	removeAuthCmd.Flags().IntVarP(&AuthID, "id", "i", -1, "Authorization ID")
	removeAuthCmd.MarkFlagRequired("id")
}

var authsCmd = &cobra.Command{
	Use:   "auths",
	Short: "Manage Arrowhead authorizations",
	Long:  "Manage Arrowhead authorizations",
}

var lsAuthsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List authorizations rules",
	Long:  "List authorizations rules",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()
		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)
		auths, err := arrowhead.Management.GetAuthorizations()
		CheckError(err)
		if len(auths) == 0 {
			log.Info("No authorizations found")
			return
		}
		printAuthsTable(auths)
	},
}

var addAuthCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an authorization rule",
	Long:  "Add an authorization rule",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()
		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)

		if ConsumerName == "" {
			CheckError(errors.New("Consumer name is required"))
		}
		if ProviderName == "" {
			CheckError(errors.New("Producer name is required"))
		}
		if ServiceDefinition == "" {
			CheckError(errors.New("Service definition is required"))
		}

		auth, err := arrowhead.Management.AddAuthorization(ConsumerName, ProviderName, ServiceDefinition)
		CheckError(err)

		log.WithFields(log.Fields{"AuthID": auth.ID}).Info("Authorization added")
	},
}

var removeAuthCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an authorization rule",
	Long:  "Remove an authorization rule",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()
		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)

		if AuthID == -1 {
			CheckError(errors.New("Authorization ID is required"))
		}

		err = arrowhead.Management.RemoveAuthorization(AuthID)
		CheckError(err)

		log.WithFields(log.Fields{"AuthID": AuthID}).Info("Authorization removed")
	},
}
