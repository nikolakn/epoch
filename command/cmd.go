package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	jsonFile    string
	userLicense string
	Verbose     bool
	Print       bool
	rootCmd     = &cobra.Command{
		Use:   "epoch",
		Short: "CLI Epoch application for history, timelines",
		Long:  `Epoch`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do stuff here
		},
	}

	tryCmd = &cobra.Command{
		Use:   "try",
		Short: "Try and possibly fail at something",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("try funkcija")
			return nil
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "version",
		Long:  `version`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Epoch v0.1")
		},
	}

	authorCmd = &cobra.Command{
		Use:   "author",
		Short: "author",
		Long:  `program author`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Nikola Knezevic")
		},
	}
)

// Execute executes the root command.
func Execute() (string, bool, error) {
	if err := rootCmd.Execute(); err != nil {
		return "", true, err
	}
	return jsonFile, Print, nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&jsonFile, "file", "f", "", "json file to load")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&Print, "print", "p", false, "print only and exit")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(authorCmd)
}
