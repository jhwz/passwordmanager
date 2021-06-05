package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/viper"
)

// CheckEnv makes sure the program has recieved both a database path, and a master password.
// Some commands require this, and should call this function at the start of execution
func CheckEnv() {
	// Make sure the database is a path, if it isn't then get it
	path := viper.GetString("database")
check:
	if path == "" {
		path = prompt.Input("Enter path to password database: ", func(d prompt.Document) []prompt.Suggest { return nil })
		viper.Set("database", path)
		goto check
	} else if info, err := os.Stat(path); err != nil || !info.IsDir() {
		response := prompt.Input(fmt.Sprintf("Create new password database at '%s'? (y/n): ", path), func(d prompt.Document) []prompt.Suggest { return nil })
		if strings.HasPrefix(strings.ToLower(response), "y") {
			os.MkdirAll(path, 0755)
		} else {
			os.Exit(0)
		}
		goto check
	}

	// Make sure a master password was provided
	if master := viper.GetString("master"); master == "" {
		master = prompt.Input("Enter master password: ", func(d prompt.Document) []prompt.Suggest { return nil })
		fmt.Println()
		viper.Set("master", master)
	}
}
