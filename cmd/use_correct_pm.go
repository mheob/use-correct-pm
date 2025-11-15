// Package cmd holds the root cobra command
package cmd

import (
	"fmt"
	"os"

	"github.com/mheob/use-correct-pm/internal/output"
	"github.com/mheob/use-correct-pm/internal/pm"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "use-correct-pm",
	Version: "1.2.0",
	Short:   "A simple check of the usage of the correct package manager",
	Long: `Check if the correct package manager is used.
Available are 'BUN', 'NPM', 'Yarn' and 'PNPM'.`,
	Run: func(_ *cobra.Command, args []string) {
		chosenPackageManager := pm.GetSelectedPackageManager(args)
		errors := pm.GetCheckLockFilesErrors(chosenPackageManager)
		hasErrors := len(errors) > 0

		if hasErrors {
			output.PrintInvalidMessage()

			for _, err := range errors {
				fmt.Println(err)
			}
			os.Exit(1)
		}

		output.PrintValidMessage()
	},
}

// Execute executes the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	for _, manager := range pm.AvailablePackageManagers {
		rootCmd.Flags().BoolP(manager.Name, manager.ShortFlag, false, manager.Description)
	}
}
