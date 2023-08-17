package cmd

import (
	"log"
	"strings"

	"github.com/jackaitken/toto/pkg/todos"

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
