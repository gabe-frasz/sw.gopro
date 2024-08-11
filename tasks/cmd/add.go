package cmd

import (
	"database/sql"
	"fmt"

	"github.com/gabe-frasz/gopro/tasks/internal/app/usecases"
	"github.com/gabe-frasz/gopro/tasks/internal/infra/database/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		db, err := sql.Open("sqlite3", "tasks.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		taskRepository := repository.NewSqlTaskRepository(db)
		usecase := usecases.NewAddNewTask(args[0], taskRepository)
		err = usecase.Execute()
		if err != nil {
			panic(err)
		}

		fmt.Println("Task added")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
