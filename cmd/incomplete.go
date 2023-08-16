package cmd

import (
	"cli-todos/pkg/todos"
	"log"

	"github.com/spf13/cobra"
)

var incompleteCmd = &cobra.Command{
	Use:     "incomplete",
	Short:   "Mark a todo as incomplete. toto <todo_id>",
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Too many arguments passed")
		}

		if err := todos.Incomplete(args[0]); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(incompleteCmd)
}
