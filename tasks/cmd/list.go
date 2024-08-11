package cmd

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/gabe-frasz/gopro/tasks/internal/app/entity"
	"github.com/gabe-frasz/gopro/tasks/internal/app/usecases"
	"github.com/gabe-frasz/gopro/tasks/internal/infra/database/repository"
	"github.com/spf13/cobra"
)

var (
	allFlag bool
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "A brief description of your command",
		Run: func(cmd *cobra.Command, args []string) {
			db, err := sql.Open("sqlite3", "tasks.db")
			if err != nil {
				panic(err)
			}
			defer db.Close()

			taskRepository := repository.NewSqlTaskRepository(db)
			var tasks []*entity.Task

			if allFlag {
				usecase := usecases.NewListAllTasks(taskRepository)
				tasks, err = usecase.Execute()
				if err != nil {
					panic(err)
				}
			} else {
				usecase := usecases.NewListUncompletedTasks(taskRepository)
				tasks, err = usecase.Execute()
				if err != nil {
					panic(err)
				}
			}

			for _, task := range tasks {
				fmt.Println(task.Description + " | " + strconv.FormatBool(task.Done) + " | " + task.CreatedAt.Local().Format(time.RFC3339))
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "list all tasks")
}
