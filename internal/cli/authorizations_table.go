package cli

import (
	"strconv"

	goprettytable "github.com/jedib0t/go-pretty/v6/table"
	"github.com/johankristianss/arrowhead/internal/table"
	"github.com/johankristianss/arrowhead/pkg/core"
	"github.com/muesli/termenv"
)

func printAuthsTable(auths []core.Authorization) {
	t := createAuthsTable(auths)
	t.Render()
}

// Generate a string of all interface names
func getInterfacesString(interfaces []core.Interface) string {
	interfacesStr := ""
	for i, intf := range interfaces {
		interfacesStr += intf.InterfaceName
		if i < len(interfaces)-1 {
			interfacesStr += ", "
		}
	}
	return interfacesStr
}

func createAuthsTable(auths []core.Authorization) *table.Table {
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
		{ID: "consumerName", Name: "Consumer System Name", SortIndex: 1, Width: 0},
		{ID: "providerName", Name: "Provider System Name", SortIndex: 2, Width: 0},
		{ID: "serviceDefinition", Name: "Service Definition", SortIndex: 3, Width: 0},
		{ID: "interfaces", Name: "Interfaces", SortIndex: 4, Width: 0},
	}

	t.SetCols(cols)

	for _, auth := range auths {
		idStr := strconv.Itoa(int(auth.ID))
		t.AddRow([]interface{}{
			termenv.String(string(idStr)).Foreground(theme.ColorCyan),
			termenv.String(auth.ConsumerSystem.SystemName).Foreground(theme.ColorViolet),
			termenv.String(auth.ProviderSystem.SystemName).Foreground(theme.ColorMagenta),
			termenv.String(auth.ServiceDefinition.ServiceDefinition).Foreground(theme.ColorGreen),
			termenv.String(getInterfacesString(auth.Interfaces)).Foreground(theme.ColorYellow),
		})
	}

	return t
}
