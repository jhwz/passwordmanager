package main

import (
	"fmt"
	"os"

	"github.com/jhwz/passwordmanager/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Adds our cobra commands to a root command and then executes it
func main() {
	root := &cobra.Command{
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	root.PersistentFlags().StringP("database", "d", "./password_db", "Path to the password database. Can be provided via environment variables. If not defined will initilise in current directory.")
	viper.BindPFlag("database", root.PersistentFlags().Lookup("database"))

	root.PersistentFlags().StringP("master", "m", "", "Master password for the database. Can provide via environment variables. Not recommended to store as plaintext anywhere!")
	viper.BindPFlag("master", root.PersistentFlags().Lookup("master"))

	viper.SetEnvPrefix("password")
	viper.AutomaticEnv()

	root.AddCommand(cmd.Retrieve)
	root.AddCommand(cmd.Generate)
	root.AddCommand(cmd.Store)

	if err := root.Execute(); err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
}
