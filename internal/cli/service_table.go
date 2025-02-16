package cli

import (
	"strconv"

	"github.com/johankristianss/arrowhead/internal/table"
	"github.com/johankristianss/arrowhead/pkg/core"
)

func printServiceTable(service core.Service) {
	theme, err := table.LoadTheme("solarized-dark")
	CheckError(err)

	printKeyValueTable("Service Detials", theme, map[string]string{
		"ID":          strconv.Itoa(service.ID),
		"Service URI": service.ServiceURI,
		"Secure":      service.Secure,
		"Version":     strconv.Itoa(service.Version),
		"Created At":  service.CreatedAt.Format(TimeLayout),
		"Updated At":  service.UpdatedAt.Format(TimeLayout),
	})

	printKeyValueTable("Service Definition", theme, map[string]string{
		"ID":                 strconv.Itoa(service.ServiceDefinition.ID),
		"Service Definition": service.ServiceDefinition.ServiceDefinition,
		"Created At":         service.ServiceDefinition.CreatedAt.Format(TimeLayout),
		"Updated At":         service.ServiceDefinition.UpdatedAt.Format(TimeLayout),
	})

	printKeyValueTable("Provider", theme, map[string]string{
		"ID":                  strconv.Itoa(service.Provider.ID),
		"System Name":         service.Provider.SystemName,
		"Address":             service.Provider.Address,
		"Port":                strconv.Itoa(service.Provider.Port),
		"Authentication Info": service.Provider.AuthenticationInfo,
		"Created At":          service.Provider.CreatedAt.Format(TimeLayout),
		"Updated At":          service.Provider.UpdatedAt.Format(TimeLayout),
	})

	// Print a separate table for each interface
	for _, iface := range service.Interfaces {
		printKeyValueTable("Interface", theme, map[string]string{
			"ID":             strconv.Itoa(iface.ID),
			"Interface Name": iface.InterfaceName,
			"Created At":     iface.CreatedAt.Format(TimeLayout),
			"Updated At":     iface.UpdatedAt.Format(TimeLayout),
		})
	}
}
