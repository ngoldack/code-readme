/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"os"

	"github.com/ngoldack/code-readme/internal"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config is used to generate a config file",
	Long:  `config is used to generate a config file`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Info("Creating config file...")

		file, err := os.Create(".code-readme.yaml")
		if _, err := os.Stat(".code-readme.yaml"); err == nil {
			forceFlag, err := cmd.Flags().GetBool("force")
			if err != nil {
				internal.Error("Could not get force flag: " + err.Error())
			}

			if !forceFlag {
				internal.Error("Config file already exists. Skipping...")
				return
			}

			internal.Warning("Config file already exists. Force flag is set. Overwriting")

		} else if !errors.Is(err, os.ErrNotExist) {
			internal.Error("Shroedingers error: " + err.Error())
			return
		}
		defer file.Close()

		if err != nil {
			internal.Error("Could not create config file: " + err.Error())
			return
		}

		cfg := internal.Config{}
		bytes, err := yaml.Marshal(cfg)
		if err != nil {
			internal.Error("Could not marshal config file: " + err.Error())
			return
		}

		_, err = file.Write(bytes)
		if err != nil {
			internal.Error("Could not write config file: " + err.Error())
			return
		}

		internal.Info("Config file created.")
		internal.Done()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	configCmd.Flags().BoolP("force", "f", false, "force overwrite of config file")
}
