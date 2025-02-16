package cli

import (
	"strconv"

	"github.com/johankristianss/arrowhead/internal/table"
	"github.com/johankristianss/arrowhead/pkg/core"
)

func printSystemTable(system core.System) {
	theme, err := table.LoadTheme("solarized-dark")
	CheckError(err)

	printKeyValueTable("System Details", theme, map[string]string{
		"ID":          strconv.Itoa(system.ID),
		"System Name": system.SystemName,
		"Address":     system.Address,
		"Port":        strconv.Itoa(system.Port),
		"Created At":  system.CreatedAt.Format(TimeLayout),
		"Updated At":  system.UpdatedAt.Format(TimeLayout),
	})
}
