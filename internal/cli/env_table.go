package cli

import (
	"strconv"

	"github.com/muesli/termenv"
)

func printEnvTable() {
	t, theme := createTable(1)

	ascii := "false"
	if ASCII {
		ascii = "true"
	}
	row := []interface{}{
		termenv.String("ASCII").Foreground(theme.ColorCyan),
		termenv.String(ascii).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	verbose := "false"
	if Verbose {
		verbose = "true"
	}
	row = []interface{}{
		termenv.String("Verbose").Foreground(theme.ColorCyan),
		termenv.String(verbose).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	row = []interface{}{
		termenv.String("KeystorePassword").Foreground(theme.ColorCyan),
		termenv.String(KeystorePassword).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	row = []interface{}{
		termenv.String("RootKeystorePath").Foreground(theme.ColorCyan),
		termenv.String(RootKeystorePath).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	row = []interface{}{
		termenv.String("RootKeystoreAlias").Foreground(theme.ColorCyan),
		termenv.String(RootKeystoreAlias).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	row = []interface{}{
		termenv.String("CloudKeystorePath").Foreground(theme.ColorCyan),
		termenv.String(CloudKeystorePath).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	row = []interface{}{
		termenv.String("CloudKeystoreAlias").Foreground(theme.ColorCyan),
		termenv.String(CloudKeystoreAlias).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	row = []interface{}{
		termenv.String("SysOpsKeystorePath").Foreground(theme.ColorCyan),
		termenv.String(SysOpsKeystorePath).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	row = []interface{}{
		termenv.String("Truststore").Foreground(theme.ColorCyan),
		termenv.String(Truststore).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	arrowheadTLSStr := "false"
	if ArrowheadTLS {
		arrowheadTLSStr = "true"
	}

	row = []interface{}{
		termenv.String("ArrowheadTLS").Foreground(theme.ColorCyan),
		termenv.String(arrowheadTLSStr).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	row = []interface{}{
		termenv.String("ArrowheadAuthorizationHost").Foreground(theme.ColorCyan),
		termenv.String(ArrowheadAuthorizationHost).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	arrowheadAuthorizationPort := strconv.Itoa(ArrowheadAuthorizationPort)
	row = []interface{}{
		termenv.String("ArrowheadAuthorizationPort").Foreground(theme.ColorCyan),
		termenv.String(arrowheadAuthorizationPort).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	row = []interface{}{
		termenv.String("ArrowheadServiceRegistryHost").Foreground(theme.ColorCyan),
		termenv.String(ArrowheadServiceRegistryHost).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	arrowheadServiceRegistryPort := strconv.Itoa(ArrowheadServiceRegistryPort)
	row = []interface{}{
		termenv.String("ArrowheadServiceRegistryPort").Foreground(theme.ColorCyan),
		termenv.String(arrowheadServiceRegistryPort).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	row = []interface{}{
		termenv.String("ArrowheadOrchestratorHost").Foreground(theme.ColorCyan),
		termenv.String(ArrowheadOrchestratorHost).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	arrowheadOrchestratorPort := strconv.Itoa(ArrowheadOrchestratorPort)
	row = []interface{}{
		termenv.String("ArrowheadOrchestratorPort").Foreground(theme.ColorCyan),
		termenv.String(arrowheadOrchestratorPort).Foreground(theme.ColorGray),
	}
	t.AddRow(row)

	t.Render()
}
