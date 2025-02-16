package cli

import (
	"strconv"

	"github.com/johankristianss/arrowhead/internal/table"
	"github.com/johankristianss/arrowhead/pkg/core"
)

func printOrchestrationResponseCompact(orchestrationResponse core.OrchestrationResponse) {
	theme, err := table.LoadTheme("solarized-dark")
	CheckError(err)

	for _, serviceResponse := range orchestrationResponse.Response {
		printKeyValueTable("Orchestration Result", theme, map[string]string{
			"Service URI":          serviceResponse.ServiceURI,
			"Service Definition":   serviceResponse.ServiceDefinition.ServiceDefinition,
			"Provider System Name": serviceResponse.Provider.SystemName,
			"Provider Address":     serviceResponse.Provider.Address,
			"Provider Port":        strconv.Itoa(serviceResponse.Provider.Port),
		})

		// Print authorization tokens if present
		if len(serviceResponse.AuthorizationTokens) > 0 {
			printKeyValueTable("Authorization Tokens", theme, serviceResponse.AuthorizationTokens)
		}

		// Print each interface separately
		for _, iface := range serviceResponse.Interfaces {
			printKeyValueTable("Interface", theme, map[string]string{
				"ID":             strconv.Itoa(iface.ID),
				"Interface Name": iface.InterfaceName,
				"Created At":     iface.CreatedAt.Format(TimeLayout),
				"Updated At":     iface.UpdatedAt.Format(TimeLayout),
			})
		}
	}
}

func printOrchestrationResponse(orchestrationResponse core.OrchestrationResponse) {
	theme, err := table.LoadTheme("solarized-dark")
	CheckError(err)

	for _, serviceResponse := range orchestrationResponse.Response {
		printKeyValueTable("Service Response", theme, map[string]string{
			"Service URI": serviceResponse.ServiceURI,
			"Secure":      serviceResponse.Secure,
			"Version":     strconv.Itoa(serviceResponse.Version),
		})

		printKeyValueTable("Service Definition", theme, map[string]string{
			"ID":                 strconv.Itoa(serviceResponse.ServiceDefinition.ID),
			"Service Definition": serviceResponse.ServiceDefinition.ServiceDefinition,
			"Created At":         serviceResponse.ServiceDefinition.CreatedAt.Format(TimeLayout),
			"Updated At":         serviceResponse.ServiceDefinition.UpdatedAt.Format(TimeLayout),
		})

		printKeyValueTable("Provider", theme, map[string]string{
			"ID":                  strconv.Itoa(serviceResponse.Provider.ID),
			"System Name":         serviceResponse.Provider.SystemName,
			"Address":             serviceResponse.Provider.Address,
			"Port":                strconv.Itoa(serviceResponse.Provider.Port),
			"Authentication Info": serviceResponse.Provider.AuthenticationInfo,
			"Created At":          serviceResponse.Provider.CreatedAt.Format(TimeLayout),
			"Updated At":          serviceResponse.Provider.UpdatedAt.Format(TimeLayout),
		})

		// Print metadata as a separate table
		if len(serviceResponse.Metadata) > 0 {
			printKeyValueTable("Metadata", theme, serviceResponse.Metadata)
		}

		// Print authorization tokens if present
		if len(serviceResponse.AuthorizationTokens) > 0 {
			printKeyValueTable("Authorization Tokens", theme, serviceResponse.AuthorizationTokens)
		}

		// Print warnings if any
		if len(serviceResponse.Warnings) > 0 {
			for i, warning := range serviceResponse.Warnings {
				printKeyValueTable("Warning "+strconv.Itoa(i+1), theme, map[string]string{"Message": warning})
			}
		}

		// Print each interface separately
		for _, iface := range serviceResponse.Interfaces {
			printKeyValueTable("Interface", theme, map[string]string{
				"ID":             strconv.Itoa(iface.ID),
				"Interface Name": iface.InterfaceName,
				"Created At":     iface.CreatedAt.Format(TimeLayout),
				"Updated At":     iface.UpdatedAt.Format(TimeLayout),
			})
		}
	}
}
