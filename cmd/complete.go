package cmd

import (
	"cli-todos/pkg/todos"
	"log"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Short:   "Mark a todo list item as complete. toto complete <todo_id>",
	Aliases: []string{"c"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Only one todo at a time can be marked as complete")
		}

		if err := todos.Complete(args[0]); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
