// Package cmd holds the root cobra command
package cmd

import (
	"fmt"
	"os"

	pm "github.com/mheob/use-correct-pm/internal/packageManager"
	"github.com/mheob/use-correct-pm/internal/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "use-correct-pm",
	Version: "1.0.1",
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
	rootCmd.Flags().BoolP(pm.NPM.Name, pm.NPM.ShortFlag, false, pm.NPM.Description)
	rootCmd.Flags().BoolP(pm.PNPM.Name, pm.PNPM.ShortFlag, false, pm.PNPM.Description)
	rootCmd.Flags().BoolP(pm.YARN.Name, pm.YARN.ShortFlag, false, pm.YARN.Description)
}
