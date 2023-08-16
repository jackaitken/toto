package cmd

import (
	"cli-todos/pkg/todos"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Show all todo items for today. toto list",
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		list, err := todos.Get()
		if err != nil {
			log.Fatal(err)
		}

		formattedList := todos.ParseList(list)

		fmt.Println(formattedList)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
