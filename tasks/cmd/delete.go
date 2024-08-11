package cmd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gabe-frasz/gopro/tasks/internal/app/usecases"
	"github.com/gabe-frasz/gopro/tasks/internal/infra/database/repository"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <taskid>",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		db, err := sql.Open("sqlite3", "tasks.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		taskRepository := repository.NewSqlTaskRepository(db)
		deleteTask := usecases.NewDeleteTask(id, taskRepository)
		err = deleteTask.Execute()
		if err != nil {
			panic(err)
		}

		fmt.Println("Task deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
