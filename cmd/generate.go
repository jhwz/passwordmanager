package cmd

import (
	"fmt"
	"strconv"

	"github.com/jhwz/passwordmanager/util"
	"github.com/spf13/cobra"
)

var Generate = &cobra.Command{
	Use:   "gen length",
	Short: "Generate random passwords",
	Long: `
Generates a random password of the given length.
The password will only contain letters, numbers and special characters [!, @, #, $, %, ^, &, /, ?, <, >, *]
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		length, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("require length of password to generate")
		}

		generated, err := util.GeneratePassword(length)
		if err != nil {
			return err
		}

		fmt.Printf("Generated password (%d): %s\n", length, generated)
		return nil
	},
}
