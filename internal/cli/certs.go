package cli

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(certsCmd)
	certsCmd.AddCommand(genCertCmd)
	genCertCmd.Flags().StringVarP(&SystemName, "name", "n", "", "System name")
	genCertCmd.MarkFlagRequired("SystemName")
}

var certsCmd = &cobra.Command{
	Use:   "certs",
	Short: "Manage Arrowhead PKCS#12 certificates",
	Long:  "Manage Arrowhead PKCS#12 certificates",
}

func isValidSystemName(systemName string) bool {
	// System name should only container letters and numbers and not be empty
	if len(systemName) == 0 {
		return false
	}

	for _, c := range systemName {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
			return false
		}
	}

	return true
}

var genCertCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a system certificate",
	Long:  "Generate a system certificate",
	Run: func(cmd *cobra.Command, args []string) {
		parseEnv()

		_, err := generateCert(SystemName)
		CheckError(err)
	},
}
