package cmd

import (
	"cli-todos/pkg/todos"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new todo item for today. toto add <new todo text>",
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		newTodo := strings.Join(args, " ")
		if err := todos.CreateTodo(newTodo); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
