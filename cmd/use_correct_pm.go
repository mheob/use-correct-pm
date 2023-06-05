// Package cmd holds the root cobra command
package cmd

import (
	"fmt"
	"os"

	"github.com/mheob/use-correct-pm/internal/pm"
	"github.com/mheob/use-correct-pm/internal/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "use-correct-pm",
	Version: "1.0.4",
	Short:   "A simple check of the usage of the correct package manager",
	Long: `Check if the correct package manager is used.
Available are 'NPM', 'Yarn' and 'PNPM'.`,
	Run: func(cmd *cobra.Command, args []string) {
		chosenPackageManager := pm.GetSelectedPackageManager(args)
		errors := pm.GetCheckLockFilesErrors(chosenPackageManager)
		hasErrors := len(errors) > 0

		if hasErrors {
			utils.PrintInvalidMessage()

			for _, err := range errors {
				fmt.Println(err)
			}
			os.Exit(1)
		}

		utils.PrintValidMessage()
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
