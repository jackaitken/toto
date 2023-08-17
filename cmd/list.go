package cmd

import (
	"fmt"
	"log"

	"github.com/jackaitken/toto/pkg/todos"

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
