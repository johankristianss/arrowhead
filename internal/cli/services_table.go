package cli

import (
	"strconv"

	goprettytable "github.com/jedib0t/go-pretty/v6/table"
	"github.com/johankristianss/arrowhead/internal/table"
	"github.com/johankristianss/arrowhead/pkg/core"
	"github.com/muesli/termenv"
)

func printServicesTable(systems []core.Service) {
	t := createServicesTable(systems)
	t.Render()
}

func createServicesTable(services []core.Service) *table.Table {
	theme, err := table.LoadTheme("solarized-dark")
	CheckError(err)

	// Set table options:
	// - Columns: indices of the fields to display.
	// - SortBy: which column index to sort by (here, 0 for ID).
	// - Style: use the rounded style from go-pretty.
	opts := table.TableOptions{
		Columns: []int{0, 1, 2, 3, 4, 5, 6},
		SortBy:  0,
		Style:   goprettytable.StyleRounded,
	}

	t := table.NewTable(theme, opts, false)

	cols := []table.Column{
		{ID: "id", Name: "ID", SortIndex: 0, Width: 0},
		{ID: "providerName", Name: "Provider Name", SortIndex: 1, Width: 0},
		{ID: "serviceURI", Name: "URI", SortIndex: 2, Width: 0},
		{ID: "serviceDefinition", Name: "Service Definition", SortIndex: 3, Width: 0},
		{ID: "address", Name: "Address", SortIndex: 4, Width: 0},
		{ID: "port", Name: "Port", SortIndex: 5, Width: 0},
		{ID: "metadata", Name: "Metadata", SortIndex: 5, Width: 0},
	}

	t.SetCols(cols)

	for _, service := range services {
		idStr := strconv.Itoa(int(service.ID))
		portStr := strconv.Itoa(service.Provider.Port)

		metadataStr := ""
		for key, value := range service.Metadata {
			metadataStr += key + ": " + value
			metadataStr += ", "
		}
		if len(metadataStr) > 0 {
			metadataStr = metadataStr[:len(metadataStr)-2]
		}

		t.AddRow([]interface{}{
			termenv.String(string(idStr)).Foreground(theme.ColorCyan),
			termenv.String(service.Provider.SystemName).Foreground(theme.ColorViolet),
			termenv.String(service.ServiceURI).Foreground(theme.ColorMagenta),
			termenv.String(service.ServiceDefinition.ServiceDefinition).Foreground(theme.ColorYellow),
			termenv.String(service.Provider.Address).Foreground(theme.ColorBlue),
			termenv.String(string(portStr)).Foreground(theme.ColorGreen),
			termenv.String(metadataStr).Foreground(theme.ColorRed),
		})
	}

	return t
}
