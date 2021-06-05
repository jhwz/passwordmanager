package cmd

import (
	"fmt"
	"strconv"

	"github.com/c-bata/go-prompt"
	"github.com/jhwz/passwordmanager/db"
	"github.com/jhwz/passwordmanager/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Store = &cobra.Command{
	Use:   "add id [password]",
	Short: "Add a new password to the database",
	Long: `
Adds a password to the manager database. 
If the password is not provided then a prompt will open asking
if you want to generate a random password as per the 'gen' command.
`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		util.CheckEnv()
		db, err := db.Open(viper.GetString("database"), viper.GetString("master"))
		if err != nil {
			return err
		}

		// still need to get a password, offer to generate one.
		if len(args) == 1 {
			fmt.Println("No password provided! Either provide a length to generate one, or press enter to exit")
			str := prompt.Input("password length: ", func(d prompt.Document) []prompt.Suggest { return nil })
			length, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("No valid length provided, exiting...")
				return nil
			}
			password, err := util.GeneratePassword(length)
			if err != nil {
				return err
			}
			fmt.Printf("Generated password (%d): %s\n", length, password)
			args = append(args, password)
		}

		if err := db.AddPassword(args[0], args[1]); err != nil {
			return err
		}

		fmt.Printf("Successfully added password for '%s'\n", args[0])
		return nil
	},
}
