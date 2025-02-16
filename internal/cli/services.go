package cli

import (
	"errors"
	"fmt"

	"github.com/johankristianss/arrowhead/pkg/rpc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(servicesCmd)
	servicesCmd.AddCommand(lsServicesCmd)
	servicesCmd.AddCommand(getServiceCmd)
	servicesCmd.AddCommand(registerServiceCmd)
	servicesCmd.AddCommand(unregisterServiceCmd)

	lsServicesCmd.Flags().StringVarP(&Filter, "filter", "", "", "Filter services by System Name")

	getServiceCmd.Flags().IntVarP(&ServiceID, "id", "i", -1, "Service ID")
	getServiceCmd.MarkFlagRequired("id")
	getServiceCmd.Flags().BoolVarP(&AuthInfo, "authinfo", "", false, "Fetch Authentication Info")

	registerServiceCmd.Flags().StringVarP(&SystemName, "system", "s", "", "System Name")
	registerServiceCmd.MarkFlagRequired("system")
	registerServiceCmd.Flags().StringVarP(&ServiceDefinition, "definition", "d", "", "Service Definition")
	registerServiceCmd.MarkFlagRequired("definition")
	registerServiceCmd.Flags().StringVarP(&ServiceURI, "uri", "u", "", "Service URI")
	registerServiceCmd.MarkFlagRequired("uri")
	registerServiceCmd.Flags().StringVarP(&HTTPMethod, "method", "m", "GET", "HTTP Method")
	registerServiceCmd.MarkFlagRequired("method")

	unregisterServiceCmd.Flags().IntVarP(&ServiceID, "id", "i", -1, "System ID")
	unregisterServiceCmd.MarkFlagRequired("id")
}

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Manage Arrowhead services",
	Long:  "Manage Arrowhead services",
}

var lsServicesCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available services",
	Long:  "List available services",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()
		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)
		services, err := arrowhead.Management.GetServices()
		CheckError(err)

		if len(services) == 0 {
			log.Info("No services found")
			return
		}

		if Filter != "" {
			services = filterServiceByName(services, Filter)
		}

		printServicesTable(services)
	},
}

var getServiceCmd = &cobra.Command{
	Use:   "get",
	Short: "Get info about a service",
	Long:  "Get info about a service",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()

		if ServiceID == -1 {
			CheckError(errors.New("Service ID is required"))
		}

		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)
		service, err := arrowhead.Management.GetServiceByID(ServiceID)
		CheckError(err)

		if AuthInfo {
			fmt.Println(service.Provider.AuthenticationInfo)
		} else {
			printServiceTable(service)
		}
	},
}

var registerServiceCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a service",
	Long:  "Register a service",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()

		if SystemName == "" {
			CheckError(errors.New("System Name is required"))
		}

		if ServiceDefinition == "" {
			CheckError(errors.New("Service Definition is required"))
		}

		if ServiceURI == "" {
			CheckError(errors.New("Service URI is required"))
		}

		if HTTPMethod != "GET" && HTTPMethod != "POST" && HTTPMethod != "PUT" && HTTPMethod != "DELETE" {
			log.WithField("HTTPMethod", HTTPMethod).Error("Invalid HTTP Method, must be GET, POST, PUT or DELETE")
			CheckError(errors.New("Invalid HTTP Method"))
		}

		var httpMethod int
		switch HTTPMethod {
		case "GET":
			httpMethod = 0
		case "POST":
			httpMethod = 1
		case "PUT":
			httpMethod = 2
		case "DELETE":
			httpMethod = 3
		}

		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)

		system, err := arrowhead.Management.GetSystemByName(SystemName)
		CheckError(err)

		_, err = arrowhead.Management.RegisterService(&system, httpMethod, ServiceDefinition, ServiceURI)
		CheckError(err)

		log.WithFields(log.Fields{"SystemName": system.SystemName, "ServiceDefinition": ServiceDefinition, "ServiceURI": ServiceURI, "HTTPMethod": HTTPMethod}).Info("Service registered")
	},
}

var unregisterServiceCmd = &cobra.Command{
	Use:   "unregister",
	Short: "Unregister a service",
	Long:  "Unregister a service",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()
		config := generateRPCSysOpsConfig()

		if ServiceID == -1 {
			CheckError(errors.New("System ID is required"))
		}

		arrowhead, err := rpc.CreateArrowhead(config)
		CheckError(err)

		err = arrowhead.Management.UnregisterService(ServiceID)
		CheckError(err)

		log.WithFields(log.Fields{"ServiceID": ServiceID}).Info("Service unregistered")
	},
}
