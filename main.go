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

		Long: `
Simple interface for adding and retrieving password from a password database.

Passwords can be generated, which is done in a cryptographically secure manner.

If performing multiple actions, recommend using environment variables to persist the path 
to your password database and master password across actions.
You should not however store the master password in a .bash_rc, .zsh_rc or using windows SETX, this
is not secure!
		
`,
	}

	root.PersistentFlags().StringP("database", "d", "", "Path to the password database. Can be provided via environment variables.")
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
