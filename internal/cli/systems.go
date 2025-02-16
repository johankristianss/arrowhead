package cli

import (
	"errors"

	"github.com/johankristianss/arrowhead/pkg/core"
	"github.com/johankristianss/arrowhead/pkg/rpc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(systemsCmd)
	systemsCmd.AddCommand(lsSystemCmd)
	systemsCmd.AddCommand(getSystemCmd)
	systemsCmd.AddCommand(registerSystemCmd)
	systemsCmd.AddCommand(unRegisterSystemCmd)

	lsSystemCmd.Flags().StringVarP(&Filter, "filter", "", "", "Filter systems by provider System Name")

	getSystemCmd.Flags().IntVarP(&SystemID, "id", "i", -1, "System ID")
	getSystemCmd.MarkFlagRequired("id")

	registerSystemCmd.Flags().StringVarP(&SystemName, "name", "n", "", "System Name")
	getSystemCmd.MarkFlagRequired("name")
	registerSystemCmd.Flags().StringVarP(&Address, "address", "a", "", "System Address")
	getSystemCmd.MarkFlagRequired("address")
	registerSystemCmd.Flags().IntVarP(&Port, "port", "p", -1, "System Port")
	getSystemCmd.MarkFlagRequired("port")

	unRegisterSystemCmd.Flags().IntVarP(&SystemID, "id", "i", -1, "System ID")
	getSystemCmd.MarkFlagRequired("id")
}

var systemsCmd = &cobra.Command{
	Use:   "systems",
	Short: "Manage Arrowhead systems",
	Long:  "Manage Arrowhead systems",
}

var lsSystemCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available systems",
	Long:  "List available systems",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()
		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)
		systems, err := arrowhead.Management.GetSystems()

		if len(systems) == 0 {
			log.Info("No systems found")
			return
		}

		if Filter != "" {
			systems = filterSystemByName(systems, Filter)
		}

		CheckError(err)
		printSystemsTable(systems)
	},
}

var getSystemCmd = &cobra.Command{
	Use:   "get",
	Short: "Get info about a system",
	Long:  "Get info about a system",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()

		if SystemID == -1 {
			CheckError(errors.New("System ID is required"))
		}

		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)
		system, err := arrowhead.Management.GetSystemByID(SystemID)
		CheckError(err)

		printSystemTable(system)
	},
}

var registerSystemCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a system",
	Long:  "Register a system",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()

		if SystemName == "" {
			CheckError(errors.New("System Name is required"))
		}

		if Address == "" {
			CheckError(errors.New("Address is required"))
		}

		if Port == -1 {
			CheckError(errors.New("Port is required"))
		}

		authInfo, err := generateCert(SystemName)
		CheckError(err)

		systemReg := core.SystemRegistration{
			Address:            Address,
			AuthenticationInfo: authInfo,
			Metadata:           map[string]string{},
			Port:               Port,
			SystemName:         SystemName,
		}

		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)

		_, err = arrowhead.Management.RegisterSystem(systemReg)
		CheckError(err)

		//err = createSystemEnvFile(SystemName, Address, Port, "./"+SystemName+".p12", KeystorePassword, "./"+SystemName+".pem", "./"+SystemName+".key")
		err = createSystemEnvFile(SystemName, Address, Port, "./"+SystemName+".p12", KeystorePassword)
		CheckError(err)

		//log.Info("System registered successfully, PKCS#12 certificate stored in ./" + SystemName + ".p12 and ./" + SystemName + ".pub and corresponding PEM certificate stored in ./" + SystemName + ".pem and ./" + SystemName + ".key, config file stored in ./" + SystemName + ".env")
		log.Info("System registered successfully, PKCS#12 certificate stored in ./" + SystemName + ".p12 and ./" + SystemName + ".pub, config file stored in ./" + SystemName + ".env")
	},
}

var unRegisterSystemCmd = &cobra.Command{
	Use:   "unregister",
	Short: "Unregister a system",
	Long:  "Unregister a system",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()

		if SystemID == -1 {
			CheckError(errors.New("System ID is required"))
		}

		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)

		err = arrowhead.Management.UnregisterSystemByID(SystemID)
		CheckError(err)

		log.Info("Successfully unregisted system with ID ", SystemID)
	},
}
