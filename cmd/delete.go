package cmd

import (
	"log"

	"github.com/jackaitken/toto/pkg/todos"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a todo. toto <todo_id>",
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Too many arguments")
		}

		if err := todos.Delete(args[0]); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
