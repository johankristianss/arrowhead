package cli

import (
	"strconv"

	goprettytable "github.com/jedib0t/go-pretty/v6/table"
	"github.com/johankristianss/arrowhead/internal/table"
	"github.com/johankristianss/arrowhead/pkg/core"
	"github.com/muesli/termenv"
)

func printSystemsTable(systems []core.System) {
	t := createSystemsTable(systems)
	t.Render()
}

func createSystemsTable(systems []core.System) *table.Table {
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
		{ID: "systemName", Name: "System Name", SortIndex: 1, Width: 0},
		{ID: "address", Name: "Address", SortIndex: 2, Width: 0},
		{ID: "port", Name: "Port", SortIndex: 3, Width: 0},
	}

	t.SetCols(cols)

	for _, system := range systems {
		idStr := strconv.Itoa(int(system.ID))
		portStr := strconv.Itoa(system.Port)
		t.AddRow([]interface{}{
			termenv.String(string(idStr)).Foreground(theme.ColorCyan),
			termenv.String(system.SystemName).Foreground(theme.ColorViolet),
			termenv.String(system.Address).Foreground(theme.ColorMagenta),
			termenv.String(string(portStr)).Foreground(theme.ColorGreen),
		})
	}

	return t
}
