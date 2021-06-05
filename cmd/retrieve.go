package cmd

import (
	"fmt"

	"github.com/jhwz/passwordmanager/db"
	"github.com/jhwz/passwordmanager/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Retrieve = &cobra.Command{
	Use:   "get id",
	Short: "Retrieve a password from the database",
	Long: `
Retrieve a password from storage. 
The ID must be the identifier for a password which has already been stored with the 'add' command.
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		util.CheckEnv()
		db, err := db.Open(viper.GetString("database"), viper.GetString("master"))
		if err != nil {
			return err
		}

		password, err := db.GetPassword(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("%s: %s\n", args[0], password)
		return nil
	},
}
