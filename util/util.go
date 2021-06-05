package util

import (
	"fmt"
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/viper"
)

// CheckEnv makes sure the program has recieved both a database path, and a master password.
// Some commands require this, and should call this function at the start of execution
func CheckEnv() {
	// Make sure the database is a path, if it isn't then get it
	path := viper.GetString("database")
	get := true
	if path == "" {
		fmt.Println("No database path provided")
	} else if info, err := os.Stat(path); err != nil || !info.IsDir() {
		fmt.Printf("Database path '%s' not found\n", path)
	} else {
		get = false
	}

	if get {
		fmt.Println("Either provide path to existing database or path to create new database at")
		path = prompt.Input("Database path: ", func(d prompt.Document) []prompt.Suggest { return nil })
		os.MkdirAll(path, 0755)
		viper.Set("database", path)
	}

	// Make sure a master password was provided
	if master := viper.GetString("master"); master == "" {
		fmt.Println("No master password provided")
		master = prompt.Input("master password: ", func(d prompt.Document) []prompt.Suggest { return nil })
		viper.Set("master", master)
	}
}
